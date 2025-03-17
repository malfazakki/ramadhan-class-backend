package main

import (
	"fmt"
	"math/rand"
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

func main() {

}
