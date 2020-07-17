package domains

var TABLE_USERAFH = "user_afh"

type UserAfh struct {
	Id     uint `json:"id"`
	UserId uint `json:"userId" db:"user_id"`
	AfhId  uint `json:"afhId" db:"afh_id"`
}