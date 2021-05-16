package domain

type RoomUser struct {
	RoomId  uint
	UserUid string
	User    User `gorm:"foreignKey"`
}
