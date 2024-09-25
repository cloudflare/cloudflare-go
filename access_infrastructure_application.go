package cloudflare

type InfraProtocol string

const (
	InfraSSH InfraProtocol = "SSH"
	RDP      InfraProtocol = "RDP"
)

// WARNING : The following table has a constraint. Each organization can specify at maximum one protocol for each of their ports.
// There can be at maximum a single (OrgID, Port, Protocol) combination, which can be shared across multiple records.
type InfraTargetContext struct {
	TargetAttributes map[string][]string `json:"target_attributes"`
	Port             int                 `json:"port"`
	Protocol         InfraProtocol       `json:"protocol"`
}

type InfrastructureApplication struct {
	// Items common to both SAML and OIDC
	ID             string                       `json:"id,omitempty"`
	Type           AccessApplicationType        `json:"type"`
	Name           string                       `json:"name"`
	Aud            string                       `json:"aud,omitempty"`
	CreatedAt      string                       `json:"created_at,omitempty"`
	UpdatedAt      string                       `json:"updated_at,omitempty"`
	ScimConfig     *AccessApplicationSCIMConfig `json:"scim_config,omitempty"`
	TargetContexts []InfraTargetContext         `json:"target_criteria"`
	Policies       []AccessPolicy               `json:"policies"`
}

// Reuse methods from access_application
