package suggestions

import (
	"fmt"
	"strings"
)

//Подсказки о адресу https://dadata.ru/api/suggest/#about-address
func (s *Suggestions) Address(address string, count int) ([]AddressItem, error) {
	var suggestions Address
	query := map[string]interface{}{
		"query": strings.Replace(address, "\\", " ", -1), //Дадата сходит с ума при нахождении в адресе \\
		"count": count,
	}
	if err := s.request("address", query, &suggestions); err != nil {
		return []AddressItem{}, fmt.Errorf(`Can't make request "%s" %d: %s`, address, count, err.Error())
	}
	return suggestions.Suggestions, nil
}

//Результат ответа подскаски адреса
type Address struct {
	Suggestions []AddressItem `json:"suggestions"`
}

//Единица результата ответа подсказки адреса
type AddressItem struct {
	Value             string          `json:"value"`
	UnrestrictedValue string          `json:"unrestricted_value"`
	Data              AddressItemData `json:"data"`
}

type AddressItemData struct {
	PostalCode           string   `json:"postal_code"`
	Country              string   `json:"country"`
	RegionFiasID         string   `json:"region_fias_id"`
	RegionKladrID        string   `json:"region_kladr_id"`
	RegionWithType       string   `json:"region_with_type"`
	RegionType           string   `json:"region_type"`
	RegionTypeFull       string   `json:"region_type_full"`
	Region               string   `json:"region"`
	AreaFiasID           string   `json:"area_fias_id"`
	AreaKladrID          string   `json:"area_kladr_id"`
	AreaWithType         string   `json:"area_with_type"`
	AreaType             string   `json:"area_type"`
	AreaTypeFull         string   `json:"area_type_full"`
	Area                 string   `json:"area"`
	CityFiasID           string   `json:"city_fias_id"`
	CityKladrID          string   `json:"city_kladr_id"`
	CityWithType         string   `json:"city_with_type"`
	CityType             string   `json:"city_type"`
	CityTypeFull         string   `json:"city_type_full"`
	City                 string   `json:"city"`
	CityArea             string   `json:"city_area"`
	CityDistrictFiasID   string   `json:"city_district_fias_id"`
	CityDistrictKladrID  string   `json:"city_district_kladr_id"`
	CityDistrictWithType string   `json:"city_district_with_type"`
	CityDistrictType     string   `json:"city_district_type"`
	CityDistrictTypeFull string   `json:"city_district_type_full"`
	CityDistrict         string   `json:"city_district"`
	SettlementFiasID     string   `json:"settlement_fias_id"`
	SettlementKladrID    string   `json:"settlement_kladr_id"`
	SettlementWithType   string   `json:"settlement_with_type"`
	SettlementType       string   `json:"settlement_type"`
	SettlementTypeFull   string   `json:"settlement_type_full"`
	Settlement           string   `json:"settlement"`
	StreetFiasID         string   `json:"street_fias_id"`
	StreetKladrID        string   `json:"street_kladr_id"`
	StreetWithType       string   `json:"street_with_type"`
	StreetType           string   `json:"street_type"`
	StreetTypeFull       string   `json:"street_type_full"`
	Street               string   `json:"street"`
	HouseFiasID          string   `json:"house_fias_id"`
	HouseKladrID         string   `json:"house_kladr_id"`
	HouseType            string   `json:"house_type"`
	HouseTypeFull        string   `json:"house_type_full"`
	House                string   `json:"house"`
	BlockType            string   `json:"block_type"`
	BlockTypeFull        string   `json:"block_type_full"`
	Block                string   `json:"block"`
	FlatType             string   `json:"flat_type"`
	FlatTypeFull         string   `json:"flat_type_full"`
	Flat                 string   `json:"flat"`
	FlatArea             float64  `json:"flat_area"`
	SquareMeterPrice     float64  `json:"square_meter_price"`
	FlatPrice            uint     `json:"flat_price"`
	PostalBox            string   `json:"postal_box"`
	FiasID               string   `json:"fias_id"`
	FiasCode             string   `json:"fias_code"`
	FiasLevel            int      `json:"fias_level"`
	FiasActualityState   string   `json:"fias_actuality_state"`
	KladrID              string   `json:"kladr_id"`
	CapitalMarker        string   `json:"capital_marker"`
	Okato                string   `json:"okato"`
	Oktmo                string   `json:"oktmo"`
	TaxOffice            string   `json:"tax_office"`
	TaxOfficeLegal       string   `json:"tax_office_legal"`
	Timezone             string   `json:"timezone"`
	GeoLat               float64  `json:"geo_lat"`
	GeoLon               float64  `json:"geo_lon"`
	BeltwayHit           string   `json:"beltway_hit"`
	BeltwayDistance      string   `json:"beltway_distance"`
	Metro                string   `json:"metro"`
	QcGeo                uint     `json:"qc_geo"`
	QcComplete           uint     `json:"qc_complete"`
	QcHouse              uint     `json:"qc_house"`
	HistoryValues        []string `json:"history_values"`
	UnparsedParts        string   `json:"unparsed_parts"`
	Source               string   `json:"source"`
	Qc                   uint     `json:"qc"`
}
