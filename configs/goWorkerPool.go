package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents the job to be run
type Job struct {
	ID      int
	Payload interface{}
}

// Worker represents the worker that executes the job
type Worker struct {
	ID          int
	JobQueue    chan Job
	WorkerQueue chan chan Job
	QuitChan    chan bool
}

// NewWorker creates a new worker
func NewWorker(id int, workerQueue chan chan Job) Worker {
	return Worker{
		ID:          id,
		JobQueue:    make(chan Job),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
	}
}

// Start method starts the worker by starting a goroutine
func (w Worker) Start(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			// Add ourselves into the worker queue
			w.WorkerQueue <- w.JobQueue

			select {
			case job := <-w.JobQueue:
				// Receive a job from the job queue
				fmt.Printf("Worker %d: Started job %d\n", w.ID, job.ID)
				// Simulate work
				// (replace this with actual work, e.g. job processing)
				fmt.Printf("Worker %d: Finished job %d\n", w.ID, job.ID)
			case <-w.QuitChan:
				// We have been asked to stop
				fmt.Printf("Worker %d stopping\n", w.ID)
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

// Dispatcher

// Dispatcher is responsible for managing the worker pool
type Dispatcher struct {
	WorkerQueue chan chan Job
	JobQueue    chan Job
	MaxWorkers  int
	Workers     []Worker
}

// NewDispatcher creates a new dispatcher
func NewDispatcher(maxWorkers int) *Dispatcher {
	workerQueue := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerQueue: workerQueue,
		JobQueue:    make(chan Job),
		MaxWorkers:  maxWorkers,
		Workers:     make([]Worker, maxWorkers),
	}
}

// Run starts the dispatcher and the workers
func (d *Dispatcher) Run() {
	var wg sync.WaitGroup
	// Start the workers
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i+1, d.WorkerQueue)
		d.Workers[i] = worker
		wg.Add(1)
		worker.Start(&wg)
	}

	// Start the dispatcher
	go d.dispatch(&wg)
}

func (d *Dispatcher) dispatch(wg *sync.WaitGroup) {
	for {
		select {
		case job := <-d.JobQueue:
			// A job request has been received
			go func(job Job) {
				// Try to obtain a worker job channel that is available.
				// This will block until a worker is idle
				jobChannel := <-d.WorkerQueue

				// Dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}

// Stop stops all the workers
func (d *Dispatcher) Stop() {
	for _, worker := range d.Workers {
		worker.Stop()
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 5

	dispatcher := NewDispatcher(numWorkers)
	dispatcher.Run()

	for i := 0; i < numJobs; i++ {
		job := Job{
			ID:      i + 1,
			Payload: fmt.Sprintf("Job payload %d", i+1),
		}
		dispatcher.JobQueue <- job
	}

	// Allow some time for jobs to be processed
	time.Sleep(2 * time.Second)

	// Stop the dispatcher and workers
	dispatcher.Stop()

	// Allow some time for workers to stop gracefully
	time.Sleep(1 * time.Second)
}
