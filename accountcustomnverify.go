// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountCustomNVerifyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountCustomNVerifyService]
// method instead.
type AccountCustomNVerifyService struct {
	Options []option.RequestOption
}

// NewAccountCustomNVerifyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountCustomNVerifyService(opts ...option.RequestOption) (r *AccountCustomNVerifyService) {
	r = &AccountCustomNVerifyService{}
	r.Options = opts
	return
}

// Verify Account Custom Nameserver Glue Records
func (r *AccountCustomNVerifyService) AccountLevelCustomNameserversVerifyAccountCustomNameserverGlueRecords(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AcnsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_ns/verify", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
