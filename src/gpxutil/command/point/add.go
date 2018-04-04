package point

import (
	"gpxutil/context"
	"errors"
	"strconv"
)

type AddCommand struct {

}

func (ac *AddCommand) Execute(gctx *context.GPXContext, params []string) (bool, error) {
	if len(params) < 3 {
		return false, errors.New("bad params")
	}

	lat, err := strconv.ParseFloat(params[1], 64)
	if err != nil {
		return false, err
	}

	log, err := strconv.ParseFloat(params[2], 64)
	if err != nil {
		return false, err
	}

	return true, gctx.AddPoint(lat, log)
}

func (ac *AddCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}
