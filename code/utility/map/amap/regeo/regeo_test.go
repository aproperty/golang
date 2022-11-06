package regeo

import (
	"testing"
)

func TestClient(t *testing.T) {

	address, province, city, err := GetReGeoByAddress("116.310003,39.991957")
	if err != nil {
		t.Errorf("GetReGeoByAddress error = %v", err)
	}

	t.Logf("%v", address)
	t.Logf("%v", province)
	t.Logf("%v", city)

}
