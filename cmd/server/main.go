package main

import (
	"fmt"
	"os"

	cmd "github.com/dhanusaputra/somewhat-server/pkg/cmd/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
