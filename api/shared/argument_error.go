package shared

import "github.com/stretchr/testify/mock"

func MockArgumentsError(arguments mock.Arguments, index int) error {
	var err error
	if arguments.Error(index) != nil {
		err = arguments.Error(index)
	}
	return err
}
