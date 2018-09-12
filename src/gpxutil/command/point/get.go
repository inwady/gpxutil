package point

import (
	"fmt"
	"gpxutil/context"
	"errors"
	"strconv"
)

type GetPointCommand struct  {
	index uint
	changedLat float64
	changedLog float64
}

func (ac *GetPointCommand) Execute(gctx *context.GPXContext, params []string) (bool, error) {
	if len(params) < 2 {
		return false, errors.New("bad params")
	}

	_i, err := strconv.ParseInt(params[1], 10, 32)
	if err != nil || _i < 0 {
		return false, errors.New("bad index")
	}

	i := uint(_i)

	lat, log, err := gctx.GetPoint(i)
	if err != nil {
		return false, err
	}

	fmt.Printf("lat: %v, log: %v\n", lat, log)

	return false, nil
}

func (ac *GetPointCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}

func (ac *GetPointCommand) Info() string {
	return "get point [index]"
}
