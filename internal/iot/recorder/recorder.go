// Package recorder contains the logic used to handle a single instance of a dog recording
package recorder

import (
	"context"
	"errors"
	"fmt"
	"github.com/mkrant/dogelistener/internal/server/api"
	"google.golang.org/grpc"
	"io"
	"log"
	"os/exec"
	"sync"
)

type Recorder struct {
	uploaderClient DataUploader
	rec            *Record

	pythonPath       string
	listenScriptPath string

	listenerCmd *exec.Cmd
	cancelCdm   context.CancelFunc

	mu sync.Mutex
}
type DataUploader interface {
	Send(*api.Request) error
}

func NewRecorder(addr string) *Recorder {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewDogeServerClient(conn)
	stream, err := client.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	recorder := &Recorder{
		uploaderClient: stream,
		rec: &Record{
			uploaderClient:    stream,
			pythonPath:        "python",
			processScriptPath: "scripts/express_stream.py",
		},
		pythonPath:       "python",
		listenScriptPath: "scripts/listen_stream.py",
	}

	go func() {
		if err := recorder.HandleStream(stream); err != nil {
			log.Printf("Recorder stream handle: %v", err)
			return
		}
	}()

	return recorder
}

func (r *Recorder) HandleStream(client api.DogeServer_ConnectClient) error {
	if err := client.Send(&api.Request{Type: &api.Request_Ping{Ping: &api.Ping{}}}); err != nil {
		return err
	}

	for {
		resp, err := client.Recv()
		if errors.Is(err, io.EOF) || errors.Is(err, context.Canceled) {
			return nil
		}

		if err != nil {
			return fmt.Errorf("stream recv: %w", err)
		}

		switch action := resp.Type.(type) {
		case *api.Response_StartRun:
			log.Println("StartRun", action)
			if err := r.StartRun(); err != nil {
				log.Printf("Failed to start recorder run: %v", err)
				continue
			}
		case *api.Response_EndRun:
			log.Println("EndRun")
			if err := r.StopRun(); err != nil {
				log.Printf("Failed to stop run: %v", err)
				continue
			}
		case *api.Response_Pong:
			log.Println("Got a pong")
		default:
			log.Printf("Got other type #T", action)
		}
	}
}

func (r *Recorder) StartRun() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.listenerCmd != nil {
		return errors.New("already running")
	}

	r.rec = &Record{
		uploaderClient:    r.uploaderClient,
		pythonPath:        "python",
		processScriptPath: "scripts/express_stream.py",
	}

	return r.startListenerScript()
}

func (r *Recorder) StopRun() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.cancelCdm != nil {
		r.cancelCdm()
	}

	r.listenerCmd = nil
	return nil
}

func (r *Recorder) startListenerScript() error {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, r.pythonPath, r.listenScriptPath)

	fmt.Println(r.pythonPath, r.listenScriptPath)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("starting listener script: %w", err)
	}

	r.cancelCdm = cancel
	r.listenerCmd = cmd

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
