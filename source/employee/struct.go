package employee

// Struct of Employeee + JSON
type Employee struct {
	FirstName string            `json:"FirstName"`
	LastName  string            `json:"LastName"`
	Age       int               `json:"Age"`
	Gender    string            `json:"Gender"`
	Skill     []string          `json:"Skill"`
	Rank      string            `json:"Rank"`
	Private   map[string]string `json:"Private"`
	ID        string            `json:"ID"`
}
