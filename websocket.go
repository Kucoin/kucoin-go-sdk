package kucoin

import "net/http"

type WebSocketTokenModel struct {
	Token   string                `json:"token"`
	Servers WebSocketServersModel `json:"instanceServers"`
}
type WebSocketServerModel struct {
	PingInterval int64  `json:"pingInterval"`
	Endpoint     string `json:"endpoint"`
	Protocol     string `json:"protocol"`
	Encrypt      bool   `json:"encrypt"`
	PingTimeout  int64  `json:"pingTimeout"`
}

type WebSocketServersModel []*WebSocketServerModel

func (as *ApiService) WebSocketPublicToken() (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/bullet-public", map[string]string{})
	return as.call(req)
}

func (as *ApiService) WebSocketPrivateToken() (*ApiResponse, error) {
	req := NewRequest(http.MethodPost, "/api/v1/bullet-private", map[string]string{})
	return as.call(req)
}
