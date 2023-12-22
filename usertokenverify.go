// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserTokenVerifyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserTokenVerifyService] method
// instead.
type UserTokenVerifyService struct {
	Options []option.RequestOption
}

// NewUserTokenVerifyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserTokenVerifyService(opts ...option.RequestOption) (r *UserTokenVerifyService) {
	r = &UserTokenVerifyService{}
	r.Options = opts
	return
}

// Test whether a token works.
func (r *UserTokenVerifyService) UserAPITokensVerifyToken(ctx context.Context, opts ...option.RequestOption) (res *ResponseSingleSegment, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/tokens/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ResponseSingleSegment struct {
	Errors   []ResponseSingleSegmentError   `json:"errors"`
	Messages []ResponseSingleSegmentMessage `json:"messages"`
	Result   ResponseSingleSegmentResult    `json:"result"`
	// Whether the API call was successful
	Success ResponseSingleSegmentSuccess `json:"success"`
	JSON    responseSingleSegmentJSON    `json:"-"`
}

// responseSingleSegmentJSON contains the JSON metadata for the struct
// [ResponseSingleSegment]
type responseSingleSegmentJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleSegment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseSingleSegmentError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    responseSingleSegmentErrorJSON `json:"-"`
}

// responseSingleSegmentErrorJSON contains the JSON metadata for the struct
// [ResponseSingleSegmentError]
type responseSingleSegmentErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleSegmentError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseSingleSegmentMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    responseSingleSegmentMessageJSON `json:"-"`
}

// responseSingleSegmentMessageJSON contains the JSON metadata for the struct
// [ResponseSingleSegmentMessage]
type responseSingleSegmentMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleSegmentMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseSingleSegmentResult struct {
	// Token identifier tag.
	ID string `json:"id,required"`
	// Status of the token.
	Status ResponseSingleSegmentResultStatus `json:"status,required"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time                       `json:"not_before" format:"date-time"`
	JSON      responseSingleSegmentResultJSON `json:"-"`
}

// responseSingleSegmentResultJSON contains the JSON metadata for the struct
// [ResponseSingleSegmentResult]
type responseSingleSegmentResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	ExpiresOn   apijson.Field
	NotBefore   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleSegmentResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the token.
type ResponseSingleSegmentResultStatus string

const (
	ResponseSingleSegmentResultStatusActive   ResponseSingleSegmentResultStatus = "active"
	ResponseSingleSegmentResultStatusDisabled ResponseSingleSegmentResultStatus = "disabled"
	ResponseSingleSegmentResultStatusExpired  ResponseSingleSegmentResultStatus = "expired"
)

// Whether the API call was successful
type ResponseSingleSegmentSuccess bool

const (
	ResponseSingleSegmentSuccessTrue ResponseSingleSegmentSuccess = true
)
