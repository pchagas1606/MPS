package report

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"mps_notas_back/internal/infra/model"
	"strconv"
	"time"
)

//
// ===== INTERFACE =====
//

type ReportGenerator interface {
	FormatData(data interface{}) interface{}
	GenerateOutput(data interface{}) ([]byte, error)
}

//
// ===== TEMPLATE METHOD =====
//

func GenerateReportTemplate(r ReportGenerator, data interface{}) ([]byte, error) {

	// 1. Formatar dados
	formatted := r.FormatData(data)

	// 2. Gerar saída
	return r.GenerateOutput(formatted)
}

//
// ===== IMPLEMENTAÇÃO JSON =====
//

type UserJSONReport struct{}

func (r *UserJSONReport) FormatData(data interface{}) interface{} {
	users := data.([]model.UserDAO)
	return map[string]interface{}{
		"report_type":  "user_list",
		"total_users":  len(users),
		"users":        users,
		"generated_at": time.Now().Format(time.RFC3339),
	}
}

func (r *UserJSONReport) GenerateOutput(data interface{}) ([]byte, error) {
	return json.MarshalIndent(data, "", "  ")
}

//
// ===== IMPLEMENTAÇÃO CSV =====
//

type UserCSVReport struct{}

func (r *UserCSVReport) FormatData(data interface{}) interface{} {
	// CSV usa os dados diretamente
	return data
}

func (r *UserCSVReport) GenerateOutput(data interface{}) ([]byte, error) {
	users := data.([]model.UserDAO)
	var buffer bytes.Buffer
	writer := csv.NewWriter(&buffer)

	// Cabeçalho
	writer.Write([]string{"ID", "Nome", "Email", "Senha em Hash", "Criado Em"})

	// Dados
	for _, user := range users {
		writer.Write([]string{
			strconv.Itoa(user.ID),
			user.Name,
			user.Email,
			user.Password_Hash,
			user.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	writer.Flush()
	return buffer.Bytes(), writer.Error()
}
