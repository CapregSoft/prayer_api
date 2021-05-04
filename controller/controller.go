package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CapregSoft/go-prayer-time/prayers"
	"github.com/CapregSoft/prayer-api/model"
	"github.com/labstack/echo/v4"
)

func GetPrayerTimePOST(c echo.Context) error {

	req := new(model.ReqBody)

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "internal server error"})
	}
	fmt.Println(req.Date)
	return c.JSON(http.StatusOK, req) //map[string]string{"OK": "success"})

}
func GetPrayerTime(c echo.Context) error {

	req := new(model.ReqBody)

	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "internal server error"})
	}
	fmt.Println(req)
	//myDate := time.Now()
	//year, month, day := myDate.Date()
	p := prayers.New()
	//latitude := 33.57368163412395

	//longitude := 73.17308661244054
	//timezone := 5
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)

	t, _ := time.Parse(layoutISO, req.Date)
	year, month, day := t.Date()
	fmt.Println(t)

	prayerData := p.GetPrayerTimesAsObject(year, int(month), day, req.Latitude, req.Longitude, req.Timezone)

	/*
		pray.TimeFormat = constants.TIME_12_NS
		pray.CalcMethod = constants.KARACHI
		pray.AsrJuristic = constants.HANAFI
		pray.AdjustHighLats = constants.ANGLE_BASED
		prayerData := pray.GetPrayerTimesAsObject(2021, 5, 6, latitude, longitude, timezone)
	*/
	//prayerTimes := pray.GetPrayerTimes(2021, 5, 6, latitude, longitude, timezone)

	return c.JSON(http.StatusOK, prayerData)
}
