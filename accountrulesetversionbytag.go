// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRulesetVersionByTagService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountRulesetVersionByTagService] method instead.
type AccountRulesetVersionByTagService struct {
	Options []option.RequestOption
}

// NewAccountRulesetVersionByTagService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountRulesetVersionByTagService(opts ...option.RequestOption) (r *AccountRulesetVersionByTagService) {
	r = &AccountRulesetVersionByTagService{}
	r.Options = opts
	return
}

// Fetches the rules of a managed account ruleset version for a given tag.
func (r *AccountRulesetVersionByTagService) Get(ctx context.Context, accountID string, rulesetID string, rulesetVersion string, ruleTag string, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s/by_tag/%s", accountID, rulesetID, rulesetVersion, ruleTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
