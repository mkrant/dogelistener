package recorder

import (
	"fmt"
	"log"
	"os/exec"
)

type Record struct {
	pythonPath        string
	processScriptPath string
}

func (r *Record) UploadData(t []float32, energy []float32, index string) error {
	return nil
}

func (r *Record) ProcessFile(f string) error {
	log.Println("Processing file " + f)
	return nil

	if err := exec.Command(r.pythonPath, r.processScriptPath, f).Run(); err != nil {
		return fmt.Errorf("running process script: %w", err)
	}

	return nil
}
