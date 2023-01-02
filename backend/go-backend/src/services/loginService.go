package services

import (
	"context"
	"encoding/json"
	"main/src/keycloak"
	"net/http"
	"time"
)

type doc struct {
	Id   string    `json:"id"`
	Num  string    `json:"num"`
	Date time.Time `json:"date"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
}

type controller struct {
	keycloak *keycloak.Keycloak
}

func NewController(keycloak *keycloak.Keycloak) *controller {
	return &controller{
		keycloak: keycloak,
	}
}

func (controller *controller) Login(writer http.ResponseWriter, request *http.Request) {

	loginRequest := &loginRequest{}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(loginRequest); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	jwt, err := controller.keycloak.Gocloak.Login(context.Background(),
		controller.keycloak.ClientId,
		controller.keycloak.ClientSecret,
		controller.keycloak.Realm,
		loginRequest.Username,
		loginRequest.Password)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusForbidden)
		return
	}

	response := &loginResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	}

	responseJson, _ := json.Marshal(response)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(responseJson)
}