package cmd

import "mps_notas_back/internal/buisness/service"

type Command interface {
	Execute() (interface{}, error)
}

type GetStatusCommand struct {
	StatusService *service.StatusService
}

// NewGetStatusCommand cria uma nova instância do comando GetStatusCommand
func NewGetStatusCommand(statusService *service.StatusService) *GetStatusCommand {
	return &GetStatusCommand{
		StatusService: statusService,
	}
}

func (c *GetStatusCommand) Execute() (interface{}, error) {
	return c.StatusService.GetStatus()
}
