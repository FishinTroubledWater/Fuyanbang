package database

type User struct {
	ID            int `gorm:"primaryKey"`
	Account       string
	Password      string
	PhoneNumber   string
	NickName      string
	Major         string
	AvatarUrl     string
	Sex           string
	Area          string
	Year          int64
	TargetCollege string
	TargetMajor   string
	Slogan        string
	Balance       int
	College       string
}

type Admin struct {
	ID          int `gorm:"primaryKey"`
	Account     string
	Password    string
	PhoneNumber string
}

type Academy struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Code     string
	Guide    string
	Type     string
	Belong   string
	Logo     string
	Level    string
	LineType string
	Profile  string
	Region   string
}

type News struct {
	ID          int `gorm:"primaryKey"`
	Author      string
	Title       string
	Content     string
	PublishTime string `gorm:"column:publishTime"`
	UserID      int    `gorm:"column:userID"`
	Type        string `gorm:"column:type"`
}
