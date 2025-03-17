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

func main() {

}
