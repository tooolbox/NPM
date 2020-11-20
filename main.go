package main

import (
	"github.com/tooolbox/firstvision/simple"
)

func main() {
	_ = simple.NewGateway(simple.Config{
		ApiKey:      "abc",
		ApiSecret:   "def",
		Environment: simple.SandboxEnvironment,
	})
}
