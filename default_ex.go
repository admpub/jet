package jet

import (
	"reflect"
)

func SetDefaultExtensions(exts ...string) {
	defaultExtensions = exts
}

func AddDefaultExtensions(exts ...string) {
	defaultExtensions = append(defaultExtensions, exts...)
}

func AddDefaultVariables(values map[string]interface{}) {
	for name, value := range values {
		defaultVariables[name] = reflect.ValueOf(value)
	}
}

func (s *Set) SetDefaultExtensions(exts ...string) *Set {
	s.defaultExtensions = exts
	return s
}

func (s *Set) AddDefaultExtensions(exts ...string) *Set {
	s.defaultExtensions = append(s.defaultExtensions, exts...)
	return s
}
