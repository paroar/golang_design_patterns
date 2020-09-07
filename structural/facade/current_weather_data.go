package facade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type CurrentWeatherData struct {
	APIKey string
}

func (c *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (weather *Weather, err error) {
	return c.doRequest(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&APPID=%s", lat, lon, c.APIKey))
}

func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *Weather, err error) {
	return c.doRequest(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weatherq=%s,%s&APPID=%s", city, countryCode, c.APIKey))
}

func (c *CurrentWeatherData) doRequest(uri string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		byt, errMsg := ioutil.ReadAll(resp.Body)
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("Status code was %d, aborting. Error message was:\n%s\n", resp.StatusCode, errMsg)
		return
	}
	weather, err = c.responseParser(resp.Body)
	resp.Body.Close()
	return
}

func getMockData() io.Reader {
	response := `{
		"coord": {
		  "lon": -122.08,
		  "lat": 37.39
		},
		"weather": [
		  {
			"id": 800,
			"main": "Clear",
			"description": "clear sky",
			"icon": "01d"
		  }
		],
		"base": "stations",
		"main": {
		  "temp": 282.55,
		  "feels_like": 281.86,
		  "temp_min": 280.37,
		  "temp_max": 284.26,
		  "pressure": 1023,
		  "humidity": 100
		},
		"visibility": 16093,
		"wind": {
		  "speed": 1.5,
		  "deg": 350
		},
		"clouds": {
		  "all": 1
		},
		"dt": 1560350645,
		"sys": {
		  "type": 1,
		  "id": 5122,
		  "message": 0.0139,
		  "country": "US",
		  "sunrise": 1560343627,
		  "sunset": 1560396563
		},
		"timezone": -25200,
		"id": 420006353,
		"name": "Mountain View",
		"cod": 200
	  }`
	r := bytes.NewReader([]byte(response))
	return r
}
