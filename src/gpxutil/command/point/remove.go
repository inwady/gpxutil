package point

import (
	"gpxutil/context"
	"errors"
	"strconv"
)

type RemoveCommand struct{}

func (ac *RemoveCommand) execute(gctx *context.GPXContext, params []string) (bool, error) {
	if len(params) < 2 {
		return false, errors.New("bad params")
	}

	i, err := strconv.ParseInt(params[1], 10, 32)
	if err != nil || i < 0 {
		return false, errors.New("bad index")
	}

	return true, gctx.RemovePoint(uint(i))
}

func (ac *RemoveCommand) unExecute(gctx *context.GPXContext) error {
	return nil
}
