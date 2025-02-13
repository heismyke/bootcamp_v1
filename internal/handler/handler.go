package handler

import (
	"bootcamp_v1/internal/database"
	"net/http"
)

type Handler struct{
  server *database.Service
}

func New(server *database.Service) *Handler{
  return &Handler{
    server:server,
  }
}

func(h *Handler) CreateNewBootcamp(w http.ResponseWriter, r *http.Request)  {
  if r.Method != http.MethodPost{
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }

  w.WriteHeader(http.StatusAccepted)
  w.Write([]byte(`{
    "success": true,
    "message": "welcome to devcamper"
    }`))
}
