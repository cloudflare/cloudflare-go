package cloudflare

type ResourceGroup struct {
	ID    string            `json:"id"`
	Name  string            `json:"name"`
	Meta  map[string]string `json:"meta"`
	Scope Scope             `json:"scope"`
}

type Scope struct {
	Key          string        `json:"key"`
	ScopeObjects []ScopeObject `json:"objects"`
}

type ScopeObject struct {
	Key string `json:"key"`
}
