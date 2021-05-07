package model

type User struct {
	ID          uint   `gorm:"PRIMARY_KEY" json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Description string `json:"description"`
}

type Profile struct {
	ID         uint   `gorm:"PRIMARY_KEY" json:"id"`
	ImaProfile string `json:"ima_profile"`
	UserID     uint   `json:"user_id"`
	User       User   `gorm:"foreignkey:UserID" json:"user"`
}

func (User) TableName() string {
	return "users"
}

func (Profile) TableName() string {
	return "profiles"
}
