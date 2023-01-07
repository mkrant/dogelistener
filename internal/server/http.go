package server

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type SessionDTO struct {
	ID        string    `json:"id"`
	IsRunning bool      `json:"is_running"`
	Data      []float32 `json:"data"`
}

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (a *API) RegisterRoutes(router *httprouter.Router) {
	router.GET("/sessions", a.ListSessions)
	router.GET("/sessions/:id", a.GetSession)
	router.POST("/sessions/:id/start", a.StartRun)
	router.POST("/sessions/:id/stop", a.StopRun)
}

func (a *API) ListSessions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sessions := a.sessManager.Sessions()

	dtos := make([]SessionDTO, len(sessions))
	for i := range sessions {
		dtos[i] = SessionDTO{
			ID:        sessions[i].ID(),
			IsRunning: sessions[i].IsRunning(),
			Data:      sessions[i].Data(),
		}
	}

	json.NewEncoder(w).Encode(dtos)
}

func (a *API) GetSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sess, ok := a.sessManager.GetSession(ps.ByName("id"))
	if !ok {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "not found"})
		return
	}

	dto := SessionDTO{
		ID:        sess.ID(),
		IsRunning: sess.IsRunning(),
		Data:      sess.Data(),
	}

	json.NewEncoder(w).Encode(dto)
}

func (a *API) StartRun(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("HTTP StartRun")
	sess, ok := a.sessManager.GetSession(ps.ByName("id"))
	if !ok {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "not found"})
		return
	}

	if sess.IsRunning() {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "already running"})
		return
	}

	if err := sess.StartRun(); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ErrorResponse{Message: fmt.Sprintf("failed to start run: %v", err)})
		return
	}

	w.WriteHeader(http.StatusNoContent)
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

	w.WriteHeader(http.StatusNoContent)
}
