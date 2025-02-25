package handler

import (
	"encoding/json"
	"mps_notas_back/internal/model"
	"mps_notas_back/internal/service"
	"net/http"
	"strconv"
)

// UserHandler lida com as requisições HTTP relacionadas a usuários
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler cria uma nova instância do manipulador de usuários
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetAllUsers retorna todos os usuários
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.userService.GetAllUsers()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID retorna um usuário específico pelo ID
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Extrair ID da URL
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	
	user := h.userService.GetUserByID(id)
	if user == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser cria um novo usuário
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input model.NewUserInput
	
	// Decodificar corpo da requisição
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	
	// Validação básica, prox atividade
	// if input.Name == "" || input.Email == "" {
	//     w.Header().Set("Content-Type", "application/json")
	//     w.WriteHeader(http.StatusBadRequest)
	//     json.NewEncoder(w).Encode(map[string]string{"error": "Name and email are required"})
	//     return
	// }
	
	user := h.userService.CreateUser(input)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}