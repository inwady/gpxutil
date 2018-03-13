package command

import (
	"gpxutil/context"
	"errors"
	"strconv"
)

type AddCommand struct{}

func (ac *AddCommand) Execute(gctx *context.GPXContext, params []string) error {
	if len(params) < 3 {
		return errors.New("bad params")
	}

	lat, err := strconv.ParseFloat(params[1], 64)
	if err != nil {
		return err
	}

	log, err := strconv.ParseFloat(params[2], 64)
	if err != nil {
		return err
	}

	return gctx.AddPoint(lat, log)
}

func (ac *AddCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}
