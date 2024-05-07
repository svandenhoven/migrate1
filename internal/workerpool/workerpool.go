package workerpool

type JobResult struct {
	Error error
}

type jobFn[T any] func(chan T)
type workFn[T any] func(T) *JobResult

func EnqueueJobs[T any](maxWorkers int, jobsCh chan T, resultsCh chan *JobResult, enqueueJobsFn jobFn[T], executeWorkFn workFn[T]) {
	// Create a worker pool
	for i := 0; i < maxWorkers; i++ {
		go func() {
			for job := range jobsCh {
				resultsCh <- executeWorkFn(job)
			}
		}()
	}

	// Enqueue jobs
	go enqueueJobsFn(jobsCh)
}
