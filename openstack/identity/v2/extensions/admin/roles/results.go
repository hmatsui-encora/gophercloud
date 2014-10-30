package roles

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

type Role struct {
	// The unique ID for the role.
	ID string

	// The human-readable name of the role.
	Name string

	// The description of the role.
	Description string
}

// RolePage is a single page of a user Role collection.
type RolePage struct {
	pagination.SinglePageBase
}

// IsEmpty determines whether or not a page of Tenants contains any results.
func (page RolePage) IsEmpty() (bool, error) {
	users, err := ExtractRoles(page)
	if err != nil {
		return false, err
	}
	return len(users) == 0, nil
}

// ExtractRoles returns a slice of roles contained in a single page of results.
func ExtractRoles(page pagination.Page) ([]Role, error) {
	casted := page.(RolePage).Body
	var response struct {
		Roles []Role `mapstructure:"roles"`
	}

	err := mapstructure.Decode(casted, &response)
	return response.Roles, err
}

type AddRoleResult struct {
	gophercloud.ErrResult
}
