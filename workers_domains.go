package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type WorkerDomainParams struct {
	ZoneID      string `json:"zone_id"`
	Hostname    string `json:"hostname"`
	Service     string `json:"service"`
	Environment string `json:"environment,omitempty"`
}

type WorkerDomainResult struct {
	ID          string `json:"id"`
	ZoneID      string `json:"zone_id"`
	ZoneName    string `json:"zone_name"`
	Hostname    string `json:"hostname"`
	Service     string `json:"service"`
	Environment string `json:"environment"`
}

type WorkerDomainResponse struct {
	Response
	WorkerDomainResult `json:"result"`
}

// AttachWorkerToDomain attaches a worker to a zone and hostname
//
// API reference: https://api.cloudflare.com/#worker-domain-attach-to-domain
func (api *API) AttachWorkerToDomain(ctx context.Context, rc *ResourceContainer, params *WorkerDomainParams) (WorkerDomainResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/domains", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return WorkerDomainResponse{}, err
	}

	var r WorkerDomainResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerDomainResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r, nil
}
