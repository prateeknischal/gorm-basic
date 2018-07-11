package models

import (
	"fmt"

	"github.com/prateeknischal/gorm-basic/db"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

func (s *Student) GetStudent() error {
	err := db.DB.Find(&s).Error
	if err != nil {
		fmt.Printf("Failed to get student; err: %s", err.Error())
		return err
	}
	return nil
}

func (s *Student) Save() error {
	err := db.DB.Save(s).Error
	if err != nil {
		fmt.Printf("Failed to persist Student; err: %s", err.Error())
		return err
	}
	return nil
}
