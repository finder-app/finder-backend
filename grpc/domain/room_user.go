package domain

type RoomUser struct {
	Id      uint64
	RoomId  uint
	UserUid string
	User    User `gorm:"foreignKey"`
}
