package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type TunnelRoute struct {
	Network    string    `json:"network"`
	TunnelId   string    `json:"tunnel_id"`
	TunnelName string    `json:"tunnel_name"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

// TunnelRouteListResponse is the API response for listing tunnel routes.
type TunnelRouteListResponse struct {
	Response
	Result []TunnelRoute `json:"result"`
}

func (api *API) TunnelRoutes(ctx context.Context) ([]TunnelRoute, error) {
	uri := fmt.Sprintf("/%s/%s/teamnet/routes", AccountRouteRoot, api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return []TunnelRoute{}, err
	}

	var resp TunnelRouteListResponse
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return []TunnelRoute{}, errors.Wrap(err, errUnmarshalError)
	}

	return resp.Result, nil
}
