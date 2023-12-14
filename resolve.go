package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/reeveci/reeve-lib/schema"
)

func (p *ConsulPlugin) Resolve(env []string) (map[string]schema.Env, error) {
	kvUrl := fmt.Sprintf("%s/v1/kv/%s?recurse=true", p.Url, url.PathEscape(p.KeyPrefix))
	req, err := http.NewRequest(http.MethodGet, kvUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("fetching variables failed - %s", err)
	}

	req.Header.Set("X-Consul-Token", p.Token)

	resp, err := p.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching variables failed - %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil
	}

	var consulResponse ConsulResponse
	err = json.NewDecoder(resp.Body).Decode(&consulResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing consul response - %s", err)
	}

	entries := make(map[string]string, len(consulResponse))
	for _, entry := range consulResponse {
		entries[strings.TrimPrefix(entry.Key, p.KeyPrefix)] = entry.Value
	}

	result := make(map[string]schema.Env, len(env))
	for _, key := range env {
		if value, ok := entries[key]; ok {
			decoded, err := base64.StdEncoding.DecodeString(value)
			if err != nil {
				p.Log.Error(fmt.Sprintf(`error decoding base64 encoded value for key "%s" - %s`, key, err))
				continue
			}

			result[key] = schema.Env{
				Value:    string(decoded),
				Priority: p.Priority,
				Secret:   p.Secret,
			}
		}
	}

	return result, nil
}
