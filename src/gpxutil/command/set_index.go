package command

import (
	"gpxutil/context"
	"errors"
	"strconv"
)

type SetIndexCommand struct{}

func (ac *SetIndexCommand) Execute(gctx *context.GPXContext, params []string) error {
	if len(params) < 2 {
		return errors.New("bad params")
	}

	index, err := strconv.ParseInt(params[1], 10, 32)
	if err != nil || index < 0 {
		return errors.New("bad index")
	}

	return gctx.SetWorkIndex(uint(index))
}

func (ac *SetIndexCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}