package responsebody

type Contact struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	IsVip     bool   `json:"isVip"`
}

type Contacts struct {
	PageNo   uint      `json:"pageNo"`
	PageSize uint      `json:"pageSize"`
	Contacts []Contact `json:"contacts"`
}
