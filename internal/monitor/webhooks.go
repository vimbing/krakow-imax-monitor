package monitor

import (
	"bytes"
	"encoding/json"
	"fmt"

	http "github.com/vimbing/fhttp"
)

func (m *Monitor) SendWebhook(entries []MovieEntry) {
	for i := 0; i < 3; i++ {
		fields := []Field{}

		for _, e := range entries {
			fields = append(fields, Field{
				Name:   e.Time,
				Value:  fmt.Sprintf("[BILETY](%s)", e.BookingLink),
				Inline: true,
			})
		}

		payload, err := json.Marshal(WebhookPayload{
			Content: nil,
			Embeds: []Embeds{
				{
					Title:  fmt.Sprintf("Nowy dzień znaleziony esssa: %s", entries[0].Day),
					Color:  nil,
					Fields: fields,
				},
			},
			Attachments: []any{},
		})

		if err != nil {
			continue
		}

		req, err := http.NewRequest("POST", "https://discord.com/api/webhooks/1135879696815493200/_EJU3Q-oQizXa5iPA62F_wVqVsGniHVEVF2gDjuNQw5p2k9cUOWaG97WXDDZW6Rf_IVM", bytes.NewBuffer(payload))

		if err != nil {
			continue
		}

		req.Header = http.Header{
			"content-type": {"application/json"},
		}

		res, err := m.Client.Do(req)

		if err != nil {
			continue
		}

		if res.StatusCode == 204 {
			return
		}
	}
}
