package models

import (
	"errors"
)

var (
	Students map[string]*Student
)

type Student struct {
	StudentId string
	Name      string
	Age       int64
}

func init() {
	Students = make(map[string]*Student)
	Students["student_a"] = &Student{"student_a", "tom", 19}
	Students["student_b"] = &Student{"student_b", "jack", 21}
}

func GetAllStudent() map[string]*Student {
	return Students
}

func GetOneStudent(StudentId string) (student *Student, err error) {
	if v, ok := Students[StudentId]; ok {
		return v, nil
	}
	return nil, errors.New("StudentId Not Exist")
}
