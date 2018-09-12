package exchange

import (
	"gpxutil/context"
	"errors"
	"gpxutil/util"
)

type ExchangeCommand struct {
	index uint
}

func (ec *ExchangeCommand) Execute(gctx *context.GPXContext, params []string) (bool, error) {
	if len(params) < 3 {
		return false, errors.New("bad params")
	}

	name := params[1]
	data := params[2]

	gpx, err := util.ImportFromPolyline(name, []byte(data))
	if err != nil {
		return false, err
	}

	ec.index = gctx.AddGPX(gpx)
	return true, nil
}

func (ec *ExchangeCommand) UnExecute(gctx *context.GPXContext) error {
	return gctx.RemoveGPX(ec.index)
}

func (ac *ExchangeCommand) Info() string {
	return "import new gpx [polyline]"
}
