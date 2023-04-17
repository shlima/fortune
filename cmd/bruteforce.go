package cmd

import (
	"fmt"
	"time"

	"github.com/google/logger"
	"github.com/shlima/fortune/internal/pkg/bruteforce"
	"github.com/shlima/fortune/internal/pkg/key"
	"github.com/shlima/fortune/internal/pkg/mapper"
	"github.com/urfave/cli/v2"
)

func Bruteforce(c *cli.Context) error {
	force := bruteforce.New(
		NewIndex(c),
		NewKeyGen(c).SetTesting(c.Args().First()),
		c.Int(FlagWorkers.Name),
	)

	force.SetNightMode(c.Bool(FlagNightMode.Name))
	return terror(c, force)
}

func terror(c *cli.Context, force *bruteforce.Executor) error {
	logger.Info(fmt.Sprintf("loaded: %d addresses", force.DataLength()))
	logger.Info(fmt.Sprintf("workers count: %d", force.WorkersCount()))
	logger.Info(fmt.Sprintf("test passed: %v", force.Get(c.String(FlagTestAddress.Name))))
	logger.Info(fmt.Sprintf("telegram enabled: %v", NewTelegram(c).IsReal()))

	go heartbit(c, force)
	go telegram(c, force)

	force.Run(onFound(c))
	return nil
}

func heartbit(c *cli.Context, force *bruteforce.Executor) {
	for range time.Tick(time.Second * time.Duration(c.Int(FlagHeartBeatSec.Name))) {
		logger.Info(force.Heartbeat().ToString())
	}
}

func telegram(c *cli.Context, force *bruteforce.Executor) {
	bot := NewTelegram(c)
	for range time.Tick(time.Second * time.Duration(c.Int(FlagTelegramPingSec.Name))) {
		if err := bot.HeartBeat(force.Heartbeat().ToString()); err != nil {
			logger.Error(fmt.Sprintf("failed to send to telegram: %s\n", err))
		}
	}
}

func onFound(c *cli.Context) bruteforce.FoundFn {
	return func(chain key.KeyChain) {
		logger.Warning(fmt.Sprintf("FOUND: %s", chain.ToString()))
		err := NewTelegram(c).KeyFound(mapper.KeyChainHTML(chain))
		logger.Info(fmt.Sprintf("Send to telegram result: %s", err))
		panic(chain.ToString())
	}
}
