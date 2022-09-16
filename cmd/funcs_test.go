package cmd

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecShell(t *testing.T) {
	assert.NoError(t, execShell(context.Background(), map[string]string{"cmd": "ls", "args": "-r -l"}))
}

func TestRenderFileByEnv(t *testing.T) {
	_ = os.Setenv("MYSQL_USER", "root")
	assert.NoError(t, renderFileByEnv(context.Background(), map[string]string{"source": "testdata/config.goyaml", "target": "testdata/config.yaml"}))
	cfgBytes, err := ioutil.ReadFile("testdata/config.yaml")
	assert.NoError(t, err)
	assert.Contains(t, string(cfgBytes), "root")
}
