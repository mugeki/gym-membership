package request

import (
	"github.com/google/uuid"
)

type Admins struct {
	ID        uint
	UUID      uuid.UUID
	Username  string `json:"username" valid:"required,minstringlength(6)"`
	Password  string `json:"password" valid:"required,minstringlength(6)"`
	Email     string `json:"email" valid:"required,email"`
	FullName  string `json:"fullname" valid:"required"`
	Gender    string `json:"gender" valid:"required"`
	Telephone string `json:"telephone" valid:"required,numeric"`
	Address   string `json:"address" valid:"-"`
}

type AdminsLogin struct {
	Username string `json:"username" valid:"required,minstringlength(6)"`
	Password string `json:"password" valid:"required,minstringlength(6)"`
}
