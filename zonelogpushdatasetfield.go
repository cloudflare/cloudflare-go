// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneLogpushDatasetFieldService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneLogpushDatasetFieldService] method instead.
type ZoneLogpushDatasetFieldService struct {
	Options []option.RequestOption
}

// NewZoneLogpushDatasetFieldService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneLogpushDatasetFieldService(opts ...option.RequestOption) (r *ZoneLogpushDatasetFieldService) {
	r = &ZoneLogpushDatasetFieldService{}
	r.Options = opts
	return
}

// Lists all fields available for a dataset. The response result is an object with
// key-value pairs, where keys are field names, and values are descriptions.
func (r *ZoneLogpushDatasetFieldService) GetZonesZoneIdentifierLogpushDatasetsDatasetFields(ctx context.Context, zoneIdentifier string, dataset string, opts ...option.RequestOption) (res *LogpushFieldResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/datasets/%s/fields", zoneIdentifier, dataset)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
