package impl

import (
	"haha/conf"

	"github.com/reugn/go-streams"
	"github.com/reugn/go-streams/flow"
)

type PassFlow struct {
}

func NewPassFlow() *PassFlow {
	return &PassFlow{}
}

func (pf *PassFlow) ToFlow(config conf.FlowConfig) streams.Flow {
	return flow.NewPassThrough()
}
