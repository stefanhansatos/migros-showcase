package functions

type Request struct {
	ClientVersion  string `json:"clientVersion"`
	ClientId       string `json:"clientId"`
	Text           string `json:"text"`
	SourceLanguage string `json:"sourceLanguage"`
	TargetLanguage string `json:"targetLanguage"`
}

type Response struct {
	TaskId         string   `json:"taskId"`
	TranslatedText string   `json:"translatedText"`
	LoadCommands   []string `json:"loadCommands"`
}

type TranslationTask struct {
	ClientVersion  string `json:"clientVersion"`
	ClientId       string `json:"clientId"`
	TaskId         string `json:"taskId"`
	Text           string `json:"text"`
	SourceLanguage string `json:"sourceLanguage"`
	TargetLanguage string `json:"targetLanguage"`
	TranslatedText string `json:"translatedText"`
	//TraceInfo      []map[string]string `json:"traceInfo"`
}

// Old types for conversion

type TranslationQuery struct {
	Uuid           string `json:"uuid"`
	Text           string `json:"text"`
	SourceLanguage string `json:"sourceLanguage"`
	TargetLanguage string `json:"targetLanguage"`
}

type Translation struct {
	TranslationQuery  *TranslationQuery `json:"translationQuery"`
	TranslatedText    string            `json:"translatedText"`
	TranslationErrors []string          `json:"translationErrors"`
}

type WrappedData struct {
	Source      string       `json:"source"`
	Translation *Translation `json:"translation"`
	LogFilter   string       `json:"logFilter"`
	Timestamp   string       `json:"timestamp,omitempty"`
	Unix        int64        `json:"unix,omitempty"` // Unix time in seconds
}
