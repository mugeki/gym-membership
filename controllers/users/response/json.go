package response

type Users struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FullName  string `json:"fullname"`
	Gender    string `json:"gender"`
	Telephone string `json:"telephone"`
	Address   string `json:"address"`
	Token     string `json:"token"`
}