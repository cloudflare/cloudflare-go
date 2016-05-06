package cloudflare

// Organization represents a multi-user organization.
type Organization struct {
	ID          string
	Name        string
	Status      string
	Permissions []string
	Roles       []string
}

// OrganizationResponse represents the response from the Organization endpoint.
type OrganizationResponse struct {
	Response
	Result     []Organization `json:"result"`
	ResultInfo ResultInfo     `json:"result_info"`
}
