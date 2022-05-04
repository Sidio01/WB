package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Worker(wg *sync.WaitGroup, jobs chan string, quit chan bool, n int) {
	wg.Add(1)
	for {
		select {
		case work := <-jobs:
			log.Printf("worker_%d doing %v\n", n, work)
			time.Sleep(3 * time.Second)
		case <-quit:
			log.Printf("worker_%d stop working\n", n)
			wg.Done()
			return
		}
	}

}
func main() {
	jobs := []string{"job_1", "job_2", "job_3", "job_4", "job_5"}
	wg := &sync.WaitGroup{}
	cancel := make(chan bool)
	jobsChan := make(chan string)
	sigChan := make(chan os.Signal, 1)
	nWorkers := 5

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for i := 0; i < nWorkers; i++ {
		go Worker(wg, jobsChan, cancel, i)
	}
	for {
		select {
		case <-sigChan:
			close(cancel)
			log.Println("stopping")
			wg.Wait()
			return
		default:
			jobsChan <- jobs[rand.Intn(5)]
		}
	}
}
