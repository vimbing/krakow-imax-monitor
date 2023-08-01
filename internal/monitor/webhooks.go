package monitor

import (
	"bytes"
	"encoding/json"
	"fmt"

	http "github.com/vimbing/fhttp"
)

func (m *Monitor) SendWebhook(entry MovieEntry) {
	for i := 0; i < 3; i++ {
		payload, err := json.Marshal(WebhookPayload{
			Content: nil,
			Embeds: []Embeds{
				{
					Title: fmt.Sprintf("Nowy seansik o id: %s", entry.Id),
					Color: nil,
					Fields: []Field{
						{
							Value:  fmt.Sprintf("[BILETY](%s)", entry.BookingLink),
							Name:   fmt.Sprintf("%s - %s", entry.Day, entry.Time),
							Inline: true,
						},
					},
				},
			},
			Attachments: []any{},
		})

		if err != nil {
			continue
		}

		// HEHE WEBHOOK IN PLAIN STRING YES I DO IT BECOUSE IM GIGACHAD PLS DON'T SPAM :(
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
