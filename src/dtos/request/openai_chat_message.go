package request

type OpenAIChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
