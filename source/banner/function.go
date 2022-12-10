package banner

import (
	"github.com/JesusKian/School-Management/source/employee"
	"github.com/JesusKian/School-Management/source/student"
	"github.com/fatih/color"

	"time"
)

func SchoolBanner() {
	hour, min, sec := time.Now().Clock()

	color.Magenta(`
	 ____       _                 _ 
	/ ___|  ___| |__   ___   ___ | |  Live Time -> %d:%d:%d
	\___ \ / __| '_ \ / _ \ / _ \| |  Students  -> %d
	 ___) | (__| | | | (_) | (_) | |  Employees -> %d
	|____/ \___|_| |_|\___/ \___/|_| 
	https://github.com/JesusKian/School-Management
	`, hour, min, sec, student.Counter(), employee.Counter())
}

func SelectPersonBanner() {
	color.Cyan(`
	[1] -> Student
	[2] -> Employee
	[0] -> Exit`)
}

func StudentBanner() {
	color.Cyan(`
	[1] -> Add Student
	[2] -> Edit Student Information
	[3] -> Remove Student
	[4] -> Show Student Information
	[5] -> Show All Students 
	[6] -> Save Students in JSON Format
	[0] -> Back To Main Menu`)
}

func EmployeeBanner() {
	color.Cyan(`
	[1] -> Add Employee
	[2] -> Edit Employee Information
	[3] -> Remove Employee
	[4] -> Show Employee Information
	[5] -> Show All Employees
	[6] -> Save Employees in JSON Format
	[0] -> Back To Main Menu`)
}

func ScoreBanner() {
	color.Yellow("[?] Do U Want To Enter Score ?")
	color.Cyan(`	[1] -> Yes
	[2] -> NO`)
}

func SkillBanner() {
	color.Yellow("[?] Do U Want To Add Skill ?")
	color.Cyan(`	[1] -> Yes
	[2] -> No`)
}

func EditSkillBanner() {
	color.Cyan(`	[1] -> Add Skill
	[2] -> Remove Skill`)
}

func EditScoreBanner() {
	color.Cyan(`	[1] -> Add Score
	[2] -> Edit Score
	[3] -> Remove Score
	`)
}

func EditPrivateBanner() {
	color.Cyan(`	[1] -> Add Private Information
	[2] -> Edit Private Information
	[3] -> Remove Private Information
	`)
}

func GenderBanner() {
	color.Yellow("[?] Select Your Gender")
	color.Cyan(`	[1] -> Male
	[2] -> Female`)
}

func ChangeStudentInformationBanner() {
	color.Cyan(`	[1] -> First Name
	[2] -> Last Name
	[3] -> Age
	[4] -> Gender
	[5] -> Skill
	[6] -> Score
	[7] -> Private
	[0] -> Back To Student's Menu
	`)
}

func ChangeEmployeeInformationBanner() {
	color.Cyan(`	[1] -> First Name
	[2] -> Last Name
	[3] -> Age
	[4] -> Gender
	[5] -> Skill
	[6] -> Rank
	[7] -> Private
	[0] -> Back To Student's Menu
	`)
}
