package main

import (
	"os"

	"github.com/Kyuubang/plz/pkg/cmd/root"
)

func main() {
	code := root.Execute()
	os.Exit(int(code))
}
