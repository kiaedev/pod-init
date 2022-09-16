package cmd

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Command struct {
	Name   string            `json:"name"`
	Params map[string]string `json:"params"`
}

var hc = resty.New()

func fetchCommands(url string) ([]Command, error) {
	commands := make([]Command, 0)
	resp, err := hc.R().SetResult(&commands).Get(url)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("%v", resp.Error())
	}

	return commands, nil
}
