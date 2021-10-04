package conf

type StreamConfig struct {
	EntityConfigs []EntityConfig `yaml:"entitys"`
	AggConfigs    []AggConfig    `yaml:"aggs"`
}

func (sc StreamConfig) FindEntityConfig(id int) (EntityConfig, bool) {
	for _, ec := range sc.EntityConfigs {
		if ec.Id == id {
			return ec, true
		}
	}
	return EntityConfig{}, false
}

type EntityConfig struct {
	Id         int    `yaml:"id"`
	EntityType string `yaml:"type"`
	Url        string `yaml:"url"`

	// kafka
	Topic   string `yaml:"topic"`
	Group   string `yaml:"group"`
	Version string `yaml:"version"`

	// mongo
}

type AggConfig struct {
	SourceId    int          `yaml:"source"`
	SinkId      int          `yaml:"sink"`
	FlowConfigs []FlowConfig `yaml:"flows"`
}

type FlowConfig struct {
	FlowType string `yaml:"flow"`
}
