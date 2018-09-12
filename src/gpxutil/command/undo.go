package command

import (
	"gpxutil/context"
)

type UndoCommand struct{}

func (ac *UndoCommand) Execute(_ *context.GPXContext, params []string) (bool, error) {
	stackCommand, err := getElementAndRevertStack(&stackChange, &stackBack)
	if err != nil {
		return false, err
	}

	return false, stackCommand.cmd.UnExecute(stackCommand.gctx)
}

func (ac *UndoCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}

func (ac *UndoCommand) Info() string {
	return "undo"
}
