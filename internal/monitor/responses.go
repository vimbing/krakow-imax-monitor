package monitor

type GetBodyResponse struct {
	Body struct {
		Cinemas []struct {
			ID          string `json:"id"`
			GroupID     string `json:"groupId"`
			DisplayName string `json:"displayName"`
			Link        string `json:"link"`
			ImageURL    string `json:"imageUrl"`
			Address     string `json:"address"`
			AddressInfo struct {
				Address1   string `json:"address1"`
				Address2   any    `json:"address2"`
				Address3   any    `json:"address3"`
				Address4   any    `json:"address4"`
				City       string `json:"city"`
				State      any    `json:"state"`
				PostalCode string `json:"postalCode"`
			} `json:"addressInfo"`
			BookingURL            string   `json:"bookingUrl"`
			BlockOnlineSales      bool     `json:"blockOnlineSales"`
			BlockOnlineSalesUntil any      `json:"blockOnlineSalesUntil"`
			Latitude              float64  `json:"latitude"`
			Longitude             float64  `json:"longitude"`
			AttributeIds          []string `json:"attributeIds"`
		} `json:"cinemas"`
		Events []struct {
			ID                   string   `json:"id"`
			FilmID               string   `json:"filmId"`
			CinemaID             string   `json:"cinemaId"`
			BusinessDay          string   `json:"businessDay"`
			EventDateTime        string   `json:"eventDateTime"`
			AttributeIds         []string `json:"attributeIds"`
			BookingLink          string   `json:"bookingLink"`
			CompositeBookingLink struct {
				Type       string `json:"type"`
				BookingURL struct {
					URL    string `json:"url"`
					Params struct {
						Lang string `json:"lang"`
						Key  string `json:"key"`
					} `json:"params"`
				} `json:"bookingUrl"`
				ObsoleteBookingURL    string `json:"obsoleteBookingUrl"`
				BlockOnlineSales      bool   `json:"blockOnlineSales"`
				BlockOnlineSalesUntil any    `json:"blockOnlineSalesUntil"`
				ServiceURL            string `json:"serviceUrl"`
			} `json:"compositeBookingLink"`
			PresentationCode   string `json:"presentationCode"`
			SoldOut            bool   `json:"soldOut"`
			Auditorium         string `json:"auditorium"`
			AuditoriumTinyName string `json:"auditoriumTinyName"`
		} `json:"events"`
	} `json:"body"`
}
