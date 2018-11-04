package main

import (
	"os"

	"github.com/cholick/pipeline-test/greet"
)

func main() {
	greeter := greet.Greeter{
		Message: "Hello pipeline\n",
	}
	err := greeter.Greet(os.Stdout)
	if err != nil {
		panic(err)
	}
}
