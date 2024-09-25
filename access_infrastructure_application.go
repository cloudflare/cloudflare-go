package cloudflare

type InfraProtocol string

const (
	InfraSSH InfraProtocol = "SSH"
	RDP      InfraProtocol = "RDP"
)

type InfraTargetContext struct {
	TargetAttributes map[string][]string `json:"target_attributes"`
	Port             int                 `json:"port"`
	Protocol         InfraProtocol       `json:"protocol"`
}
