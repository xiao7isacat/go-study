package models

import "test/database"

func AutoMigrat() {
	database.DB.AutoMigrate(&Test{})

	//database.DB.AutoMigrate(&IpAddrInfo{})
	//database.DB.AutoMigrate(&HostJobRelation{})
}
