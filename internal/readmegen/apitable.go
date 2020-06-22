package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func processSupportedAPIRegion(w *bytes.Buffer) error {
	b, err := ioutil.ReadFile("internal/.repo-metadata-full.json")
	if err != nil {
		return fmt.Errorf("unable to read .repo-metadata file: %v", err)
	}
	var apiMetadata map[string]json.RawMessage
	if err := json.Unmarshal(b, &apiMetadata); err != nil {
		return fmt.Errorf("unable to unmarshal api metadata: %v", err)
	}
	type tableRow struct {
		name          string
		cloudSiteLink string
		status        string
		pkgName       string
		pkgDocLink    string
		libType       string
	}
	var rows map[string]tableRow
	for _, v := range apiMetadata {
		var api struct {
			DistributionName string `json:"distribution_name"`
			Description      string `json:"description"`
			DocsURL          string `json:"docs_url"`
			ReleaseLevel     string `json:"release_level"`
			LibType          string `json:"client_library_type"`
		}
		if err := json.Unmarshal(v, &api); err != nil {
			return fmt.Errorf("unable to unmarshal api metadata: %v", err)
		}
		tb := tableRow{
			name:          api.Description,
			cloudSiteLink: "https://cloud.google.com", // TODO(codyoss): where can we get this...
			status:        toStatus(api.ReleaseLevel),
			pkgName:       api.DistributionName,
			pkgDocLink:    api.DocsURL,
			libType:       api.LibType,
		}
		key := pkgKey(api.DistributionName)
		if _, ok := rows[key]; ok {
			//TODO(codyoss): something
			continue
		}
		rows[key] = tb
	}
	for _, v := range rows {
		log.Println(v.name, v.status)
	}
	w.WriteString("Found the Supported apis region tag\n")
	return nil
}

func toStatus(releaseLevel string) string {
	if releaseLevel == "ga" {
		return "stable"
	}
	return releaseLevel
}

func pkgKey(s string) string {
	key := strings.TrimPrefix(s, "cloud.google.com/go")
	return key
}
