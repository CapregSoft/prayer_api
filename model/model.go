package model

type ReqBody struct {
	Date           string  `json:"date"`
	Calculation    int     `json:"calculation"`
	Juristic       int     `json:"juristic"`
	HigherAltitude int     `json:"higherAltitude"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	TimeZone       int     `json:"timeZone"`
	TimeFormat     int     `json:"timeFormat"`
}
