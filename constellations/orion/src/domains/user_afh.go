package domains

var TABLE_USERAFH = "user_afh"

type UserAFH struct {
	Id     uint `json:"id"`
	UserId uint `json:"id"`
	AfhId  uint `json:"id"`
}

func (userAFH *UserAFH) Validate() error {
	id := userAFH.Id
	userId := userAFH.UserId
	afhId := userAFH.AfhId

}
