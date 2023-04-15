package cmd

import (
	"io/ioutil"
	"log"

	"github.com/google/logger"
	"github.com/urfave/cli/v2"
)

func InitLogger(c *cli.Context) error {
	logger.Init(c.App.Name, true, true, ioutil.Discard)
	logger.SetLevel(logger.Level(0))
	logger.SetFlags(log.LstdFlags)

	return nil
}
