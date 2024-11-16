package request

type OpenAIChatHistory struct {
	Messages []OpenAIChatMessage `json:"messages"`
}
