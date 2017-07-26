package main

import (
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/codegangsta/cli"
)

func zoneLockdownCreate(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	if err := checkFlags(c, "zone", "url"); err != nil {
		return
	}
	zone := c.String("zone")
	desc := c.String("description")
	urls := c.StringSlice("url")
	ips := c.StringSlice("ip")
	ipRanges := c.StringSlice("ip-range")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO(matt): for _, ip := ranges ip -> append to configurations
	// TODO(matt): for _, ipRange := ranges ipRanges -> append to configurations
	// TODO(matt): break into helper -> "buildConfigs"

	var ipConfigs []struct {
		Target string
		Value  string
	}

	var ipRangeConfigs []struct {
		Target string
		Value  string
	}

	for _, ip := range ips {
		ipConfigs = append(ipConfigs, {Target: "ip", Value: ip})
	}

	for _, range := range ipRanges {
		ipRangeConfigs = append(ipRangeConfigs, {Target: "ip_range", Value: ip})
	}

	rule := &cloudflare.ZoneLockdown{
		Description: desc,
		URLs:        urls,
	}

	// TODO: Print the result.
	_, err = api.CreateDNSRecord(zoneID, record)
	if err != nil {
		fmt.Println("Error creating DNS record:", err)
	}
}

func zoneLockdownUpdate(c *cli.Context) {}

func zoneLockdown(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	if err := checkFlags(c, "zone", "id"); err != nil {
		return
	}
	zone := c.String("zone")
	id := c.String("id")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := api.ZoneLockdown(zoneID, id)
	if err != nil {
		fmt.Println("Error creating fetching Zone Lockdown rules:", err)
	}

	lockdowns := []cloudflare.ZoneLockdown{resp.Result}
	output := printLockdowns(lockdowns)
	makeTable(output, "id", "description", "urls", "configurations")
}

func zoneLockdownList(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	if err := checkFlags(c, "zone", "page"); err != nil {
		return
	}
	zone := c.String("zone")
	page := c.Int("page")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := api.ListZoneLockdowns(zoneID, page)
	if err != nil {
		fmt.Println("Error creating fetching Zone Lockdown rules:", err)
	}

	output := printLockdowns(resp.Result)
	makeTable(output, "id", "description", "urls", "configurations")
}

func printLockdowns(rules []cloudflare.ZoneLockdown) []table {
	var output []table
	for _, rule := range rules {
		output = append(output, table{
			"id":             rule.ID,
			"description":    rule.Description,
			"urls":           strings.Join(rule.URLs, ","),
			"configurations": fmt.Sprintf("%s", rule.Configurations),
		})
	}

	return output
}
