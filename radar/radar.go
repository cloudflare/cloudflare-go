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

type UnnamedSchemaRef16e559c45a31db5480e21fbe904b2e42 struct {
	Code string                                               `json:"code,required"`
	Name string                                               `json:"name,required"`
	JSON unnamedSchemaRef16e559c45a31db5480e21fbe904b2e42JSON `json:"-"`
}

// unnamedSchemaRef16e559c45a31db5480e21fbe904b2e42JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef16e559c45a31db5480e21fbe904b2e42]
type unnamedSchemaRef16e559c45a31db5480e21fbe904b2e42JSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef16e559c45a31db5480e21fbe904b2e42) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef16e559c45a31db5480e21fbe904b2e42JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef4124a22436f90127c7fa2c4543219752 struct {
	ClientASN    int64                                                `json:"clientASN,required"`
	ClientAsName string                                               `json:"clientASName,required"`
	Value        string                                               `json:"value,required"`
	JSON         unnamedSchemaRef4124a22436f90127c7fa2c4543219752JSON `json:"-"`
}

// unnamedSchemaRef4124a22436f90127c7fa2c4543219752JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef4124a22436f90127c7fa2c4543219752]
type unnamedSchemaRef4124a22436f90127c7fa2c4543219752JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UnnamedSchemaRef4124a22436f90127c7fa2c4543219752) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef4124a22436f90127c7fa2c4543219752JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef75bae70cf28e6bcef364b9840db3bdeb struct {
	Timestamps []time.Time                                          `json:"timestamps,required" format:"date-time"`
	Values     []string                                             `json:"values,required"`
	JSON       unnamedSchemaRef75bae70cf28e6bcef364b9840db3bdebJSON `json:"-"`
}

// unnamedSchemaRef75bae70cf28e6bcef364b9840db3bdebJSON contains the JSON metadata
// for the struct [UnnamedSchemaRef75bae70cf28e6bcef364b9840db3bdeb]
type unnamedSchemaRef75bae70cf28e6bcef364b9840db3bdebJSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef75bae70cf28e6bcef364b9840db3bdeb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef75bae70cf28e6bcef364b9840db3bdebJSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55 struct {
	Timestamps  []string                                             `json:"timestamps,required"`
	ExtraFields map[string][]string                                  `json:"-,extras"`
	JSON        unnamedSchemaRef7826220e105d84352ba1108d9ed88e55JSON `json:"-"`
}

// unnamedSchemaRef7826220e105d84352ba1108d9ed88e55JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55]
type unnamedSchemaRef7826220e105d84352ba1108d9ed88e55JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef7826220e105d84352ba1108d9ed88e55JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef83a14d589e799bc901b9ccc870251d09 struct {
	ClientCountryAlpha2 string                                               `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                               `json:"clientCountryName,required"`
	Value               string                                               `json:"value,required"`
	JSON                unnamedSchemaRef83a14d589e799bc901b9ccc870251d09JSON `json:"-"`
}

// unnamedSchemaRef83a14d589e799bc901b9ccc870251d09JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef83a14d589e799bc901b9ccc870251d09]
type unnamedSchemaRef83a14d589e799bc901b9ccc870251d09JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *UnnamedSchemaRef83a14d589e799bc901b9ccc870251d09) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef83a14d589e799bc901b9ccc870251d09JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef8b383e904d9fb02521257ef9cc77d297 struct {
	IPv4 string                                               `json:"IPv4,required"`
	IPv6 string                                               `json:"IPv6,required"`
	JSON unnamedSchemaRef8b383e904d9fb02521257ef9cc77d297JSON `json:"-"`
}

// unnamedSchemaRef8b383e904d9fb02521257ef9cc77d297JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef8b383e904d9fb02521257ef9cc77d297]
type unnamedSchemaRef8b383e904d9fb02521257ef9cc77d297JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef8b383e904d9fb02521257ef9cc77d297) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef8b383e904d9fb02521257ef9cc77d297JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef9002274ed7cb7f3dc567421e31529a3a struct {
	IPv4       []string                                             `json:"IPv4,required"`
	IPv6       []string                                             `json:"IPv6,required"`
	Timestamps []string                                             `json:"timestamps,required"`
	JSON       unnamedSchemaRef9002274ed7cb7f3dc567421e31529a3aJSON `json:"-"`
}

// unnamedSchemaRef9002274ed7cb7f3dc567421e31529a3aJSON contains the JSON metadata
// for the struct [UnnamedSchemaRef9002274ed7cb7f3dc567421e31529a3a]
type unnamedSchemaRef9002274ed7cb7f3dc567421e31529a3aJSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef9002274ed7cb7f3dc567421e31529a3a) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef9002274ed7cb7f3dc567421e31529a3aJSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRefB5f3bd1840490bc487ffef84567807b1 struct {
	DataSource      string                                               `json:"dataSource,required"`
	Description     string                                               `json:"description,required"`
	EventType       string                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                               `json:"linkedUrl"`
	StartTime       time.Time                                            `json:"startTime" format:"date-time"`
	JSON            unnamedSchemaRefB5f3bd1840490bc487ffef84567807b1JSON `json:"-"`
}

// unnamedSchemaRefB5f3bd1840490bc487ffef84567807b1JSON contains the JSON metadata
// for the struct [UnnamedSchemaRefB5f3bd1840490bc487ffef84567807b1]
type unnamedSchemaRefB5f3bd1840490bc487ffef84567807b1JSON struct {
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

func (r *UnnamedSchemaRefB5f3bd1840490bc487ffef84567807b1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRefB5f3bd1840490bc487ffef84567807b1JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRefBaac9d7da12de53e99142f8ecd3982e5 struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      unnamedSchemaRefBaac9d7da12de53e99142f8ecd3982e5JSON `json:"-"`
}

// unnamedSchemaRefBaac9d7da12de53e99142f8ecd3982e5JSON contains the JSON metadata
// for the struct [UnnamedSchemaRefBaac9d7da12de53e99142f8ecd3982e5]
type unnamedSchemaRefBaac9d7da12de53e99142f8ecd3982e5JSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRefBaac9d7da12de53e99142f8ecd3982e5) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRefBaac9d7da12de53e99142f8ecd3982e5JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRefC5858f1f916a921846e0b6159af470a7 struct {
	DataTime   string                                               `json:"data_time,required"`
	QueryTime  string                                               `json:"query_time,required"`
	TotalPeers int64                                                `json:"total_peers,required"`
	JSON       unnamedSchemaRefC5858f1f916a921846e0b6159af470a7JSON `json:"-"`
}

// unnamedSchemaRefC5858f1f916a921846e0b6159af470a7JSON contains the JSON metadata
// for the struct [UnnamedSchemaRefC5858f1f916a921846e0b6159af470a7]
type unnamedSchemaRefC5858f1f916a921846e0b6159af470a7JSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRefC5858f1f916a921846e0b6159af470a7) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRefC5858f1f916a921846e0b6159af470a7JSON) RawJSON() string {
	return r.raw
}
