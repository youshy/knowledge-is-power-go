// Based on https://www.opsdash.com/blog/job-queues-in-go.html
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	Text string
}

var wg sync.WaitGroup

func newJob(text string) Job {
	var j Job
	j.Text = text
	return j
}

func process(j Job) {
	fmt.Printf("Job text: %v\n", j.Text)
}

func worker(jobChan <-chan Job) {
	defer wg.Done()

	for job := range jobChan {
		i := time.Duration(rand.Intn(100))
		time.Sleep(time.Millisecond * i)
		process(job)
	}
}

// Worker but with a setup for cancellation
func cancelworker(ctx context.Context, jobChan <-chan Job) {
	for {
		select {
		case <-ctx.Done():
			return
		case job := <-jobChan:
			process(job)
		}
	}
}

// TryEnqueue tries to enqueue a job to the given job channel.
// Returns true if the operation was successful, and false
// if enqueuing would not have been possible without blocking.
// Job is not enqueueud in the latter case.
/*
func TryEnqueue(job Job, jobChan <-chan Job) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}
*/

// WaitTimeout does a Wait of a sync.WaitGroup object but with a specified
// timeout. Returns true if the wait completed without timing out, false otherwise.
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	ch := make(chan struct{})
	go func() {
		wg.Wait()
		close(ch)
	}()
	select {
	case <-ch:
		return true
	case <-time.After(timeout):
		return false
	}
}

func main() {
	jobChan := make(chan Job, 100)

	// wg will wait for the worker to finish, then close
	wg.Add(1)
	go worker(jobChan)

	one := newJob("Me")
	two := newJob("You")

	jobChan <- one
	jobChan <- two
	jobChan <- one
	jobChan <- one
	jobChan <- two

	// closing the channel after finishing
	close(jobChan)
	WaitTimeout(&wg, 5*time.Second)

	time.Sleep(time.Second * 10)
}
