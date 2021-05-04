package controller

type Context interface {
	Param(string) string
	Status(int)
	BindJSON(interface{}) error
	Value(interface{}) interface{}
	JSON(int, interface{})
	AbortWithStatusJSON(int, interface{})
}
