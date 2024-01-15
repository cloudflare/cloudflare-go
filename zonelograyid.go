// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZoneLogRayidService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneLogRayidService] method
// instead.
type ZoneLogRayidService struct {
	Options []option.RequestOption
}

// NewZoneLogRayidService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneLogRayidService(opts ...option.RequestOption) (r *ZoneLogRayidService) {
	r = &ZoneLogRayidService{}
	r.Options = opts
	return
}

// The `/rayids` api route allows lookups by specific rayid. The rayids route will
// return zero, one, or more records (ray ids are not unique).
func (r *ZoneLogRayidService) Get(ctx context.Context, zoneIdentifier string, rayIdentifier string, query ZoneLogRayidGetParams, opts ...option.RequestOption) (res *ZoneLogRayidGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logs/rayids/%s", zoneIdentifier, rayIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Union satisfied by [shared.UnionString] or [ZoneLogRayidGetResponseUnknown].
type ZoneLogRayidGetResponse interface {
	ImplementsZoneLogRayidGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneLogRayidGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneLogRayidGetParams struct {
	// The `/received` route by default returns a limited set of fields, and allows
	// customers to override the default field set by specifying individual fields. The
	// reasons for this are: 1. Most customers require only a small subset of fields,
	// but that subset varies from customer to customer; 2. Flat schema is much easier
	// to work with downstream (importing into BigTable etc); 3. Performance (time to
	// process, file size). If `?fields=` is not specified, default field set is
	// returned. This default field set may change at any time. When `?fields=` is
	// provided, each record is returned with the specified fields. `fields` must be
	// specified as a comma separated list without any whitespaces, and all fields must
	// exist. The order in which fields are specified does not matter, and the order of
	// fields in the response is not specified.
	Fields param.Field[string] `query:"fields"`
	// By default, timestamps in responses are returned as Unix nanosecond integers.
	// The `?timestamps=` argument can be set to change the format in which response
	// timestamps are returned. Possible values are: `unix`, `unixnano`, `rfc3339`.
	// Note that `unix` and `unixnano` return timestamps as integers; `rfc3339` returns
	// timestamps as strings.
	Timestamps param.Field[ZoneLogRayidGetParamsTimestamps] `query:"timestamps"`
}

// URLQuery serializes [ZoneLogRayidGetParams]'s query parameters as `url.Values`.
func (r ZoneLogRayidGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// By default, timestamps in responses are returned as Unix nanosecond integers.
// The `?timestamps=` argument can be set to change the format in which response
// timestamps are returned. Possible values are: `unix`, `unixnano`, `rfc3339`.
// Note that `unix` and `unixnano` return timestamps as integers; `rfc3339` returns
// timestamps as strings.
type ZoneLogRayidGetParamsTimestamps string

const (
	ZoneLogRayidGetParamsTimestampsUnix     ZoneLogRayidGetParamsTimestamps = "unix"
	ZoneLogRayidGetParamsTimestampsUnixnano ZoneLogRayidGetParamsTimestamps = "unixnano"
	ZoneLogRayidGetParamsTimestampsRfc3339  ZoneLogRayidGetParamsTimestamps = "rfc3339"
)
