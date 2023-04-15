package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"

	"github.com/shlima/fortune/cmd"
	"github.com/urfave/cli/v2"
)

// Version is a build time constant
var Version = ":VERSION:"

func main() {
	app := &cli.App{
		Name:      "fortune",
		Usage:     "bitcoin wallet cracker",
		Version:   Version,
		Copyright: "© github.com/shlima/fortune",
		Flags: []cli.Flag{
			cmd.FlagFiles,
			cmd.FlagWorkers,
			cmd.FlagNightMode,
			cmd.FlagTestAddress,
			cmd.FlagHeartBeatSec,
			cmd.FlagTelegramPingSec,
			cmd.FlagTelegramToken,
			cmd.FlagTelegramChannel,
		},
		Before: func(c *cli.Context) error {
			return each(c, cmd.InitLogger)
		},
		Commands: []*cli.Command{
			{
				Name:   "bruteforce",
				Action: cmd.Bruteforce,
				Subcommands: []*cli.Command{
					{
						Name:   "test",
						Usage:  "test 1LQoWist8KkaUXSPKZHNvEyfrEkPHzSsCd",
						Action: cmd.BruteforceTest,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func each(c *cli.Context, fns ...func(c *cli.Context) error) error {
	for _, fn := range fns {
		if err := fn(c); err != nil {
			name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
			return fmt.Errorf("failed to execute <%s>: %w", name, err)
		}
	}

	return nil
}
