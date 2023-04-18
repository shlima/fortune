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
		Copyright: "© github.com/shlima/fortune",
		Version:   Version,
		Suggest:   true,
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
				Usage:  "run bruteforce against the dataset of rich addresses",
			}, {
				Name:   "random",
				Action: cmd.Random,
				Usage:  "prints random address from the dataset files",
			}, {
				Name:   "generate",
				Action: cmd.Generate,
				Usage:  "random bitcoin address",
			}, {
				Name:   "brain",
				Action: cmd.Brain,
				Usage:  "generate brain wallet (based on a password as a first argument) and check it's current balance online",
			}, {
				Name:   "brainforce",
				Action: cmd.BrainForce,
				Usage:  "run bruteforce with alphabetical passwords permutations against the dataset of rich addresses",
				Flags: []cli.Flag{
					cmd.FlagPassState,
					cmd.FlagPassLength,
					cmd.FlagPassAlphabet,
					cmd.FlagPassShuffle,
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
