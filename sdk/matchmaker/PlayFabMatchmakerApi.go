package matchmaker

// This code was generated by a tool. Any changes may be overwritten

import (
	"encoding/json"

	playfab "github.com/Myfreedom614/playfab-go-sdk/sdk"

	"github.com/mitchellh/mapstructure"
)

// AuthUser validates a user with the PlayFab service
// https://api.playfab.com/Documentation/Matchmaker/method/AuthUser
func AuthUser(settings *playfab.Settings, postData *AuthUserRequestModel, developerSecretKey string) (*AuthUserResponseModel, error) {
	if developerSecretKey == "" {
		return nil, playfab.NewCustomError("developerSecretKey should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Matchmaker/AuthUser", "X-SecretKey", developerSecretKey)
	if err != nil {
		return nil, err
	}

	result := &AuthUserResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, errDecoding := mapstructure.NewDecoder(&config)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	errDecoding = decoder.Decode(sourceMap)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	return result, nil
}

// PlayerJoined informs the PlayFab game server hosting service that the indicated user has joined the Game Server Instance specified
// https://api.playfab.com/Documentation/Matchmaker/method/PlayerJoined
func PlayerJoined(settings *playfab.Settings, postData *PlayerJoinedRequestModel, developerSecretKey string) (*PlayerJoinedResponseModel, error) {
	if developerSecretKey == "" {
		return nil, playfab.NewCustomError("developerSecretKey should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Matchmaker/PlayerJoined", "X-SecretKey", developerSecretKey)
	if err != nil {
		return nil, err
	}

	result := &PlayerJoinedResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, errDecoding := mapstructure.NewDecoder(&config)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	errDecoding = decoder.Decode(sourceMap)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	return result, nil
}

// PlayerLeft informs the PlayFab game server hosting service that the indicated user has left the Game Server Instance specified
// https://api.playfab.com/Documentation/Matchmaker/method/PlayerLeft
func PlayerLeft(settings *playfab.Settings, postData *PlayerLeftRequestModel, developerSecretKey string) (*PlayerLeftResponseModel, error) {
	if developerSecretKey == "" {
		return nil, playfab.NewCustomError("developerSecretKey should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Matchmaker/PlayerLeft", "X-SecretKey", developerSecretKey)
	if err != nil {
		return nil, err
	}

	result := &PlayerLeftResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, errDecoding := mapstructure.NewDecoder(&config)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	errDecoding = decoder.Decode(sourceMap)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	return result, nil
}

// StartGame instructs the PlayFab game server hosting service to instantiate a new Game Server Instance
// https://api.playfab.com/Documentation/Matchmaker/method/StartGame
func StartGame(settings *playfab.Settings, postData *StartGameRequestModel, developerSecretKey string) (*StartGameResponseModel, error) {
	if developerSecretKey == "" {
		return nil, playfab.NewCustomError("developerSecretKey should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Matchmaker/StartGame", "X-SecretKey", developerSecretKey)
	if err != nil {
		return nil, err
	}

	result := &StartGameResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, errDecoding := mapstructure.NewDecoder(&config)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	errDecoding = decoder.Decode(sourceMap)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	return result, nil
}

// UserInfo retrieves the relevant details for a specified user, which the external match-making service can then use to compute
// effective matches
// https://api.playfab.com/Documentation/Matchmaker/method/UserInfo
func UserInfo(settings *playfab.Settings, postData *UserInfoRequestModel, developerSecretKey string) (*UserInfoResponseModel, error) {
	if developerSecretKey == "" {
		return nil, playfab.NewCustomError("developerSecretKey should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Matchmaker/UserInfo", "X-SecretKey", developerSecretKey)
	if err != nil {
		return nil, err
	}

	result := &UserInfoResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, errDecoding := mapstructure.NewDecoder(&config)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	errDecoding = decoder.Decode(sourceMap)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	return result, nil
}
