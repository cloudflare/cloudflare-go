// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountStreamTokenService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamTokenService] method
// instead.
type AccountStreamTokenService struct {
	Options []option.RequestOption
}

// NewAccountStreamTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamTokenService(opts ...option.RequestOption) (r *AccountStreamTokenService) {
	r = &AccountStreamTokenService{}
	r.Options = opts
	return
}

// Creates a signed URL token for a video. If a body is not provided in the
// request, a token is created with default values.
func (r *AccountStreamTokenService) StreamVideosNewSignedURLTokensForVideos(ctx context.Context, accountIdentifier string, identifier string, body AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParams, opts ...option.RequestOption) (res *SignedTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/%s/token", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SignedTokenResponse struct {
	Errors   []SignedTokenResponseError   `json:"errors"`
	Messages []SignedTokenResponseMessage `json:"messages"`
	Result   SignedTokenResponseResult    `json:"result"`
	// Whether the API call was successful
	Success SignedTokenResponseSuccess `json:"success"`
	JSON    signedTokenResponseJSON    `json:"-"`
}

// signedTokenResponseJSON contains the JSON metadata for the struct
// [SignedTokenResponse]
type signedTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SignedTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SignedTokenResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    signedTokenResponseErrorJSON `json:"-"`
}

// signedTokenResponseErrorJSON contains the JSON metadata for the struct
// [SignedTokenResponseError]
type signedTokenResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SignedTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SignedTokenResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    signedTokenResponseMessageJSON `json:"-"`
}

// signedTokenResponseMessageJSON contains the JSON metadata for the struct
// [SignedTokenResponseMessage]
type signedTokenResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SignedTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SignedTokenResponseResult struct {
	// The signed token used with the signed URLs feature.
	Token string                        `json:"token"`
	JSON  signedTokenResponseResultJSON `json:"-"`
}

// signedTokenResponseResultJSON contains the JSON metadata for the struct
// [SignedTokenResponseResult]
type signedTokenResponseResultJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SignedTokenResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SignedTokenResponseSuccess bool

const (
	SignedTokenResponseSuccessTrue SignedTokenResponseSuccess = true
)

type AccountStreamTokenStreamVideosNewSignedURLTokensForVideosParams struct {
}
