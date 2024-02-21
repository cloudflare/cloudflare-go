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
	Options             []option.RequestOption
	Annotations         *RadarAnnotationService
	BGP                 *RadarBGPService
	Datasets            *RadarDatasetService
	DNS                 *RadarDNSService
	Netflows            *RadarNetflowService
	Search              *RadarSearchService
	VerifiedBots        *RadarVerifiedBotService
	As112               *RadarAs112Service
	ConnectionTampering *RadarConnectionTamperingService
	Email               *RadarEmailService
	Attacks             *RadarAttackService
	Emails              *RadarEmailService
	Entities            *RadarEntityService
	HTTP                *RadarHTTPService
	Quality             *RadarQualityService
	Ranking             *RadarRankingService
	TrafficAnomalies    *RadarTrafficAnomalyService
}

// NewRadarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRadarService(opts ...option.RequestOption) (r *RadarService) {
	r = &RadarService{}
	r.Options = opts
	r.Annotations = NewRadarAnnotationService(opts...)
	r.BGP = NewRadarBGPService(opts...)
	r.Datasets = NewRadarDatasetService(opts...)
	r.DNS = NewRadarDNSService(opts...)
	r.Netflows = NewRadarNetflowService(opts...)
	r.Search = NewRadarSearchService(opts...)
	r.VerifiedBots = NewRadarVerifiedBotService(opts...)
	r.As112 = NewRadarAs112Service(opts...)
	r.ConnectionTampering = NewRadarConnectionTamperingService(opts...)
	r.Email = NewRadarEmailService(opts...)
	r.Attacks = NewRadarAttackService(opts...)
	r.Emails = NewRadarEmailService(opts...)
	r.Entities = NewRadarEntityService(opts...)
	r.HTTP = NewRadarHTTPService(opts...)
	r.Quality = NewRadarQualityService(opts...)
	r.Ranking = NewRadarRankingService(opts...)
	r.TrafficAnomalies = NewRadarTrafficAnomalyService(opts...)
	return
}
