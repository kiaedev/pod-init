package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	testCmds := []Command{
		{Name: "renderFileByEnv", Params: map[string]string{"source": "", "target": ""}},
		{Name: "shell", Params: map[string]string{"cmd": "ls", "args": "-r -l"}},
		{Name: "shell", Params: map[string]string{"cmd": "pwd", "args": ""}},
	}
	http.Handle("/pod-init-commands", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		v, _ := json.Marshal(testCmds)
		_, _ = w.Write(v)
	}))
	s := httptest.NewServer(http.DefaultServeMux)
	commands, err := fetchCommands(s.URL + "/pod-init-commands")
	assert.NoError(t, err)
	assert.Equal(t, len(testCmds), len(commands))
	for i, command := range commands {
		assert.Equal(t, command, testCmds[i])
	}
}
