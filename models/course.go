package models

import (
	"fmt"

	"github.com/prateeknischal/gorm-basic/db"
)

type Course struct {
	Id      int
	Name    string
	Credits int
}

func (c *Course) GetCourse() error {
	err := db.DB.Find(&c).Error
	if err != nil {
		fmt.Printf("Failed to get Course; err: %s", err.Error())
		return err
	}
	return nil
}

func (c *Course) Save() error {
	err := db.DB.Save(c).Error
	if err != nil {
		fmt.Printf("Failed to persist Course; err: %s", err.Error())
		return err
	}
	return nil
}
