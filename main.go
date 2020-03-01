package main

import (
	"context"
	"log"
	"os"

	"github.com/int128/transerr/pkg/cmd"
)

func init() {
	log.SetFlags(0)
}

func main() {
	os.Exit(cmd.Run(context.Background(), os.Args))
}
