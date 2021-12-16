package classifications

type Classifications struct {
	ID   int    `gorm:"primarykey"`
	Name string `gorm:"column:name"`
}