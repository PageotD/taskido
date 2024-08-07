package libtaskido

// Task structure corresponds to the JSON object
type Task struct {
	ID            int      `json:"id"`
	UUID          string   `json:"uuid"`
	Description   string   `json:"description"`
	Projects      []string `json:"projects"`
	Contexts      []string `json:"contexts"`
	Due           string   `json:"due"`
	Status        string   `json:"status"`
	Priority      int      `json:"priority"`
	CreatedAt     string   `json:"createdAt`
	UpdatedAt     string   `json:"createdAt`
}