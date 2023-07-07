package server

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type SessionsDTO struct {
	Data []SessionDTO `json:"data"`
}

type SessionDTO struct {
	ID        string `json:"id"`
	IsRunning bool   `json:"is_running"`
}

type SessionDetailDTO struct {
	Data SessionDetail `json:"data"`
}

type SessionDetail struct {
	ID        string   `json:"id"`
	IsRunning bool     `json:"is_running"`
	Runs      []RunDTO `json:"runs"`
}

type RunDTO struct {
	ID              string    `json:"id"`
	StartTime       time.Time `json:"start_time"`
	DurationSeconds int       `json:"duration_seconds"`
	Live            bool      `json:"live"`
}

type RunDetailDTO struct {
	Data RunDetail `json:"data"`
}

type RunDetail struct {
	ID              string    `json:"id"`
	StartTime       time.Time `json:"start_time"`
	DurationSeconds int       `json:"duration_seconds"`
	Live            bool      `json:"live"`
	Data            []float32 `json:"data"`
}

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func CORSMiddleware(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", "GET")
		header.Set("Access-Control-Allow-Origin", "*")

		fn(w, req, ps)
	}
}

func (a *API) RegisterRoutes(router *httprouter.Router) {
	router.POST("/api/login", CORSMiddleware(a.ListSessions))
	router.GET("/api/sessions", CORSMiddleware(a.ListSessions))
	router.GET("/api/sessions/:id", CORSMiddleware(a.GetSession))
	router.POST("/api/sessions/:id/start", CORSMiddleware(a.StartRun))
	router.POST("/api/sessions/:id/stop", CORSMiddleware(a.StopRun))
	router.GET("/api/sessions/:id/runs/:run_id", CORSMiddleware(a.GetRunDetail))
}

func (a *API) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (a *API) ListSessions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sessions := a.sessManager.Sessions()

	resp := SessionsDTO{}
	dtos := make([]SessionDTO, len(sessions))
	for i := range sessions {
		dtos[i] = SessionDTO{
			ID:        sessions[i].ID(),
			IsRunning: sessions[i].IsRunning(),
		}
	}

	resp.Data = dtos

	json.NewEncoder(w).Encode(resp)
}

func (a *API) GetSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess, ok := a.sessManager.GetSession(ps.ByName("id"))
	if !ok {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "not found", StatusCode: 404})
		return
	}

	dto := SessionDetail{
		ID:        sess.ID(),
		IsRunning: sess.IsRunning(),
		Runs:      []RunDTO{},
	}

	for _, run := range sess.Runs() {
		dur := time.Since(run.StartTime)
		live := true
		if !run.EndTime.IsZero() {
			dur = run.EndTime.Sub(run.StartTime)
			live = false
		}
		dto.Runs = append(dto.Runs, RunDTO{
			ID:              run.ID,
			StartTime:       run.StartTime,
			DurationSeconds: int(dur.Seconds()),
			Live:            live,
		})
	}

	json.NewEncoder(w).Encode(SessionDetailDTO{Data: dto})
}

func (a *API) GetRunDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess, ok := a.sessManager.GetSession(ps.ByName("id"))
	if !ok {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "session not found", StatusCode: 404})
		return
	}

	run, ok := sess.GetRun(ps.ByName("run_id"))
	if !ok {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "run not found", StatusCode: 404})
		return
	}

	dur := time.Since(run.StartTime)
	live := true
	if !run.EndTime.IsZero() {
		dur = run.EndTime.Sub(run.StartTime)
		live = false
	}

	fmt.Println(run)

	dto := RunDetail{
		ID:              sess.ID(),
		StartTime:       run.StartTime,
		DurationSeconds: int(dur.Seconds()),
		Live:            live,
	}

	for _, df := range run.Data() {
		dto.Data = append(dto.Data, df.Value)
	}

	json.NewEncoder(w).Encode(RunDetailDTO{Data: dto})
}

func (a *API) StartRun(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("HTTP StartRun")
	sess, ok := a.sessManager.GetSession(ps.ByName("id"))
	if !ok {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "not found", StatusCode: 404})
		return
	}

	if sess.IsRunning() {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "already running", StatusCode: 400})
		return
	}

	if err := sess.StartRun(); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorResponse{Message: fmt.Sprintf("failed to start run: %v", err)})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{})
}

func (a *API) StopRun(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("HTTP StopRun")
	sess, ok := a.sessManager.GetSession(ps.ByName("id"))
	if !ok {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "not found", StatusCode: 404})
		return
	}

	if !sess.IsRunning() {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "not running", StatusCode: 400})
		return
	}

	if err := sess.StopRun(); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorResponse{Message: fmt.Sprintf("failed to stop run: %v", err)})
		return
	}

	json.NewEncoder(w).Encode(map[string]any{})
}
