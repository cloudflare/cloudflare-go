package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// LogpushJob describes a Logpush job.
type LogpushJob struct {
	ID                 int        `json:"id,omitempty"`
	Dataset            string     `json:"dataset"`
	Enabled            bool       `json:"enabled"`
	Name               string     `json:"name"`
	LogpullOptions     string     `json:"logpull_options"`
	DestinationConf    string     `json:"destination_conf"`
	OwnershipChallenge string     `json:"ownership_challenge,omitempty"`
	LastComplete       *time.Time `json:"last_complete,omitempty"`
	LastError          *time.Time `json:"last_error,omitempty"`
	ErrorMessage       string     `json:"error_message,omitempty"`
}

// LogpushJobsResponse is the API response, containing an array of Logpush Jobs.
type LogpushJobsResponse struct {
	Response
	Result []LogpushJob `json:"result"`
}

// LogpushJobDetailsResponse is the API response, containing a single Logpush Job.
type LogpushJobDetailsResponse struct {
	Response
	Result LogpushJob `json:"result"`
}

// LogpushFieldsResponse is the API response for a datasets fields
type LogpushFieldsResponse struct {
	Response
	Result LogpushFields `json:"result"`
}

// LogpushFields is a map of available Logpush field names & descriptions
type LogpushFields map[string]string

// LogpushGetOwnershipChallenge describes a ownership validation.
type LogpushGetOwnershipChallenge struct {
	Filename string `json:"filename"`
	Valid    bool   `json:"valid"`
	Message  string `json:"message"`
}

// LogpushGetOwnershipChallengeResponse is the API response, containing a ownership challenge.
type LogpushGetOwnershipChallengeResponse struct {
	Response
	Result LogpushGetOwnershipChallenge `json:"result"`
}

// LogpushGetOwnershipChallengeRequest is the API request for get ownership challenge.
type LogpushGetOwnershipChallengeRequest struct {
	DestinationConf string `json:"destination_conf"`
}

// LogpushOwnershipChallengeValidationResponse is the API response,
// containing a ownership challenge validation result.
type LogpushOwnershipChallengeValidationResponse struct {
	Response
	Result struct {
		Valid bool `json:"valid"`
	}
}

// LogpushValidateOwnershipChallengeRequest is the API request for validate ownership challenge.
type LogpushValidateOwnershipChallengeRequest struct {
	DestinationConf    string `json:"destination_conf"`
	OwnershipChallenge string `json:"ownership_challenge"`
}

// LogpushDestinationExistsResponse is the API response,
// containing a destination exists check result.
type LogpushDestinationExistsResponse struct {
	Response
	Result struct {
		Exists bool `json:"exists"`
	}
}

// LogpushDestinationExistsRequest is the API request for check destination exists.
type LogpushDestinationExistsRequest struct {
	DestinationConf string `json:"destination_conf"`
}

func getRouteRoot (identifierType string) {
	if identifierType == "account" {
		return AccountRouteRoot
	} else {
		return ZoneRouteRoot
	}
}

// CreateLogpushJob creates a new LogpushJob for a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-create-logpush-job
func (api *API) CreateLogpushJob(ctx context.Context, identifier string, job LogpushJob, identifierType string) (*LogpushJob, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/jobs", getRouteRoot(identifierType), identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, job)
	if err != nil {
		return nil, err
	}
	var r LogpushJobDetailsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return &r.Result, nil
}
func (api *API) CreateAccountLogpushJob(ctx context.Context, accountID string, job LogpushJob) (*LogpushJob, error) {
	return api.CreateLogpushJob(ctx, accountID, job, "account")
}
func (api *API) CreateZoneLogpushJob(ctx context.Context, zoneID string, job LogpushJob) (*LogpushJob, error) {
	return api.CreateLogpushJob(ctx, zoneID, job, "zone")
}

