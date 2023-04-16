package cmd

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"github.com/shlima/fortune/internal/pkg/pass"
	"github.com/urfave/cli/v2"
)

var FlagWorkers = &cli.IntFlag{
	Name:    "workers",
	Usage:   "number of workers for parallel execution",
	EnvVars: []string{"WORKERS"},
	Value:   1,
}

var FlagHeartBeatSec = &cli.IntFlag{
	Name:    "heartbit-sec",
	Usage:   "print status each N seconds to STDOUT",
	EnvVars: []string{"HEARTBEAT_SEC"},
	Value:   10,
}

var FlagTestAddress = &cli.StringFlag{
	Name:    "test-address",
	Usage:   "address to test dataset before running brute force",
	EnvVars: []string{"TEST_ADDRESS"},
	Value:   "1LQoWist8KkaUXSPKZHNvEyfrEkPHzSsCd",
}

var FlagNightMode = &cli.BoolFlag{
	Name:    "night",
	Usage:   "night or silent mode (reduced CPU usage)",
	EnvVars: []string{"NIGHT"},
	Value:   false,
}

var FlagFiles = &cli.StringSliceFlag{
	Name:    "file",
	Usage:   "a file with a custom dictionary",
	EnvVars: []string{"FILE"},
	Value: cli.NewStringSlice(
		"addresses/Bitcoin/2023/04/p2pkh_Rich_Max_1.txt",
		"addresses/Bitcoin/2023/04/p2pkh_Rich_Max_10.txt",
		"addresses/Bitcoin/2023/04/p2pkh_Rich_Max_100.txt",
		"addresses/Bitcoin/2023/04/p2pkh_Rich_Max_1000.txt",
		"addresses/Bitcoin/2023/04/p2pkh_Rich_Max_10000.txt",
		"addresses/Bitcoin/2023/04/p2pkh_Rich_Max_100000.txt",
	),
}

var FlagPassState = &cli.StringFlag{
	Name:    "pass-state",
	Usage:   "last state of the previous password brain force (to continue from where you stopped)",
	EnvVars: []string{"PASS_STATE"},
	Value:   "",
}

var FlagPassLength = &cli.IntFlag{
	Name:    "pass-length",
	Usage:   "the length of the passwords to be picked",
	EnvVars: []string{"PASS_LENGTH"},
	Value:   5,
}

var FlagPassAlphabet = &cli.StringSliceFlag{
	Name:    "pass-alphabet",
	Usage:   fmt.Sprintf("one of %s or any characters separated by comma", strings.Join(lo.Keys(pass.Dictionary), ", ")),
	EnvVars: []string{"ALPHABET"},
	Value:   cli.NewStringSlice(lo.Keys(pass.Dictionary)...),
}
