package models

import (
	"fmt"
	"gorm.io/gorm"
	"test/database"
)

type Test struct {
	gorm.Model
	TestStr  string
	TestInt  int64
	TestInt1 int64
}

func (this *Test) TableName() string {
	return "test_t"
}

func (this *Test) CreateOne() (uint, error) {
	table := database.DB.Table(this.TableName())

	if err := table.Debug().Create(this).Error; err != nil {
		return this.ID, err
	}

	return this.ID, nil
}

func (this *Test) Update() error {
	table := database.DB.Table(this.TableName())
	if this.TestInt1 != 0 {
		table = table.Where("test_int1 = ?", this.TestInt1)

	}
	fmt.Println(this.TestStr)
	if err := table.Model(this).Select("test_int", "test_str", "").Debug().Updates(&this).Error; err != nil {
		return err
	}
	return nil
}
