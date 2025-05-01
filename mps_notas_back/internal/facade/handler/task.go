package handler

import (
	"encoding/json"
	"mps_notas_back/internal/buisness/service"
	"mps_notas_back/internal/infra/model"
	"net/http"
	"strconv"
)

// TaskHandler lida com as requisições HTTP relacionadas a tarefas
type TaskHandler struct {
	taskService *service.TaskService
}

// NewTaskHandler cria uma nova instância do manipulador de tarefas
func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

// GetAll retorna todas as tarefas
func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.taskService.GetAllTask()
	if err != nil {
		http.Error(w, "Error Interno"+ err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GetTaskByID retorna uma tarefa específica pelo ID
func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// Extrair ID da URL
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	task, err := h.taskService.GetTaskByID(id)
	if task == nil || err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "task not found" + err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// CreateTask cria uma nova tarefa
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var input model.NewTaskInput

	// Decodificar corpo da requisição
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	err = h.taskService.CreateTask(input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error while trying to process the entity!" + err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Tarefa criada com sucesso!"})

}

// UpdateTask atualiza uma tarefa
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var input model.NewTaskInput
	// Extrair ID da URL
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	// Decodificar corpo da requisição
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	task, err := h.taskService.UpdateTask(id, input)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error while trying to process the entity!" + err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// GetTaskByID retorna uma tarefa específica pelo ID
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Extrair ID da URL
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.taskService.DeleteTask(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "task not found" + err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "task Deleted Successfully"})
}
