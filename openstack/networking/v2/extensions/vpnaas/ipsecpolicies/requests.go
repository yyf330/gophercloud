package ipsecpolicies

import "github.com/gophercloud/gophercloud"

type TransformProtocol string
type AuthAlgorithm string
type EncapsulationMode string
type EncryptionAlgorithm string
type PFS string
type Unit string

const (
	TransformProtocolESP       TransformProtocol   = "esp"
	TransformProtocolAH        TransformProtocol   = "ah"
	TransformProtocolAHESP     TransformProtocol   = "ah-esp"
	AuthAlgorithmSHA1          AuthAlgorithm       = "sha1"
	AuthAlgorithmSHA256        AuthAlgorithm       = "sha256"
	AuthAlgorithmSHA384        AuthAlgorithm       = "sha384"
	AuthAlgorithmSHA512        AuthAlgorithm       = "sha512"
	EncryptionAlgorithm3DES    EncryptionAlgorithm = "3des"
	EncryptionAlgorithmAES128  EncryptionAlgorithm = "aes-128"
	EncryptionAlgorithmAES256  EncryptionAlgorithm = "aes-256"
	EncryptionAlgorithmAES192  EncryptionAlgorithm = "aes-192"
	EncapsulationModeTunnel    EncapsulationMode   = "tunnel"
	EncapsulationModeTransport EncapsulationMode   = "transport"
	UnitSeconds                Unit                = "seconds"
	UnitKilobytes              Unit                = "kilobytes"
	PFSGroup2                  PFS                 = "group2"
	PFSGroup5                  PFS                 = "group5"
	PFSGroup14                 PFS                 = "group14"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new IPSec policy
type CreateOpts struct {
	// TenantID specifies a tenant to own the IPSec policy. The caller must have
	// an admin role in order to set this. Otherwise, this field is left unset
	// and the caller will be the owner.
	TenantID string `json:"tenant_id,omitempty"`

	// Description is the human readable description of the policy.
	Description string `json:"description,omitempty"`

	// Name is the human readable name of the policy.
	// Does not have to be unique.
	Name string `json:"name,omitempty"`

	// AuthAlgorithm is the authentication hash algorithm.
	// Valid values are sha1, sha256, sha384, sha512.
	// The default is sha1.
	AuthAlgorithm AuthAlgorithm `json:"auth_algorithm,omitempty"`

	// EncapsulationMode is the encapsulation mode.
	// A valid value is tunnel or transport.
	// Default is tunnel.
	EncapsulationMode EncapsulationMode `json:"encapsulation_mode,omitempty"`

	// EncryptionAlgorithm is the encryption algorithm.
	// A valid value is 3des, aes-128, aes-192, aes-256, and so on.
	// Default is aes-128.
	EncryptionAlgorithm EncryptionAlgorithm `json:"encryption_algorithm,omitempty"`

	// PFS is the Perfect forward secrecy mode.
	// A valid value is Group2, Group5, Group14, and so on.
	// Default is Group5.
	PFS PFS `json:"pfs,omitempty"`

	// TransformProtocol is the transform protocol.
	// A valid value is ESP, AH, or AH- ESP.
	// Default is ESP.
	TransformProtocol TransformProtocol `json:"transform_protocol,omitempty"`

	//Lifetime is the lifetime of the security association
	Lifetime *LifetimeCreateOpts `json:"lifetime,omitempty"`
}

// The lifetime consists of a unit and integer value
// You can omit either the unit or value portion of the lifetime
type LifetimeCreateOpts struct {
	// Units is the units for the lifetime of the security association
	// Default unit is seconds
	Units Unit `json:"units,omitempty"`

	// The lifetime value.
	// Must be a positive integer.
	// Default value is 3600.
	Value int `json:"value,omitempty"`
}

// ToPolicyCreateMap casts a CreateOpts struct to a map.
func (opts CreateOpts) ToPolicyCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "ipsecpolicy")
}

// Create accepts a CreateOpts struct and uses the values to create a new
// IPSec policy
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPolicyCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, nil)
	return
}
