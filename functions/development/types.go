package functions

type Request struct {
	ClientVersion  string `json:"clientVersion"`
	ClientId       string `json:"clientId"`
	Text           string `json:"text"`
	SourceLanguage string `json:"sourceLanguage"`
	TargetLanguage string `json:"targetLanguage"`
}

type Response struct {
	TaskId         string `json:"taskId"`
	TranslatedText string `json:"translatedText"`
	LoadCommand    string `json:"loadCommand"`
}

type TranslationTask struct {
	ClientVersion  string              `json:"clientVersion"`
	ClientId       string              `json:"clientId"`
	TaskId         string              `json:"taskId"`
	Text           string              `json:"text"`
	SourceLanguage string              `json:"sourceLanguage"`
	TargetLanguage string              `json:"targetLanguage"`
	TranslatedText string              `json:"translatedText"`
	TraceInfo      []map[string]string `json:"traceInfo"`
}
