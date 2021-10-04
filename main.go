package main

import (
	"errors"
	"fmt"
	"os"

	"haha/conf"
	"haha/entity"

	"github.com/reugn/go-streams"
	"github.com/reugn/go-streams/flow"
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

func run(config conf.StreamConfig) error {
	for _, agg := range config.AggConfigs {
		sourceConfig, ok := config.FindEntityConfig(agg.SourceId)
		if !ok {
			return errors.New("source not found")
		}
		sinkConfig, ok := config.FindEntityConfig(agg.SinkId)
		if !ok {
			return errors.New("sink not found")
		}
		source, ok := entity.GetSourceEntity(sourceConfig)
		if !ok {
			return errors.New("source not found")
		}
		sink, ok := entity.GetSinkEntity(sinkConfig)
		if !ok {
			return errors.New("sink not found")
		}

		var (
			f streams.Flow = flow.NewPassThrough()
		)
		for _, flowConfig := range agg.FlowConfigs {
			f, ok = entity.GetFlow(flowConfig)
			if !ok {
				return errors.New("source not found")
			}
			f = f.Via(f)
		}

		source.Via(f).To(sink)
	}
	return nil
}
