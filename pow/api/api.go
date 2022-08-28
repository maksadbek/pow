package api

import (
	"errors"
	"pow/pow"
)

type API struct {
	algPow pow.Pow
}

func (api *API) HandleVerify(s string) (string, error) {
	if !api.algPow.Verify(s) {
		return "", errors.New("invalid token")
	}

	payload := api.algPow.Parse(s)
	if payload[4] == "john@doe.com" {
		return "success", nil
	}

	return "", errors.New("invalid user")
}
