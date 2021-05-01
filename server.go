package main

import (
	"finder/infrastructure"
	"fmt"
)

func main() {
	// validationもinfrastructureに作れ
	dbConn := infrastructure.NewGormConnect()
	defer dbConn.Close()
	fmt.Println(dbConn)
}
