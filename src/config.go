package src

import (
	"gopkg.in/yaml.v3"
	"os"
)

// ReadConfigFromYaml New function to read yaml into config struct
func (m *Moujin) ReadConfigFromYaml(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &m.C)
	return err
}
