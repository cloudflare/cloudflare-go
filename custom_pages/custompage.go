// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// CustomPageService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomPageService] method instead.
type CustomPageService struct {
	Options []option.RequestOption
}

// NewCustomPageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomPageService(opts ...option.RequestOption) (r *CustomPageService) {
	r = &CustomPageService{}
	r.Options = opts
	return
}

// Updates the configuration of an existing custom page.
func (r *CustomPageService) Update(ctx context.Context, identifier string, params CustomPageUpdateParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env CustomPageUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/custom_pages/%s", accountOrZone, accountOrZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all the custom pages.
func (r *CustomPageService) List(ctx context.Context, query CustomPageListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CustomPageListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/custom_pages", accountOrZone, accountOrZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches all the custom pages.
func (r *CustomPageService) ListAutoPaging(ctx context.Context, query CustomPageListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CustomPageListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Fetches the details of a custom page.
func (r *CustomPageService) Get(ctx context.Context, identifier string, query CustomPageGetParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env CustomPageGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/custom_pages/%s", accountOrZone, accountOrZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CustomPageListResponse = interface{}

type CustomPageUpdateParams struct {
	// The custom page state.
	State param.Field[CustomPageUpdateParamsState] `json:"state,required"`
	// The URL associated with the custom page.
	URL param.Field[string] `json:"url,required" format:"uri"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r CustomPageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The custom page state.
type CustomPageUpdateParamsState string

const (
	CustomPageUpdateParamsStateDefault    CustomPageUpdateParamsState = "default"
	CustomPageUpdateParamsStateCustomized CustomPageUpdateParamsState = "customized"
)

func (r CustomPageUpdateParamsState) IsKnown() bool {
	switch r {
	case CustomPageUpdateParamsStateDefault, CustomPageUpdateParamsStateCustomized:
		return true
	}
	return false
}

type CustomPageUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success CustomPageUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    customPageUpdateResponseEnvelopeJSON    `json:"-"`
}

// customPageUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomPageUpdateResponseEnvelope]
type customPageUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPageUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomPageUpdateResponseEnvelopeSuccess bool

const (
	CustomPageUpdateResponseEnvelopeSuccessTrue CustomPageUpdateResponseEnvelopeSuccess = true
)

func (r CustomPageUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomPageUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomPageListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type CustomPageGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type CustomPageGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success CustomPageGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    customPageGetResponseEnvelopeJSON    `json:"-"`
}

// customPageGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomPageGetResponseEnvelope]
type customPageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomPageGetResponseEnvelopeSuccess bool

const (
	CustomPageGetResponseEnvelopeSuccessTrue CustomPageGetResponseEnvelopeSuccess = true
)

func (r CustomPageGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomPageGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
