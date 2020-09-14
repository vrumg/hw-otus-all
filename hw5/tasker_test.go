package tasker

import (
	"errors"
	"testing"
)

func TestRunNoErrors(t *testing.T) {
	tasks := make([]func() error, 0, 6)

	var task1 = func() error {
		return nil
	}
	var task2 = func() error {
		return nil
	}
	var task3 = func() error {
		return nil
	}
	var task4 = func() error {
		return nil
	}
	var task5 = func() error {
		return nil
	}
	var task6 = func() error {
		return nil
	}
	tasks = append(tasks, task1)
	tasks = append(tasks, task2)
	tasks = append(tasks, task3)
	tasks = append(tasks, task4)
	tasks = append(tasks, task5)
	tasks = append(tasks, task6)

	threadsNum := 3
	errorsMax := 0
	wantErr := false

	t.Run("TestRunNoErrors", func(t *testing.T) {
		if err := Run(tasks, threadsNum, errorsMax); (err != nil) != wantErr {
			t.Errorf("Run() error = %v, wantErr %v", err, wantErr)
		}
	})

}

func TestRunWithErrors(t *testing.T) {
	tasks := make([]func() error, 0, 6)

	var task1 = func() error {
		return errors.New("task1 err")
	}
	var task2 = func() error {
		return errors.New("task2 err")
	}
	var task3 = func() error {
		return errors.New("task3 err")
	}
	var task4 = func() error {
		return errors.New("task4 err")
	}
	var task5 = func() error {
		return errors.New("task5 err")
	}
	var task6 = func() error {
		return errors.New("task6 err")
	}

	tasks = append(tasks, task1)
	tasks = append(tasks, task2)
	tasks = append(tasks, task3)
	tasks = append(tasks, task4)
	tasks = append(tasks, task5)
	tasks = append(tasks, task6)

	threadsNum := 3
	errorsMax := 2
	wantErr := true

	t.Run("TestRunWithErrors", func(t *testing.T) {
		if err := Run(tasks, threadsNum, errorsMax); (err != nil) != wantErr {
			t.Errorf("Run() error = %v, wantErr %v", err, wantErr)
		}
	})

}

func TestRunWrongThreads(t *testing.T) {
	tasks := make([]func() error, 0, 6)

	var task1 = func() error {
		return nil
	}
	var task2 = func() error {
		return nil
	}

	tasks = append(tasks, task1)
	tasks = append(tasks, task2)

	threadsNum := 0
	errorsMax := 0
	wantErr := true

	t.Run("TestRunWrongThreads", func(t *testing.T) {
		if err := Run(tasks, threadsNum, errorsMax); (err != nil) != wantErr {
			t.Errorf("Run() error = %v, wantErr %v", err, wantErr)
		}
	})

}

func TestRunWrongTasks(t *testing.T) {
	tasks := make([]func() error, 0, 6)

	threadsNum := 1
	errorsMax := 0
	wantErr := true

	t.Run("TestRunWrongTasks", func(t *testing.T) {
		if err := Run(tasks, threadsNum, errorsMax); (err != nil) != wantErr {
			t.Errorf("Run() error = %v, wantErr %v", err, wantErr)
		}
	})

}

func TestRunTasksLessThreads(t *testing.T) {
	tasks := make([]func() error, 0, 6)

	var task1 = func() error {
		return nil
	}
	var task2 = func() error {
		return nil
	}

	tasks = append(tasks, task1)
	tasks = append(tasks, task2)

	threadsNum := 10
	errorsMax := 0
	wantErr := false

	t.Run("TestRunTasksLessThreads", func(t *testing.T) {
		if err := Run(tasks, threadsNum, errorsMax); (err != nil) != wantErr {
			t.Errorf("Run() error = %v, wantErr %v", err, wantErr)
		}
	})

}
