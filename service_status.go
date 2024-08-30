package kucoin

import (
	"context"
	"net/http"
)

// A ServiceStatusModel represents the structure of service status.
type ServiceStatusModel struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

// ServiceStatus returns the service status.
func (as *ApiService) ServiceStatus(ctx context.Context) (*ApiResponse, error) {
	req := NewRequest(http.MethodGet, "/api/v1/status", nil)
	return as.Call(ctx, req)
}
