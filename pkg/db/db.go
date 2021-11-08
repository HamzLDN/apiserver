package db

import (
	"context"

	"github.com/yubo/golib/orm"
	"github.com/yubo/golib/util/errors"
)

const (
	DefaultName = "__default__"
)

type DB interface {
	orm.DB

	GetDB(name string) DB // panic if db[name] is not exist
}

type serverDB struct {
	name string
	orm.DB
	dbs    map[string]orm.DB
	ctx    context.Context
	cancel context.CancelFunc
}

func NewDB(ctx context.Context, config *Config) (DB, error) {
	ret := &serverDB{
		dbs: make(map[string]orm.DB),
	}
	ret.ctx, ret.cancel = context.WithCancel(ctx)

	for _, cf := range config.Databases {
		if cf.Dsn == "" || cf.Driver == "" {
			continue
		}
		opts := []orm.Option{
			orm.WithContext(ctx),
		}

		if cf.WithoutPing {
			opts = append(opts, orm.WithoutPing())
		}
		if cf.IgnoreNotFound {
			opts = append(opts, orm.WithIgnoreNotFound())
		}
		if cf.MaxRows > 0 {
			opts = append(opts, orm.WithMaxRows(cf.MaxRows))
		}
		if cf.MaxIdleCount > 0 {
			opts = append(opts, orm.WithMaxIdleCount(cf.MaxIdleCount))
		}
		if cf.MaxOpenConns > 0 {
			opts = append(opts, orm.WithMaxOpenConns(cf.MaxOpenConns))
		}
		if !cf.ConnMaxLifetime.IsZero() {
			opts = append(opts, orm.WithConnMaxLifetime(cf.ConnMaxLifetime.Duration))
		}
		if !cf.ConnMaxIdletime.IsZero() {
			opts = append(opts, orm.WithConnMaxLifetime(cf.ConnMaxIdletime.Duration))
		}

		if db, err := orm.Open(cf.Driver, cf.Dsn, opts...); err != nil {
			ret.cancel()
			return nil, err
		} else {
			ret.dbs[cf.Name] = db
		}
	}

	if db, ok := ret.dbs[DefaultName]; ok {
		ret.name = DefaultName
		ret.DB = db
	}

	return ret, nil
}

func (p *serverDB) Close() error {
	var errs []error
	for _, db := range p.dbs {
		if err := db.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.NewAggregate(errs)
}

func (p *serverDB) GetDB(name string) DB {
	if p == nil {
		panic("nil db")
	}

	if db, ok := p.dbs[name]; !ok {
		panic("nil db")
	} else {
		return &serverDB{
			name: name,
			DB:   db,
			dbs:  p.dbs,
		}
	}
}
