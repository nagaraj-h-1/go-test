package structs

type Intent struct {
	StartEpoch  float64 `json:"start_epoch"`
	Duration    float64 `json:"duration_s"`
	IntentType  string  `json:"intent_type"`
	TriggerType string  `json:"trigger_type"`
}

type Event struct {
	EventID   int    `json:"event_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	EventType string `json:"event_type"`
	Intents   Intent `json:"intents"`
}

type GitHubContent struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	DownloadURL string `json:"download_url"`
}
