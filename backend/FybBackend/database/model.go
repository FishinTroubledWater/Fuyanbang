package database

type User struct {
	ID            int64 `gorm:"primaryKey"`
	Account       string
	Password      string
	PhoneNumber   string `gorm:"column:phoneNumber"`
	NickName      string `gorm:"column:nickName"`
	Major         string
	AvatarUrl     string `gorm:"column:avatarUrl"`
	Sex           string
	Area          string
	Year          int64
	TargetCollege string `gorm:"column:targetCollege"`
	TargetMajor   string `gorm:"column:targetMajor"`
	Slogan        string
	Balance       int64
	College       string
}

type Admin struct {
	ID          int64 `gorm:"primaryKey"`
	Account     string
	Password    string
	PhoneNumber string `gorm:"phoneNumber"`
	Token       string
}

type Academy struct {
	ID       int64 `gorm:"primaryKey"`
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
	ID          int64 `gorm:"primaryKey"`
	Author      string
	Title       string
	Content     string
	PublishTime string `gorm:"column:publishTime"`
	UserID      int64  `gorm:"column:userID"`
	Type        string `gorm:"column:type"`
}

type Major struct {
	ID                    int64  `gorm:"primaryKey;comment:专业id"`
	Name                  string `gorm:"column:_name;comment:专业名称"`
	Code                  string `gorm:"column:_code;comment:专业代码"`
	Profile               string `gorm:"comment:专业简介"`
	JobOrientation        string `gorm:"column:jobOrient;comment:专业就业方向"`
	JobProspect           string `gorm:"column:jobProspect;comment:专业就业前景"`
	SubjectCategory       string `gorm:"column:subjectCategory;comment:专业学科门类"`
	FirstLevelDiscipline  string `gorm:"column:firstLevelDiscipline;comment:专业一级学科"`
	SecondLevelDiscipline string `gorm:"column:secondLevelDiscipline;comment:专业二级学课"`
}

type Post struct {
}
