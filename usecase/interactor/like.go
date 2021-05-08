package interactor

import (
	"finder/domain"
	"finder/interface/repository"
)

type LikeInteractor interface {
	CreateLike(sentUesrUid string, recievedUserUid string) (*domain.Like, error)
	GetOldestLikeByUid(currentUserUid string) (*domain.Like, error)
	GetNextUserByUid(recievedUserUid string, sentUesrUid string) (*domain.Like, error)
	Consent(recievedUserUid string, sentUesrUid string) error
}

type likeInteractor struct {
	likeRepository repository.LikeRepository
}

func NewLikeInteractor(lr repository.LikeRepository) *likeInteractor {
	return &likeInteractor{
		likeRepository: lr,
	}
}

func (i *likeInteractor) CreateLike(sentUesrUid string, recievedUserUid string) (*domain.Like, error) {
	like, err := i.likeRepository.CreateLike(sentUesrUid, recievedUserUid)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (i *likeInteractor) GetOldestLikeByUid(currentUserUid string) (*domain.Like, error) {
	like, err := i.likeRepository.GetOldestLikeByUid(currentUserUid)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (i *likeInteractor) GetNextUserByUid(recievedUserUid string, sentUesrUid string) (*domain.Like, error) {
	if err := i.likeRepository.NopeUserByUid(recievedUserUid, sentUesrUid); err != nil {
		return nil, err
	}
	like, err := i.likeRepository.GetOldestLikeByUid(recievedUserUid)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (i *likeInteractor) Consent(recievedUserUid string, sentUesrUid string) error {
	if err := i.likeRepository.Consent(recievedUserUid, sentUesrUid); err != nil {
		return err
	}
	return nil
}
