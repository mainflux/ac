package ac

// Action represents any action.
type Action int64

// Predefined list of actions.
const (
	Any Action = iota
	Create
	Update
	Delete
	Read
	Write
)

// Rule represents Policy rule.
type Rule struct {
	Actions      []Action
	ResourceType string
	// True = allow
	Effect  bool
	Matcher Matcher
}

// Matcher represents Rule matcher.
type Matcher interface {
	Match(subject, object Resource) bool
}
