package libtaskido

// Task structure corresponds to the JSON object
type Task struct {
	ID            int      `json:"id"`
	UUID          string   `json:"uuid"`
	Subject       string   `json:"subject"`
	Projects      []string `json:"projects"`
	Contexts      []string `json:"contexts"`
	Due           string   `json:"due"`
	Completed     bool     `json:"completed"`
	CompletedDate string   `json:"completedDate"`
	Archived      bool     `json:"archived"`
	Priority      int      `json:"priority"`
	CreatedAt     string   `json:"createdAt`
	UpdatedAt     string   `json:"createdAt`
}