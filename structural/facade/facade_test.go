package facade

import "testing"

func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{APIKey: ""}
	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}
	if weather.ID != 420006353 {
		t.Errorf("Expected: 420006353\nActual: %d\n", weather.ID)
	}
}
