package openai

const (
	completionsPath = "/v1/chat/completions"
)

type textMessage struct {
	Role    string    `json:"role"`
	Content []content `json:"content"`
}

type TextCompletionInput struct {
	Prompt           string                `json:"prompt"`
	Images           []string              `json:"images"`
	ChatHistory      []*textMessage        `json:"chat-history,omitempty"`
	Model            string                `json:"model"`
	SystemMessage    *string               `json:"system-message,omitempty"`
	Temperature      *float32              `json:"temperature,omitempty"`
	TopP             *float32              `json:"top-p,omitempty"`
	N                *int                  `json:"n,omitempty"`
	Stop             *string               `json:"stop,omitempty"`
	MaxTokens        *int                  `json:"max-tokens,omitempty"`
	PresencePenalty  *float32              `json:"presence-penalty,omitempty"`
	FrequencyPenalty *float32              `json:"frequency-penalty,omitempty"`
	ResponseFormat   *responseFormatStruct `json:"response-format,omitempty"`
}

type responseFormatStruct struct {
	Type string `json:"type,omitempty"`
}

type TextCompletionOutput struct {
	Texts []string `json:"texts"`
	Usage usage    `json:"usage"`
}

type textCompletionReq struct {
	Model            string                `json:"model"`
	Messages         []interface{}         `json:"messages"`
	Temperature      *float32              `json:"temperature,omitempty"`
	TopP             *float32              `json:"top_p,omitempty"`
	N                *int                  `json:"n,omitempty"`
	Stop             *string               `json:"stop,omitempty"`
	MaxTokens        *int                  `json:"max_tokens,omitempty"`
	PresencePenalty  *float32              `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float32              `json:"frequency_penalty,omitempty"`
	ResponseFormat   *responseFormatStruct `json:"response_format,omitempty"`
}

type multiModalMessage struct {
	Role    string    `json:"role"`
	Content []content `json:"content"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type imageURL struct {
	URL string `json:"url"`
}

type content struct {
	Type     string    `json:"type"`
	Text     *string   `json:"text,omitempty"`
	ImageURL *imageURL `json:"image_url,omitempty"`
}

type textCompletionResp struct {
	ID      string      `json:"id"`
	Object  string      `json:"object"`
	Created int         `json:"created"`
	Choices []choices   `json:"choices"`
	Usage   usageOpenAI `json:"usage"`
}

type outputMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type choices struct {
	Index        int           `json:"index"`
	FinishReason string        `json:"finish_reason"`
	Message      outputMessage `json:"message"`
}

type usageOpenAI struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type usage struct {
	PromptTokens     int `json:"prompt-tokens"`
	CompletionTokens int `json:"completion-tokens"`
	TotalTokens      int `json:"total-tokens"`
}
