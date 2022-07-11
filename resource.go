package cloudflare

// RouteLevel holds the "level" where the resource resides.
type RouteLevel string

const (
	AccountRouteLevel RouteLevel = "accounts"
	ZoneRouteLevel    RouteLevel = "zones"
	UserRouteLevel    RouteLevel = "user"
)

// Resource defines an API resource you wish to target. Should not be used
// directly, use `UserIdentifer`, `ZoneIdentifier` and `AccountIdentifier`
// instead.
type Resource struct {
	Level      RouteLevel
	Identifier string
}

// UserIdentifer returns a user level *Resource.
func UserIdentifer(id string) *Resource {
	return &Resource{
		Level:      UserRouteLevel,
		Identifier: id,
	}
}

// ZoneIdentifer returns a zone level *Resource.
func ZoneIdentifer(id string) *Resource {
	return &Resource{
		Level:      ZoneRouteLevel,
		Identifier: id,
	}
}

// AccountIdentifer returns an account level *Resource.
func AccountIdentifer(id string) *Resource {
	return &Resource{
		Level:      AccountRouteLevel,
		Identifier: id,
	}
}
