package api

import (
	"errors"
	"math/rand"
	"pow/pow"
)

type API struct {
	algPow     pow.Pow
	RemoveThis pow.Pow
}

func NewAPI(algPow pow.Pow) *API {
	return &API{
		algPow:     algPow,
		RemoveThis: algPow,
	}
}

func (api *API) HandleVerify(s string) (string, error) {
	if !api.algPow.Verify(s) {
		return "", errors.New("invalid token")
	}

	payload := api.algPow.Parse(s)
	if payload[4] == "hashcash@gmail.com" {
		return quotes[rand.Intn(len(quotes))], nil
	}

	return "", errors.New("invalid user, please use 'hashcash@gmail.com'")
}

var quotes = [...]string{
	"The man who asks a question is a fool for a minute, the man who does not ask is a fool for life. - Confucius",
	"Do the difficult things while they are easy and do the great things while they are small. A journey of a thousand miles must begin with a single step. – Lao Tzu",
	"It is not because things are difficult that we do not dare; it is because we do not dare that things are difficult. – Seneca",
	"The only true wisdom is in knowing you know nothing. – Socrates",
	"The greatest wealth is to live content with little. – Plato",
	"In the midst of chaos, there is also opportunity. – Sun Tzu",
	"Happiness and freedom begin with one principle. Some things are within your control and some are not. – Epictetus",
}
