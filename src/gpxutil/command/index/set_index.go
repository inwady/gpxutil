package index

import (
	"gpxutil/context"
	"errors"
	"strconv"
)

type SetIndexCommand struct {
	fromIndex uint
}

func (sc *SetIndexCommand) Execute(gctx *context.GPXContext, params []string) (bool, error) {
	if len(params) < 2 {
		return false, errors.New("bad params")
	}

	index, err := strconv.ParseInt(params[1], 10, 32)
	if err != nil || index < 0 {
		return false, errors.New("bad index")
	}

	sc.fromIndex = gctx.GetIndex()

	err = gctx.SetWorkIndex(uint(index))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (sc *SetIndexCommand) UnExecute(gctx *context.GPXContext) error {
	err := gctx.SetWorkIndex(sc.fromIndex)
	if err != nil {
		return err
	}

	return nil
}

func (ac *SetIndexCommand) Info() string {
	return "set gpx index [index]"
}
