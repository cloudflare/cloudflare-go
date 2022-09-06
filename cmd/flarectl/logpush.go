package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/cloudflare/cloudflare-go"
	"github.com/urfave/cli/v2"
)

func accountLogpushJobs(c *cli.Context) error {
	if err := checkFlags(c, "account-id"); err != nil {
		return err
	}
	accountID := c.String("account-id")

	jobs, err := api.ListAccountLogpushJobs(context.Background(), accountID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	output, cols, err := outputLogpushJobs(jobs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	writeTable(c, output, cols...)
	return nil
}

func accountCreateLogpushJob(c *cli.Context) error {
	if err := checkFlags(c, "dataset", "destination-conf", "account-id"); err != nil {
		cli.ShowSubcommandHelp(c) //nolint
		return err
	}
	accountID := c.String("account-id")

	job := cloudflare.LogpushJob{
		Dataset:         c.String("dataset"),
		DestinationConf: c.String("destination-conf"),
		Enabled:         c.Bool("enabled"),
	}

	var filter cloudflare.LogpushJobFilters
	if f := c.String("filter"); f != "" {
		err := json.Unmarshal([]byte(f), &filter)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}
		job.Filter = &filter
	}

	if k := c.String("kind"); k != "" {
		job.Kind = k
	}

	if l := c.String("logpull-options"); l != "" {
		job.LogpullOptions = l
	}

	if m := c.Int("max-upload-bytes"); m > 0 {
		job.MaxUploadBytes = m
	}

	if m := c.Int("max-upload-records"); m > 0 {
		job.MaxUploadRecords = m
	}

	if n := c.String("name"); n != "" {
		job.Name = n
	}

	returnedJob, err := api.CreateAccountLogpushJob(context.Background(), accountID, job)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	output, cols, err := outputLogpushJobs([]cloudflare.LogpushJob{*returnedJob})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	writeTable(c, output, cols...)
	return nil
}

func accountDeleteLogpushJob(c *cli.Context) error {
	if err := checkFlags(c, "account-id"); err != nil {
		cli.ShowSubcommandHelp(c) //nolint
		return err
	}
	accountID := c.String("account-id")

	jobID := c.Int("id")
	if jobID <= 0 {
		err := fmt.Errorf("error: the required flag 'id' was invalid or not provided")
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	err := api.DeleteAccountLogpushJob(context.Background(), accountID, jobID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}

func zoneLogpushJobs(c *cli.Context) error {
	if err := checkFlags(c, "zone"); err != nil {
		cli.ShowSubcommandHelp(c) //nolint
		return err
	}

	zoneID, err := api.ZoneIDByName(c.String("zone"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	jobs, err := api.ListZoneLogpushJobs(context.Background(), zoneID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	output, cols, err := outputLogpushJobs(jobs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	writeTable(c, output, cols...)
	return nil
}

func zoneCreateLogpushJob(c *cli.Context) error {
	if err := checkFlags(c, "dataset", "destination-conf", "zone"); err != nil {
		cli.ShowSubcommandHelp(c) //nolint
		return err
	}

	zoneID, err := api.ZoneIDByName(c.String("zone"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	job := cloudflare.LogpushJob{
		Dataset:         c.String("dataset"),
		DestinationConf: c.String("destination-conf"),
		Enabled:         c.Bool("enabled"),
	}

	var filter cloudflare.LogpushJobFilters
	if f := c.String("filter"); f != "" {
		err := json.Unmarshal([]byte(f), &filter)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}
		job.Filter = &filter
	}

	if k := c.String("kind"); k != "" {
		job.Kind = k
	}

	if l := c.String("logpull-options"); l != "" {
		job.LogpullOptions = l
	}

	if m := c.Int("max-upload-bytes"); m > 0 {
		job.MaxUploadBytes = m
	}

	if m := c.Int("max-upload-records"); m > 0 {
		job.MaxUploadRecords = m
	}

	if n := c.String("name"); n != "" {
		job.Name = n
	}

	returnedJob, err := api.CreateZoneLogpushJob(context.Background(), zoneID, job)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	output, cols, err := outputLogpushJobs([]cloudflare.LogpushJob{*returnedJob})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	writeTable(c, output, cols...)
	return nil
}

func zoneDeleteLogpushJob(c *cli.Context) error {
	if err := checkFlags(c, "zone"); err != nil {
		cli.ShowSubcommandHelp(c) //nolint
		return err
	}

	jobID := c.Int("id")
	if jobID <= 0 {
		err := fmt.Errorf("error: the required flag 'id' was invalid or not provided")
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	zoneID, err := api.ZoneIDByName(c.String("zone"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	err = api.DeleteZoneLogpushJob(context.Background(), zoneID, jobID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}

func outputLogpushJobs(jobs []cloudflare.LogpushJob) (output [][]string, cols []string, err error) {
	for _, job := range jobs {
		filter, err := json.Marshal(job.Filter)
		if err != nil {
			return [][]string{}, []string{}, err
		}
		var lastComplete, lastError string
		if job.LastComplete != nil {
			lastComplete = job.LastComplete.String()
		}
		if job.LastError != nil {
			lastError = job.LastError.String()
		}
		output = append(output, []string{
			fmt.Sprintf("%d", job.ID),
			job.Name,
			job.Dataset,
			job.Frequency,
			string(filter),
			job.Kind,
			job.DestinationConf,
			job.LogpullOptions,
			strconv.FormatBool(job.Enabled),
			lastComplete,
			lastError,
			job.ErrorMessage,
		})
	}
	cols = []string{
		"ID", "Name", "Dataset", "Frequency", "Filter", "Kind", "DestinationConf",
		"LogpullOptions", "Enabled", "LastComplete", "LastError", "ErrorMessage",
	}
	return output, cols, nil
}
