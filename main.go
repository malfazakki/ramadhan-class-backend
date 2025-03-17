package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type Task interface {
	Execute() error
	Info() string
}

type BaseTask struct {
	ID          string
	CreatedAt   time.Time
	Description string
}

type EmailTask struct {
	BaseTask
}

type SMSTask struct {
	BaseTask
}

type ReportTask struct {
	BaseTask
}

type TaskManager struct {
	tasks  []Task
	failed []Task
}

func (e *EmailTask) Execute() error {
	fmt.Printf("Executing Email Task: %s\n", e.ID)

	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	if rand.Float32() < .2 {
		return fmt.Errorf("EmailTask %s failed", e.ID)
	}

	return nil
}

func (e *EmailTask) Info() string {
	return fmt.Sprintf("EmailTask %s - %s", e.ID, e.Description)
}

func (s *SMSTask) Execute() error {
	fmt.Printf("Executing SMSTask: %s\n", s.ID)

	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	if rand.Float32() < .2 {
		return fmt.Errorf("SMSTask %s failed", s.ID)
	}

	return nil
}

func (s *SMSTask) Info() string {
	return fmt.Sprintf("SMSTask %s - %s", s.ID, s.Description)
}

func (r *ReportTask) Execute() error {
	fmt.Printf("Executing ReportTask: %s\n", r.ID)

	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	if rand.Float32() < .2 {
		return fmt.Errorf("ReportTask %s failed", r.ID)
	}

	return nil
}

func (tm *TaskManager) AddTask(t Task) {
	tm.tasks = append(tm.tasks, t)
}

func (tm *TaskManager) ExecuteAll() {
	for _, task := range tm.tasks {
		start := time.Now()
		err := task.Execute()
		duration := time.Since(start)
		typeName := reflect.TypeOf(task).Elem().Name()

		if err != nil {
			fmt.Printf("[ERROR] %s failed: %s\n", typeName, err)
			tm.failed = append(tm.failed, task)
		} else {
			fmt.Printf("[SUCCESS] %s completed in %v\n", typeName, duration)
		}
	}
}

func (tm *TaskManager) RetryFailedTask() {
	if len(tm.failed) == 0 {
		fmt.Println("No failed tasks to retry")
		return
	}

	fmt.Println("Retrying failed tasks...")
	tm.tasks = tm.failed
	tm.failed = nil
	tm.ExecuteAll()
}

func main() {

}
