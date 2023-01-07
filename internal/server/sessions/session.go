package sessions

import "github.com/mkrant/dogelistener/internal/server/api"

type Session struct {
	id        string
	isRunning bool
	data      []float32
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
	return s.stream.Send(&api.Response{Type: &api.Response_StartRun{StartRun: &api.StartRun{}}})
}

func (s *Session) StopRun() error {
	s.isRunning = false
	return s.stream.Send(&api.Response{Type: &api.Response_EndRun{EndRun: &api.EndRun{}}})
}

func (s *Session) AddData(data []float32) error {
	s.data = append(s.data, data...)
	return nil
}

func (s *Session) ID() string {
	return s.id
}

func (s *Session) IsRunning() bool {
	return s.isRunning
}

func (s *Session) Data() []float32 {
	return s.data
}

func (s *Session) SetRunning(b bool) {
	s.isRunning = b
}
