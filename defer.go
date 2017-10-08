package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	defer fmt.Println("the end")

	fmt.Println("The time is", time.Now())
	handleCommandFlags()
	envs()
}

func handleCommandFlags() {
	flagStringPtr := flag.String("word", "default word", "give word")
	flagCountPtr := flag.Int("count", 4, "give count")

	flag.Parse()

	fmt.Println("word:", *flagStringPtr)
	fmt.Println("count:", *flagCountPtr)

	for i := 0; i < *flagCountPtr; i++ {
		defer fmt.Println(i)
		fmt.Println(*flagStringPtr)
	}
}

func envs() {
	fmt.Println("FOO:", os.Getenv("GOPATH"))
	fmt.Println("USER:", os.Getenv("USER"))
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair)
	}
}
