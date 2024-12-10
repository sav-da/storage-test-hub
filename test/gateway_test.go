package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestProxyToAuthService(t *testing.T) {
	body := map[string]string{"worker_id": "test_worker"}
	bodyJSON, _ := json.Marshal(body)

	resp, err := http.Post("http://localhost:8080/auth/login", "application/json", bytes.NewBuffer(bodyJSON))
	if err != nil {
		t.Fatalf("Failed to proxy to Auth Service: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", resp.StatusCode)
	}
}

func TestProxyToTestManagementService(t *testing.T) {
	body := map[string]string{"test_name": "example_test"}
	bodyJSON, _ := json.Marshal(body)

	resp, err := http.Post("http://localhost:8080/tests", "application/json", bytes.NewBuffer(bodyJSON))
	if err != nil {
		t.Fatalf("Failed to proxy to Test Management Service: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", resp.StatusCode)
	}
}
