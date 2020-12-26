package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

// ProcessLog a
type ProcessLog struct {
	st time.Time
	ed time.Time
}

// Meatadata a
type Meatadata struct {
	info string
}

// Message a
type Message struct {
	image string
	logs  []ProcessLog
	meta  Meatadata
}

func getRand(n, base int) int {
	return base + rand.Intn(n)
}

// downloads a image and sends the msg to the out channel.
func downloadImage(done <-chan bool, imageID int, out chan<- Message, w *sync.WaitGroup) {
	img := fmt.Sprintf("Image %v", imageID)
	msg := Message{
		image: img,
		logs:  make([]ProcessLog, 0),
	}
	processingTime := time.Duration(getRand(2000, 0)) * time.Millisecond
	msg.logs = append(msg.logs, ProcessLog{
		st: time.Now(),
		ed: time.Now().Add(processingTime),
	})
	fmt.Printf("[Download image-%v starts]Its gonna take %v \n", imageID, processingTime)
	time.Sleep(processingTime)
	fmt.Printf("[Download image-%v finishes]\n", imageID)
	w.Done()
}

// loads data from disk. each operation takes ~250 - 500 ms each.
// slow hard disk, large image what can I say!!!
func downoadImagesFromS3(done <-chan bool, imageIds ...int) chan Message {
	out := make(chan Message)
	go func() {
		var w sync.WaitGroup
		defer close(out)
		for _, imageID := range imageIds {
			w.Add(1)
			go downloadImage(done, imageID, out, &w)
		}
		// waiting until all downloads are finished, we can then close out channel
		w.Wait()
		fmt.Println("Download go routines are spawned and working in the background(producing message for which main is waiting)")
	}()
	return out
}

func extractMetadata(msg Message, out chan<- Message, w *sync.WaitGroup) {
	processingTime := time.Duration(getRand(750, 750)) * time.Millisecond
	msg.logs = append(msg.logs, ProcessLog{
		st: time.Now(),
		ed: time.Now().Add(processingTime),
	})
	msg.meta = Meatadata{
		info: fmt.Sprintf("Image meta data for %s. size = %v MB", msg.image, getRand(5, 5)),
	}
	fmt.Printf("[Extract metadata from-%v starts]Its gonna take %v \n", msg.image, processingTime)
	time.Sleep(processingTime)
	fmt.Printf("[Download-%v finishes]\n", msg.image)
	out <- msg
	w.Done()
}

func extractMetadataBatch(done <-chan bool, ch <-chan Message) <-chan Message {
	out := make(chan Message)
	go func() {
		var w sync.WaitGroup
		defer close(out)
		for msg := range ch {
			w.Add(1)
			go extractMetadata(msg, out, &w)
		}
		w.Wait()
	}()
	return out
}

func persist(msg Message, out chan<- Message, w *sync.WaitGroup) {
	processingTime := time.Duration(getRand(125, 125)) * time.Millisecond
	msg.logs = append(msg.logs, ProcessLog{
		st: time.Now(),
		ed: time.Now().Add(processingTime),
	})
	fmt.Printf("[Persist-%v starts]Its gonna take %v \n", msg.image, processingTime)
	time.Sleep(processingTime)
	fmt.Printf("[Persist-%v finishes]\n", msg.image)
	out <- msg
	w.Done()
}
func persistBatch(done <-chan bool, ch <-chan Message) <-chan Message {
	out := make(chan Message)
	go func() {
		var w sync.WaitGroup
		defer close(out)
		for msg := range ch {
			w.Add(1)
			go persist(msg, out, &w)
		}
		w.Wait()
	}()
	return out
}

func merge(done <-chan bool, channels ...<-chan Message) <-chan Message {
	out := make(chan Message)
	go func() {
		defer close(out)
		for _, ch := range channels {
			for msg := range ch {
				out <- msg
			}
		}
	}()
	return out
}
func main() {
	done := make(chan bool)

	for message := range persistBatch(done,
		merge(done,
			extractMetadataBatch(done, downoadImagesFromS3(done, 1, 2, 3)),
			extractMetadataBatch(done, downoadImagesFromS3(done, 4, 5, 6)),
			extractMetadataBatch(done, downoadImagesFromS3(done, 7, 8, 9)),
		)) {
		fmt.Printf("Finished processing for message %s", message.image)
		if strings.Contains(message.image, "4") {
			done <- true
		}

	}

	// for range persistBatch(extractMetadataBatch(downoadImagesFromS3(1, 2, 3, 4, 5, 6, 7, 8, 9))) {

	// }

	fmt.Printf("Number of active go routine %v\n", runtime.NumGoroutine())
}
