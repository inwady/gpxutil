package command

import (
	"gpxutil/context"
	"fmt"
	"gpxutil/command/point"
	"gpxutil/command/index"
	"gpxutil/command/info"
	"gpxutil/command/exchange"
	"gpxutil/model"
)

type StackCommand struct {
	gctx *context.GPXContext
	cmd model.Command
	argv []string
}

var (
	commandTable = map[string]func() model.Command{
		"add": func() model.Command { return new(point.AddCommand) },
		"remove": func() model.Command { return new(point.RemoveCommand) },
		"change": func() model.Command { return new(point.ChangeCommand) },
		"index": func() model.Command { return new(index.SetIndexCommand) },

		"list": func() model.Command { return new(info.ListCommand) },

		"import": func() model.Command { return new(exchange.ExchangeCommand) },

		"undo": func() model.Command { return new(UndoCommand) },
		"redo": func() model.Command { return new(RedoCommand) },
		"info": func() model.Command { return new(InfoCommand) },
		"help": func() model.Command { return new(HelpCommand) },
		"exit": func() model.Command { return new(ExitCommand) },
	}

	stackChange []StackCommand
	stackBack []StackCommand
)

func getElementAndRevertStack(stackFrom *[]StackCommand, stackTo *[]StackCommand) (StackCommand, error)  {
	if len(*stackFrom) <= 0 {
		return StackCommand{}, fmt.Errorf("no elements in stack")
	}

	lastElement := (*stackFrom)[len(*stackFrom) - 1]

	fmt.Println(stackFrom)
	*stackFrom = (*stackFrom)[:len(*stackFrom) - 1]
	fmt.Println(stackFrom)

	*stackTo = append(*stackTo, lastElement)
	return lastElement, nil
}

func addCommand(gctx *context.GPXContext, cmd model.Command, argv []string) {
	stackChange = append(stackChange, StackCommand{
		gctx: gctx,
		cmd: cmd,
		argv: argv,
	})
}

func Execute(gctx *context.GPXContext, argv []string) error {
	buildCommandFunc, ok := commandTable[argv[0]]
	if !ok {
		return fmt.Errorf("unknown command")
	}

	c := buildCommandFunc()
	undoFlag, err := c.Execute(gctx, argv)
	if err != nil {
		return err
	}

	if undoFlag {
		addCommand(gctx, c, argv)
	}

	return nil
}
