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

// IntelASNSubnetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIntelASNSubnetService] method
// instead.
type IntelASNSubnetService struct {
	Options []option.RequestOption
}

// NewIntelASNSubnetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntelASNSubnetService(opts ...option.RequestOption) (r *IntelASNSubnetService) {
	r = &IntelASNSubnetService{}
	r.Options = opts
	return
}

// Get ASN Subnets
func (r *IntelASNSubnetService) List(ctx context.Context, accountID string, asn int64, opts ...option.RequestOption) (res *IntelASNSubnetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/asn/%v/subnets", accountID, asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type IntelASNSubnetListResponse struct {
	ASN int64 `json:"asn"`
	// Total results returned based on your search parameters.
	Count        float64 `json:"count"`
	IPCountTotal int64   `json:"ip_count_total"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64                        `json:"per_page"`
	Subnets []string                       `json:"subnets"`
	JSON    intelASNSubnetListResponseJSON `json:"-"`
}

// intelASNSubnetListResponseJSON contains the JSON metadata for the struct
// [IntelASNSubnetListResponse]
type intelASNSubnetListResponseJSON struct {
	ASN          apijson.Field
	Count        apijson.Field
	IPCountTotal apijson.Field
	Page         apijson.Field
	PerPage      apijson.Field
	Subnets      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IntelASNSubnetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
