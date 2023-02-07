package main

import (
	"log"
	"os"

	"github.com/jmoffatt32/assignment-mover/arguments"
	"github.com/jmoffatt32/assignment-mover/worker"
)

func main() {

	rawArgs := os.Args
	args, err := arguments.NewArguments(rawArgs)
	if err != nil {
		log.Fatal(err)
	}

	wkr, err := worker.NewWorker(args)
	if err != nil {
		log.Fatal(err)
	}

	wkr.MoveFile()

}
