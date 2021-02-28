package it_test

import (
	"bytes"
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
)

type ItTestSuite struct {
	suite.Suite
	baseUrl string
}

func (s *ItTestSuite) SetupTest() {
	s.baseUrl = "http://localhost:8081"
}

func (s *ItTestSuite) TestUserFlow() {

	userFirstName := faker.FirstName()
	userLastName := faker.LastName()
	userEmail := faker.Email()
	userPassword := faker.Password()

	payload, _ := json.Marshal(map[string]string{
		"first_name": userFirstName,
		"last_name":  userLastName,
		"email":      userEmail,
		"password":   userPassword,
	})

	client := http.DefaultClient

	// Create user
	createUser, _ := client.Post(s.baseUrl+"/users", "application/json", bytes.NewBuffer(payload))
	assert.Equal(s.T(), http.StatusCreated, createUser.StatusCode, "Should create createUserResponseBody successfully")
	createUserResponseBody := map[string]string{}
	body, _ := ioutil.ReadAll(createUser.Body)
	json.Unmarshal(body, &createUserResponseBody)
	userId := createUserResponseBody["id"]
	createUser.Body.Close()

	// Get User without auth
	getUserWithoutAuthorization, _ := client.Get(s.baseUrl + "/users/" + userId)
	assert.Equal(s.T(), http.StatusUnauthorized, getUserWithoutAuthorization.StatusCode)
	getUserWithoutAuthorization.Body.Close()

	// Login
	loginUserRequest, _ := http.NewRequest("POST", s.baseUrl+"/login", nil)
	loginUserRequest.SetBasicAuth(userEmail, userPassword)
	loginUserResponse, _ := client.Do(loginUserRequest)
	assert.Equal(s.T(), http.StatusOK, loginUserResponse.StatusCode)
	body, _ = ioutil.ReadAll(loginUserResponse.Body)
	loginResponseBody := map[string]string{}
	json.Unmarshal(body, &loginResponseBody)
	authToken := loginResponseBody["token"]
	loginUserResponse.Body.Close()

	// Get user with auth
	getUserResponseWithAuthorizationJWTRequest, _ := http.NewRequest("GET", s.baseUrl+"/users/"+userId, nil)
	getUserResponseWithAuthorizationJWTRequest.Header.Set("Authorization", "Bearer "+authToken)
	getUserWithAuthorizationResponse, _ := client.Do(getUserResponseWithAuthorizationJWTRequest)
	assert.Equal(s.T(), http.StatusOK, getUserWithAuthorizationResponse.StatusCode)
	getUserWithAuthorizationResponse.Body.Close()
}

func TestItTestSuite(t *testing.T) {
	suite.Run(t, new(ItTestSuite))
}
