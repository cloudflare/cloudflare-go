// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAccessKeyRotateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountAccessKeyRotateService]
// method instead.
type AccountAccessKeyRotateService struct {
	Options []option.RequestOption
}

// NewAccountAccessKeyRotateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessKeyRotateService(opts ...option.RequestOption) (r *AccountAccessKeyRotateService) {
	r = &AccountAccessKeyRotateService{}
	r.Options = opts
	return
}

// Perfoms a key rotation for an account.
func (r *AccountAccessKeyRotateService) AccessKeyConfigurationRotateAccessKeys(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *KeysSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/keys/rotate", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
