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
