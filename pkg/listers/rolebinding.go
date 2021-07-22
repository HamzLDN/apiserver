package listers

import (
	"github.com/yubo/apiserver/pkg/storage"
	"github.com/yubo/apiserver/pkg/api/rbac"
	"github.com/yubo/golib/staging/labels"
	"github.com/yubo/golib/orm"
)

// RoleBindingLister helps list Roles.
// All objects returned here must be treated as read-only.
type RoleBindingLister interface {
	// List lists all Roles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*rbac.RoleBinding, err error)
	Get(name string) (*rbac.RoleBinding, error)
}

// roleBindingLister implements the RoleBindingLister interface.
type roleBindingLister struct {
	db *orm.DB
}

// NewRoleBindingLister returns a new RoleBindingLister.
func NewRoleBindingLister(db *orm.DB) RoleBindingLister {
	return &roleBindingLister{db: db}
}

// List lists all Roles in the indexer.
func (s *roleBindingLister) List(selector labels.Selector) (ret []*rbac.RoleBinding, err error) {
	err = storage.List(s.db, "cluster_role", selector, &ret)
	return
}

// Get retrieves the RoleBinding from the db for a given name.
func (s *roleBindingLister) Get(name string) (ret *rbac.RoleBinding, err error) {
	err = storage.Get(s.db, "cluster_role", name, &ret)
	return
}
