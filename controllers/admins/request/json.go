package request

type Admins struct {
	Username     string `json:"username" valid:"required,minstringlength(6)"`
	Password     string `json:"password" valid:"required,minstringlength(6)"`
	Email        string `json:"email" valid:"required,email"`
	FullName     string `json:"fullname" valid:"required"`
	Gender       string `json:"gender" valid:"required"`
	Telephone    string `json:"telephone" valid:"required,numeric"`
	Address      string `json:"address" valid:"-"`
	UrlImage     string
	IsSuperAdmin bool `json:"is_super_admin" valid:"-"`
}

type AdminsLogin struct {
	Username string `json:"username" valid:"required,minstringlength(6)"`
	Password string `json:"password" valid:"required,minstringlength(6)"`
}
