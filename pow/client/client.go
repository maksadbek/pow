package client

import (
	"pow/pow"
)

type Client struct {
	pow pow.Pow
}

func NewClient(pow pow.Pow) *Client {
	return &Client{
		pow: pow,
	}
}

func (c *Client) Generate(id string) string {
	return c.pow.Generate(id)
}
