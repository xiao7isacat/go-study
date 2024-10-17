package main

import (
	"k8s.io/klog/v2"
	"test/database"
	"test/models"
)

func main() {
	if err := database.ConnectDb("sqlite"); err != nil {
		klog.Errorln("database:", "sqlite", "error: ", err)
		return
	}

	models.AutoMigrat()

	var test models.Test
	/*test.TestInt = 1
	test.TestInt1 = 2
	test.TeatStr = "a"
	test.CreateOne()*/

	test.TestInt1 = 1
	test.TestStr = "as"
	test.Update()
}
