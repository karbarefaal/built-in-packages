package main

import (
	"fmt"

	"github.com/karbarefaal/buit-in-packages/config"
)

func main() {
	config.LoadEnv()
	fmt.Println(config.GlobalConfig.PostgresqlUsername)
}
