package suggestions

import (
	"context"
	"fmt"
	"github.com/dmalykh/dadata/request"
	"net/url"
)

//Город по IP https://dadata.ru/api/detect_address_by_ip/
func (s *Suggestions) Iplocate(ctx context.Context, ip string) (IpLocate, error) {
	var iplocate IpLocate

	var u = url.Values{}
	u.Set("ip", ip)

	err := s.makeRequest(
		ctx,
		"iplocate/address",
		request.Request{
			Method:      request.GET,
			QueryParams: u,
		},
		&iplocate,
	)

	if err != nil {
		return iplocate, fmt.Errorf(`Can't make makeRequest "%s": %s`, ip, err.Error())
	}
	return iplocate, nil
}

type IpLocate struct {
	Location struct {
		Value             string `json:"value"`
		UnrestrictedValue string `json:"unrestricted_value"`
		Data              struct {
			PostalCode           string      `json:"postal_code"`
			Country              string      `json:"country"`
			CountryIsoCode       string      `json:"country_iso_code"`
			FederalDistrict      string      `json:"federal_district"`
			RegionFiasID         string      `json:"region_fias_id"`
			RegionKladrID        string      `json:"region_kladr_id"`
			RegionIsoCode        string      `json:"region_iso_code"`
			RegionWithType       string      `json:"region_with_type"`
			RegionType           string      `json:"region_type"`
			RegionTypeFull       string      `json:"region_type_full"`
			Region               string      `json:"region"`
			AreaFiasID           interface{} `json:"area_fias_id"`
			AreaKladrID          interface{} `json:"area_kladr_id"`
			AreaWithType         interface{} `json:"area_with_type"`
			AreaType             interface{} `json:"area_type"`
			AreaTypeFull         interface{} `json:"area_type_full"`
			Area                 interface{} `json:"area"`
			CityFiasID           string      `json:"city_fias_id"`
			CityKladrID          string      `json:"city_kladr_id"`
			CityWithType         string      `json:"city_with_type"`
			CityType             string      `json:"city_type"`
			CityTypeFull         string      `json:"city_type_full"`
			City                 string      `json:"city"`
			CityArea             interface{} `json:"city_area"`
			CityDistrictFiasID   interface{} `json:"city_district_fias_id"`
			CityDistrictKladrID  interface{} `json:"city_district_kladr_id"`
			CityDistrictWithType interface{} `json:"city_district_with_type"`
			CityDistrictType     interface{} `json:"city_district_type"`
			CityDistrictTypeFull interface{} `json:"city_district_type_full"`
			CityDistrict         interface{} `json:"city_district"`
			SettlementFiasID     interface{} `json:"settlement_fias_id"`
			SettlementKladrID    interface{} `json:"settlement_kladr_id"`
			SettlementWithType   interface{} `json:"settlement_with_type"`
			SettlementType       interface{} `json:"settlement_type"`
			SettlementTypeFull   interface{} `json:"settlement_type_full"`
			Settlement           interface{} `json:"settlement"`
			StreetFiasID         interface{} `json:"street_fias_id"`
			StreetKladrID        interface{} `json:"street_kladr_id"`
			StreetWithType       interface{} `json:"street_with_type"`
			StreetType           interface{} `json:"street_type"`
			StreetTypeFull       interface{} `json:"street_type_full"`
			Street               interface{} `json:"street"`
			HouseFiasID          interface{} `json:"house_fias_id"`
			HouseKladrID         interface{} `json:"house_kladr_id"`
			HouseType            interface{} `json:"house_type"`
			HouseTypeFull        interface{} `json:"house_type_full"`
			House                interface{} `json:"house"`
			BlockType            interface{} `json:"block_type"`
			BlockTypeFull        interface{} `json:"block_type_full"`
			Block                interface{} `json:"block"`
			FlatType             interface{} `json:"flat_type"`
			FlatTypeFull         interface{} `json:"flat_type_full"`
			Flat                 interface{} `json:"flat"`
			FlatArea             interface{} `json:"flat_area"`
			SquareMeterPrice     interface{} `json:"square_meter_price"`
			FlatPrice            interface{} `json:"flat_price"`
			PostalBox            interface{} `json:"postal_box"`
			FiasID               string      `json:"fias_id"`
			FiasCode             string      `json:"fias_code"`
			FiasLevel            string      `json:"fias_level"`
			FiasActualityState   string      `json:"fias_actuality_state"`
			KladrID              string      `json:"kladr_id"`
			GeonameID            string      `json:"geoname_id"`
			CapitalMarker        string      `json:"capital_marker"`
			Okato                string      `json:"okato"`
			Oktmo                string      `json:"oktmo"`
			TaxOffice            string      `json:"tax_office"`
			TaxOfficeLegal       string      `json:"tax_office_legal"`
			Timezone             interface{} `json:"timezone"`
			GeoLat               string      `json:"geo_lat"`
			GeoLon               string      `json:"geo_lon"`
			BeltwayHit           interface{} `json:"beltway_hit"`
			BeltwayDistance      interface{} `json:"beltway_distance"`
			Metro                interface{} `json:"metro"`
			QcGeo                string      `json:"qc_geo"`
			QcComplete           interface{} `json:"qc_complete"`
			QcHouse              interface{} `json:"qc_house"`
			HistoryValues        interface{} `json:"history_values"`
			UnparsedParts        interface{} `json:"unparsed_parts"`
			Source               interface{} `json:"source"`
			Qc                   interface{} `json:"qc"`
		} `json:"data"`
	} `json:"location"`
}
