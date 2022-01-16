package response

type Admins struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	FullName     string `json:"fullname"`
	Gender       string `json:"gender"`
	Telephone    string `json:"telephone"`
	Address      string `json:"address"`
	UrlImage     string `json:"url_image"`
	Token        string `json:"token"`
	IsSuperAdmin bool   `json:"is_super_admin"`
}