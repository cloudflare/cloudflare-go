// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// BrandProtectionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrandProtectionService] method instead.
type BrandProtectionService struct {
	Options     []option.RequestOption
	Queries     *QueryService
	Matches     *MatchService
	Logos       *LogoService
	LogoMatches *LogoMatchService
}

// NewBrandProtectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBrandProtectionService(opts ...option.RequestOption) (r *BrandProtectionService) {
	r = &BrandProtectionService{}
	r.Options = opts
	r.Queries = NewQueryService(opts...)
	r.Matches = NewMatchService(opts...)
	r.Logos = NewLogoService(opts...)
	r.LogoMatches = NewLogoMatchService(opts...)
	return
}

// Return new URL submissions
func (r *BrandProtectionService) Submit(ctx context.Context, body BrandProtectionSubmitParams, opts ...option.RequestOption) (res *BrandProtectionSubmitResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/submit", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Return submitted URLs based on ID
func (r *BrandProtectionService) URLInfo(ctx context.Context, query BrandProtectionURLInfoParams, opts ...option.RequestOption) (res *BrandProtectionURLInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/url-info", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BrandProtectionSubmitResponse struct {
	// Error code
	Code int64 `json:"code"`
	// Errors
	Errors map[string]interface{} `json:"errors"`
	// Error message
	Message string `json:"message"`
	// Error name
	Status string                            `json:"status"`
	JSON   brandProtectionSubmitResponseJSON `json:"-"`
}

// brandProtectionSubmitResponseJSON contains the JSON metadata for the struct
// [BrandProtectionSubmitResponse]
type brandProtectionSubmitResponseJSON struct {
	Code        apijson.Field
	Errors      apijson.Field
	Message     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionSubmitResponseJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionURLInfoResponse struct {
	// Error code
	Code int64 `json:"code"`
	// Errors
	Errors map[string]interface{} `json:"errors"`
	// Error message
	Message string `json:"message"`
	// Error name
	Status string                             `json:"status"`
	JSON   brandProtectionURLInfoResponseJSON `json:"-"`
}

// brandProtectionURLInfoResponseJSON contains the JSON metadata for the struct
// [BrandProtectionURLInfoResponse]
type brandProtectionURLInfoResponseJSON struct {
	Code        apijson.Field
	Errors      apijson.Field
	Message     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionURLInfoResponseJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionSubmitParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type BrandProtectionURLInfoParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
