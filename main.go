package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "conf", Usage: "配置", Value: "./stream.yaml"},
	}
	app.Action = do
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("err", err)
	}
}

func do(ctx *cli.Context) error {
	confPath := ctx.String("conf")
	if len(confPath) == 0 {
		return errors.New("no config path")
	}
	return nil
}
