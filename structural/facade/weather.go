package facade

type Weather struct {
	Coord struct {
		lon float32 `json:"lon"`
		lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temps     float32 `json:"temps"`
		FeelsLike float32 `json:"feels_like"`
		TempMin   float32 `json:"temp_min"`
		TempMax   float32 `json:"temp_max"`
		Pressure  float32 `json:"pressure"`
		Humidity  float32 `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float32 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     int    `json:"code"`
}
