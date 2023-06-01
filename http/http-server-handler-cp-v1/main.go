package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now()
		fmt.Println(now)
		hari := ""
		switch now.Weekday() {
		case time.Sunday:
			hari = "Sunday"
		case time.Monday:
			hari = "Monday"
		case time.Tuesday:
			hari = "Tuesday"
		case time.Wednesday:
			hari = "Wednesday"
		case time.Thursday:
			hari = "Thursday"
		case time.Friday:
			hari = "Friday"
		case time.Saturday:
			hari = "Saturday"
		}
		fmt.Println(hari)
		day := now.Day()
		month := now.Month()
		year := now.Year()
		fmt.Println(day, month, year)
		result := fmt.Sprintf("%s, %d %s %d", hari, day, month, year)
		writer.Write([]byte(result))
	} // TODO: replace this
}

func main() {
	err := http.ListenAndServe("localhost:8080", GetHandler())
	if err != nil {
		fmt.Println(err)
	}
}
