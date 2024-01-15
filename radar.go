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
	Options              []option.RequestOption
	As112                *RadarAs112Service
	Quality              *RadarQualityService
	TrafficAnomalies     *RadarTrafficAnomalyService
	Annotations          *RadarAnnotationService
	Attacks              *RadarAttackService
	Bgps                 *RadarBgpService
	Datasets             *RadarDatasetService
	DNS                  *RadarDNSService
	Email                *RadarEmailService
	Entities             *RadarEntityService
	HTTP                 *RadarHTTPService
	Netflows             *RadarNetflowService
	Ranking              *RadarRankingService
	Searches             *RadarSearchService
	VerifiedBots         *RadarVerifiedBotService
	ConnectionTamperings *RadarConnectionTamperingService
}

// NewRadarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRadarService(opts ...option.RequestOption) (r *RadarService) {
	r = &RadarService{}
	r.Options = opts
	r.As112 = NewRadarAs112Service(opts...)
	r.Quality = NewRadarQualityService(opts...)
	r.TrafficAnomalies = NewRadarTrafficAnomalyService(opts...)
	r.Annotations = NewRadarAnnotationService(opts...)
	r.Attacks = NewRadarAttackService(opts...)
	r.Bgps = NewRadarBgpService(opts...)
	r.Datasets = NewRadarDatasetService(opts...)
	r.DNS = NewRadarDNSService(opts...)
	r.Email = NewRadarEmailService(opts...)
	r.Entities = NewRadarEntityService(opts...)
	r.HTTP = NewRadarHTTPService(opts...)
	r.Netflows = NewRadarNetflowService(opts...)
	r.Ranking = NewRadarRankingService(opts...)
	r.Searches = NewRadarSearchService(opts...)
	r.VerifiedBots = NewRadarVerifiedBotService(opts...)
	r.ConnectionTamperings = NewRadarConnectionTamperingService(opts...)
	return
}
