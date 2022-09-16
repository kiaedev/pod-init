package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var commandFuncs = map[string]func(ctx context.Context, args map[string]string) error{
	"renderFileByEnv": renderFileByEnv,
	"shell":           execShell,
}

func execCommands(ctx context.Context, commands []Command) {
	for _, command := range commands {
		fn, ok := commandFuncs[command.Name]
		if !ok {
			log.Printf("command %s not found", command.Name)
			continue
		}
		if err := fn(ctx, command.Params); err != nil {
			log.Printf("command %s: %s", command.Name, err)
		}
	}
}

func renderFileByEnv(ctx context.Context, args map[string]string) error {
	t, err := template.ParseFiles(args["source"])
	if err != nil {
		return err
	}

	tp := map[string]string{}
	for _, env := range os.Environ() {
		envItems := strings.Split(env, "=")
		tp[envItems[0]] = envItems[1]
	}

	file, err := os.Create(args["target"])
	if err != nil {
		return err
	}

	return t.Execute(file, tp)
}

func execShell(ctx context.Context, args map[string]string) error {
	cmdArgs := strings.Split(args["args"], " ")
	c := exec.CommandContext(ctx, args["cmd"], cmdArgs...)
	fmt.Println(c.String())
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
