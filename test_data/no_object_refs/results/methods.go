package vkapiclient

import (
	"encoding/json"
)

// Completes a user's registration (begun with the [vk.com/dev/auth.signup|auth.signup] method) using an authorization code.
func AuthConfirm(
	clientId int,
	clientSecret string,
	phone string,
	code string,
	password string,
	testMode bool,
	intro int) (response *AuthConfirmResponse, err error) {
	request := NewRequest("auth.confirm")
	request.AddIntParam("client_id", clientId)
	request.AddStrParam("client_secret", clientSecret)
	request.AddStrParam("phone", phone)
	request.AddStrParam("code", code)
	request.AddStrParam("password", password)
	request.AddBoolParam("test_mode", testMode)
	request.AddIntParam("intro", intro)
	byteResponse := request.Send()
	err = json.Unmarshal(byteResponse, &response)
	return
}
