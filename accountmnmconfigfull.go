// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountMnmConfigFullService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountMnmConfigFullService]
// method instead.
type AccountMnmConfigFullService struct {
	Options []option.RequestOption
}

// NewAccountMnmConfigFullService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMnmConfigFullService(opts ...option.RequestOption) (r *AccountMnmConfigFullService) {
	r = &AccountMnmConfigFullService{}
	r.Options = opts
	return
}

// Lists default sampling, router IPs, and rules for account.
func (r *AccountMnmConfigFullService) MagicNetworkMonitoringConfigurationListRulesAndAccountConfiguration(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/config/full", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
