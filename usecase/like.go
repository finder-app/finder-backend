package usecase

import (
	"finder/domain"
	"finder/infrastructure/repository"
	"fmt"
)

type LikeUsecase interface {
	CreateLike(sentUesrUid string, recievedUserUid string) (*domain.Like, error)
	GetOldestLikeByUid(currentUserUid string) (*domain.Like, error)
	GetNextUserByUid(recievedUserUid string, sentUesrUid string) (*domain.Like, error)
	Consent(recievedUserUid string, sentUesrUid string) error
}

type likeUsecase struct {
	likeRepository repository.LikeRepository
	roomRepository repository.RoomRepository
}

func NewLikeUsecase(lr repository.LikeRepository, rr repository.RoomRepository) *likeUsecase {
	return &likeUsecase{
		likeRepository: lr,
		roomRepository: rr,
	}
}

func (u *likeUsecase) CreateLike(sentUesrUid string, recievedUserUid string) (*domain.Like, error) {
	like, err := u.likeRepository.CreateLike(sentUesrUid, recievedUserUid)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (u *likeUsecase) GetOldestLikeByUid(currentUserUid string) (*domain.Like, error) {
	like, err := u.likeRepository.GetOldestLikeByUid(currentUserUid)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (u *likeUsecase) GetNextUserByUid(recievedUserUid string, sentUesrUid string) (*domain.Like, error) {
	if err := u.likeRepository.NopeUserByUid(recievedUserUid, sentUesrUid); err != nil {
		return nil, err
	}
	like, err := u.likeRepository.GetOldestLikeByUid(recievedUserUid)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (u *likeUsecase) Consent(recievedUserUid string, sentUesrUid string) error {
	tx := u.likeRepository.Begin()
	if err := u.likeRepository.Consent(tx, recievedUserUid, sentUesrUid); err != nil {
		tx.Rollback()
		return err
	}
	room, err := u.roomRepository.CreateRoom(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(room)
	tx.Rollback()
	// tx.Commit()
	return nil
}
