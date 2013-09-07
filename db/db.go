package db

import (
//"encoding/xml"
)

type Student struct {
	ID         string `json:"student_id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Birth      string `json:"dayofbirth"`
	Class      string `json:"classname"`
}

type CourseInfo struct {
	ID          string    `json:"course_id"`
	Students    []Student `json:"students,omitempty"`
	SubjectName string    `json:"subject"`
	SubjectID   string    `json:"subject_id"`
	Teacher     string    `json:"teacher"`
	TeacherID   string    `json:"teacher_id"`
	Faculty     string    `json:"faculty"`
	FacultyID   string    `json:"faculty_id"`
	SchoolName  string    `json:"school_name"`
	Term        string    `json:"term"`
	Year        string    `json:"year"`
	Credit      string    `json:"credits"`
	Percent     string    `json:"percent"`
	Description string    `json:"description"`
}
