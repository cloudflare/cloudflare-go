package cloudflare

import "fmt"

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

func NewResourceGroupForZone(zone Zone) ResourceGroup {
	return NewResourceGroup(fmt.Sprintf("com.cloudflare.api.account.zone.%s", zone.ID))
}

func NewResourceGroupForAccount(account Account) ResourceGroup {
	return NewResourceGroup(fmt.Sprintf("com.cloudflare.api.account.%s", account.ID))
}

func NewResourceGroup(key string) ResourceGroup {
	scope := Scope{
		Key: key,
		ScopeObjects: []ScopeObject{
			ScopeObject{
				Key: "*",
			},
		},
	}
	resourceGroup := ResourceGroup{
		ID:   "",
		Name: key,
		Meta: map[string]string{
			"editable": "false",
		},
		Scope: scope,
	}
	return resourceGroup
}
