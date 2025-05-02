package handler

import (
	"encoding/json"
	"net/http"

	"mps_notas_back/internal/cmd"
)

type StatusHandler struct {
	Command cmd.Command
}

// NewStatusHandler cria uma nova instância do manipulador de status
func NewStatusHandler(statusService *cmd.GetStatusCommand) *StatusHandler {
	return &StatusHandler{
		Command: statusService,
	}
}

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result, err := h.Command.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
