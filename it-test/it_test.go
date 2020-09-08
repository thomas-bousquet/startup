package it_test

import (
	"net/http"
	"testing"
)

func TestApplication(t *testing.T) {

	// TODO: try to run command from the test so it can be used in a beforeEach()
	//exec.Command("bash", "-c", "make docker-up").Run()

	// WAIT HERE
	for {
		_, err := http.Get("http://localhost:8080/health")

		if err == nil {
			break
		}
	}

	// THEN START TEST
}
