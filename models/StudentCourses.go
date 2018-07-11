package models

import (
  "fmt"
  "github.com/prateeknischal/gorm-basic/db"
)

type StudentCourses struct {
  Id int
  StudentId int
  CourseId int
}

func (sc *StudentCourses) Save() error {
  err := db.DB.Save(sc).Error
  if err != nil {
    fmt.Printf("Failed to save StudentCourses; err: %s", err.Error())
    return err
  }
  return nil
}
