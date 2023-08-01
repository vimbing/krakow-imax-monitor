package monitor

type WebhookPayload struct {
	Content     any      `json:"content"`
	Embeds      []Embeds `json:"embeds"`
	Attachments []any    `json:"attachments"`
}
type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
type Embeds struct {
	Title  string  `json:"title"`
	Color  any     `json:"color"`
	Fields []Field `json:"fields"`
}
