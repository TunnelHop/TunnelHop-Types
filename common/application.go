// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package common

import "fmt"
import "encoding/json"

type Application struct {
	// ApplicationName corresponds to the JSON schema field "application_name".
	ApplicationName *string `json:"application_name,omitempty" yaml:"application_name,omitempty"`

	// Id corresponds to the JSON schema field "id".
	Id *int `json:"id,omitempty" yaml:"id,omitempty"`

	// InternalPort corresponds to the JSON schema field "internal_port".
	InternalPort int `json:"internal_port" yaml:"internal_port"`

	// IsActive corresponds to the JSON schema field "is_active".
	IsActive *bool `json:"is_active,omitempty" yaml:"is_active,omitempty"`

	// Protocol corresponds to the JSON schema field "protocol".
	Protocol string `json:"protocol" yaml:"protocol"`

	// RequireAuth corresponds to the JSON schema field "require_auth".
	RequireAuth *bool `json:"require_auth,omitempty" yaml:"require_auth,omitempty"`

	// RequireEncryption corresponds to the JSON schema field "require_encryption".
	RequireEncryption *bool `json:"require_encryption,omitempty" yaml:"require_encryption,omitempty"`

	// TcpPort corresponds to the JSON schema field "tcp_port".
	TcpPort *int `json:"tcp_port,omitempty" yaml:"tcp_port,omitempty"`

	// UserEmail corresponds to the JSON schema field "user_email".
	UserEmail string `json:"user_email" yaml:"user_email"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Application) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["internal_port"]; !ok || v == nil {
		return fmt.Errorf("field internal_port in Application: required")
	}
	if v, ok := raw["protocol"]; !ok || v == nil {
		return fmt.Errorf("field protocol in Application: required")
	}
	if v, ok := raw["user_email"]; !ok || v == nil {
		return fmt.Errorf("field user_email in Application: required")
	}
	type Plain Application
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Application(plain)
	return nil
}
