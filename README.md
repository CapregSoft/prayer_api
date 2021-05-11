# prayer_api
Golang API for the prayers times.

## [Prayer_api](/)

Golang API for the calculating the prayers time.

	go get github.com/CapregSoft/prayer_api
 ## Run
	 go run .\starter\main.go
	 
 ## Examples
```
	salah := server.New()
	salah.Start(":8080")
```
**POST /api/prayertime/**

* **Request Body**
```
{
	"date":string,
	"calculation": int,
	"juristic": int,
	"higherAltitude": int,
	"timeFormat": int,
	"latitude":float64,
	"longitude":float64,
	"timeZone":int
}
```
timeFormat 

	0 = TIME_24        // 12-hour format
	1 = TIME_12        // 24-hour format
	2 = TIME_12_NS // 12-hour format with no suffix
	3 = FLOATING    // floating point number

calculation

	JAFARI  = 0
	KARACHI = 1  // University of Islamic Sciences, Karachi
	ISNA    = 2  // Islamic Society of North America (ISNA)
	MWL     = 3  // Muslim World League (MWL)
	MAKKAH  = 4  // Umm al-Qura, Makkah
	EGYPT   = 5  // Egyptian General Authority of Survey
	CUSTOM  = 6  // Custom Setting
	TEHRAN  = 7  // Institute of Geophysics, University of Tehran
  
juristic:

	SHAFII int = 0  // Shafii (standard)
	HANAFI int = 1  // Hanafi

higherAltitude:

		NONE 		= 0  // No adjustment
		MID_NIGHT   = 1  // middle of night
		ONE_SEVENTH = 2  // 1/7th of night
		ANGLE_BASED = 3  // angle/60th of night

* **Success Response:**

  * **Code:** 200
   * **Response Body:**
    ```
	 {
	    "fajar": string,
	    "sunrise": string,
	    "dhuhr": string,
	    "asr": string,
	    "sunset": string,
	    "maghrib": string,
	    "isha": string
	}
    ```
    
 
* **Error Response:**

  * **Code:** 400 Bad Request<br />
    **Content:** `{ error : "" }`