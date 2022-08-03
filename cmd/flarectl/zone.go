package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/urfave/cli/v2"
)

func zoneCerts(*cli.Context) error {
	return nil
}

func zoneKeyless(*cli.Context) error {
	return nil
}

func zoneRailgun(*cli.Context) error {
	return nil
}

func zoneCreate(c *cli.Context) error {
	if err := checkFlags(c, "zone"); err != nil {
		return err
	}
	zone := c.String("zone")
	jumpstart := c.Bool("jumpstart")
	accountID := c.String("account-id")
	zoneType := c.String("type")
	var account cloudflare.Account
	if accountID != "" {
		account.ID = accountID
	}

	if zoneType != "partial" {
		zoneType = "full"
	}

	_, err := api.CreateZone(context.Background(), zone, jumpstart, account, zoneType)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		return err
	}

	return nil
}

func zoneCheck(c *cli.Context) error {
	if err := checkFlags(c, "zone"); err != nil {
		return err
	}
	zone := c.String("zone")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return err
	}

	res, err := api.ZoneActivationCheck(context.Background(), zoneID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("%s\n", res.Messages[0].Message)

	return nil
}

func zoneList(c *cli.Context) error {
	zones, err := api.ListZones(context.Background())
	nbDaysStats := c.Int("with-days-stats")
	if err != nil {
		fmt.Println(err)
		return err
	}
	output := make([][]string, 0, len(zones))
	cols := []string{"ID", "Name", "Plan", "Status"}
	if nbDaysStats != 0 {
		cols = append(cols, "Stats")
	}
	for _, z := range zones {
		// Prepare the row of data to add
		outputBuffer := make([][]string, 0, 1)
		outputBuffer = append(outputBuffer, []string{
			z.ID,
			z.Name,
			z.Plan.Name,
			z.Status,
		})
		if nbDaysStats != 0 {
			zoneStats(nbDaysStats, z.ID, &outputBuffer, 1, -1)
		}

		output = append(output, outputBuffer[0])
	}

	writeTable(c, output, cols...)

	return nil
}

