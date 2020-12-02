package it_test

import (
	"bytes"
	"encoding/json"
	faker "github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type ItTestSuite struct {
	suite.Suite
	baseUrl string
}

func (s *ItTestSuite) SetupTest() {
	s.baseUrl = "http://localhost:8080"

	for {
		resp, err := http.Get(s.baseUrl +"/admin/health")

		if resp != nil {
			defer resp.Body.Close()
		}

		if err == nil {
			break
		}
	}
}

func (s *ItTestSuite) TestUserFlow() {
	userFirstName := faker.FirstName()
	userLastName := faker.LastName()
	userEmail := faker.Email()
	userPassword := faker.Password()

	payload, err := json.Marshal(map[string]string{
		"first_name": userFirstName,
		"last_name": userLastName,
		"email": userEmail,
		"password": userPassword,
	})

	if err != nil {
		s.Suite.T().Fatal(err)
	}

	createUserResponse, err := http.Post(s.baseUrl +"/users", "application/json", bytes.NewBuffer(payload))

	assert.Equal(s.T(), nil, err)
	assert.Equal(s.T(), http.StatusCreated, createUserResponse.StatusCode)

	//resp := map[string]string{}
	//userId := createUserResponse.Body.Read(&resp)
	//
	//readUser, err := http.Post(s.baseUrl +"/users" + , "application/json", bytes.NewBuffer(payload))
}

func TestItTestSuite(t *testing.T) {
	suite.Run(t, new(ItTestSuite))
}
