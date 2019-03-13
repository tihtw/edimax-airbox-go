package airbox

type DeviceResponse struct {
	Status  string   `json:"status"`
	Devices []Device `json:"devices"`
}

type Device struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	Pm25      float64 `json:"pm25"`
	Pm10      float64 `json:"pm10"`
	Pm1       float64 `json:"pm1"`
	Co2       float64 `json:"co2"`
	Hcho      float64 `json:"hcho"`
	Tvoc      float64 `json:"tvoc"`
	Co        float64 `json:"co"`
	T         float64 `json:"t"`
	H         float64 `json:"h"`
	Time      float64 `json:"time"`
	UtcTime   string  `json:"utc_time"`
	Org       string  `json:"org"`
	Area      string  `json:"area"`
	Type      string  `json:"type"`
	Odm       string  `json:"odm"`
	Status    string  `json:"status"`
	AdfStatus float64 `json:"adf_status"`
}
