// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// MiscategorizationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMiscategorizationService] method instead.
type MiscategorizationService struct {
	Options []option.RequestOption
}

// NewMiscategorizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMiscategorizationService(opts ...option.RequestOption) (r *MiscategorizationService) {
	r = &MiscategorizationService{}
	r.Options = opts
	return
}

// Allows you to submit requests to change a domain’s category.
func (r *MiscategorizationService) New(ctx context.Context, params MiscategorizationNewParams, opts ...option.RequestOption) (res *MiscategorizationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/miscategorization", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type MiscategorizationNewResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success MiscategorizationNewResponseSuccess `json:"success,required"`
	JSON    miscategorizationNewResponseJSON    `json:"-"`
}

// miscategorizationNewResponseJSON contains the JSON metadata for the struct
// [MiscategorizationNewResponse]
type miscategorizationNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MiscategorizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r miscategorizationNewResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MiscategorizationNewResponseSuccess bool

const (
	MiscategorizationNewResponseSuccessTrue MiscategorizationNewResponseSuccess = true
)

func (r MiscategorizationNewResponseSuccess) IsKnown() bool {
	switch r {
	case MiscategorizationNewResponseSuccessTrue:
		return true
	}
	return false
}

type MiscategorizationNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Content category IDs to add.
	ContentAdds param.Field[[]int64] `json:"content_adds"`
	// Content category IDs to remove.
	ContentRemoves param.Field[[]int64]                                 `json:"content_removes"`
	IndicatorType  param.Field[MiscategorizationNewParamsIndicatorType] `json:"indicator_type"`
	// Provide only if indicator_type is `ipv4` or `ipv6`.
	IP param.Field[string] `json:"ip"`
	// Security category IDs to add.
	SecurityAdds param.Field[[]int64] `json:"security_adds"`
	// Security category IDs to remove.
	SecurityRemoves param.Field[[]int64] `json:"security_removes"`
	// Provide only if indicator_type is `domain` or `url`. Example if indicator_type
	// is `domain`: `example.com`. Example if indicator_type is `url`:
	// `https://example.com/news/`.
	URL param.Field[string] `json:"url"`
}

func (r MiscategorizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MiscategorizationNewParamsIndicatorType string

const (
	MiscategorizationNewParamsIndicatorTypeDomain MiscategorizationNewParamsIndicatorType = "domain"
	MiscategorizationNewParamsIndicatorTypeIPV4   MiscategorizationNewParamsIndicatorType = "ipv4"
	MiscategorizationNewParamsIndicatorTypeIPV6   MiscategorizationNewParamsIndicatorType = "ipv6"
	MiscategorizationNewParamsIndicatorTypeURL    MiscategorizationNewParamsIndicatorType = "url"
)

func (r MiscategorizationNewParamsIndicatorType) IsKnown() bool {
	switch r {
	case MiscategorizationNewParamsIndicatorTypeDomain, MiscategorizationNewParamsIndicatorTypeIPV4, MiscategorizationNewParamsIndicatorTypeIPV6, MiscategorizationNewParamsIndicatorTypeURL:
		return true
	}
	return false
}
