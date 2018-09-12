package command

import (
	"gpxutil/context"
	"fmt"
)

type InfoCommand struct{}

func (ac *InfoCommand) Execute(gctx *context.GPXContext, params []string) (bool, error) {
	fmt.Printf("main info: \n")
	fmt.Printf("current index %d\n", gctx.GetIndex())
	fmt.Printf("current points %v\n", gctx.SizePoint())
	fmt.Printf("undo stack %v\n", stackChange)
	fmt.Printf("redo stack %v\n", stackBack)
	return false, nil
}

func (ac *InfoCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}

func (ac *InfoCommand) Info() string {
	return "basic info"
}
