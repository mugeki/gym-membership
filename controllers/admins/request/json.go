package request

type Admins struct {
	Username     string `json:"username" valid:"required,minstringlength(6)"`
	Password     string `json:"password" valid:"required,minstringlength(6)"`
	Email        string `json:"email" valid:"required,email"`
	FullName     string `json:"fullname" valid:"required"`
	Gender       string `json:"gender" valid:"required"`
	Telephone    string `json:"telephone" valid:"required,numeric"`
	Address      string `json:"address" valid:"-"`
	UrlImage     string `json:"url_image" valid:"url"`
	IsSuperAdmin bool   `json:"is_super_admin" valid:"-"`
}

type AdminsUpdate struct {
	Username     string `json:"username" valid:"required,minstringlength(6)"`
	Password     string `json:"password,omitempty" valid:"-"`
	Email        string `json:"email" valid:"required,email"`
	FullName     string `json:"fullname" valid:"required"`
	Gender       string `json:"gender" valid:"required"`
	Telephone    string `json:"telephone" valid:"required,numeric"`
	Address      string `json:"address" valid:"-"`
	UrlImage     string `json:"url_image" valid:"url"`
	IsSuperAdmin bool   `json:"is_super_admin" valid:"-"`
}

type AdminsLogin struct {
	Username string `json:"username" valid:"required,minstringlength(6)"`
	Password string `json:"password" valid:"required,minstringlength(6)"`
}
