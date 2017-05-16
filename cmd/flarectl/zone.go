package main

import (
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/codegangsta/cli"
)

func zoneCerts(*cli.Context) {
}

func zoneKeyless(*cli.Context) {
}

func zoneRailgun(*cli.Context) {
}

func zoneCreate(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	if err := checkFlags(c, "zone"); err != nil {
		return
	}
	zone := c.String("zone")
	jumpstart := c.Bool("jumpstart")
	orgID := c.String("org-id")
	var org cloudflare.Organization
	if orgID != "" {
		org.ID = orgID
	}
	api.CreateZone(zone, jumpstart, org)
}

func zoneCheck(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	if err := checkFlags(c, "zone"); err != nil {
		return
	}
	zone := c.String("zone")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := api.ZoneActivationCheck(zoneID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", res.Messages[0].Message)
}

func zoneList(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	zones, err := api.ListZones()
	if err != nil {
		fmt.Println(err)
		return
	}
	var output []table
	for _, z := range zones {
		output = append(output, table{
			"ID":     z.ID,
			"Name":   z.Name,
			"Plan":   z.Plan.Name,
			"Status": z.Status,
		})
	}
	makeTable(output, "ID", "Name", "Plan", "Status")
}

func zoneInfo(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	var zone string
	if len(c.Args()) > 0 {
		zone = c.Args()[0]
	} else if c.String("zone") != "" {
		zone = c.String("zone")
	} else {
		cli.ShowSubcommandHelp(c)
		return
	}
	zones, err := api.ListZones(zone)
	if err != nil {
		fmt.Println(err)
		return
	}
	var output []table
	for _, z := range zones {
		var nameservers []string
		if len(z.VanityNS) > 0 {
			nameservers = z.VanityNS
		} else {
			nameservers = z.NameServers
		}
		output = append(output, table{
			"ID":           z.ID,
			"Zone":         z.Name,
			"Plan":         z.Plan.Name,
			"Status":       z.Status,
			"Name Servers": strings.Join(nameservers, ", "),
			"Paused":       fmt.Sprintf("%t", z.Paused),
			"Type":         z.Type,
		})
	}
	makeTable(output, "ID", "Zone", "Plan", "Status", "Name Servers", "Paused", "Type")
}

func zonePlan(*cli.Context) {
}

func zoneSettings(*cli.Context) {
}

func zoneRecords(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}
	var zone string
	if len(c.Args()) > 0 {
		zone = c.Args()[0]
	} else if c.String("zone") != "" {
		zone = c.String("zone")
	} else {
		cli.ShowSubcommandHelp(c)
		return
	}

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a an empty record for searching for records
	rr := cloudflare.DNSRecord{}
	var records []cloudflare.DNSRecord
	if c.String("id") != "" {
		rec, err := api.DNSRecord(zoneID, c.String("id"))
		if err != nil {
			fmt.Println(err)
			return
		}
		records = append(records, rec)
	} else {
		if c.String("name") != "" {
			rr.Name = c.String("name")
		}
		if c.String("content") != "" {
			rr.Name = c.String("content")
		}
		var err error
		records, err = api.DNSRecords(zoneID, rr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	var output []table
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
		output = append(output, table{
			"ID":      r.ID,
			"Type":    r.Type,
			"Name":    r.Name,
			"Content": r.Content,
			"Proxied": fmt.Sprintf("%t", r.Proxied),
			"TTL":     fmt.Sprintf("%d", r.TTL),
		})
	}
	makeTable(output, "ID", "Type", "Name", "Content", "Proxied", "TTL")
}

func zoneCustHostList(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}

	var zone string
	if c.String("zone") != "" {
		zone = c.String("zone")
	} else {
		cli.ShowSubcommandHelp(c)
		return
	}
	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	customhostnamesResult, _, err := api.CustomHostnames(zoneID, 0, cloudflare.CustomHostname{})
	if err != nil {
		fmt.Println(err)
		return
	}

	var output []table
	for _, z := range customhostnamesResult {
		output = append(output, table{
			"Hostname": z.Hostname,
			"ID":       z.ID,
		})
	}
	makeTable(output, "Hostname", "ID")
}

func zoneCustHostInfo(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}

	var zone string
	if c.String("zone") != "" {
		zone = c.String("zone")
	} else {
		cli.ShowSubcommandHelp(c)
		return
	}
	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	var hostname string
	if c.String("hostname") != "" {
		hostname = c.String("hostname")
	} else {
		cli.ShowSubcommandHelp(c)
		return
	}
	hostnameID, err := api.CustomHostnameIDByName(zoneID, hostname)
	if err != nil {
		fmt.Println(err)
		return
	}

	customname, err := api.CustomHostname(zoneID, hostnameID)
	if err != nil {
		fmt.Println(err)
		return
	}

	var output []table
	output = append(output, table{
		"Hostname":   customname.Hostname,
		"ID":         customname.ID,
		"SSL Status": customname.SSL.Status,
		"SSL Method": customname.SSL.Method,
		"SSL Type":   customname.SSL.Type,
	})
	makeTable(output, "Hostname", "ID", "SSL Status", "SSL Method", "SSL Type")
}

func zoneCustHostCreate(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}

	var zone string
	if c.String("zone") != "" {
		zone = c.String("zone")
	} else {
		cli.ShowSubcommandHelp(c)
		return
	}
	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	var hostname string
	if c.String("hostname") != "" {
		hostname = c.String("hostname")
	} else {
		cli.ShowSubcommandHelp(c)
		return
	}

	var sslMethod string
	if c.String("ssl-method") != "" {
		sslMethod = c.String("ssl-method")
	} else {
		cli.ShowSubcommandHelp(c)
		return
	}

	var sslType string
	if c.String("ssl-type") != "" {
		sslType = c.String("ssl-type")
	} else {
		cli.ShowSubcommandHelp(c)
		return
	}

	newCustHost := cloudflare.CustomHostname{Hostname: hostname, SSL: cloudflare.CustomHostnameSSL{Method: sslMethod, Type: sslType}}
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := api.CreateCustomHostname(zoneID, newCustHost)
	if err != nil {
		fmt.Println(err)
		return
	}

	var output []table
	output = append(output, table{
		"Created":    fmt.Sprintf("%t", res.Response.Success),
		"Hostname":   res.Result.Hostname,
		"ID":         res.Result.ID,
		"SSL Status": res.Result.SSL.Status,
		"SSL Method": res.Result.SSL.Method,
		"SSL Type":   res.Result.SSL.Type,
	})
	makeTable(output, "Created", "Hostname", "ID", "SSL Status", "SSL Method", "SSL Type")
}

func zoneCustHostDelete(c *cli.Context) {
	if err := checkEnv(); err != nil {
		fmt.Println(err)
		return
	}

	var zone string
	if c.String("zone") != "" {
		zone = c.String("zone")
	} else {
		cli.ShowSubcommandHelp(c)
		return
	}
	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return
	}

	var hostname string
	if c.String("hostname") != "" {
		hostname = c.String("hostname")
	} else {
		cli.ShowSubcommandHelp(c)
		return
	}
	hostnameID, err := api.CustomHostnameIDByName(zoneID, hostname)
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = api.DeleteCustomHostname(zoneID, hostnameID)
	if err != nil {
		fmt.Println(err)
		return
	}

}
