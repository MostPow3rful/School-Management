package student

// Struct Of Student + JSON
type Student struct {
	FirstName string            `json:"FirstName"`
	LastName  string            `json:"LastName"`
	Age       int               `json:"Age"`
	Gender    string            `json:"Gender"`
	Skill     []string          `json:"Skill"`
	Score     map[string]string `json:"Score"`
	Private   map[string]string `json:"Private"`
	ID        string            `json:"ID"`
}

