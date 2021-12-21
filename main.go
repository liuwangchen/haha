package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Action = do
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("err", err)
	}
}

// 解析配置，run
func do(c *cli.Context) error {
	var (
		i1  = 0
		i2  = 0
		err error
	)
	if len(c.Args()) == 0 {
		return errors.New("args len != 2")
	}
	i1, err = strconv.Atoi(c.Args()[0])
	if err != nil {
		return err
	}

	all, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	split := bytes.Split(all, []byte("\n"))
	i2 = len(split)

	if c.NArg() == 2 {
		i2, err = strconv.Atoi(c.Args()[1])
		if err != nil {
			return err
		}
	}

	buf := new(bytes.Buffer)
	for _, v := range split[i1:i2] {
		buf.Write(v)
		buf.Write([]byte("\n"))
	}
	fmt.Fprint(c.App.Writer, buf.String())
	return nil
}
