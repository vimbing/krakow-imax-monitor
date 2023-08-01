package monitor

import (
	"cinemacity/internal/utils"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	http "github.com/vimbing/fhttp"
	"golang.org/x/exp/slices"
)

func waitTillNextDay() {
	lastDay := time.Now().Day()

	for {
		time.Sleep(time.Second * 1)
		if time.Now().Day() != lastDay {
			return
		}
	}
}

func handleDate(date string) string {
	if strings.HasPrefix(date, "0") {
		return date
	}

	return fmt.Sprintf("0%s", date)
}

func (m *Monitor) GetData(date string) ([]MovieEntry, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://www.cinema-city.pl/pl/data-api-service/v1/quickbook/10103/cinema-events/in-group/krakow/with-film/5297s2r/at-date/2023-08-%s?attr=&lang=pl_PL", handleDate(date)),
		nil,
	)

	if err != nil {
		return []MovieEntry{}, err
	}

	req.Header = http.Header{
		"authority":          {"www.cinema-city.pl"},
		"accept":             {"application/json;charset=utf-8"},
		"accept-language":    {"pl-PL,pl;q=0.9,en-US;q=0.8,en;q=0.7,la;q=0.6,de;q=0.5"},
		"dnt":                {"1"},
		"sec-ch-ua":          {"\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\""},
		"sec-ch-ua-mobile":   {"?0"},
		"sec-ch-ua-platform": {"\"Linux\""},
		"sec-fetch-dest":     {"empty"},
		"sec-fetch-mode":     {"cors"},
		"sec-fetch-site":     {"same-origin"},
		"user-agent":         {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"},
	}

	res, err := m.Client.Do(req)

	if err != nil {
		return []MovieEntry{}, err
	}

	body, err := utils.GetResponseBody(res)

	if err != nil {
		return []MovieEntry{}, err
	}

	var resBody GetBodyResponse
	json.Unmarshal([]byte(body), &resBody)

	if len(resBody.Body.Events) == 0 {
		fmt.Println("No events for this day...")

		return []MovieEntry{}, nil
	}

	entries := make([]MovieEntry, 0)

	for _, e := range resBody.Body.Events {
		if strings.EqualFold(e.FilmID, OPPENHIMER_ID) && slices.Contains(e.AttributeIds, "imax") {
			timeString := "UNKNOWN"
			time := strings.Split(e.EventDateTime, "T")

			if len(time) >= 2 {
				timeString = time[1]
			}

			entries = append(entries, MovieEntry{
				Id:          e.ID,
				Day:         time[0],
				Time:        timeString,
				BookingLink: e.BookingLink,
			})
		}
	}

	return entries, nil
}

func (m *Monitor) Monitor() {
	for {
		time.Sleep(time.Second * 2)

		data, err := m.GetData(fmt.Sprint(time.Now().Day() + 3))

		if err != nil {
			time.Sleep(time.Minute * 1)
			continue
		}

		if len(data) == 0 {
			continue
		}

		m.SendWebhook(data)
		waitTillNextDay()
	}
}
