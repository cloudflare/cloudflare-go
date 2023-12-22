// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneRailgunService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneRailgunService] method
// instead.
type ZoneRailgunService struct {
	Options   []option.RequestOption
	Diagnoses *ZoneRailgunDiagnosisService
}

// NewZoneRailgunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneRailgunService(opts ...option.RequestOption) (r *ZoneRailgunService) {
	r = &ZoneRailgunService{}
	r.Options = opts
	r.Diagnoses = NewZoneRailgunDiagnosisService(opts...)
	return
}

// Lists details about a specific Railgun.
func (r *ZoneRailgunService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *RailgunResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/railguns/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Connect or disconnect a Railgun.
func (r *ZoneRailgunService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneRailgunUpdateParams, opts ...option.RequestOption) (res *RailgunResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/railguns/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// A list of available Railguns the zone can use.
func (r *ZoneRailgunService) RailgunConnectionsForAZoneListAvailableRailguns(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *SchemasRailgunResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/railguns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SchemasRailgunResponseCollection struct {
	Errors     []SchemasRailgunResponseCollectionError    `json:"errors"`
	Messages   []SchemasRailgunResponseCollectionMessage  `json:"messages"`
	Result     []interface{}                              `json:"result"`
	ResultInfo SchemasRailgunResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SchemasRailgunResponseCollectionSuccess `json:"success"`
	JSON    schemasRailgunResponseCollectionJSON    `json:"-"`
}

// schemasRailgunResponseCollectionJSON contains the JSON metadata for the struct
// [SchemasRailgunResponseCollection]
type schemasRailgunResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasRailgunResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasRailgunResponseCollectionError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    schemasRailgunResponseCollectionErrorJSON `json:"-"`
}

// schemasRailgunResponseCollectionErrorJSON contains the JSON metadata for the
// struct [SchemasRailgunResponseCollectionError]
type schemasRailgunResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasRailgunResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasRailgunResponseCollectionMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    schemasRailgunResponseCollectionMessageJSON `json:"-"`
}

// schemasRailgunResponseCollectionMessageJSON contains the JSON metadata for the
// struct [SchemasRailgunResponseCollectionMessage]
type schemasRailgunResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasRailgunResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasRailgunResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       schemasRailgunResponseCollectionResultInfoJSON `json:"-"`
}

// schemasRailgunResponseCollectionResultInfoJSON contains the JSON metadata for
// the struct [SchemasRailgunResponseCollectionResultInfo]
type schemasRailgunResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasRailgunResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasRailgunResponseCollectionSuccess bool

const (
	SchemasRailgunResponseCollectionSuccessTrue SchemasRailgunResponseCollectionSuccess = true
)

type ZoneRailgunUpdateParams struct {
	// A flag indicating whether the given zone is connected to the Railgun.
	Connected param.Field[bool] `json:"connected,required"`
}

func (r ZoneRailgunUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
