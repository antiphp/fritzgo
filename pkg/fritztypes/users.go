// Package fritztypes contains fritz types.
package fritztypes

// User represents a fritz user.
type User struct {
	// Name contains the user name.
	Name string

	// Default determines whether it is considered the default user, or not.
	Default bool
}
