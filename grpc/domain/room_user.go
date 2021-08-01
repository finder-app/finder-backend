package domain

type RoomUser struct {
	Id      uint64
	RoomId  uint64
	UserUid string
	User    User `gorm:"foreignKey"`
}
