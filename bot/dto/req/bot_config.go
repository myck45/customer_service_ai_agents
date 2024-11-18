package req

type BotConfig struct {
	BotName         string `json:"bot_name"`
	BotIdentity     string `json:"bot_identity"`
	SemanticContext string `json:"semantic_context"`
}
