package entity

import (
	"haha/conf"
	"haha/entity/impl"

	"github.com/reugn/go-streams"
)

type ISourceEntify interface {
	ToSource(config conf.EntityConfig) streams.Source
}

type ISinkEntify interface {
	ToSink(config conf.EntityConfig) streams.Sink
}

type IFlow interface {
	ToFlow(config conf.FlowConfig) streams.Flow
}

var (
	sourceEntitys = map[string]ISourceEntify{
		"kafka": impl.NewKafkaEntity(),
	}
	sinkEntitys = map[string]ISinkEntify{
		"kafka": impl.NewKafkaEntity(),
	}
	flows = map[string]IFlow{
		"pass": impl.NewPassFlow(),
	}
)

func GetSourceEntity(config conf.EntityConfig) (streams.Source, bool) {
	sourceEntity, ok := sourceEntitys[config.EntityType]
	if !ok {
		return nil, false
	}
	return sourceEntity.ToSource(config), ok
}

func GetSinkEntity(config conf.EntityConfig) (streams.Sink, bool) {
	sinkEntity, ok := sinkEntitys[config.EntityType]
	if !ok {
		return nil, false
	}
	return sinkEntity.ToSink(config), ok
}

func GetFlow(config conf.FlowConfig) (streams.Flow, bool) {
	f, ok := flows[config.FlowType]
	if !ok {
		return nil, false
	}
	return f.ToFlow(config), ok
}
