// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneRulesetVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneRulesetVersionService] method
// instead.
type ZoneRulesetVersionService struct {
	Options []option.RequestOption
}

// NewZoneRulesetVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneRulesetVersionService(opts ...option.RequestOption) (r *ZoneRulesetVersionService) {
	r = &ZoneRulesetVersionService{}
	r.Options = opts
	return
}

// Fetches a specific version of a zone ruleset.
func (r *ZoneRulesetVersionService) Get(ctx context.Context, zoneID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s/versions/%s", zoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an existing version of a zone ruleset.
func (r *ZoneRulesetVersionService) Delete(ctx context.Context, zoneID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s/versions/%s", zoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches the versions of a zone ruleset.
func (r *ZoneRulesetVersionService) ZoneRulesetsListAZoneRulesetSVersions(ctx context.Context, zoneID string, rulesetID string, opts ...option.RequestOption) (res *RulesetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s/versions", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
