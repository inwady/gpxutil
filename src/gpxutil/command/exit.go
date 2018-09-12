package command

import (
	"fmt"
	"gpxutil/context"
	"os"
)

type ExitCommand struct {}

func (sc *ExitCommand) Execute(gctx *context.GPXContext, params []string) (bool, error) {
	fmt.Printf("Have a good day!\n")

	os.Exit(0)

	return false, nil
}

func (sc *ExitCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}
