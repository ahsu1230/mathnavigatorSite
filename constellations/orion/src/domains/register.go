package domains

type RegisterBody struct {
	Student  User `json:"student"`
	Guardian User `json:"guardian"`
}
