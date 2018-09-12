package command

import (
	"fmt"
	"gpxutil/context"
)

type HelpCommand struct {}

func (sc *HelpCommand) Execute(gctx *context.GPXContext, params []string) (bool, error) {
	fmt.Printf("Commands:\n")
	for key, _ := range commandTable {
		fmt.Printf("- %s\n", key)
	}

	return false, nil
}

func (sc *HelpCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}