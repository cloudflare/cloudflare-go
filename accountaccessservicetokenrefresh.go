// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAccessServiceTokenRefreshService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAccessServiceTokenRefreshService] method instead.
type AccountAccessServiceTokenRefreshService struct {
	Options []option.RequestOption
}

// NewAccountAccessServiceTokenRefreshService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessServiceTokenRefreshService(opts ...option.RequestOption) (r *AccountAccessServiceTokenRefreshService) {
	r = &AccountAccessServiceTokenRefreshService{}
	r.Options = opts
	return
}

// Refreshes the expiration of a service token.
func (r *AccountAccessServiceTokenRefreshService) AccessServiceTokensRefreshAServiceToken(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ServiceTokensSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/refresh", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
