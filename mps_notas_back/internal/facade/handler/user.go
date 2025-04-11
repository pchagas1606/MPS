package handler

import (
	"encoding/json"
	"fmt"
	"mps_notas_back/internal/buisness/report"
	"mps_notas_back/internal/buisness/service"
	"mps_notas_back/internal/infra/model"
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
	users, err := h.userService.GetAllUsers()
	if err != nil {
		http.Error(w, "Error Interno", http.StatusInternalServerError)
		return
	}

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

	user, err := h.userService.GetUserByID(id)
	if user == nil || err != nil {
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

	user, err := h.userService.CreateUser(input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error while trying to process the entity!" + err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Login autentica um usuario na API e retorna um token
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input model.AuthUserInput

	// Decodificar corpo da requisição
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	token, err := h.userService.Login(input)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error while trying to authenticate"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}

// CreateUser cria um novo usuário
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input model.NewUserInput
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

	user, err := h.userService.UpdateUser(id, input)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error while trying to process the entity!" + err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUserByID retorna um usuário específico pelo ID
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Extrair ID da URL
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.userService.Delete(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found" + err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User Deleted Successfully"})
}

func (h *UserHandler) GenerateReport(w http.ResponseWriter, r *http.Request) {
	// Obter o formato do relatório da query string
	format := r.URL.Query().Get("format")
	if format == "" {
		format = "json" // formato padrão
	}

	// Obter todos os usuários
	users, err := h.userService.GetAllUsers()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Criar o gerador de relatório apropriado
	var generator report.ReportGenerator
	switch format {
	case "json":
		generator = &report.UserJSONReport{}
	case "csv":
		generator = &report.UserCSVReport{}
	default:
		http.Error(w, "Formato de relatório não suportado", http.StatusBadRequest)
		return
	}

	// Gerar o relatório usando o Template Method corretamente
	reportData, err := report.GenerateReportTemplate(generator, users)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao gerar relatório: %v", err), http.StatusInternalServerError)
		return
	}

	// Definir o tipo de conteúdo apropriado
	contentType := "application/json"
	if format == "csv" {
		contentType = "text/csv"
		w.Header().Set("Content-Disposition", "attachment; filename=users_report.csv")
	}

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(reportData)
}
