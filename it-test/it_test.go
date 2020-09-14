package it_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"os/exec"
	"testing"
)

type ItTestSuite struct {
	suite.Suite
	baseUrl string
}

func (s *ItTestSuite) SetupTest() {
	s.baseUrl = "http://localhost:8080"

	exec.Command("bash", "-c", "cd .. && make docker-up").Start()

	for {
		resp, err := http.Get(s.baseUrl +"/health")

		if resp != nil {
			defer resp.Body.Close()
		}

		if err == nil {
			break
		}
	}
}

func (s *ItTestSuite) TearDownTest() {
	exec.Command("bash", "-c", "cd .. && make docker-down").Run()
}

func (s *ItTestSuite) TestUserFlow() {
	payload, err := json.Marshal(map[string]string{
		"first_name": "John",
		"last_name": "Doe",
		"email": "john.doe@test.com",
		"password": "12345678",
	})

	if err != nil {
		s.Suite.T().Fatal(err)
	}

	createUserResponse, err := http.Post(s.baseUrl +"/users", "application/json", bytes.NewBuffer(payload))

	assert.Equal(s.T(), nil, err)
	assert.Equal(s.T(), 200, createUserResponse.StatusCode)
}

func (s *ItTestSuite) AfterTest(_, _ string) {
	if s.Suite.T().Failed() {
		logs, _ := exec.Command("bash", "-c", "cd .. && make docker-logs").Output()
		s.Suite.T().Log(string(logs))
	}
}

func TestItTestSuite(t *testing.T) {
	suite.Run(t, new(ItTestSuite))
}
