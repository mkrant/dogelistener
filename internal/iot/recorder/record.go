package recorder

import (
	"fmt"
	"github.com/mkrant/dogelistener/internal/server/api"
	"log"
	"os/exec"
)

type Record struct {
	pythonPath        string
	processScriptPath string

	uploaderClient DataUploader
}

func (r *Record) UploadData(t []float32, energy []float32, index string) error {
	log.Println("Uploading data " + index)

	err := r.uploaderClient.Send(&api.Request{Type: &api.Request_RunData{RunData: &api.RunData{}}})
	if err != nil {
		return err
	}

	return nil
}

func (r *Record) ProcessFile(f string) error {
	log.Println("Processing file " + f)

	if err := exec.Command(r.pythonPath, r.processScriptPath, "scripts/tmp/"+f).Run(); err != nil {
		return fmt.Errorf("running process script: %w", err)
	}

	return nil
}
