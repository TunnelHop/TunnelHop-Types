// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package common

import "fmt"
import "encoding/json"

type ServerJson struct {
	// Name corresponds to the JSON schema field "name".
	Name string `json:"name" yaml:"name"`

	// Port corresponds to the JSON schema field "port".
	Port int `json:"port" yaml:"port"`

	// Status corresponds to the JSON schema field "status".
	Status bool `json:"status" yaml:"status"`

	// Url corresponds to the JSON schema field "url".
	Url string `json:"url" yaml:"url"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ServerJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in ServerJson: required")
	}
	if v, ok := raw["port"]; !ok || v == nil {
		return fmt.Errorf("field port in ServerJson: required")
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status in ServerJson: required")
	}
	if v, ok := raw["url"]; !ok || v == nil {
		return fmt.Errorf("field url in ServerJson: required")
	}
	type Plain ServerJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ServerJson(plain)
	return nil
}
