package gen

import (
	"context"

	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/entitysql"

	"taurus_go_demo/entity/template/internal"
	"taurus_go_demo/entity/template/user"
)

const UserTag = "usertag"

// User  is an struct of the database
type User struct {
	*internal.Config
	tracker entity.Tracker
	Users   *user.UserEntityBuilder
}

// NewUser creates a new User instance.
func NewUser() (*User, error) {
	config, err := internal.NewConfig(UserTag)
	if err != nil {
		return nil, err
	}
	user := &User{
		Config:  config,
		tracker: &entity.Tracking{},
	}
	user.init()
	return user, nil
}

// Close closes the database.
func (d *User) Close() error {
	return d.Driver.Close()
}

// Save saves all changes to the database.
func (d *User) Save(ctx context.Context) error {
	tx, err := d.Config.MayTx(ctx)
	if err != nil {
		return err
	}
	if err := func() error {
		for _, m := range d.tracker.Mutators() {
			if err := m.Exec(ctx, tx); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return entitysql.Rollback(tx, err)
	}
	return tx.Commit()
}

func (d *User) init() {
	d.Users = user.NewUserEntityBuilder(d.Config, d.tracker)
	d.tracker.Add(d.Users)
}
