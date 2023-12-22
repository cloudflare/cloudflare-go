// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneLogpushDatasetJobService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneLogpushDatasetJobService]
// method instead.
type ZoneLogpushDatasetJobService struct {
	Options []option.RequestOption
}

// NewZoneLogpushDatasetJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneLogpushDatasetJobService(opts ...option.RequestOption) (r *ZoneLogpushDatasetJobService) {
	r = &ZoneLogpushDatasetJobService{}
	r.Options = opts
	return
}

// Lists Logpush jobs for a zone for a dataset.
func (r *ZoneLogpushDatasetJobService) GetZonesZoneIdentifierLogpushDatasetsDatasetJobs(ctx context.Context, zoneIdentifier string, dataset string, opts ...option.RequestOption) (res *LogpushJobResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/datasets/%s/jobs", zoneIdentifier, dataset)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
