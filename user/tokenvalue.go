// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// TokenValueService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenValueService] method instead.
type TokenValueService struct {
	Options []option.RequestOption
}

// NewTokenValueService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTokenValueService(opts ...option.RequestOption) (r *TokenValueService) {
	r = &TokenValueService{}
	r.Options = opts
	return
}

// Roll the token secret.
func (r *TokenValueService) Update(ctx context.Context, tokenID interface{}, body TokenValueUpdateParams, opts ...option.RequestOption) (res *Value, err error) {
	opts = append(r.Options[:], opts...)
	var env TokenValueUpdateResponseEnvelope
	path := fmt.Sprintf("user/tokens/%v/value", tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Value = string

type TokenValueUpdateParams struct {
	Body interface{} `json:"body,required"`
}

func (r TokenValueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type TokenValueUpdateResponseEnvelope struct {
	// The token value.
	Result Value                                `json:"result"`
	JSON   tokenValueUpdateResponseEnvelopeJSON `json:"-"`
}

// tokenValueUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TokenValueUpdateResponseEnvelope]
type tokenValueUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenValueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenValueUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
