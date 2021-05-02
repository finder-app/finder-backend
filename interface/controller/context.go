package controller

type Context interface {
	Param(string) string
	Status(int)
	BindJSON(interface{}) error
	JSON(int, interface{})
	AbortWithStatusJSON(int, interface{})
}
