package main

import (
	"os"

	fc "github.com/contribsys/faktory/client"
)

func main() {
	os.Setenv("faktory_address", "tcp://:Teamwork1@127.0.0.1:7419")
	os.Setenv("FAKTORY_PROVIDER", "faktory_address")
	println("Sending a job to the faktory")
	client, err := fc.Open()
	if err != nil {
		panic(err)
	}
	job := fc.NewJob("Test Job", "hello world")
	err = client.Push(job)

	if err != nil {
		panic(err)
	}
	println("Job Sent")
}
