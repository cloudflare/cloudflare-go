// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// DomainService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainService] method instead.
type DomainService struct {
	Options []option.RequestOption
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r *DomainService) {
	r = &DomainService{}
	r.Options = opts
	return
}

// Update individual domain.
func (r *DomainService) Update(ctx context.Context, domainName string, params DomainUpdateParams, opts ...option.RequestOption) (res *DomainUpdateResponseUnion, err error) {
	var env DomainUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/registrar/domains/%s", params.AccountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List domains handled by Registrar.
func (r *DomainService) List(ctx context.Context, query DomainListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DomainListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/registrar/domains", query.AccountID)
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

// List domains handled by Registrar.
func (r *DomainService) ListAutoPaging(ctx context.Context, query DomainListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DomainListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Show individual domain.
func (r *DomainService) Get(ctx context.Context, domainName string, query DomainGetParams, opts ...option.RequestOption) (res *DomainGetResponseUnion, err error) {
	var env DomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/registrar/domains/%s", query.AccountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [registrar.DomainUpdateResponseUnknown],
// [registrar.DomainUpdateResponseArray] or [shared.UnionString].
type DomainUpdateResponseUnion interface {
	ImplementsRegistrarDomainUpdateResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DomainUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DomainUpdateResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DomainUpdateResponseArray []interface{}

func (r DomainUpdateResponseArray) ImplementsRegistrarDomainUpdateResponseUnion() {}

type DomainListResponse struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   DomainListResponseResultUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DomainListResponseSuccess    `json:"success,required"`
	ResultInfo DomainListResponseResultInfo `json:"result_info"`
	JSON       domainListResponseJSON       `json:"-"`
}

// domainListResponseJSON contains the JSON metadata for the struct
// [DomainListResponse]
type domainListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainListResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [registrar.DomainListResponseResultUnknown],
// [registrar.DomainListResponseResultArray] or [shared.UnionString].
type DomainListResponseResultUnion interface {
	ImplementsRegistrarDomainListResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DomainListResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DomainListResponseResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DomainListResponseResultArray []interface{}

func (r DomainListResponseResultArray) ImplementsRegistrarDomainListResponseResultUnion() {}

// Whether the API call was successful
type DomainListResponseSuccess bool

const (
	DomainListResponseSuccessTrue DomainListResponseSuccess = true
)

func (r DomainListResponseSuccess) IsKnown() bool {
	switch r {
	case DomainListResponseSuccessTrue:
		return true
	}
	return false
}

type DomainListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                          `json:"total_count"`
	JSON       domainListResponseResultInfoJSON `json:"-"`
}

// domainListResponseResultInfoJSON contains the JSON metadata for the struct
// [DomainListResponseResultInfo]
type domainListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [registrar.DomainGetResponseUnknown],
// [registrar.DomainGetResponseArray] or [shared.UnionString].
type DomainGetResponseUnion interface {
	ImplementsRegistrarDomainGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DomainGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DomainGetResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DomainGetResponseArray []interface{}

func (r DomainGetResponseArray) ImplementsRegistrarDomainGetResponseUnion() {}

type DomainUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Auto-renew controls whether subscription is automatically renewed upon domain
	// expiration.
	AutoRenew param.Field[bool] `json:"auto_renew"`
	// Shows whether a registrar lock is in place for a domain.
	Locked param.Field[bool] `json:"locked"`
	// Privacy option controls redacting WHOIS information.
	Privacy param.Field[bool] `json:"privacy"`
}

func (r DomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DomainUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   DomainUpdateResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success DomainUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    domainUpdateResponseEnvelopeJSON    `json:"-"`
}

// domainUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainUpdateResponseEnvelope]
type domainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainUpdateResponseEnvelopeSuccess bool

const (
	DomainUpdateResponseEnvelopeSuccessTrue DomainUpdateResponseEnvelopeSuccess = true
)

func (r DomainUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DomainListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DomainGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DomainGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   DomainGetResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success DomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    domainGetResponseEnvelopeJSON    `json:"-"`
}

// domainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelope]
type domainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainGetResponseEnvelopeSuccess bool

const (
	DomainGetResponseEnvelopeSuccessTrue DomainGetResponseEnvelopeSuccess = true
)

func (r DomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
