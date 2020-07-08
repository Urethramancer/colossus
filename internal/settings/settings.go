package settings

// Settings is a map of string variables.
type Settings struct {
	defaults map[string]string
	settings map[string]string
}

func (s *Settings) InitVars(def map[string]string) {
	s.settings = make(map[string]string)
	s.defaults = make(map[string]string)
	for k, v := range def {
		s.defaults[k] = v
	}
}

// Set variable.
func (s *Settings) Set(k, v string) {
	s.settings[k] = v
}

// Get variable.
func (s *Settings) Get(k string) string {
	v, ok := s.settings[k]
	if ok {
		return v
	}

	v, ok = s.defaults[k]
	if ok {
		return v
	}

	return ""
}
