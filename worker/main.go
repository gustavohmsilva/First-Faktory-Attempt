package main

import (
	"context"
	"fmt"
	"os"

	fw "github.com/contribsys/faktory_worker_go"
)

func main() {
	os.Setenv("faktory_address", "tcp://:password@127.0.0.1:7419")
	os.Setenv("FAKTORY_PROVIDER", "faktory_address")
	manager := fw.NewManager()
	manager.Register("Test Job", worker)
	manager.Run()
}

func worker(ctx context.Context, args ...interface{}) error {
	msg, ok := args[0].(string)
	if !ok {
		return fmt.Errorf(
			"oh oh, something happened! The string could not be parsed",
		)
	}
	fmt.Printf("Job Executed! Message: %v", msg)
	return nil
}
