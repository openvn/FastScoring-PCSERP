package db

import (
	"strconv"
)

type DBContext interface {
	GetCourseList(faculty string) []CourseInfo
	GetCourse(id string) CourseInfo
	SaveScores(id string, scores []float32)
}

type TestDBContext struct {
	test int
}

func (t *TestDBContext) GetCourseList(faculty string) []CourseInfo {
	cis := make([]CourseInfo, 5, 5)
	inf := CourseInfo{}
	inf.Credit = "3"
	inf.Description = "ascasc asdasd"
	inf.Faculty = "aasdqweasd"
	inf.FacultyID = "asdasd asd asd asd asd"
	inf.ID = "123 123 123"
	inf.Percent = "50"
	inf.SchoolName = "32212423423"
	inf.SubjectID = "123 123 123 123"
	inf.SubjectName = "aaaaaaaaaa"
	inf.Teacher = "asdasdasd"
	inf.TeacherID = "123123123"
	inf.Term = "1"
	inf.Year = "2312"

	cis[0] = inf
	inf.SubjectName = "bbbbbbbbbbbbbbb"
	cis[1] = inf
	inf.SubjectName = "ccccccccccccccccc"
	cis[2] = inf
	inf.SubjectName = "dddddddddddddddd"
	cis[3] = inf
	inf.SubjectName = "eeeeeeeeeeeee"
	cis[4] = inf
	return cis
}

func (t *TestDBContext) GetCourse(id string) CourseInfo {
	inf := CourseInfo{}
	inf.Credit = "3"
	inf.Description = "ascasc asdasd"
	inf.Faculty = "aasdqweasd"
	inf.FacultyID = "asdasd asd asd asd asd"
	inf.ID = "123 123 123"
	inf.Percent = "50"
	inf.SchoolName = "32212423423"
	inf.SubjectID = "123 123 123 123"
	inf.SubjectName = "aaaaaaaaaa"
	inf.Teacher = "asdasdasd"
	inf.TeacherID = "123123123"
	inf.Term = "1"
	inf.Year = "2312"

	sts := make([]Student, 45, 45)
	for i := 0; i < 45; i++ {
		sts[i].ID = "00000000" + strconv.Itoa(i)
		sts[i].LastName = "ABSVA"
		sts[i].MiddleName = "MASIWD"
		sts[i].FirstName = "OAO" + strconv.Itoa(i)
		sts[i].Birth = "12/12/1212"
	}

	inf.Students = sts
	return inf
}

func (t *TestDBContext) SaveScores(id string, scores []float32) {

}
