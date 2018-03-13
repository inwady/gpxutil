package command

import (
	"gpxutil/context"
	"errors"
	"strconv"
)

type RemoveCommand struct{}

func (ac *RemoveCommand) Execute(gctx *context.GPXContext, params []string) error {
	if len(params) < 2 {
		return errors.New("bad params")
	}

	i, err := strconv.ParseInt(params[1], 10, 32)
	if err != nil || i < 0 {
		return errors.New("bad index")
	}

	return gctx.RemovePoint(uint(i))
}

func (ac *RemoveCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}