func zoneDelete(c *cli.Context) error {
	if err := checkFlags(c, "zone"); err != nil {
		return err
	}

	zoneID, err := api.ZoneIDByName(c.String("zone"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	_, err = api.DeleteZone(context.Background(), zoneID)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		return err
	}

	return nil
}

func zoneCreateLockdown(c *cli.Context) error {
	if err := checkFlags(c, "zone", "urls", "targets", "values"); err != nil {
		return err
	}
	zoneID, err := api.ZoneIDByName(c.String("zone"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	targets := c.StringSlice("targets")
	values := c.StringSlice("values")
	if len(targets) != len(values) {
		cli.ShowCommandHelp(c, "targets and values does not match") //nolint
		return nil
	}
	var zonelockdownconfigs = []cloudflare.ZoneLockdownConfig{}
	for index := 0; index < len(targets); index++ {
		zonelockdownconfigs = append(zonelockdownconfigs, cloudflare.ZoneLockdownConfig{
			Target: c.StringSlice("targets")[index],
			Value:  c.StringSlice("values")[index],
		})
	}
	lockdown := cloudflare.ZoneLockdown{
		Description:    c.String("description"),
		URLs:           c.StringSlice("urls"),
		Configurations: zonelockdownconfigs,
	}

	var resp *cloudflare.ZoneLockdownResponse

	resp, err = api.CreateZoneLockdown(context.Background(), zoneID, lockdown)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating ZONE lock down: ", err)
		return err
	}
	output := make([][]string, 0, 1)
	output = append(output, formatLockdownResponse(resp))

	writeTable(c, output, "ID")

	return nil
}

func zoneInfo(c *cli.Context) error {
	var zone string
	if c.NArg() > 0 {
		zone = c.Args().First()
	} else if c.String("zone") != "" {
		zone = c.String("zone")
	} else {
		cli.ShowSubcommandHelp(c) //nolint
		return nil
	}
	zones, err := api.ListZones(context.Background(), zone)
	if err != nil {
		fmt.Println(err)
		return err
	}
	output := make([][]string, 0, len(zones))
	for _, z := range zones {
		var nameservers []string
		if len(z.VanityNS) > 0 {
			nameservers = z.VanityNS
		} else {
			nameservers = z.NameServers
		}
		output = append(output, []string{
			z.ID,
			z.Name,
			z.Plan.Name,
			z.Status,
			strings.Join(nameservers, ", "),
			fmt.Sprintf("%t", z.Paused),
			z.Type,
		})
	}
	writeTable(c, output, "ID", "Zone", "Plan", "Status", "Name Servers", "Paused", "Type")

	return nil
}

func zonePlan(*cli.Context) error {
	return nil
}

func zoneSettings(*cli.Context) error {
	return nil
}

func zoneCachePurge(c *cli.Context) error {
	if err := checkFlags(c, "zone"); err != nil {
		cli.ShowSubcommandHelp(c) //nolint
		return err
	}

	zoneName := c.String("zone")
	zoneID, err := api.ZoneIDByName(c.String("zone"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	var resp cloudflare.PurgeCacheResponse

	// Purge everything
	if c.Bool("everything") {
		resp, err = api.PurgeEverything(context.Background(), zoneID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error purging all from zone %q: %s\n", zoneName, err)
			return err
		}
	} else {
		var (
			files    = c.StringSlice("files")
			tags     = c.StringSlice("tags")
			hosts    = c.StringSlice("hosts")
			prefixes = c.StringSlice("prefixes")
		)

		if len(files) == 0 && len(tags) == 0 && len(hosts) == 0 && len(prefixes) == 0 {
			fmt.Fprintln(os.Stderr, "You must provide at least one of the --files, --tags, --prefixes or --hosts flags")
			return nil
		}

		// Purge selectively
		purgeReq := cloudflare.PurgeCacheRequest{
			Files:    c.StringSlice("files"),
			Tags:     c.StringSlice("tags"),
			Hosts:    c.StringSlice("hosts"),
			Prefixes: c.StringSlice("prefixes"),
		}

		resp, err = api.PurgeCache(context.Background(), zoneID, purgeReq)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error purging the cache from zone %q: %s\n", zoneName, err)
			return err
		}
	}

	output := make([][]string, 0, 1)
	output = append(output, formatCacheResponse(resp))

	writeTable(c, output, "ID")

	return nil
}

func zoneStats(nbDaysStats int, zoneId string, output *[][]string, colDnsNameIndex int, colDnsTypeIndex int) {
	dateTo := time.Now()
	dateFrom := dateTo.AddDate(0, 0, -1*nbDaysStats)
	batchSize := 100 // Number of GraphQL objects being queried by query
	nbElems := len(*output)
	iBatch := 0
	for iBatch*batchSize < nbElems {
		zoneStatsFromTo(zoneId, dateFrom, dateTo, output, colDnsNameIndex, colDnsTypeIndex, iBatch*batchSize, (iBatch+1)*batchSize)
		iBatch = iBatch + 1
	}
}

func zoneStatsFromTo(zoneId string, dateFrom time.Time, dateTo time.Time, output *[][]string, colDnsNameIndex int, colDnsTypeIndex int, iFrom int, iTo int) {
	// Build the zone-scoped GraphQL query
	totalGraphQuery := fmt.Sprintf("query { viewer { zones(filter: { zoneTag: \"%s\" }) { ", zoneId)
	bAnyQuery := false
	for i, r := range *output {
		if i >= iFrom && i < iTo {
			if colDnsTypeIndex == -1 || (*output)[i][colDnsTypeIndex] == "CNAME" || (*output)[i][colDnsTypeIndex] == "A" {
				partGraphQuery := fmt.Sprintf(
					"e%d: httpRequestsAdaptiveGroups(filter:{datetime_gt: \"%s\",datetime_lt: \"%s\",OR:[{clientRequestHTTPHost_like: \"%s\"},{clientRequestHTTPHost_like: \"%s\"}]}, limit: 10, orderBy: [count_DESC]) { count }",
					i,
					dateFrom.Format("2006-01-02T15:04:05Z"),
					dateTo.Format("2006-01-02T15:04:05Z"),
					"%."+strings.Replace(r[colDnsNameIndex], "*.", "", -1),
					r[colDnsNameIndex])
				totalGraphQuery = fmt.Sprintf("%s %s", totalGraphQuery, partGraphQuery)
				bAnyQuery = true
			}
		}
	}
	if !bAnyQuery {
		return
	}
	totalGraphQuery = fmt.Sprintf("%s}}}", totalGraphQuery)

	// Post the GraphQL query
	payload := strings.NewReader(fmt.Sprintf("{\"query\":\"%s\",\"variables\":{}}", strings.Replace(totalGraphQuery, "\"", "\\\"", -1)))
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.cloudflare.com/client/v4/graphql", payload)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	apiToken := os.Getenv("CF_API_TOKEN")
	if apiToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	} else {
		req.Header.Add("X-Auth-Email", os.Getenv("CF_API_EMAIL"))
		req.Header.Add("X-Auth-Key", os.Getenv("CF_API_KEY"))
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer res.Body.Close()
	bodyResponse, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Parse the response to inject the event counts into the output
	var resultStats map[string]interface{}
	var zone map[string]interface{}
	json.Unmarshal([]byte(bodyResponse), &resultStats)
	if resultStats["data"] == nil { // This zone is not having data
		return
	}
	zone = resultStats["data"].(map[string]interface{})["viewer"].(map[string]interface{})["zones"].([]interface{})[0].(map[string]interface{})

	for i := iFrom; i < iTo; i++ {
		if i >= len(*output) {
			return
		}
		var zoneEid string
		var zoneId []interface{}
		injectedStat := ""
		zoneEid = fmt.Sprintf("e%d", i)
		if reflect.TypeOf(zone[zoneEid]) == reflect.TypeOf(zoneId) {
			zoneId = zone[zoneEid].([]interface{})
			if zoneId != nil && len(zoneId) > 0 {
				countVal := zoneId[0].(map[string]interface{})["count"].(float64)
				injectedStat = strconv.FormatFloat(countVal, 'f', -1, 64)
			}
		}
		(*output)[i] = append((*output)[i], injectedStat)
	}
}

func zoneRecords(c *cli.Context) error {
	var zone string
	if c.NArg() > 0 {
		zone = c.Args().First()
	} else if c.String("zone") != "" {
		zone = c.String("zone")
	} else {
		cli.ShowSubcommandHelp(c) //nolint
		return nil
	}
	nbDaysStats := c.Int("with-days-stats")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create a an empty record for searching for records
	rr := cloudflare.DNSRecord{}
	var records []cloudflare.DNSRecord
	if c.String("id") != "" {
		rec, err := api.DNSRecord(context.Background(), zoneID, c.String("id"))
		if err != nil {
			fmt.Println(err)
			return err
		}
		records = append(records, rec)
	} else {
		if c.String("type") != "" {
			rr.Type = c.String("type")
		}
		if c.String("name") != "" {
			rr.Name = c.String("name")
		}
		if c.String("content") != "" {
			rr.Content = c.String("content")
		}
		var err error
		records, err = api.DNSRecords(context.Background(), zoneID, rr)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	output := make([][]string, 0, len(records))
	for _, r := range records {
		switch r.Type {
		case "MX":
			r.Content = fmt.Sprintf("%d %s", r.Priority, r.Content)
		case "SRV":
			dp := r.Data.(map[string]interface{})
			r.Content = fmt.Sprintf("%.f %s", dp["priority"], r.Content)
			// Cloudflare's API, annoyingly, automatically prepends the weight
			// and port into content, separated by tabs.
			// XXX: File this as a bug. LOC doesn't do this.
			r.Content = strings.Replace(r.Content, "\t", " ", -1)
		}
		output = append(output, []string{
			r.ID,
			r.Type,
			r.Name,
			r.Content,
			strconv.FormatBool(*r.Proxied),
			fmt.Sprintf("%d", r.TTL),
		})
	}
	cols := []string{"ID", "Type", "Name", "Content", "Proxied", "TTL"}
	if nbDaysStats != 0 {
		zoneStats(nbDaysStats, zoneID, &output, 2, 1)
		cols = append(cols, "Stats")
	}
	writeTable(c, output, cols...)

	return nil
}

func formatCacheResponse(resp cloudflare.PurgeCacheResponse) []string {
	return []string{
		resp.Result.ID,
	}
}

func formatLockdownResponse(resp *cloudflare.ZoneLockdownResponse) []string {
	return []string{
		resp.Result.ID,
	}
}

func zoneExport(c *cli.Context) error {
	var zone string
	if c.NArg() > 0 {
		zone = c.Args().First()
	} else if c.String("zone") != "" {
		zone = c.String("zone")
	} else {
		cli.ShowSubcommandHelp(c) //nolint
		return nil
	}

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return err
	}

	res, err := api.ZoneExport(context.Background(), zoneID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Print(res)

	return nil
}
