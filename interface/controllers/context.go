package controllers

type Context interface {
	Param(string) string
	Status(int)
	JSON(int, interface{})
	BindJSON(interface{}) error
}
