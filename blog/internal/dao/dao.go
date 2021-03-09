package dao

import (
	"context"

	"github.com/ccsunnyfd/practice/blog/ent"
	"google.golang.org/protobuf/internal/errors"
)

// Dao Dao
type Dao interface {
	Close()
}

// dao dao
type dao struct {
	*ent.Client
}

// NewDao NewDao
func NewDao() (dao Dao, cf func(), err error) {
	dao, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		err = errors.Wrap("failed opening connection to sqlite", err)
	}
	cf = dao.Close
	// Run the auto migration tool.
	if err := dao.Schema.Create(context.Background()); err != nil {
		err = errors.Wrap("failed creating schema resources", err)
	}
}
