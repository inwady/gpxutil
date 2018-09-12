package command

import (
"fmt"
"gpxutil/context"
)

type HelpCommand struct {}

func (sc *HelpCommand) Execute(gctx *context.GPXContext, params []string) (bool, error) {
	fmt.Printf("Commands GPXUtil:\n")
	for key, v := range commandTable {
		// TODO a.naberezhnyi
		fmt.Printf("         - %s (%s) \n", key, v().Info())
	}

	return false, nil
}

func (sc *HelpCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}

func (ac *HelpCommand) Info() string {
	return "help command"
}
