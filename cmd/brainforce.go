package cmd

import (
	"fmt"
	"time"

	"github.com/google/logger"
	"github.com/shlima/fortune/internal/pkg/brainforce"
	"github.com/shlima/fortune/internal/pkg/pass"
	"github.com/urfave/cli/v2"
)

func BrainForce(c *cli.Context) error {
	state, err := pass.UnmarshalState(c.String(FlagPassState.Name))
	if err != nil {
		return fmt.Errorf("failed to unmargsall state: %w", err)
	}

	opts := pass.GenOpts{
		Alphabet: pass.ParseAlphabets(c.StringSlice(FlagPassAlphabet.Name)),
		State:    state,
		Length:   c.Int(FlagPassLength.Name),
	}

	force := brainforce.New(
		NewIndex(c),
		NewKeyGen(c).SetTesting(c.Args().First()),
		pass.New(opts),
	)

	debugBrainforce(c, force)
	go brainForceHeartbit(c, force)
	go brainForceTelegram(c, force)
	return force.Generate(onFound(c))
}

func brainForceHeartbit(c *cli.Context, force *brainforce.Force) {
	for range time.Tick(time.Second * time.Duration(c.Int(FlagHeartBeatSec.Name))) {
		logger.Info(force.PassGen().Heartbeat().ToString())
	}
}

func debugBrainforce(c *cli.Context, force *brainforce.Force) {
	logger.Info(fmt.Sprintf("loaded: %d addresses", force.DataLength()))
	logger.Info(fmt.Sprintf("test passed: %v", force.Get(c.String(FlagTestAddress.Name))))
	logger.Info(fmt.Sprintf("possible permutations: %d", force.PassGen().Permutations()))
	logger.Info(fmt.Sprintf("state: %s", pass.MarshallState(force.PassGen().Opts().State)))
	logger.Info(fmt.Sprintf("password length: %d", force.PassGen().Opts().Length))
	logger.Info(fmt.Sprintf("alphabet: %s", pass.MarshallAlphabet(force.PassGen().Opts().Alphabet)))
	logger.Info(fmt.Sprintf("telegram enabled: %v", NewTelegram(c).IsReal()))
}

func brainForceTelegram(c *cli.Context, force *brainforce.Force) {
	bot := NewTelegram(c)
	for range time.Tick(time.Second * time.Duration(c.Int(FlagTelegramPingSec.Name))) {
		if err := bot.HeartBeat(force.PassGen().Heartbeat().ToString()); err != nil {
			logger.Error(fmt.Sprintf("failed to send to telegram: %s\n", err))
		}
	}
}
