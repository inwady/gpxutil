package point

import (
	"gpxutil/context"
	"errors"
	"strconv"
)

type RemoveCommand struct  {
	index uint
	changedLat float64
	changedLog float64
}

func (ac *RemoveCommand) Execute(gctx *context.GPXContext, params []string) (bool, error) {
	if len(params) < 2 {
		return false, errors.New("bad params")
	}

	_i, err := strconv.ParseInt(params[1], 10, 32)
	if err != nil || _i < 0 {
		return false, errors.New("bad index")
	}

	i := uint(_i)

	_lat, _log, err := gctx.GetPoint(i)
	ac.index = i
	ac.changedLat = _lat
	ac.changedLog = _log

	return true, gctx.RemovePoint(i)
}

func (ac *RemoveCommand) UnExecute(gctx *context.GPXContext) error {
	return gctx.AddPoint(ac.index - 1, ac.changedLat, ac.changedLog)
}
