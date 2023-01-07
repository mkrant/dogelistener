package iot

import (
	"bufio"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"github.com/mkrant/dogelistener/internal/iot/recorder"
	"log"
	"net"
	"os"
)

const (
	ProcessFileType = "process_file"
	UploadDataType  = "upload_data"
)

type Message struct {
	Type    string         `json:"type"`
	Payload map[string]any `json:"payload"`
}

type ProcessFilePayload struct {
	Filename string `json:"filename" mapstructure:"filename"`
}

type UploadDataPayload struct {
	Time   []float32 `json:"time" mapstructure:"time"`
	Energy []float32 `json:"energy" mapstructure:"energy"`
	Index  string    `json:"index" mapstructure:"index"`
}

type SocketServer struct {
	sockAddr string
	recorder *recorder.Recorder
}

func NewSocketServer(sock, serverAddr string) *SocketServer {
	return &SocketServer{
		sockAddr: sock,
		recorder: recorder.NewRecorder(serverAddr),
	}
}

func (s *SocketServer) Run() {
	if err := os.RemoveAll(s.sockAddr); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", s.sockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	for {
		// Accept new connections, dispatching them to echoServer
		// in a goroutine.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go s.handleConn(conn)
	}
}

func (s *SocketServer) handleConn(c net.Conn) {
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
	defer c.Close()

	record, ok := s.recorder.CurrentRecord()
	if !ok {
		log.Print("Socket conn received, but currently not recording")
		return
	}

	for {
		scanner := bufio.NewScanner(c)
		for scanner.Scan() {
			msg := &Message{}
			if err := json.Unmarshal(scanner.Bytes(), msg); err != nil {
				log.Printf("Invalid socket message json: %v", err)
				continue
			}

			log.Printf("Received message of type %q", msg.Type)

			switch msg.Type {
			case ProcessFileType:
				// Kick off a python script to process the file
				payload := &ProcessFilePayload{}
				if err := mapstructure.Decode(msg.Payload, payload); err != nil {
					log.Printf("Invalid socket message payload for type %s: %v", msg.Type, err)
					continue
				}

				go func() {
					if err := record.ProcessFile(payload.Filename); err != nil {
						log.Printf("Failed to process file %s: %v", payload.Filename, err)
					}
				}()
			case UploadDataType:
				// Process the data, add to run
				payload := &UploadDataPayload{}
				if err := mapstructure.Decode(msg.Payload, payload); err != nil {
					log.Printf("Invalid socket message payload for type %s: %v", msg.Type, err)
					continue
				}

				if err := record.UploadData(payload.Time, payload.Energy, payload.Index); err != nil {
					log.Printf("Failed to upload data for index %s: %v", payload.Index, err)
				}
			default:
				log.Printf("Invalid socket message type: %v", msg.Type)
				continue
			}
		}
	}
}
