package command

import (
	"gpxutil/context"
	"log"
)

type RedoCommand struct{}

func (ac *RedoCommand) Execute(_ *context.GPXContext, params []string) (bool, error) {
	stackCommand, err := getElementAndRevertStack(&stackBack, &stackChange)
	if err != nil {
		return false, err
	}

	_, err = stackCommand.cmd.Execute(stackCommand.gctx, stackCommand.argv)
	if err != nil {
		log.Fatalf("redo fail, %v", err)
	}

	return false, nil
}

func (ac *RedoCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}
