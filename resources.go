package ac

// Resource represents any resource in the system.
// All the actions are executed by Resource and against Resource.
type Resource interface {
	// ID field is likely to be taken.
	ResourceID() string
	Type() string
	Attributes() map[string]string
}

// Constraint is Policy constraint.
type Constraint interface {
	Validate(r Resource) bool
}
