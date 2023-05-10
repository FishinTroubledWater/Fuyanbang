package database

import "time"

type User struct {
	ID            int64 `gorm:"column:ID;primaryKey"`
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
	Role          string
	State         int64     `gorm:"column:state"`
	registerTime  time.Time `gorm:"column:registerTime"`
}

type Admin struct {
	ID          int64 `gorm:"column:ID;primaryKey"`
	Account     string
	Password    string
	PhoneNumber string `gorm:"phoneNumber"`
	Token       string
}

type Academy struct {
	ID       int64 `gorm:"column:ID;primaryKey"`
	Name     string
	Code     string
	Guide    string
	Type     string
	Belong   string
	Logo     string
	Level    string
	LineType string `gorm:"column:lineType"`
	Profile  string
	Region   string
}

type News struct {
	ID          int64 `gorm:"column:ID;primaryKey"`
	Author      string
	Title       string
	Content     string
	PublishTime string `gorm:"column:publishTime"`
	UserID      int64  `gorm:"column:userID"`
	Type        string `gorm:"column:type"`
}

type Major struct {
	ID                    int64  `gorm:"primaryKey;comment:专业id"`
	Name                  string `gorm:"column:name;comment:专业名称"`
	Code                  string `gorm:"column:code;comment:专业代码"`
	Profile               string `gorm:"column:profile;comment:专业简介"`
	MajorDirection        string `gorm:"column:majorDirection;comment:专业培养方向"`
	JobDirection          string `gorm:"column:jobDirection;comment:专业就业方向"`
	JobProspect           string `gorm:"column:jobProspect;comment:专业就业前景"`
	SubjectCategory       string `gorm:"column:subjectCategory;comment:专业学科门类"`
	FirstLevelDiscipline  string `gorm:"column:firstLevelDiscipline;comment:专业一级学科"`
	SecondLevelDiscipline string `gorm:"column:secondLevelDiscipline;comment:专业二级学课"`
	ScoreUrl              string `gorm:"column:scoreUrl;comment:考研分数线url"`
	ForeignType           string `gorm:"column:foreignType"`
	MathType              string `gorm:"column:mathType"`
}
type Part struct {
	ID       int64  `gorm:"column:ID;primaryKey"`
	PartName string `gorm:"column:partName"`
}
type Post struct {
	ID          int64     `gorm:"column:ID;primaryKey"`
	Summary     string    `gorm:"column:summary"`
	PartID      int64     `gorm:"column:partID"`
	AuthorID    int64     `gorm:"column:authorID"`
	Favorite    int64     `gorm:"column:favorite"`
	Like        int64     `gorm:"column:like"`
	PublishTime time.Time `gorm:"column:publishTime"`
	State       string    `gorm:"column:state"`

	Author   User      `gorm:"foreignKey:AuthorID"`
	PostImgs []PostImg `gorm:"foreignKey:PostID"`
}

type PostImg struct {
	PostID int64  `gorm:"primaryKey"`
	Img    string `gorm:"primaryKey"`
}

type Recipe struct {
	ID          int64     `gorm:"column:ID;primaryKey"`
	Author      string    `gorm:"column:author"`
	Title       string    `gorm:"column:title"`
	Content     string    `gorm:"column:content"`
	Favorite    int64     `gorm:"column:favorite"`
	Like        int64     `gorm:"column:like"`
	PublishTime time.Time `gorm:"column:publishTime"`
}

type Comment struct {
	ID          int64     `gorm:"column:ID;primaryKey"`
	Author      string    `gorm:"column:author"`
	Title       string    `gorm:"column:title"`
	Content     string    `gorm:"column:content"`
	Favorite    int64     `gorm:"column:favorite"`
	Like        int64     `gorm:"column:like"`
	PublishTime time.Time `gorm:"column:publishTime"`
}
