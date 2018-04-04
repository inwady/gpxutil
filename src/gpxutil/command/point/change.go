package point

import (
	"gpxutil/context"
	"errors"
	"strconv"
)

type ChangeCommand struct{}

func (ac *ChangeCommand) execute(gctx *context.GPXContext, params []string) (bool, error) {
	if len(params) < 4 {
		return false, errors.New("bad params")
	}

	i, err := strconv.ParseInt(params[1], 10, 32)
	if err != nil || i < 0 {
		return false, errors.New("bad index")
	}

	lat, err := strconv.ParseFloat(params[2], 64)
	if err != nil {
		return false, err
	}

	log, err := strconv.ParseFloat(params[3], 64)
	if err != nil {
		return false, err
	}

	return true, gctx.ChangePoint(uint(i), lat, log)
}

func (ac *ChangeCommand) unExecute(gctx *context.GPXContext) error {
	return nil
}
