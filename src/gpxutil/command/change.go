package command

import (
	"gpxutil/context"
	"errors"
	"strconv"
)

type ChangeCommand struct{}

func (ac *ChangeCommand) Execute(gctx *context.GPXContext, params []string) error {
	if len(params) < 4 {
		return errors.New("bad params")
	}

	i, err := strconv.ParseInt(params[1], 10, 32)
	if err != nil || i < 0 {
		return errors.New("bad index")
	}

	lat, err := strconv.ParseFloat(params[2], 64)
	if err != nil {
		return err
	}

	log, err := strconv.ParseFloat(params[3], 64)
	if err != nil {
		return err
	}

	return gctx.ChangePoint(uint(i), lat, log)
}

func (ac *ChangeCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}
