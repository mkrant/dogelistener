// Package recorder contains the logic used to handle a single instance of a dog recording
package recorder

import "sync"

type Recorder struct {
	rec *Record
	mu  sync.Mutex
}

func NewRecorder() *Recorder {
	return &Recorder{
		rec: &Record{
			pythonPath:        "python3",
			processScriptPath: "cmd/client/client.py",
		},
	}
}

// Start starts a new run
func (r *Recorder) Start() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.rec = &Record{}

	return nil
}

func (r *Recorder) Stop() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return nil
}

func (r *Recorder) CurrentRecord() (*Record, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.rec == nil {
		return nil, false
	}

	return r.rec, true
}
