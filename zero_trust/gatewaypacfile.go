// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// GatewayPacfileService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayPacfileService] method instead.
type GatewayPacfileService struct {
	Options []option.RequestOption
}

// NewGatewayPacfileService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayPacfileService(opts ...option.RequestOption) (r *GatewayPacfileService) {
	r = &GatewayPacfileService{}
	r.Options = opts
	return
}

// Create a new Zero Trust Gateway PAC file.
func (r *GatewayPacfileService) New(ctx context.Context, params GatewayPacfileNewParams, opts ...option.RequestOption) (res *GatewayPacfileNewResponse, err error) {
	var env GatewayPacfileNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/pacfiles", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a configured Zero Trust Gateway PAC file.
func (r *GatewayPacfileService) Update(ctx context.Context, pacfileID string, params GatewayPacfileUpdateParams, opts ...option.RequestOption) (res *GatewayPacfileUpdateResponse, err error) {
	var env GatewayPacfileUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pacfileID == "" {
		err = errors.New("missing required pacfile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/pacfiles/%s", params.AccountID, pacfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Zero Trust Gateway PAC files for an account.
func (r *GatewayPacfileService) List(ctx context.Context, query GatewayPacfileListParams, opts ...option.RequestOption) (res *pagination.SinglePage[GatewayPacfileListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/pacfiles", query.AccountID)
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

// List all Zero Trust Gateway PAC files for an account.
func (r *GatewayPacfileService) ListAutoPaging(ctx context.Context, query GatewayPacfileListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[GatewayPacfileListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a configured Zero Trust Gateway PAC file.
func (r *GatewayPacfileService) Delete(ctx context.Context, pacfileID string, body GatewayPacfileDeleteParams, opts ...option.RequestOption) (res *GatewayPacfileDeleteResponse, err error) {
	var env GatewayPacfileDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pacfileID == "" {
		err = errors.New("missing required pacfile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/pacfiles/%s", body.AccountID, pacfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a single Zero Trust Gateway PAC file.
func (r *GatewayPacfileService) Get(ctx context.Context, pacfileID string, query GatewayPacfileGetParams, opts ...option.RequestOption) (res *GatewayPacfileGetResponse, err error) {
	var env GatewayPacfileGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pacfileID == "" {
		err = errors.New("missing required pacfile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/pacfiles/%s", query.AccountID, pacfileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayPacfileNewResponse struct {
	ID string `json:"id"`
	// Actual contents of the PAC file
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Detailed description of the PAC file.
	Description string `json:"description"`
	// Name of the PAC file.
	Name string `json:"name"`
	// URL-friendly version of the PAC file name.
	Slug      string    `json:"slug"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Unique URL to download the PAC file.
	URL  string                        `json:"url"`
	JSON gatewayPacfileNewResponseJSON `json:"-"`
}

// gatewayPacfileNewResponseJSON contains the JSON metadata for the struct
// [GatewayPacfileNewResponse]
type gatewayPacfileNewResponseJSON struct {
	ID          apijson.Field
	Contents    apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayPacfileNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayPacfileNewResponseJSON) RawJSON() string {
	return r.raw
}

type GatewayPacfileUpdateResponse struct {
	ID string `json:"id"`
	// Actual contents of the PAC file
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Detailed description of the PAC file.
	Description string `json:"description"`
	// Name of the PAC file.
	Name string `json:"name"`
	// URL-friendly version of the PAC file name.
	Slug      string    `json:"slug"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Unique URL to download the PAC file.
	URL  string                           `json:"url"`
	JSON gatewayPacfileUpdateResponseJSON `json:"-"`
}

// gatewayPacfileUpdateResponseJSON contains the JSON metadata for the struct
// [GatewayPacfileUpdateResponse]
type gatewayPacfileUpdateResponseJSON struct {
	ID          apijson.Field
	Contents    apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayPacfileUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayPacfileUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type GatewayPacfileListResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Detailed description of the PAC file.
	Description string `json:"description"`
	// Name of the PAC file.
	Name string `json:"name"`
	// URL-friendly version of the PAC file name.
	Slug      string    `json:"slug"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Unique URL to download the PAC file.
	URL  string                         `json:"url"`
	JSON gatewayPacfileListResponseJSON `json:"-"`
}

// gatewayPacfileListResponseJSON contains the JSON metadata for the struct
// [GatewayPacfileListResponse]
type gatewayPacfileListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayPacfileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayPacfileListResponseJSON) RawJSON() string {
	return r.raw
}

type GatewayPacfileDeleteResponse = interface{}

type GatewayPacfileGetResponse struct {
	ID string `json:"id"`
	// Actual contents of the PAC file
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Detailed description of the PAC file.
	Description string `json:"description"`
	// Name of the PAC file.
	Name string `json:"name"`
	// URL-friendly version of the PAC file name.
	Slug      string    `json:"slug"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Unique URL to download the PAC file.
	URL  string                        `json:"url"`
	JSON gatewayPacfileGetResponseJSON `json:"-"`
}

// gatewayPacfileGetResponseJSON contains the JSON metadata for the struct
// [GatewayPacfileGetResponse]
type gatewayPacfileGetResponseJSON struct {
	ID          apijson.Field
	Contents    apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	UpdatedAt   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayPacfileGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayPacfileGetResponseJSON) RawJSON() string {
	return r.raw
}

type GatewayPacfileNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Actual contents of the PAC file
	Contents param.Field[string] `json:"contents,required"`
	// Name of the PAC file.
	Name param.Field[string] `json:"name,required"`
	// Detailed description of the PAC file.
	Description param.Field[string] `json:"description"`
	// URL-friendly version of the PAC file name. If not provided, it will be
	// auto-generated
	Slug param.Field[string] `json:"slug"`
}

func (r GatewayPacfileNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayPacfileNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayPacfileNewResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayPacfileNewResponse                `json:"result"`
	JSON    gatewayPacfileNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayPacfileNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayPacfileNewResponseEnvelope]
type gatewayPacfileNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayPacfileNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayPacfileNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type GatewayPacfileNewResponseEnvelopeSuccess bool

const (
	GatewayPacfileNewResponseEnvelopeSuccessTrue GatewayPacfileNewResponseEnvelopeSuccess = true
)

func (r GatewayPacfileNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayPacfileNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayPacfileUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Actual contents of the PAC file
	Contents param.Field[string] `json:"contents,required"`
	// Detailed description of the PAC file.
	Description param.Field[string] `json:"description,required"`
	// Name of the PAC file.
	Name param.Field[string] `json:"name,required"`
}

func (r GatewayPacfileUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayPacfileUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayPacfileUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayPacfileUpdateResponse                `json:"result"`
	JSON    gatewayPacfileUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayPacfileUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayPacfileUpdateResponseEnvelope]
type gatewayPacfileUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayPacfileUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayPacfileUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type GatewayPacfileUpdateResponseEnvelopeSuccess bool

const (
	GatewayPacfileUpdateResponseEnvelopeSuccessTrue GatewayPacfileUpdateResponseEnvelopeSuccess = true
)

func (r GatewayPacfileUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayPacfileUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayPacfileListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayPacfileDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayPacfileDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayPacfileDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayPacfileDeleteResponse                `json:"result"`
	JSON    gatewayPacfileDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayPacfileDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayPacfileDeleteResponseEnvelope]
type gatewayPacfileDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayPacfileDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayPacfileDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type GatewayPacfileDeleteResponseEnvelopeSuccess bool

const (
	GatewayPacfileDeleteResponseEnvelopeSuccessTrue GatewayPacfileDeleteResponseEnvelopeSuccess = true
)

func (r GatewayPacfileDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayPacfileDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayPacfileGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayPacfileGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayPacfileGetResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayPacfileGetResponse                `json:"result"`
	JSON    gatewayPacfileGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayPacfileGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayPacfileGetResponseEnvelope]
type gatewayPacfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayPacfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayPacfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type GatewayPacfileGetResponseEnvelopeSuccess bool

const (
	GatewayPacfileGetResponseEnvelopeSuccessTrue GatewayPacfileGetResponseEnvelopeSuccess = true
)

func (r GatewayPacfileGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayPacfileGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
