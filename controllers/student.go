package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"mybeeapi/models"
)

type StudentController struct {
	beego.Controller
}

// @router / [get]
func (s *StudentController) GetAll() {
	data := models.GetAllStudent()
	s.Data["json"] = data
	s.ServeJSON()
}

// @router /:studentId [get]
func (s *StudentController) GetOne() {
	studentId := s.Ctx.Input.Param(":studentId")
	if studentId != "" {
		student, err := models.GetOneStudent(studentId)
		if err != nil {
			s.Data["json"] = err.Error()
		} else {
			s.Data["json"] = student
		}
	}
	s.ServeJSON()
}
