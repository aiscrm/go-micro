package memory

import (
	"context"

	"github.com/aiscrm/go-micro/v2/config/source"
)

type changeSetKey struct{}

func WithData(d []byte, f string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, changeSetKey{}, &source.ChangeSet{
			Data:   d,
			Format: f,
		})
	}
}

// WithChangeSet allows a changeset to be set
func WithChangeSet(cs *source.ChangeSet) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, changeSetKey{}, cs)
	}
}

// WithJSON allows the source data to be set to json
func WithJSON(d []byte) source.Option {
	return WithData(d, "json")
}

// WithYAML allows the source data to be set to yaml
func WithYAML(d []byte) source.Option {
	return WithData(d, "yaml")
}

// WithTOML allows the source data to be set to toml
func WithTOML(d []byte) source.Option {
	return WithData(d, "toml")
}