// LogpushJobs returns all Logpush Jobs for a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-list-logpush-jobs
func (api *API) LogpushJobs(ctx context.Context, identifier string, identifierType string) ([]LogpushJob, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/jobs", getRouteRoot(identifierType), identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []LogpushJob{}, err
	}
	var r LogpushJobsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []LogpushJob{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
func (api *API) ListAccountLogpushJobs(ctx context.Context, accountID string) ([]LogpushJob, error) {
	return api.LogpushJobs(ctx, accountID, job, "account")
}
func (api *API) ListZoneLogpushJobs(ctx context.Context, zoneID string) ([]LogpushJob, error) {
	return api.LogpushJobs(ctx, zoneID, job, "zone")
}

// LogpushJobsForDataset returns all Logpush Jobs for a dataset in a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-list-logpush-jobs-for-a-dataset
func (api *API) LogpushJobsForDataset(ctx context.Context, identifier, dataset string, identifierType string) ([]LogpushJob, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/datasets/%s/jobs", getRouteRoot(identifierType), identifier, dataset)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []LogpushJob{}, err
	}
	var r LogpushJobsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []LogpushJob{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
func (api *API) ListAccountLogpushJobsForDataset(ctx context.Context, accountID, dataset string) ([]LogpushJob, error) {
	return api.LogpushJobsForDataset(ctx, accountID, dataset, "account")
}
func (api *API) ListZoneLogpushJobsForDataset(ctx context.Context, zoneID, dataset string) ([]LogpushJob, error) {
	return api.LogpushJobsForDataset(ctx, zoneID, dataset, "zone")
}

// LogpushFields returns fields for a given dataset.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-list-logpush-jobs
func (api *API) LogpushFields(ctx context.Context, identifier, dataset string, identifierType string) (LogpushFields, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/datasets/%s/fields", getRouteRoot(identifierType), identifier, dataset)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LogpushFields{}, err
	}
	var r LogpushFieldsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return LogpushFields{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
func (api *API) GetAccountLogpushFields(ctx context.Context, accountID, dataset string) (LogpushFields, error) {
	return api.LogpushFields(ctx, accountID, dataset, "account")
}
func (api *API) GetZoneLogpushFields(ctx context.Context, zoneID, dataset string) (LogpushFields, error) {
	return api.LogpushFields(ctx, zoneID, dataset, "zone")
}

// LogpushJob fetches detail about one Logpush Job for a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-logpush-job-details
func (api *API) LogpushJob(ctx context.Context, identifier string, jobID int, identifierType string) (LogpushJob, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/jobs/%d", getRouteRoot(identifierType), identifier, jobID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LogpushJob{}, err
	}
	var r LogpushJobDetailsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return LogpushJob{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
func (api *API) GetAccountLogpushJob(ctx context.Context, accountID string, jobID int) (LogpushJob, error) {
	return api.LogpushJob(ctx, accountID, jobID, "account")
}
func (api *API) GetZoneLogpushJob(ctx context.Context, zoneID string, jobID int) (LogpushJob, error) {
	return api.LogpushJob(ctx, zoneID, jobID, "zone")
}

// UpdateLogpushJob lets you update a Logpush Job.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-update-logpush-job
func (api *API) UpdateLogpushJob(ctx context.Context, identifier string, jobID int, job LogpushJob, identifierType string) error {
	uri := fmt.Sprintf("/%s/%s/logpush/jobs/%d", getRouteRoot(identifierType), identifier, jobID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, job)
	if err != nil {
		return err
	}
	var r LogpushJobDetailsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}
	return nil
}
func (api *API) UpdateAccountLogpushJob(ctx context.Context, accountID string, jobID int, job LogpushJob) error {
	return api.UpdateLogpushJob(ctx, accountID, jobID, job, "account")
}
func (api *API) UpdateZoneLogpushJob(ctx context.Context, zoneID string, jobID int, job LogpushJob) error {
	return api.UpdateLogpushJob(ctx, zoneID, jobID, job, "zone")
}

// DeleteLogpushJob deletes a Logpush Job for a zone.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-delete-logpush-job
func (api *API) DeleteLogpushJob(ctx context.Context, identifier string, jobID int, identifierType string) error {
	uri := fmt.Sprintf("/%s/%s/logpush/jobs/%d", getRouteRoot(identifierType), identifier, jobID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	var r LogpushJobDetailsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}
	return nil
}
func (api *API) DeleteAccountLogpushJob(ctx context.Context, accountID string, jobID int) error {
	return api.DeleteLogpushJob(ctx, accountID, jobID, "account")
}
func (api *API) DeleteZoneLogpushJob(ctx context.Context, zoneID string, jobID int) error {
	return api.DeleteLogpushJob(ctx, zoneID, jobID, "zone")
}

// GetLogpushOwnershipChallenge returns ownership challenge.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-get-ownership-challenge
func (api *API) GetLogpushOwnershipChallenge(ctx context.Context, identifier, destinationConf string, identifierType string) (*LogpushGetOwnershipChallenge, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/ownership", getRouteRoot(identifierType), identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, LogpushGetOwnershipChallengeRequest{
		DestinationConf: destinationConf,
	})
	if err != nil {
		return nil, err
	}
	var r LogpushGetOwnershipChallengeResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	if !r.Result.Valid {
		return nil, errors.New(r.Result.Message)
	}

	return &r.Result, nil
}
func (api *API) GetAccountLogpushOwnershipChallenge(ctx context.Context, accountID, destinationConf string) (*LogpushGetOwnershipChallenge, error) {
	return api.GetLogpushOwnershipChallenge(ctx, accountID, destinationConf, "account")
}
func (api *API) GetZoneLogpushOwnershipChallenge(ctx context.Context, zoneID, destinationConf string) (*LogpushGetOwnershipChallenge, error) {
	return api.GetLogpushOwnershipChallenge(ctx, zoneID, destinationConf, "zone")
}

// ValidateLogpushOwnershipChallenge returns ownership challenge validation result.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-validate-ownership-challenge
func (api *API) ValidateLogpushOwnershipChallenge(ctx context.Context, identifier, destinationConf, ownershipChallenge string, identifierType string) (bool, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/ownership/validate", getRouteRoot(identifierType), identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, LogpushValidateOwnershipChallengeRequest{
		DestinationConf:    destinationConf,
		OwnershipChallenge: ownershipChallenge,
	})
	if err != nil {
		return false, err
	}
	var r LogpushGetOwnershipChallengeResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return false, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result.Valid, nil
}
func (api *API) ValidateAccountLogpushOwnershipChallenge(ctx context.Context, accountID, destinationConf, ownershipChallenge string) (bool, error) {
	return api.ValidateLogpushOwnershipChallenge(ctx, accountID, destinationConf, ownershipChallenge, "account")
}
func (api *API) ValidateZoneLogpushOwnershipChallenge(ctx context.Context, zoneID, destinationConf, ownershipChallenge string) (bool, error) {
	return api.ValidateLogpushOwnershipChallenge(ctx, zoneID, destinationConf, ownershipChallenge, "zone")
}

// CheckLogpushDestinationExists returns destination exists check result.
//
// API reference: https://api.cloudflare.com/#logpush-jobs-check-destination-exists
func (api *API) CheckLogpushDestinationExists(ctx context.Context, identifier, destinationConf string, identifierType string) (bool, error) {
	uri := fmt.Sprintf("/%s/%s/logpush/validate/destination/exists", getRouteRoot(identifierType), identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, LogpushDestinationExistsRequest{
		DestinationConf: destinationConf,
	})
	if err != nil {
		return false, err
	}
	var r LogpushDestinationExistsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return false, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result.Exists, nil
}
func (api *API) CheckAccountLogpushDestinationExists(ctx context.Context, accountID, destinationConf string) (bool, error) {
	return api.CheckLogpushDestinationExists(ctx, accountID, destinationConf, "account")
}
func (api *API) CheckZoneLogpushDestinationExists(ctx context.Context, zoneID, destinationConf string) (bool, error) {
	return api.CheckLogpushDestinationExists(ctx, zoneID, destinationConf, "zone")
}
