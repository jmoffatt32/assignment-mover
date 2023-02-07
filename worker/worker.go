package worker

import (
	"os"

	"github.com/jmoffatt32/assignment-mover/arguments"
)

const SOURCE_DIR = "/Users/jackmoffatt/Downloads"
const DEST_DIR = "/Users/jackmoffatt/Documents/SPRING23"

type Worker interface {
	MoveFile() error
}

type WorkerInstance struct {
	Arguments arguments.ArgumentsController
}

func NewWorker(args arguments.ArgumentsController) (*WorkerInstance, error) {
	return &WorkerInstance{
		Arguments: args,
	}, nil
}

func (w *WorkerInstance) MoveFile() error {

	data, err := os.ReadFile(SOURCE_DIR + "/" + w.Arguments.GetFilename())
	if err != nil {
		return err
	}

	err = os.WriteFile(DEST_DIR+"/"+w.Arguments.GetClassname()+"/"+w.Arguments.GetFilename(), data, 0644)
	if err != nil {
		return err
	}

	return nil

}
