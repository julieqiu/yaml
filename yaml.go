package yaml

import "gopkg.in/yaml.v3"

func Unmarshal(in []byte, out interface{}) (err error) {
	return yaml.Unmarshal(in, out)
}
