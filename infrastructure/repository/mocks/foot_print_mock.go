package mocks

import (
	"finder/domain"
	"finder/shared"

	"github.com/stretchr/testify/mock"
)

type FootPrintRepository struct {
	mock.Mock
}

func (r *FootPrintRepository) GetFootPrintsByUid(currentUserUid string) ([]domain.FootPrint, error) {
	footPrints := []domain.FootPrint{}
	return footPrints, nil
}

func (r *FootPrintRepository) CreateFootPrint(currentUserUid string, visitorUid string) error {
	arguments := r.Called(currentUserUid, visitorUid)

	err := shared.MockArgumentsError(arguments, 0)
	return err
}

func (r *FootPrintRepository) UpdateToAlreadyRead(currentUserUid string) error {
	return nil
}
