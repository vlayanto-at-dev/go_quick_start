package main

import (
	"log"
	"os"

	_ "github.com/vlayanto-at-dev/lib/sample/matchers"
	"github.com/vlayanto-at-dev/lib/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("President")
}
