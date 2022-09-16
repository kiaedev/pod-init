package cmd

import (
	"context"
	"log"
	"os"
)

func Execute() {
	commands, err := fetchCommands(os.Getenv("POD_INIT_FETCH_URL"))
	if err != nil {
		log.Fatalf("fetch commands failed: %v", err)
		return
	}

	execCommands(context.Background(), commands)
}
