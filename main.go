package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/urfave/cli"
)

type Config struct {
	a []int
}

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "p", Usage: "端口", Value: ""},
		cli.StringFlag{Name: "a", Usage: "密码", Value: ""},
	}
	app.Action = do
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("err", err)
	}
}

// 解析配置，run
func do(c *cli.Context) error {
	port := c.String("p")
	password := c.String("a")
	redisClient := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("127.0.0.1:%s", port),
		Password:    password,
		ReadTimeout: time.Second * 100,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}

	for i := 0; i < 100; i++ {
		t := 1642971600 - i*86400*7
		keyP := fmt.Sprintf("cross:gvg:%d*", t)
		keys, err := redisClient.Keys(keyP).Result()
		if err != nil {
			return err
		}
		if len(keys) == 0 {
			continue
		}
		count, err := redisClient.Del(keys...).Result()
		if err != nil {
			return err
		}
		fmt.Println("delete key success", keyP, count)
		time.Sleep(time.Millisecond * 100)
	}
	return nil
}
