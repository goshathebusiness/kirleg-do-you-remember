package cfg

import (
	"os"

	"gopkg.in/yaml.v3"
)

// UnmarshalYAML unmarshals a .yaml file at path to destination dest
func UnmarshalYAML(path string, dest any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, dest)
}
