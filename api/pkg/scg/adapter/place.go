package adapter

import (
	"context"

	"github.com/soulski/test-scg/pkg/scg/config"
	"googlemaps.github.io/maps"
)

type LatLng struct {
	Lat float64
	Lng float64
}

func NewLatLng(location *maps.LatLng) *LatLng {
	return &LatLng{Lat: location.Lat, Lng: location.Lng}
}

type Place struct {
	Id       string
	Location *LatLng
	Name     string
}

type PlaceSearchResponse struct {
	Places []*Place
}

type PlaceAdapter struct {
	client *maps.Client
}

func NewPlaceAdapter(config *config.PlaceConfig) *PlaceAdapter {
	cli, err := maps.NewClient(maps.WithAPIKey(config.ApiKey))
	if err != nil {
		panic(err)
	}

	return &PlaceAdapter{client: cli}
}

func (p *PlaceAdapter) FindNearbyPlace() (*PlaceSearchResponse, error) {
	bangsueLocation := &maps.LatLng{
		Lat: 13.828253,
		Lng: 100.5284507,
	}

	req := &maps.NearbySearchRequest{
		Location: bangsueLocation,
		Radius:   500,
		Keyword:  "restaurants",
	}

	result, err := p.client.NearbySearch(context.Background(), req)
	if err != nil {
		return nil, err
	}

	places := make([]*Place, len(result.Results))
	for index, placeSearch := range result.Results {
		places[index] = &Place{
			Id:       placeSearch.PlaceID,
			Name:     placeSearch.Name,
			Location: NewLatLng(&placeSearch.Geometry.Location),
		}
	}

	return &PlaceSearchResponse{Places: places}, nil
}
