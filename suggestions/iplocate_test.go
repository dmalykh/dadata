package suggestions

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuggestions_Iplocate(t *testing.T) {

	var r = getTestInstance()
	var ctx = context.Background()

	type testCase struct {
		Ip       string
		Need     string
		Response string
	}

	var cases = []testCase{
		{
			Ip:       "46.226.227.20",
			Need:     "г Краснодар",
			Response: `{ "location": { "value": "г Краснодар", "unrestricted_value": "350000, Краснодарский край, г Краснодар", "data": { "postal_code": "350000", "country": "Россия", "country_iso_code": "RU", "federal_district": "Южный", "region_fias_id": "d00e1013-16bd-4c09-b3d5-3cb09fc54bd8", "region_kladr_id": "2300000000000", "region_iso_code": "RU-KDA", "region_with_type": "Краснодарский край", "region_type": "край", "region_type_full": "край", "region": "Краснодарский", "area_fias_id": null, "area_kladr_id": null, "area_with_type": null, "area_type": null, "area_type_full": null, "area": null, "city_fias_id": "7dfa745e-aa19-4688-b121-b655c11e482f", "city_kladr_id": "2300000100000", "city_with_type": "г Краснодар", "city_type": "г", "city_type_full": "город", "city": "Краснодар", "city_area": null, "city_district_fias_id": null, "city_district_kladr_id": null, "city_district_with_type": null, "city_district_type": null, "city_district_type_full": null, "city_district": null, "settlement_fias_id": null, "settlement_kladr_id": null, "settlement_with_type": null, "settlement_type": null, "settlement_type_full": null, "settlement": null, "street_fias_id": null, "street_kladr_id": null, "street_with_type": null, "street_type": null, "street_type_full": null, "street": null, "house_fias_id": null, "house_kladr_id": null, "house_type": null, "house_type_full": null, "house": null, "block_type": null, "block_type_full": null, "block": null, "flat_type": null, "flat_type_full": null, "flat": null, "flat_area": null, "square_meter_price": null, "flat_price": null, "postal_box": null, "fias_id": "7dfa745e-aa19-4688-b121-b655c11e482f", "fias_code": "23000001000000000000000", "fias_level": "4", "fias_actuality_state": "0", "kladr_id": "2300000100000", "geoname_id": "542420", "capital_marker": "2", "okato": "03401000000", "oktmo": "03701000001", "tax_office": "2300", "tax_office_legal": "2300", "timezone": null, "geo_lat": "45.0401604", "geo_lon": "38.9759647", "beltway_hit": null, "beltway_distance": null, "metro": null, "qc_geo": "4", "qc_complete": null, "qc_house": null, "history_values": null, "unparsed_parts": null, "source": null, "qc": null } } }`,
		},
		{
			Ip:       "185.174.128.235",
			Need:     "г Москва",
			Response: `{ "location": { "value": "г Москва", "unrestricted_value": "101000, г Москва", "data": { "postal_code": "101000", "country": "Россия", "country_iso_code": "RU", "federal_district": "Центральный", "region_fias_id": "0c5b2444-70a0-4932-980c-b4dc0d3f02b5", "region_kladr_id": "7700000000000", "region_iso_code": "RU-MOW", "region_with_type": "г Москва", "region_type": "г", "region_type_full": "город", "region": "Москва", "area_fias_id": null, "area_kladr_id": null, "area_with_type": null, "area_type": null, "area_type_full": null, "area": null, "city_fias_id": "0c5b2444-70a0-4932-980c-b4dc0d3f02b5", "city_kladr_id": "7700000000000", "city_with_type": "г Москва", "city_type": "г", "city_type_full": "город", "city": "Москва", "city_area": null, "city_district_fias_id": null, "city_district_kladr_id": null, "city_district_with_type": null, "city_district_type": null, "city_district_type_full": null, "city_district": null, "settlement_fias_id": null, "settlement_kladr_id": null, "settlement_with_type": null, "settlement_type": null, "settlement_type_full": null, "settlement": null, "street_fias_id": null, "street_kladr_id": null, "street_with_type": null, "street_type": null, "street_type_full": null, "street": null, "house_fias_id": null, "house_kladr_id": null, "house_type": null, "house_type_full": null, "house": null, "block_type": null, "block_type_full": null, "block": null, "flat_type": null, "flat_type_full": null, "flat": null, "flat_area": null, "square_meter_price": null, "flat_price": null, "postal_box": null, "fias_id": "0c5b2444-70a0-4932-980c-b4dc0d3f02b5", "fias_code": "77000000000000000000000", "fias_level": "1", "fias_actuality_state": "0", "kladr_id": "7700000000000", "geoname_id": "524901", "capital_marker": "0", "okato": "45000000000", "oktmo": "45000000", "tax_office": "7700", "tax_office_legal": "7700", "timezone": null, "geo_lat": "55.7540471", "geo_lon": "37.620405", "beltway_hit": null, "beltway_distance": null, "metro": null, "qc_geo": "4", "qc_complete": null, "qc_house": null, "history_values": null, "unparsed_parts": null, "source": null, "qc": null } } }`,
		},
		{
			Ip:       "897.44.32.44",
			Need:     "",
			Response: `{"location":null}`,
		},
	}

	for _, c := range cases {

		var server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/suggestions/iplocate/address" {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(c.Response))
			}
		}))

		DADATA_SUGGESTIONS_URL = server.URL + "/suggestions/%s"

		item, err := r.Iplocate(ctx, c.Ip)
		if err != nil {
			t.Errorf(`Can't receive "%s": %s`, c.Ip, err.Error())
		}
		if item.Location.Value != c.Need {
			t.Errorf(`Waiting for "%s", got "%s".`, item.Location.Value, c.Need)
		}

		server.Close()
	}
}
