package cmd

import "github.com/urfave/cli/v2"

var FlagWorkers = &cli.IntFlag{
	Name:    "workers",
	EnvVars: []string{"WORKERS"},
	Value:   1,
}

var FlagHeartBeatSec = &cli.IntFlag{
	Name:    "heartbit",
	Usage:   "print status info each n seconds",
	EnvVars: []string{"HEARTBEAT_SEC"},
	Value:   10,
}

var FlagTestAddress = &cli.StringFlag{
	Name:    "test-address",
	EnvVars: []string{"TEST_ADDRESS"},
	Value:   "1LQoWist8KkaUXSPKZHNvEyfrEkPHzSsCd",
}

var FlagNightMode = &cli.BoolFlag{
	Name:    "night",
	EnvVars: []string{"NIGHT"},
	Value:   false,
}
