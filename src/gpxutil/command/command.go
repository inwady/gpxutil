package command

import "gpxutil/context"

type Command interface {
	Execute(gctx *context.GPXContext, params []string) error
	UnExecute(gctx *context.GPXContext) error
}

var (
	CommandTable = map[string]Command{
		"add": &AddCommand{},
		"remove": nil,
		"change": nil,
		"index": &SetIndexCommand{},

		"list": &ListCommand{},

		"import": &ImportCommand{},
	}
)
