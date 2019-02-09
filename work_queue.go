package work_queue



type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs         chan Worker
	Results      chan interface{}
	StopRequests chan string
	NumWorkers   uint
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
//adapted from: https://gobyexample.com/worker-pools
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	// TODO: initialize struct; start nWorkers workers as goroutines
	q.Jobs = make(chan Worker, maxJobs)
	q.Results = make(chan interface{}, nWorkers)
	q.StopRequests = make(chan string, nWorkers)
	q.NumWorkers = nWorkers
	var int_worker int=int(nWorkers)
	for i := 0; i < int_worker; i++ {
		go q.worker();
	}
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	
	var running bool= true
	// Run tasks from the Jobs channel, unless we have been asked to stop.

	//while loop...
	for running {
		//select is used for the Non-Blocking Channel Operations.
		//if there is no work on the work queue then the next case will be executed. Cause basic sends and receives
		//channels are blocking

		select {
		// TODO: listen on the .Jobs channel for incoming tasks
		// TODO: run tasks by calling .Run()
		// TODO: send the return value back on Results channel
		case work := <- queue.Jobs:
			queue.Results <- work.Run()
		// TODO: exit (return) when a signal is sent on StopRequests			
		case <- queue.StopRequests: 
			running = false
			
		}
	
	}
}

func (queue WorkQueue) Enqueue(work Worker) {
	// TODO: put the work into the Jobs channel so a worker can find it and start the task.
	queue.Jobs <- work
}

func (queue WorkQueue) Shutdown() {
	// TODO: tell workers to stop processing tasks.
	for i := 0; i < int(queue.NumWorkers); i++ {
		queue.StopRequests <- "STOP"
	}
}
