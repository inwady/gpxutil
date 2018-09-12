package info

import (
	"gpxutil/context"
	"errors"
	"strconv"
	"fmt"
)

type ListCommand struct{}

func (ac *ListCommand) Execute(gctx *context.GPXContext, params []string) (bool, error) {
	var (
		index uint
	)

	if len(params) > 1 {
		tempIndex, err := strconv.ParseInt(params[1], 10, 32)
		if err != nil && tempIndex < 0 {
			return false, errors.New("bad index")
		}
		index = uint(tempIndex)
	} else {
		index = gctx.GetIndex()
	}

	data, err := gctx.GetListInfo(index)
	if err != nil {
		return false, err
	}

	fmt.Printf("%s", data)
	return false, nil
}

func (ac *ListCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}

func (ac *ListCommand) Info() string {
	return "info about current gpx"
}
