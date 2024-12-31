package main

import (
	"fmt"
	"time"
)

func main() {
	//in go lang the time package have some special out of the box things like
	// the format is 2006-01-02 2006 jan 2 yyyy=2006 ,month=jan, day=02  is the standard of date formatting
	// like day the time has also the same
	// for 24h time the format is 15:04:05  15 is hour, 04 is muntin and 05 is second
	// and for 12h timing 3:04 PM is the format and PM id fix ....

	currentTime := time.Now()

	fmt.Print(currentTime)                           //2024-12-31 19:43:01.3671232 +0530 IST m=+0.001738001
	fmt.Printf("type of time is %T \n", currentTime) // time.Time  the data type Time is provided by the time package

	//lets format the  date and time

	formatted := currentTime.Format("2006-01-02 ,Monday, 3:04:05 PM") //2024-12-31 ,Tuesday, 7:54:09
	fmt.Println(formatted)                                            //12-31-2024    "2006-01-02" == 2024-12-31

	//convert the string to date format

	day := "2025-06-06" //this is a type string

	//the layout string is must be type of how the data comes like if dd-mm-yyyy is the date then the layout string must be same
	layoutString := "2006-01-02"

	convertDate, _ := time.Parse(layoutString, day) // time.parse return two things error and value

	fmt.Println(convertDate.Format("2006-01-02"))

	//add days in current date
	new_Date := currentTime.Add(24 * time.Hour)

	fmt.Print(new_Date.Format("2006-01-02,Monday,15:04:05"))

}
