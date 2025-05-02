package service

import (
	"runtime"
	"time"
)

type Status struct {
	Uptime       string `json:"uptime"`
	GoVersion    string `json:"go_version"`
	NumGoroutine int    `json:"num_goroutine"`
}

type StatusService struct {
	startTime time.Time
}

func NewStatusService() *StatusService {
	return &StatusService{startTime: time.Now()}
}

func (s *StatusService) GetStatus() (*Status, error) {
	uptime := time.Since(s.startTime).String()
	return &Status{
		Uptime:       uptime,
		GoVersion:    runtime.Version(),
		NumGoroutine: runtime.NumGoroutine(),
	}, nil
}
