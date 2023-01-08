package sessions

import (
	"errors"
	"github.com/mkrant/dogelistener/internal/server/api"
	"time"
)

type Session struct {
	id        string
	isRunning bool
	runs      []*DataRun
	stream    api.DogeServer_ConnectServer
}

func NewSession(id string, stream api.DogeServer_ConnectServer) *Session {
	return &Session{
		id:     id,
		stream: stream,
	}
}

func (s *Session) StartRun() error {
	s.isRunning = true
	s.runs = append([]*DataRun{NewDataRun()}, s.runs...)

	return s.stream.Send(&api.Response{Type: &api.Response_StartRun{StartRun: &api.StartRun{}}})
}

func (s *Session) StopRun() error {
	if !s.isRunning {
		return errors.New("not running")
	}

	s.isRunning = false

	s.runs[0].EndTime = time.Now()

	return s.stream.Send(&api.Response{Type: &api.Response_EndRun{EndRun: &api.EndRun{}}})
}

func (s *Session) GetRun(id string) (*DataRun, bool) {
	var run *DataRun
	for _, r := range s.runs {
		if r.ID == id {
			run = r
			break
		}
	}

	return run, run != nil
}

func (s *Session) Runs() []*DataRun {
	return s.runs
}

func (s *Session) CurrentRun() (*DataRun, bool) {
	if !s.isRunning {
		return nil, false
	}

	return s.runs[0], true
}

func (s *Session) ID() string {
	return s.id
}

func (s *Session) IsRunning() bool {
	return s.isRunning
}

func (s *Session) SetRunning(b bool) {
	s.isRunning = b
}
