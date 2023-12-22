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
	Options       []option.RequestOption
	Annotations   *RadarAnnotationService
	As112s        *RadarAs112Service
	Attacks       *RadarAttackService
	Bgps          *RadarBgpService
	Datasets      *RadarDatasetService
	DNS           *RadarDNSService
	Emails        *RadarEmailService
	Entities      *RadarEntityService
	HTTPs         *RadarHTTPService
	Netflows      *RadarNetflowService
	Rankings      *RadarRankingService
	Searches      *RadarSearchService
	Specialevents *RadarSpecialeventService
	VerifiedBots  *RadarVerifiedBotService
}

// NewRadarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRadarService(opts ...option.RequestOption) (r *RadarService) {
	r = &RadarService{}
	r.Options = opts
	r.Annotations = NewRadarAnnotationService(opts...)
	r.As112s = NewRadarAs112Service(opts...)
	r.Attacks = NewRadarAttackService(opts...)
	r.Bgps = NewRadarBgpService(opts...)
	r.Datasets = NewRadarDatasetService(opts...)
	r.DNS = NewRadarDNSService(opts...)
	r.Emails = NewRadarEmailService(opts...)
	r.Entities = NewRadarEntityService(opts...)
	r.HTTPs = NewRadarHTTPService(opts...)
	r.Netflows = NewRadarNetflowService(opts...)
	r.Rankings = NewRadarRankingService(opts...)
	r.Searches = NewRadarSearchService(opts...)
	r.Specialevents = NewRadarSpecialeventService(opts...)
	r.VerifiedBots = NewRadarVerifiedBotService(opts...)
	return
}
