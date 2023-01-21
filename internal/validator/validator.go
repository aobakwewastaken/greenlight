package validator

import "regexp"

// custom validator type which contains a map of errors
// we are able to conditionally add errors through the Check() method,
// return whether the map is empty or not, Valid()
type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// First check if the error does not already exist in the map
func (v *Validator) AddErrors(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddErrors(key, message)
	}
}

func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func Unique(values []string) bool {
	uniquValues := make(map[string]bool)

	for _, value := range values {
		uniquValues[value] = true
	}

	return len(values) == len(uniquValues)
}
