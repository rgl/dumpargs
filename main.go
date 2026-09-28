package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// dump all the environment variables.
	envs := os.Environ()
	sort.Slice(envs, func(i, j int) bool {
		nameI, _, _ := strings.Cut(envs[i], "=")
		nameJ, _, _ := strings.Cut(envs[j], "=")
		return nameI < nameJ
	})
	for _, env := range envs {
		fmt.Printf("ENV: %s\n", env)
	}

	// dump all the cli arguments.
	for i, arg := range os.Args {
		fmt.Printf("ARG%d: %s\n", i, arg)
	}
}
