// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneRulesetPhaseEntrypointVersionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneRulesetPhaseEntrypointVersionService] method instead.
type ZoneRulesetPhaseEntrypointVersionService struct {
	Options []option.RequestOption
}

// NewZoneRulesetPhaseEntrypointVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneRulesetPhaseEntrypointVersionService(opts ...option.RequestOption) (r *ZoneRulesetPhaseEntrypointVersionService) {
	r = &ZoneRulesetPhaseEntrypointVersionService{}
	r.Options = opts
	return
}

// Fetches a specific version of a zone entry point ruleset.
func (r *ZoneRulesetPhaseEntrypointVersionService) Get(ctx context.Context, zoneID string, rulesetPhase string, rulesetVersion string, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/%s/entrypoint/versions/%s", zoneID, rulesetPhase, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the versions of a zone entry point ruleset.
func (r *ZoneRulesetPhaseEntrypointVersionService) ZoneRulesetsListAZoneEntryPointRulesetSVersions(ctx context.Context, zoneID string, rulesetPhase string, opts ...option.RequestOption) (res *RulesetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/%s/entrypoint/versions", zoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
