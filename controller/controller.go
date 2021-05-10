package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CapregSoft/go-prayer-time/prayers"
	"github.com/CapregSoft/prayer_api/model"
	"github.com/labstack/echo/v4"
)

const (
	JAFARI  int = 0
	KARACHI int = 1 // University of Islamic Sciences, Karachi
	ISNA    int = 2 // Islamic Society of North America (ISNA)
	MWL     int = 3 // Muslim World League (MWL)
	MAKKAH  int = 4 // Umm al-Qura, Makkah
	EGYPT   int = 5 // Egyptian General Authority of Survey
	CUSTOM  int = 6 // Custom Setting
	TEHRAN  int = 7 // Institute of Geophysics, University of Tehran
)

// Juristic Methods
const (
	SHAFII int = 0 // Shafii (standard)
	HANAFI int = 1 // Hanafi
)

// Adjusting Methods for Higher Latitudes
const (
	NONE        int = 0 // No adjustment
	MID_NIGHT   int = 1 // middle of night
	ONE_SEVENTH int = 2 // 1/7th of night
	ANGLE_BASED int = 3 // angle/60th of night
)

// Time Formats
const (
	TIME_24    int = 0 // 24-hour format
	TIME_12    int = 1 // 12-hour format
	TIME_12_NS int = 2 // 12-hour format with no suffix
	FLOATING   int = 3 // floating point number
)

const (
	INVALID_TIME string = "----"
)

func isValidRequest(req *model.ReqBody) (map[string]string, bool) {

	errorDetail := make(map[string]string)

	if req.TimeFormat < 0 || req.TimeFormat > 3 {
		errorDetail["error_time_format"] = "INVALID TIME FORMAT"
	}

	if req.HigherAltitude < 0 || req.HigherAltitude > 3 {
		errorDetail["error_higher_altitude"] = "INVALID ALTITUDE METHOD"

	}

	if req.Juristic < 0 || req.Juristic > 1 {
		errorDetail["error_juristic"] = "INVALID JURITIC VALUE 0 FOR SHAAFI 1 FOR HANFI"

	}

	if req.Calculation < 0 || req.Calculation > 7 {
		errorDetail["error_juristic"] = "INVALID CALCULATION METHOD"

	}
	const (
		layoutISO = "2006-01-02"
	)

	_, err := time.Parse(layoutISO, req.Date)

	if err != nil {
		errorDetail["error_date"] = "Date must be in format YYYY-MM-DD  i.e 2006-01-02"
	}
	if len(errorDetail) != 0 {
		return errorDetail, false
	}

	return errorDetail, true
}

func GetPrayerTimePOST(c echo.Context) error {

	req := new(model.ReqBody)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request parameter"})
	}
	fmt.Print(req)
	if m, b := isValidRequest(req); !b {
		return c.JSON(http.StatusBadRequest, m)
	}

	p := prayers.New()
	const (
		layoutISO = "2006-01-02"
	)
	t, _ := time.Parse(layoutISO, req.Date)
	year, month, day := t.Date()
	p.TimeFormat = req.TimeFormat
	p.CalcMethod = req.Calculation
	p.AsrJuristic = req.Juristic
	p.AdjustHighLats = req.HigherAltitude

	prayerData := p.GetPrayerTimesAsObject(year, int(month), day, req.Latitude, req.Longitude, req.TimeZone)

	//fmt.Println(req.Date)
	return c.JSON(http.StatusOK, prayerData) //map[string]string{"OK": "success"})

}
func GetPrayerTime(c echo.Context) error {

	req := new(model.ReqBody)

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "internal server error"})
	}
	fmt.Println(req)
	p := prayers.New()
	const (
		layoutISO = "2006-01-02"
	)

	t, _ := time.Parse(layoutISO, req.Date)
	year, month, day := t.Date()
	p.TimeFormat = req.TimeFormat
	p.CalcMethod = req.Calculation
	p.AsrJuristic = req.Juristic
	p.AdjustHighLats = req.HigherAltitude
	prayerData := p.GetPrayerTimesAsObject(year, int(month), day, req.Latitude, req.Longitude, req.TimeZone)

	/*prayerData := p.GetPrayerTimesAsObject(2021, 5, 6, latitude, longitude, timezone)

	pray.TimeFormat = constants.TIME_12_NS
	pray.CalcMethod = constants.KARACHI
	pray.AsrJuristic = constants.HANAFI
	pray.AdjustHighLats = constants.ANGLE_BASED
	prayerData := pray.GetPrayerTimesAsObject(2021, 5, 6, latitude, longitude, timezone)
	*/
	//prayerTimes := pray.GetPrayerTimes(2021, 5, 6, latitude, longitude, timezone)

	return c.JSON(http.StatusOK, prayerData)
}
