package command

import (
	"gpxutil/context"
	"errors"
	"gpxutil/util"
)

type ImportCommand struct{}

func (ac *ImportCommand) Execute(gctx *context.GPXContext, params []string) error {
	if len(params) < 3 {
		return errors.New("bad params")
	}

	name := params[1]
	data := params[2]

	gpx, err := util.ImportFromPolyline(name, []byte(data))
	if err != nil {
		return err
	}

	gctx.AddGPX(gpx)
	return nil
}

func (ac *ImportCommand) UnExecute(gctx *context.GPXContext) error {
	return nil
}