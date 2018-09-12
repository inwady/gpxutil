package point

import (
	"gpxutil/context"
	"errors"
	"strconv"
)

type AddCommand struct {
	index uint
}

func (ac *AddCommand) Execute(gctx *context.GPXContext, params []string) (bool, error) {
	if len(params) < 3 {
		return false, errors.New("bad params")
	}

	_i, err := strconv.ParseInt(params[1], 10, 32)
	if err != nil || _i < 0 {
		return false, errors.New("bad index")
	}

	i := uint(_i)

	lat, err := strconv.ParseFloat(params[2], 64)
	if err != nil {
		return false, err
	}

	log, err := strconv.ParseFloat(params[3], 64)
	if err != nil {
		return false, err
	}

	ac.index = i

	err = gctx.AddPoint(i, lat, log)
	return true, err
}

func (ac *AddCommand) UnExecute(gctx *context.GPXContext) error {
	return gctx.RemovePoint(ac.index)
}
