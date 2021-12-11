package service

import (
	"context"
	"time"
)

type Health struct {
}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) Get(ctx context.Context) *HealthResponse {
	return &HealthResponse{
		Message:   "OK",
		Timestamp: time.Now().Unix(),
	}
}

type HealthResponse struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
