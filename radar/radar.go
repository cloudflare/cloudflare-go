// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RadarService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRadarService] method instead.
type RadarService struct {
	Options             []option.RequestOption
	Annotations         *AnnotationService
	BGP                 *BGPService
	Datasets            *DatasetService
	DNS                 *DNSService
	Netflows            *NetflowService
	Search              *SearchService
	VerifiedBots        *VerifiedBotService
	AS112               *AS112Service
	ConnectionTampering *ConnectionTamperingService
	Email               *EmailService
	Attacks             *AttackService
	Entities            *EntityService
	HTTP                *HTTPService
	Quality             *QualityService
	Ranking             *RankingService
	TrafficAnomalies    *TrafficAnomalyService
}

// NewRadarService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRadarService(opts ...option.RequestOption) (r *RadarService) {
	r = &RadarService{}
	r.Options = opts
	r.Annotations = NewAnnotationService(opts...)
	r.BGP = NewBGPService(opts...)
	r.Datasets = NewDatasetService(opts...)
	r.DNS = NewDNSService(opts...)
	r.Netflows = NewNetflowService(opts...)
	r.Search = NewSearchService(opts...)
	r.VerifiedBots = NewVerifiedBotService(opts...)
	r.AS112 = NewAS112Service(opts...)
	r.ConnectionTampering = NewConnectionTamperingService(opts...)
	r.Email = NewEmailService(opts...)
	r.Attacks = NewAttackService(opts...)
	r.Entities = NewEntityService(opts...)
	r.HTTP = NewHTTPService(opts...)
	r.Quality = NewQualityService(opts...)
	r.Ranking = NewRankingService(opts...)
	r.TrafficAnomalies = NewTrafficAnomalyService(opts...)
	return
}

type UnnamedSchemaRef106 struct {
	Code string                  `json:"code,required"`
	Name string                  `json:"name,required"`
	JSON unnamedSchemaRef106JSON `json:"-"`
}

// unnamedSchemaRef106JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef106]
type unnamedSchemaRef106JSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef106) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef106JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef128 struct {
	IPv4 string                  `json:"IPv4,required"`
	IPv6 string                  `json:"IPv6,required"`
	JSON unnamedSchemaRef128JSON `json:"-"`
}

// unnamedSchemaRef128JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef128]
type unnamedSchemaRef128JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef128) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef128JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef129 struct {
	Timestamps []time.Time             `json:"timestamps,required" format:"date-time"`
	Values     []string                `json:"values,required"`
	JSON       unnamedSchemaRef129JSON `json:"-"`
}

// unnamedSchemaRef129JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef129]
type unnamedSchemaRef129JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef129) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef129JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef130 struct {
	DataTime   string                  `json:"data_time,required"`
	QueryTime  string                  `json:"query_time,required"`
	TotalPeers int64                   `json:"total_peers,required"`
	JSON       unnamedSchemaRef130JSON `json:"-"`
}

// unnamedSchemaRef130JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef130]
type unnamedSchemaRef130JSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef130) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef130JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef153 struct {
	Timestamps  []string                `json:"timestamps,required"`
	ExtraFields map[string][]string     `json:"-,extras"`
	JSON        unnamedSchemaRef153JSON `json:"-"`
}

// unnamedSchemaRef153JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef153]
type unnamedSchemaRef153JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef153) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef153JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef154 struct {
	ClientASN    int64                   `json:"clientASN,required"`
	ClientAsName string                  `json:"clientASName,required"`
	Value        string                  `json:"value,required"`
	JSON         unnamedSchemaRef154JSON `json:"-"`
}

// unnamedSchemaRef154JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef154]
type unnamedSchemaRef154JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UnnamedSchemaRef154) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef154JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef160 struct {
	ClientCountryAlpha2 string                  `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                  `json:"clientCountryName,required"`
	Value               string                  `json:"value,required"`
	JSON                unnamedSchemaRef160JSON `json:"-"`
}

// unnamedSchemaRef160JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef160]
type unnamedSchemaRef160JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *UnnamedSchemaRef160) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef160JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef174 struct {
	DataSource      string                  `json:"dataSource,required"`
	Description     string                  `json:"description,required"`
	EventType       string                  `json:"eventType,required"`
	IsInstantaneous interface{}             `json:"isInstantaneous,required"`
	EndTime         time.Time               `json:"endTime" format:"date-time"`
	LinkedURL       string                  `json:"linkedUrl"`
	StartTime       time.Time               `json:"startTime" format:"date-time"`
	JSON            unnamedSchemaRef174JSON `json:"-"`
}

// unnamedSchemaRef174JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef174]
type unnamedSchemaRef174JSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *UnnamedSchemaRef174) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef174JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef175 struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time               `json:"startTime,required" format:"date-time"`
	JSON      unnamedSchemaRef175JSON `json:"-"`
}

// unnamedSchemaRef175JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef175]
type unnamedSchemaRef175JSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef175) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef175JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef53 struct {
	IPv4       []string               `json:"IPv4,required"`
	IPv6       []string               `json:"IPv6,required"`
	Timestamps []string               `json:"timestamps,required"`
	JSON       unnamedSchemaRef53JSON `json:"-"`
}

// unnamedSchemaRef53JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef53]
type unnamedSchemaRef53JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef53) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef53JSON) RawJSON() string {
	return r.raw
}
