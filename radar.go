// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRadarService] method instead.
type RadarService struct {
	Options          []option.RequestOption
	Attacks          *RadarAttackService
	Emails           *RadarEmailService
	Entities         *RadarEntityService
	HTTP             *RadarHTTPService
	Quality          *RadarQualityService
	Ranking          *RadarRankingService
	TrafficAnomalies *RadarTrafficAnomalyService
}

// NewRadarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRadarService(opts ...option.RequestOption) (r *RadarService) {
	r = &RadarService{}
	r.Options = opts
	r.Attacks = NewRadarAttackService(opts...)
	r.Emails = NewRadarEmailService(opts...)
	r.Entities = NewRadarEntityService(opts...)
	r.HTTP = NewRadarHTTPService(opts...)
	r.Quality = NewRadarQualityService(opts...)
	r.Ranking = NewRadarRankingService(opts...)
	r.TrafficAnomalies = NewRadarTrafficAnomalyService(opts...)
	return
}
