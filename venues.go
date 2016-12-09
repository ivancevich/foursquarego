package foursquarego

import (
	"errors"
	"net/url"
)

func (a FoursquareApi) GetVenue(id string) (Venue, error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id, url.Values{}, &data, _GET, response_ch}
	return data.Venue, (<-response_ch).err
}

func (a FoursquareApi) GetVenuePhotos(id string, uv url.Values) ([]Photo, error) {
	uv = cleanValues(uv)
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/photos", uv, &data, _GET, response_ch}
	return data.Photos.Items, (<-response_ch).err
}

func (a FoursquareApi) GetVenueEvents(id string) ([]Event, error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/events", url.Values{}, &data, _GET, response_ch}
	return data.Events.Items, (<-response_ch).err
}

func (a FoursquareApi) GetVenueHereNow(id string, uv url.Values) (HereNow, error) {
	uv = cleanValues(uv)
	if uv.Get("oauth_token") == "" {
		return HereNow{}, errors.New("Requires Acting User")
	}
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/herenow", url.Values{}, &data, _GET, response_ch}
	return data.HereNow, (<-response_ch).err
}

func (a FoursquareApi) GetVenueHours(id string) (HoursResponse, error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/hours", url.Values{}, &data, _GET, response_ch}
	return HoursResponse{data.Hours, data.Popular}, (<-response_ch).err
}

func (a FoursquareApi) GetVenueLikes(id string) (LikesResponse, error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/likes", url.Values{}, &data, _GET, response_ch}
	return LikesResponse{data.Likes, data.Like}, (<-response_ch).err
}

func (a FoursquareApi) GetVenueLinks(id string) (LinksResponse, error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/links", url.Values{}, &data, _GET, response_ch}
	return data.Links, (<-response_ch).err
}

func (a FoursquareApi) GetVenueListed(id string, uv url.Values) (Listed, error) {
	uv = cleanValues(uv)
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/listed", uv, &data, _GET, response_ch}
	return data.Lists, (<-response_ch).err
}

func (a FoursquareApi) GetVenueMenu(id string) (MenuResponse, error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/menu", url.Values{}, &data, _GET, response_ch}
	return data.Menu, (<-response_ch).err
}

func (a FoursquareApi) GetVenueSimilar(id string, uv url.Values) (SimilarVenueResponse, error) {
	uv = cleanValues(uv)
	if uv.Get("oauth_token") == "" {
		return SimilarVenueResponse{}, errors.New("Requires Acting User")
	}
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/similar", url.Values{}, &data, _GET, response_ch}
	return data.SimilarVenues, (<-response_ch).err
}

func (a FoursquareApi) GetCategories() ([]Category, error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/categories", url.Values{}, &data, _GET, response_ch}
	return data.Categories, (<-response_ch).err
}

func (a FoursquareApi) Search(uv url.Values) ([]Venue, error) {
	uv = cleanValues(uv)
	if uv.Get("ll") == "" && uv.Get("near") == "" && uv.Get("intent") != "global" {
		return []Venue{}, errors.New("ll or near values required")
	}
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/search", uv, &data, _GET, response_ch}
	return data.Venues, (<-response_ch).err
}
