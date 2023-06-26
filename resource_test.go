package cloudflare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourcURLFragment(t *testing.T) {
	tests := map[string]struct {
		container *ResourceContainer
		want      string
	}{
		"account resource": {container: AccountIdentifier("foo"), want: "accounts/foo"},
		"zone resource":    {container: ZoneIdentifier("foo"), want: "zones/foo"},
		// this is pretty well deprecated in favour of `AccountIdentifier` but
		// here for completeness.
		"user level resource":    {container: UserIdentifier("foo"), want: "user"},
		"missing level resource": {container: &ResourceContainer{Level: "", Identifier: "foo"}, want: "foo"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.container.URLFragment()
			assert.Equal(t, tc.want, got)
		})
	}
}
