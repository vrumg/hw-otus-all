package tasker

import (
	"fmt"
	"sync"
)

//Run func
func Run(tasks []func() error, threadsNum int, errorsMax int) error {

	//check incoming number of routines
	if threadsNum < 1 {
		return fmt.Errorf("NUMBER OF ROUTINES IS LESS THAN 1. Aborting")
	}

	tasksLen := len(tasks)
	//check incoming number of tasks
	if tasksLen < 1 {
		return fmt.Errorf("NUMBER OF TASKS IS LESS THAN 1. Aborting")
	}
	//adjust number of goroutines
	if tasksLen < threadsNum {
		//make shorter chan
		threadsNum = tasksLen
	}

	//check incoming number of errors
	if errorsMax <= 0 {
		errorsMax = -1
	}

	//channel for errors return
	chErrors := make(chan error, threadsNum)

	//channel for queue
	chQueue := make(chan struct{}, threadsNum)

	//channel for tasks creation termination
	chQuit := make(chan struct{})

	//wait group for active tasks
	var wg sync.WaitGroup

	//errors counter
	errorsCount := 0

	//start tasks execution in goroutine with chQueue
	go func() {

		//terminate condition
		done := false

		for _, task := range tasks {

			select {
			//check of chQueue is ready to recieve new task
			case chQueue <- struct{}{}:
				//inc wait group before starting next task
				wg.Add(1)
				//start next task
				go func(currTask func() error, wg *sync.WaitGroup) {
					defer wg.Done()
					err := currTask()
					//write return error to chErrors
					chErrors <- err
				}(task, &wg)

			//terminate channel
			case <-chQuit:
				done = true
			}

			//terminate routine
			if done {
				break
			}
		}
	}()

	for i := 0; i < tasksLen; i++ {
		//break rule for the task error reading loop
		breakRule := false

		select {
		//read next error from finished task
		case err := <-chErrors:
			if err != nil {
				errorsCount++
			}
			//check max number of errors and break
			if errorsCount == errorsMax {
				breakRule = true
			}

			//free one chQueue position
			<-chQueue
		}

		//terminate error reading loop on max errors
		if breakRule {
			chQuit <- struct{}{}
			break
		}
	}

	//wait for all active tasks to finish
	//and close channels
	wg.Wait()
	close(chQueue)
	close(chErrors)

	//check errors and return
	if errorsCount == errorsMax {
		return fmt.Errorf("MAX ERRORS LIMIT: %d", errorsCount)
	}
	return nil
}
