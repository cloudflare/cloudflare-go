package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/urfave/cli/v2"
)

func formatDNSRecord(record cloudflare.DNSRecord) []string {
	return []string{
		record.ID,
		record.Name,
		record.Type,
		record.Content,
		strconv.FormatInt(int64(record.TTL), 10),
		strconv.FormatBool(record.Proxiable),
		strconv.FormatBool(*record.Proxied),
		strconv.FormatBool(record.Locked),
	}
}

func dnsCreate(c *cli.Context) error {
	if err := checkFlags(c, "zone", "name", "type", "content"); err != nil {
		return err
	}
	zone := c.String("zone")
	name := c.String("name")
	rtype := c.String("type")
	content := c.String("content")
	ttl := c.Int("ttl")
	proxy := c.Bool("proxy")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return err
	}

	record := cloudflare.DNSRecord{
		Name:    name,
		Type:    strings.ToUpper(rtype),
		Content: content,
		TTL:     ttl,
		Proxied: &proxy,
	}
	resp, err := api.CreateDNSRecord(context.Background(), zoneID, record)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating DNS record: ", err)
		return err
	}

	output := [][]string{
		formatDNSRecord(resp.Result),
	}

	writeTable(c, output, "ID", "Name", "Type", "Content", "TTL", "Proxiable", "Proxy", "Locked")

	return nil
}

func dnsCreateOrUpdate(c *cli.Context) error {
	if err := checkFlags(c, "zone", "name", "type", "content"); err != nil {
		fmt.Println(err)
		return err
	}
	zone := c.String("zone")
	name := c.String("name")
	rtype := strings.ToUpper(c.String("type"))
	content := c.String("content")
	ttl := c.Int("ttl")
	proxy := c.Bool("proxy")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error updating DNS record: ", err)
		return err
	}

	// Look for an existing record
	rr := cloudflare.DNSRecord{
		Name: name + "." + zone,
	}
	records, err := api.DNSRecords(context.Background(), zoneID, rr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error fetching DNS records: ", err)
		return err
	}

	var resp *cloudflare.DNSRecordResponse
	if len(records) > 0 {
		// Record exists - find the ID and update it.
		// This is imprecise without knowing the original content; if a label
		// has multiple RRs we'll just update the first one.
		for _, r := range records {
			if r.Type == rtype {
				rr.ID = r.ID
				rr.Type = r.Type
				rr.Content = content
				rr.TTL = ttl
				rr.Proxied = &proxy

				err := api.UpdateDNSRecord(context.Background(), zoneID, r.ID, rr)
				if err != nil {
					fmt.Println("Error updating DNS record:", err)
					return err
				}
				resp = &cloudflare.DNSRecordResponse{
					Result: rr,
				}
			}
		}
	} else {
		// Record doesn't exist - create it
		rr.Type = rtype
		rr.Content = content
		rr.TTL = ttl
		rr.Proxied = &proxy
		// TODO: Print the response.
		resp, err = api.CreateDNSRecord(context.Background(), zoneID, rr)
		if err != nil {
			fmt.Println("Error creating DNS record:", err)
			return err
		}

	}

	output := [][]string{
		formatDNSRecord(resp.Result),
	}

	writeTable(c, output, "ID", "Name", "Type", "Content", "TTL", "Proxiable", "Proxy", "Locked")

	return nil
}

func dnsUpdate(c *cli.Context) error {
	if err := checkFlags(c, "zone", "id"); err != nil {
		fmt.Println(err)
		return err
	}
	zone := c.String("zone")
	recordID := c.String("id")
	name := c.String("name")
	rtype := c.String("type")
	content := c.String("content")
	ttl := c.Int("ttl")
	proxy := c.Bool("proxy")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return err
	}

	record := cloudflare.DNSRecord{
		ID:      recordID,
		Name:    name,
		Type:    strings.ToUpper(rtype),
		Content: content,
		TTL:     ttl,
		Proxied: &proxy,
	}
	err = api.UpdateDNSRecord(context.Background(), zoneID, recordID, record)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error updating DNS record: ", err)
		return err
	}

	return nil
}

func dnsDelete(c *cli.Context) error {
	if err := checkFlags(c, "zone", "id"); err != nil {
		fmt.Println(err)
		return err
	}
	zone := c.String("zone")
	recordID := c.String("id")

	zoneID, err := api.ZoneIDByName(zone)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = api.DeleteDNSRecord(context.Background(), zoneID, recordID)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error deleting DNS record: ", err)
		return err
	}

	return nil
}
