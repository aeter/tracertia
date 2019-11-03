package ipdb

import "testing"

func TestGetCountry(t *testing.T) {
	github_ip_address_in_usa := "140.82.118.4"
	if GetCountry(github_ip_address_in_usa) != "US" {
		t.Errorf("Incorrect IP address to country lookup.")
	}
}
