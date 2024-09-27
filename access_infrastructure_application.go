package cloudflare

type AccessInfrastructureProtocol string

const (
	AccessInfrastructureSSH AccessInfrastructureProtocol = "SSH"
	AccessInfrastructureRDP AccessInfrastructureProtocol = "RDP"
)

type AccessInfrastructureTargetContext struct {
	TargetAttributes map[string][]string          `json:"target_attributes"`
	Port             int                          `json:"port"`
	Protocol         AccessInfrastructureProtocol `json:"protocol"`
}

type AccessInfrastructureConnectionRulesSSH struct {
	Usernames []string `json:"usernames"`
}

type AccessInfrastructureConnectionRules struct {
	SSH *AccessInfrastructureConnectionRulesSSH `json:"ssh,omitempty"`
}
