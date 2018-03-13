package command

import (
	"gpxutil/context"
	"errors"
	"strconv"
	"fmt"
)

type ListCommand struct{}

func (ac *ListCommand) Execute(gctx *context.GPXContext, params []string) error {
	var (
		index uint
	)

	if len(params) > 1 {
		tempIndex, err := strconv.ParseInt(params[1], 10, 32)
		if err != nil && tempIndex < 0 {
			return errors.New("bad index")
		}
		index = uint(tempIndex)
	} else {
		index = gctx.GetIndex()
	}

	data, err := gctx.GetListInfo(index)
	if err != nil {
		return err
	}

	fmt.Printf("%s", data)
	return nil
}

func (ac *ListCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}