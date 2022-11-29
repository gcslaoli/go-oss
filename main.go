package main

import (
	_ "github.com/gcslaoli/go-oss/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gcslaoli/go-oss/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
