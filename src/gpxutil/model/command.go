package model

import "gpxutil/context"

type Command interface {
	Execute(gctx *context.GPXContext, params []string) (bool, error)
	UnExecute(gctx *context.GPXContext) error
	Info() string
}
