package entity

// TWrongInfo
type TWrongInfo struct {
	ClassID  string `gorm:"column:classId"`
	OpenID   string `gorm:"column:openId"`
	WrongIDs string `gorm:"column:wrongIds"`
}
