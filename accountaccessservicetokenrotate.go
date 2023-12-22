// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAccessServiceTokenRotateService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAccessServiceTokenRotateService] method instead.
type AccountAccessServiceTokenRotateService struct {
	Options []option.RequestOption
}

// NewAccountAccessServiceTokenRotateService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessServiceTokenRotateService(opts ...option.RequestOption) (r *AccountAccessServiceTokenRotateService) {
	r = &AccountAccessServiceTokenRotateService{}
	r.Options = opts
	return
}

// Generates a new Client Secret for a service token and revokes the old one.
func (r *AccountAccessServiceTokenRotateService) AccessServiceTokensRotateAServiceToken(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *CreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/rotate", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
