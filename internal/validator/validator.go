package validator

// Validator is a type which contains map of validation errors.
type Validator struct {
	Errors map[string]string
}

// New is a helper method which create new Validator instance with empty errors map.
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// Valid return true if the errors map does not contain any entry.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError adds an error message to the map (so long as no entry already exists for
// the given key).
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check adds an error message to the map only if a validation check is not 'ok'.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}
