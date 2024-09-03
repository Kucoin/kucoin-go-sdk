package kucoin

import (
	"context"
	"testing"
)

func TestApiService_ServiceStatus(t *testing.T) {
	s := NewApiServiceFromEnv()

	rsp, err := s.ServiceStatus(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	var ss ServiceStatusModel
	if err := rsp.ReadData(&ss); err != nil {
		t.Fatal(err)
	}
	if ss.Status == "" {
		t.Fatal("empty status")
	}
}
