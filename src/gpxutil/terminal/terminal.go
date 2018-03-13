package terminal

import (
	"bufio"
	"os"
	"gpxutil/context"
	"fmt"
	"strings"
	"gpxutil/command"
)

func InitTerminal(gctx *context.GPXContext) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		c, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		if strings.HasSuffix(c, "\n") {
			c = c[:len(c) - 1]
		}

		commands := strings.Split(c, " ")
		if len(commands) == 0 || commands[0] == "" {
			continue
		}

		concreteCommand, ok := command.CommandTable[commands[0]]
		if !ok {
			fmt.Printf("unknown command\n")
			continue
		}

		err = concreteCommand.Execute(gctx, commands)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
	}

	return nil
}