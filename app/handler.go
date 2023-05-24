package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
	"udemySomeTimeInc/model"
)

func getTime(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	tzParam := request.URL.Query().Get("tz")
	var response interface{}

	if tzParam == "" {
		localTime := time.Now().Format("2006-01-02 15:04:05 -0700 MST")

		currentTime := model.UserLocation{CurrentTime: localTime}

		response = currentTime
	} else {
		timeZones := strings.Split(tzParam, ",")

		data := make(map[string]string, len(timeZones))

		for _, tz := range timeZones {
			location, err := time.LoadLocation(tz)
			if err != nil {
				http.Error(writer, "invalid timezone", http.StatusNotFound)
				return
			}

			userLocationWithTimeZone := time.Now().In(location).Format("2006-01-02 15:04:05 -0700 MST")

			data[tz] = userLocationWithTimeZone
		}
		response = data
	}
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error in GetTime: %v\n", err)
	}
}
