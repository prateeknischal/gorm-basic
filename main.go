package main

import (
	"encoding/json"
	"fmt"

	"github.com/prateeknischal/gorm-basic/db"
	"github.com/prateeknischal/gorm-basic/models"
)

func main() {
	// fmt.Println(db.DB.CreateTable(&models.Course{}, &models.Student{}, &models.StudentCourses{}).Error)
	test()
	txSaveTest()
	findTest()
}

func test() {
	// courseList := make([]models.Course, 0)
	// courseList = append(courseList, models.Course{Name: "Topology", Credits: 5})
	// courseList = append(courseList, models.Course{Name: "Multivariate calculus", Credits: 10})
	// courseList = append(courseList, models.Course{Name: "Number Fields", Credits: 3})
	// courseList = append(courseList, models.Course{Name: "Cryptography", Credits: 10})
	//
	// for _, c := range courseList {
	// 	c.Save()
	// }

	// st := models.Student{Name: "Jacob", Age: 21}
	// st.Save()

	// t := models.StudentCourses{StudentId: 1, CourseId: 2}
	// t.Save()
	// tt := models.StudentCourses{StudentId: 1, CourseId: 4}
	// tt.Save()

	// s := &models.Student{Name: "Jacob"}
	// db.DB.Debug().Find(&s)
	// v, _ := json.MarshalIndent(s, "", "  ")
	// fmt.Println(string(v))

}

func txSaveTest() {
	s := models.Student{Name: "John", Age: 21}
	tx := db.DB.Begin()
	err := tx.Debug().Save(&s).Error
	if err != nil {
		fmt.Printf("Failed to save student; err: %s", err.Error())
		tx.Rollback()
	}
	tx.Commit()
}

func findTest() {
	s := models.Student{Name: "Jacob"}
	err := db.DB.Find(&s).Error
	if err != nil {
		fmt.Printf("Failed to get; err: %s", err.Error())
		return
	}
	v, _ := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(v))
}
