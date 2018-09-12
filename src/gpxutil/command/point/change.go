package point

import (
	"gpxutil/context"
	"errors"
	"strconv"
)

type ChangeCommand struct {
	index uint
	changedLat float64
	changedLog float64
}

func (ac *ChangeCommand) Execute(gctx *context.GPXContext, params []string) (bool, error) {
	if len(params) < 4 {
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

	_lat, _log, err := gctx.GetPoint(i)
	ac.index = i
	ac.changedLat = _lat
	ac.changedLog = _log

	err = gctx.ChangePoint(i, lat, log)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (ac *ChangeCommand) UnExecute(gctx *context.GPXContext) error {
	return gctx.ChangePoint(ac.index, ac.changedLat, ac.changedLog)
}

func (ac *ChangeCommand) Info() string {
	return "change point [index, lat, log]"
}
