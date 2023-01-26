package dao

type User struct {
	ID       int    `gorm:"column:id; primary_key; not null" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password;->:false" json:"-"`
	Status   int    `gorm:"column:status" json:"status"`
	RoleID   int    `gorm:"column:role_id;not null" json:"-"`
	Role     Role   `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	BaseModel
}
