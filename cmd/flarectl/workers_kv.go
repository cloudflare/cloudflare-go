package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/codegangsta/cli"
)

var kvMaxCharView = 100 // Maximum chars of a value to preview

func kvPair(ns string, pair *cloudflare.WorkersKVPair) []string {
	expires := "Never"
	if pair.Expiration > 0 {
		expires = strconv.FormatUint(uint64(pair.Expiration), 10)
	}
	value := string(pair.Value)
	if len(value) > kvMaxCharView {
		value = value[0:kvMaxCharView] + "..."
	}
	return []string{
		ns,
		pair.Name,
		expires,
		value,
		strconv.FormatInt(int64(len(pair.Value)), 10),
	}
}

func kvNamespace(c *cli.Context) string {
	if err := checkFlags(c, "namespace"); err != nil {
		return ""
	}

	split := strings.SplitN(c.String("namespace"), ":", 2)
	prefix := split[0]
	if len(split) < 2 || !(prefix == "name" || prefix == "id") {
		fmt.Println("error: namespace must be in format id:<namespace id> or name:<namespace title>")
		return ""
	}
	title := split[1]
	if prefix == "id" {
		return title
	}

	list, err := api.ListWorkersKVNamespaces(context.Background())
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for i := 0; i < len(list.Result); i++ {
		ns := list.Result[i]
		if ns.Title == title {
			return ns.ID
		}
	}

	res, err := api.CreateWorkersKVNamespace(context.Background(),
		&cloudflare.WorkersKVNamespaceRequest{Title: title})
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return res.Result.ID
}

func kvRead(c *cli.Context) {
	if err := checkFlags(c, "key"); err != nil {
		return
	}
	keys := c.StringSlice("key")

	ns := kvNamespace(c)
	if ns == "" {
		return
	}

	raw := c.Bool("raw")
	if raw && len(keys) > 1 {
		fmt.Println("error: cannot combine --raw with multiple keys")
		return
	}

	output := [][]string{}
	for _, key := range keys {
		pair, err := api.ReadWorkersKVPair(context.Background(), ns, key)
		if err != nil {
			fmt.Println(err)
			return
		}

		if raw {
			fmt.Print(string(pair.Value))
			return
		}
		output = append(output, kvPair(ns, &pair))
	}

	writeTable(output, "Namespace ID", "Key", "Expiration", "Value", "Bytes")
}

func kvWrite(c *cli.Context) {
	ns := kvNamespace(c)
	if ns == "" {
		return
	}

	if err := checkFlags(c, "key", "value"); err != nil {
		return
	}

	pair := &cloudflare.WorkersKVPair{
		Name:          c.String("key"),
		Value:         c.String("value"),
		Expiration:    uint(c.Int("expiration")),
		ExpirationTTL: uint(c.Duration("ttl").Seconds()),
	}
	_, err := api.WriteWorkersKVPair(context.Background(), ns, pair)
	if err != nil {
		fmt.Println(err)
		return
	}

	kvRead(c)
}

func kvDelete(c *cli.Context) {
	if err := checkFlags(c, "key"); err != nil {
		return
	}
	key := c.String("key")

	ns := kvNamespace(c)
	if ns == "" {
		return
	}

	_, err := api.DeleteWorkersKVPair(context.Background(), ns, key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func kvList(c *cli.Context) {
	ns := kvNamespace(c)
	if ns == "" {
		return
	}
	page := cloudflare.WorkersKVKeyPagination{
		Cursor: c.String("cursor"),
		Count:  uint(c.Int("limit")),
	}
	res, err := api.ListWorkersKVKeys(context.Background(), ns, &page)
	if err != nil {
		fmt.Println(err)
		return
	}

	output := [][]string{
		[]string{
			ns,
			strconv.FormatUint(uint64(res.ResultInfo.Count), 10),
			res.ResultInfo.Cursor,
		},
	}
	writeTable(output, "Namespace ID", "Count", "Cursor")

	values := c.Bool("values")
	output = [][]string{}
	for _, key := range res.Result {
		if values {
			pair, err := api.ReadWorkersKVPair(context.Background(), ns, key.Name)
			if err != nil {
				fmt.Println(err)
				return
			}
			output = append(output, kvPair(ns, &pair)[1:5])
		} else {
			pair := cloudflare.WorkersKVPair{Name: key.Name, Expiration: key.Expiration}
			output = append(output, kvPair(ns, &pair)[1:3])
		}
	}

	if values {
		writeTable(output, "Key", "Expiration", "Value", "Bytes")
	} else {
		writeTable(output, "Key", "Expiration")
	}
}

// Technically the official limits are larger than this,
// however the payload size becomes larger due to string encoding.
//
// With discretion, use --unsafe to get around these limits.

var kvMaxRequestSize int64 = 2.097e+7 // 20MB approximately
var kvMaxValueSize int64 = 1.049e+6   // 1MB per value
var kvMaxKeyAmount int64 = 1e4        // 10k keys

func kvUpload(c *cli.Context) {
	ns := kvNamespace(c)
	if err := checkFlags(c, "path"); err != nil || ns == "" {
		return
	}

	unsafe := c.Bool("unsafe")
	root := c.String("path")
	os.Chdir(root)

	var statOk int64     // Number of uploaded files
	var statIgnore int64 // Number of ignored files
	var statSize int64   // Total bytes so far in this batch
	var statTotal int64  // Total bytes in all batches
	var statBatch int64  // Number of batches

	pairs := []*cloudflare.WorkersKVPair{}
	var path string
	err := filepath.Walk(root, func(pth string, info os.FileInfo, err error) error {
		path = pth
		if err != nil {
			return err
		}

		match, _ := regexp.MatchString(c.String("filter"), path)
		if info.IsDir() {
			if root != path && !c.Bool("recursive") {
				return filepath.SkipDir
			}
		} else if match {
			key := c.String("key")
			if strings.Count(key, "%s") != 0 {
				key = fmt.Sprintf(c.String("key"), strings.Replace(path, root, "", 1))
			}

			var size = int64(info.Size())
			if !unsafe && size > kvMaxValueSize {
				statIgnore++
				return nil
			}

			if !unsafe && (int64(len(pairs)) >= kvMaxKeyAmount || statSize >= kvMaxRequestSize) {
				_, err := api.WriteWorkersKVPairBulk(context.Background(), ns, pairs)
				if err != nil {
					return err
				}

				pairs = []*cloudflare.WorkersKVPair{}
				statBatch++
				statTotal += statSize
				statSize = 0
			}

			data, err := ioutil.ReadFile(path)
			size = int64(len(data))
			if err != nil {
				fmt.Println(err)
				return nil
			}

			pairs = append(pairs, &cloudflare.WorkersKVPair{
				Name:          key,
				Value:         string(data),
				Expiration:    uint(c.Int("expiration")),
				ExpirationTTL: uint(c.Duration("ttl").Seconds()),
			})
			statSize += size + int64(len(key))
			statOk++
		}

		return nil
	})
	if err != nil {
		fmt.Printf("upload error @ %s:\n", path)
		fmt.Println(err)
		return
	}

	if len(pairs) > 0 {
		_, err = api.WriteWorkersKVPairBulk(context.Background(), ns, pairs)
		if err != nil {
			fmt.Println(err)
			return
		}
		statBatch++
		statTotal += statSize
	}

	output := [][]string{{
		strconv.FormatInt(statOk, 10),
		strconv.FormatInt(statIgnore, 10),
		strconv.FormatInt(statTotal, 10),
		strconv.FormatInt(statBatch, 10),
	}}
	writeTable(output, "Ok", "Ignored", "Bytes", "Requests")
}
