package responsebody

type Contact struct {
	ID        uint
	FirstName string
	LastName  string
	Phone     string
	IsVip     bool
}

type Contacts struct {
	PageNo   uint
	PageSize uint
	Contacts []Contact
}
