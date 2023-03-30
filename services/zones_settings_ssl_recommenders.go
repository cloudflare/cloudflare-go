package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsSslRecommenderService struct {
	Options []options.RequestOption
}

func NewZonesSettingsSslRecommenderService(opts ...options.RequestOption) (r *ZonesSettingsSslRecommenderService) {
	r = &ZonesSettingsSslRecommenderService{}
	r.Options = opts
	return
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
func (r *ZonesSettingsSslRecommenderService) Update(ctx context.Context, zone_identifier string, body *requests.SslRecommenderUpdateParams, opts ...options.RequestOption) (res *responses.SslRecommenderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ssl_recommender", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
func (r *ZonesSettingsSslRecommenderService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.SslRecommenderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ssl_recommender", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
