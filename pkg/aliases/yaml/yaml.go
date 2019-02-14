package yaml

import (
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Unmarshal parses YAML-encoded data and returns config specification.
func Unmarshal(buf []byte) (*ConfigSpec, error) {
	spec := ConfigSpec{}
	if err := yaml.UnmarshalStrict(buf, &spec); err != nil {
		if e, ok := err.(*yaml.TypeError); ok {
			message := e.Errors[0]
			message = strings.Replace(message, " into yaml.DependencySpec", "", 1)
			message = strings.Replace(message, " into yaml.OptionSpec", "", 1)
			message = strings.Replace(message, " into yaml.ConfigSpec", "", 1)
			message = strings.Replace(message, " in type yaml.DependencySpec", "", 1)
			message = strings.Replace(message, " in type yaml.OptionSpec", "", 1)
			message = strings.Replace(message, " in type yaml.ConfigSpec", "", 1)
			return nil, Error(message)
		}
		return nil, Error(err)
	}

	v := NewValidator()
	if err := spec.DepthWalk(func(path SpecPath, current OptionSpec) (*OptionSpec, error) {
		if err := v.Struct(current); err != nil {
			return nil, Errorf("%s in `%s`", err, path)
		}
		for i, d := range current.Dependencies {
			if d.IsConfig() {
				continue
			}
			_, ok := spec[d.String()]
			if !ok {
				return nil, Errorf("invalid parameter `%s` for `dependencies[%d]` is an undefined dependency in `%s`", d.String(), i, path)
			}
		}
		return &current, nil
	}); err != nil {
		return nil, err
	}

	return &spec, nil
}
