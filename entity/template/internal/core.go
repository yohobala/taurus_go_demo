package internal

import (
	"context"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/dialect"
)

type Config struct {
	Tag    string
	Driver dialect.Driver
}

// NewConfig creates a new Config.
func NewConfig(tag string) (*Config, error) {
	c := &Config{
		Tag:    tag,
		Driver: nil,
	}
	err := c.initDriver()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) initDriver() error {
	driver, err := entity.GetConnection(c.Tag)
	if err != nil {
		return err
	}
	c.Driver = driver
	return nil
}

func (b *Config) MayTx(ctx context.Context) (dialect.Tx, error) {
	tx, err := b.Driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func SetEntityState(m *entity.Mutation, state entity.EntityState) error {
	current := m.State()
	switch state {
	case entity.Unchanged:
		m.SetState(state)
	case entity.Added:
		if current == entity.Detached {
			m.SetState(state)
		} else {
			return entity.Err_0100030003.Sprintf("Added", "Detached")
		}
	case entity.Modified:
		if current == entity.Unchanged {
			m.SetState(state)
		} else if current == entity.Added {
			return nil
		} else {
			return entity.Err_0100030003.Sprintf("Modified", "Unchanged")
		}
	case entity.Deleted:
		if current == entity.Unchanged || current == entity.Modified || current == entity.Added {
			if current == entity.Unchanged || current == entity.Modified {
				m.SetState(state)
			} else {
				m.SetState(entity.Detached)
			}
		} else {
			return entity.Err_0100030003.Sprintf("Deleted", "Unchanged 、 Modified or Added")
		}
	case entity.Detached:
		if current == entity.Deleted {
			m.SetState(state)
		}
	}
	return nil
}
