// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRulesetPhaseEntrypointVersionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountRulesetPhaseEntrypointVersionService] method instead.
type AccountRulesetPhaseEntrypointVersionService struct {
	Options []option.RequestOption
}

// NewAccountRulesetPhaseEntrypointVersionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountRulesetPhaseEntrypointVersionService(opts ...option.RequestOption) (r *AccountRulesetPhaseEntrypointVersionService) {
	r = &AccountRulesetPhaseEntrypointVersionService{}
	r.Options = opts
	return
}

// Fetches a specific version of an account entry point ruleset.
func (r *AccountRulesetPhaseEntrypointVersionService) Get(ctx context.Context, accountID string, rulesetPhase string, rulesetVersion string, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/phases/%s/entrypoint/versions/%s", accountID, rulesetPhase, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the versions of an account entry point ruleset.
func (r *AccountRulesetPhaseEntrypointVersionService) AccountRulesetsListAnAccountEntryPointRulesetSVersions(ctx context.Context, accountID string, rulesetPhase string, opts ...option.RequestOption) (res *RulesetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/phases/%s/entrypoint/versions", accountID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
