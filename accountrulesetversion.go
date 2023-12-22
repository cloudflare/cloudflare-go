// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRulesetVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRulesetVersionService]
// method instead.
type AccountRulesetVersionService struct {
	Options []option.RequestOption
	ByTags  *AccountRulesetVersionByTagService
}

// NewAccountRulesetVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRulesetVersionService(opts ...option.RequestOption) (r *AccountRulesetVersionService) {
	r = &AccountRulesetVersionService{}
	r.Options = opts
	r.ByTags = NewAccountRulesetVersionByTagService(opts...)
	return
}

// Fetches a specific version of an account ruleset.
func (r *AccountRulesetVersionService) Get(ctx context.Context, accountID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s", accountID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an existing version of an account ruleset.
func (r *AccountRulesetVersionService) Delete(ctx context.Context, accountID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s", accountID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches the versions of an account ruleset.
func (r *AccountRulesetVersionService) AccountRulesetsListAnAccountRulesetSVersions(ctx context.Context, accountID string, rulesetID string, opts ...option.RequestOption) (res *RulesetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
