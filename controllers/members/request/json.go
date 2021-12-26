package request

type Members struct {
	ID uint
	// UUID      uuid.UUID
	UserID 			uint `json:"userId" valid:"required"`
	ExpiredDate		string `json:"expiredDate" valid:"required"`
}
