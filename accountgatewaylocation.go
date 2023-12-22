// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountGatewayLocationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountGatewayLocationService]
// method instead.
type AccountGatewayLocationService struct {
	Options []option.RequestOption
}

// NewAccountGatewayLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountGatewayLocationService(opts ...option.RequestOption) (r *AccountGatewayLocationService) {
	r = &AccountGatewayLocationService{}
	r.Options = opts
	return
}

// Fetch a single Zero Trust Gateway Location.
func (r *AccountGatewayLocationService) Get(ctx context.Context, identifier interface{}, uuid interface{}, opts ...option.RequestOption) (res *SchemasSingleResponseC3wrwEik, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a configured Zero Trust Gateway Location.
func (r *AccountGatewayLocationService) Update(ctx context.Context, identifier interface{}, uuid interface{}, body AccountGatewayLocationUpdateParams, opts ...option.RequestOption) (res *SchemasSingleResponseC3wrwEik, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a Zero Trust Gateway Location.
func (r *AccountGatewayLocationService) Delete(ctx context.Context, identifier interface{}, uuid interface{}, opts ...option.RequestOption) (res *EmptyResponseArJnNcMr, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new Zero Trust Gateway Location.
func (r *AccountGatewayLocationService) ZeroTrustGatewayLocationsNewZeroTrustGatewayLocation(ctx context.Context, identifier interface{}, body AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParams, opts ...option.RequestOption) (res *SchemasSingleResponseC3wrwEik, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/locations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Zero Trust Gateway Locations for an account.
func (r *AccountGatewayLocationService) ZeroTrustGatewayLocationsListZeroTrustGatewayLocations(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *SchemasResponseCollectionAPv5oOc1, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/locations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SchemasResponseCollectionAPv5oOc1 struct {
	Errors     []SchemasResponseCollectionAPv5oOc1Error    `json:"errors"`
	Messages   []SchemasResponseCollectionAPv5oOc1Message  `json:"messages"`
	Result     []SchemasResponseCollectionAPv5oOc1Result   `json:"result"`
	ResultInfo SchemasResponseCollectionAPv5oOc1ResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SchemasResponseCollectionAPv5oOc1Success `json:"success"`
	JSON    schemasResponseCollectionAPv5oOc1JSON    `json:"-"`
}

// schemasResponseCollectionAPv5oOc1JSON contains the JSON metadata for the struct
// [SchemasResponseCollectionAPv5oOc1]
type schemasResponseCollectionAPv5oOc1JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionAPv5oOc1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionAPv5oOc1Error struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    schemasResponseCollectionAPv5oOc1ErrorJSON `json:"-"`
}

// schemasResponseCollectionAPv5oOc1ErrorJSON contains the JSON metadata for the
// struct [SchemasResponseCollectionAPv5oOc1Error]
type schemasResponseCollectionAPv5oOc1ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionAPv5oOc1Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionAPv5oOc1Message struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    schemasResponseCollectionAPv5oOc1MessageJSON `json:"-"`
}

// schemasResponseCollectionAPv5oOc1MessageJSON contains the JSON metadata for the
// struct [SchemasResponseCollectionAPv5oOc1Message]
type schemasResponseCollectionAPv5oOc1MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionAPv5oOc1Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionAPv5oOc1Result struct {
	ID interface{} `json:"id"`
	// Set if the location is the default one.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS Over HTTPS domain to send DNS requests to. (auto-generated).
	DohSubdomain string `json:"doh_subdomain"`
	// Set if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. (auto-generated).
	IP string `json:"ip"`
	// The name of the Location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []string                                    `json:"networks"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	JSON      schemasResponseCollectionAPv5oOc1ResultJSON `json:"-"`
}

// schemasResponseCollectionAPv5oOc1ResultJSON contains the JSON metadata for the
// struct [SchemasResponseCollectionAPv5oOc1Result]
type schemasResponseCollectionAPv5oOc1ResultJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SchemasResponseCollectionAPv5oOc1Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionAPv5oOc1ResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       schemasResponseCollectionAPv5oOc1ResultInfoJSON `json:"-"`
}

// schemasResponseCollectionAPv5oOc1ResultInfoJSON contains the JSON metadata for
// the struct [SchemasResponseCollectionAPv5oOc1ResultInfo]
type schemasResponseCollectionAPv5oOc1ResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionAPv5oOc1ResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasResponseCollectionAPv5oOc1Success bool

const (
	SchemasResponseCollectionAPv5oOc1SuccessTrue SchemasResponseCollectionAPv5oOc1Success = true
)

type SchemasSingleResponseC3wrwEik struct {
	Errors   []SchemasSingleResponseC3wrwEikError   `json:"errors"`
	Messages []SchemasSingleResponseC3wrwEikMessage `json:"messages"`
	Result   SchemasSingleResponseC3wrwEikResult    `json:"result"`
	// Whether the API call was successful
	Success SchemasSingleResponseC3wrwEikSuccess `json:"success"`
	JSON    schemasSingleResponseC3wrwEikJSON    `json:"-"`
}

// schemasSingleResponseC3wrwEikJSON contains the JSON metadata for the struct
// [SchemasSingleResponseC3wrwEik]
type schemasSingleResponseC3wrwEikJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseC3wrwEik) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseC3wrwEikError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    schemasSingleResponseC3wrwEikErrorJSON `json:"-"`
}

// schemasSingleResponseC3wrwEikErrorJSON contains the JSON metadata for the struct
// [SchemasSingleResponseC3wrwEikError]
type schemasSingleResponseC3wrwEikErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseC3wrwEikError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseC3wrwEikMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    schemasSingleResponseC3wrwEikMessageJSON `json:"-"`
}

// schemasSingleResponseC3wrwEikMessageJSON contains the JSON metadata for the
// struct [SchemasSingleResponseC3wrwEikMessage]
type schemasSingleResponseC3wrwEikMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseC3wrwEikMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseC3wrwEikResult struct {
	ID interface{} `json:"id"`
	// Set if the location is the default one.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS Over HTTPS domain to send DNS requests to. (auto-generated).
	DohSubdomain string `json:"doh_subdomain"`
	// Set if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. (auto-generated).
	IP string `json:"ip"`
	// The name of the Location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []string                                `json:"networks"`
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	JSON      schemasSingleResponseC3wrwEikResultJSON `json:"-"`
}

// schemasSingleResponseC3wrwEikResultJSON contains the JSON metadata for the
// struct [SchemasSingleResponseC3wrwEikResult]
type schemasSingleResponseC3wrwEikResultJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SchemasSingleResponseC3wrwEikResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasSingleResponseC3wrwEikSuccess bool

const (
	SchemasSingleResponseC3wrwEikSuccessTrue SchemasSingleResponseC3wrwEikSuccess = true
)

type AccountGatewayLocationUpdateParams struct {
	// The name of the Location.
	Name param.Field[string] `json:"name,required"`
	// Set if the location is the default one.
	ClientDefault param.Field[bool] `json:"client_default"`
	// Set if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// A list of network ranges that requests from this location would originate from.
	Networks param.Field[[]string] `json:"networks"`
}

func (r AccountGatewayLocationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParams struct {
	// The name of the Location.
	Name param.Field[string] `json:"name,required"`
	// Set if the location is the default one.
	ClientDefault param.Field[bool] `json:"client_default"`
	// Set if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// A list of network ranges that requests from this location would originate from.
	Networks param.Field[[]string] `json:"networks"`
}

func (r AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
