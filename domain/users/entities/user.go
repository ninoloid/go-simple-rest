package users

type User struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"type:varchar(50)" json:"firstName"`
	LastName  string `gorm:"type:varchar(50)" json:"lastName"`
}
