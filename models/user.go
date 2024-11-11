package models

/**
Exporting:
- Capitalize names to export functions, types and variables.
- Lowercase names are private to the packages.
*/

type User struct {
	_id             int
	FirstName       string
	LastName        string
	UserName        string
	DisplayName     string
	MobileNumber    string
	PrimaryEmail    string
	SecondaryEmails []string
	Role            string
}

func CreateNewUser(_id int, FirstName string, LastName string, UserName string, DisplayName string, MobileNumber string, PrimaryEmail string, SecondaryEmails []string, role string) *User {
	return &User{
		_id:             _id,
		FirstName:       FirstName,
		LastName:        LastName,
		UserName:        UserName,
		DisplayName:     DisplayName,
		MobileNumber:    MobileNumber,
		PrimaryEmail:    PrimaryEmail,
		SecondaryEmails: SecondaryEmails,
		Role:            role,
	}
}
