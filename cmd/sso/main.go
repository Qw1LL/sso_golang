package main

import (
	"fmt"
	"github.com/Qw1LL/sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
