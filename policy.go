package cloudflare

type Policy struct {
	ID               string            `json:"id"`
	PermissionGroups []PermissionGroup `json:"permission_groups"`
	ResourceGroups   []ResourceGroup   `json:"resource_groups"`
	Access           string            `json:"access"`
}
type PermissionGroup struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Meta        map[string]string `json:"meta"`
	Permissions []Permission      `json:"permissions"`
}
type ResourceGroup struct {
	ID    string            `json:"id"`
	Name  string            `json:"name"`
	Meta  map[string]string `json:"meta"`
	Scope Scope             `json:"scope"`
}
type Permission struct {
	ID         string            `json:"id"`
	Key        string            `json:"key"`
	Attributes map[string]string `json:"attributes,omitempty"` // same as Meta in other structs
}
type Scope struct {
	Key          string        `json:"key"`
	ScopeObjects []ScopeObject `json:"objects"`
}
type ScopeObject struct {
	Key string `json:"key"`
}
