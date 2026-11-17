package geoapify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scientist-v08/bphs/initializers"
	"github.com/scientist-v08/bphs/models"
)

func SearchPlaces(
	ctx context.Context,
	query string,
) (*models.GeoapifyResponse, error) {

	if strings.TrimSpace(query) == "" {
		return nil, fmt.Errorf("search query cannot be empty")
	}

	if initializers.GeoapifyAPIKey == "" {
		return nil, fmt.Errorf("Geoapify API key is not configured")
	}

	endpoint := "https://api.geoapify.com/v1/geocode/autocomplete"

	params := url.Values{}
	params.Set("text", query)
	params.Set("format", "geojson")
	params.Set("limit", "8")
	params.Set("apiKey", initializers.GeoapifyAPIKey)

	reqURL := endpoint + "?" + params.Encode()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		reqURL,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create Geoapify request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Geoapify request failed: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"Geoapify returned HTTP status %d",
			response.StatusCode,
		)
	}

	var result models.GeoapifyResponse

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode Geoapify response: %w", err)
	}

	return &result, nil
}