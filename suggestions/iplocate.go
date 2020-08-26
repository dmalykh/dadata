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
		Data              AddressItemData `json:"data"`
	} `json:"location"`
}
