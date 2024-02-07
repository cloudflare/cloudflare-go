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

// PageShieldConnectionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPageShieldConnectionService]
// method instead.
type PageShieldConnectionService struct {
	Options []option.RequestOption
}

// NewPageShieldConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageShieldConnectionService(opts ...option.RequestOption) (r *PageShieldConnectionService) {
	r = &PageShieldConnectionService{}
	r.Options = opts
	return
}

// Fetches a connection detected by Page Shield by connection ID.
func (r *PageShieldConnectionService) Get(ctx context.Context, zoneID string, connectionID string, opts ...option.RequestOption) (res *PageShieldConnectionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/connections/%s", zoneID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PageShieldConnectionGetResponse struct {
	ID                      interface{}                         `json:"id"`
	AddedAt                 interface{}                         `json:"added_at"`
	DomainReportedMalicious interface{}                         `json:"domain_reported_malicious"`
	FirstPageURL            interface{}                         `json:"first_page_url"`
	FirstSeenAt             interface{}                         `json:"first_seen_at"`
	Host                    interface{}                         `json:"host"`
	LastSeenAt              interface{}                         `json:"last_seen_at"`
	PageURLs                interface{}                         `json:"page_urls"`
	URL                     interface{}                         `json:"url"`
	URLContainsCdnCgiPath   interface{}                         `json:"url_contains_cdn_cgi_path"`
	JSON                    pageShieldConnectionGetResponseJSON `json:"-"`
}

// pageShieldConnectionGetResponseJSON contains the JSON metadata for the struct
// [PageShieldConnectionGetResponse]
type pageShieldConnectionGetResponseJSON struct {
	ID                      apijson.Field
	AddedAt                 apijson.Field
	DomainReportedMalicious apijson.Field
	FirstPageURL            apijson.Field
	FirstSeenAt             apijson.Field
	Host                    apijson.Field
	LastSeenAt              apijson.Field
	PageURLs                apijson.Field
	URL                     apijson.Field
	URLContainsCdnCgiPath   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PageShieldConnectionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
