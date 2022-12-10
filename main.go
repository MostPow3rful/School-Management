package main

import (
	"github.com/JesusKian/School-Management/source/banner"
	"github.com/JesusKian/School-Management/source/config"
	"github.com/JesusKian/School-Management/source/employee"
	"github.com/JesusKian/School-Management/source/student"
	"github.com/fatih/color"

	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

// ------------------------------------------- Global Variable -------------------------------------------

const (
	// Detecting Client Platform
	platform = runtime.GOOS
)

var (
	// Error
	err error

	// Log Info
	flags = log.Lshortfile
	// Sturct For Information Log
	infoLog = log.New(os.Stdout, "[?] Information -> ", flags)
	// Sturct For Warning Log
	warnLog = log.New(os.Stdout, "[*] Warning -> ", flags)
	// Sturct For Error Log
	errorLog = log.New(os.Stdout, "[!] Error -> ", flags)
	// Struct For Default Log
	defaultLog = log.New(os.Stderr, "[#] Default Log -> ", flags)
)

// ------------------------------------------- MAIN -------------------------------------------

func main() {
	// Temp Variable To Get Information And Pass it To Function Parameter
	var (
		tempString       string            = ""
		tempStringLesson string            = ""
		ID               string            = ""
		tempInt          int               = 0
		tempSliceString  []string          = []string{}
		tempSliceSkill   []string          = []string{}
		tempSliceInt     []int             = []int{}
		tempMap          map[string]string = make(map[string]string)
		tempMapPriv8     map[string]string = make(map[string]string)
	)

	// Cheking . . .
	err := config.Config()
	if err != nil {
		config.SetLog("D", err.Error())
		log.Fatal(err)
	}
	config.SetLog("W", "App Started")
	config.SetLog("I", "config.Config() -> Succesful")
	config.SetLog("I", "platform -> "+platform)

main_menu:
	config.Clear()
	banner.SchoolBanner()
	banner.SelectPersonBanner()
	fmt.Print("\t\t\t[?] -> ")
	fmt.Scan(&tempString)

	for {
		switch tempString {
		case "1": // Student
		student:
			config.Clear()
			banner.SchoolBanner()
			banner.StudentBanner()
			fmt.Print("\t\t\t[?] -> ")
			fmt.Scan(&tempString)
			config.Clear()
			banner.SchoolBanner()

			switch tempString {
			case "1": // Add
				color.Yellow("[?] Enter Student's First Name [Type : String] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				tempSliceString = append(tempSliceString, tempString)

				color.Yellow("[?] Enter Student's Last Name [Type : String] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				tempSliceString = append(tempSliceString, tempString)

				color.Yellow("[?] Enter Student's Age [Type : Int] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempInt)
				tempSliceInt = append(tempSliceInt, tempInt)

				banner.GenderBanner()
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				switch tempString {
				case "1": // Male
					tempSliceString = append(tempSliceString, "Male")

				case "2": // Female
					tempSliceString = append(tempSliceString, "Female")

				default: // None Of Them
					tempSliceString = append(tempSliceString, "None")
				}

				banner.ScoreBanner()
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)

				switch tempString {
				case "1": // YES
					color.Yellow("[?] How Many Lessons Do U Want To Enter [Type : Int] ")
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempInt)

					for i := 0; i < tempInt; i++ {
						color.Yellow("[?] Enter Lesson Name [Type : String] ")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)

						color.Yellow("[?] Enter Lesson Score [Type : String] ")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempStringLesson)
						tempMap[tempString] = tempStringLesson // 431
					}

				default: // 2,NO
					tempMap["NONE"] = "NONE"
				}

				banner.SkillBanner()
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)

				switch tempString {
				case "1": // YES
					color.Yellow("[?] How Many Skills Do U Want To Enter [Type : Int] ")
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempInt)

					for i := 0; i < tempInt; i++ {
						color.Yellow("[?] Enter Skill Name [Type : String] ")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)

						tempSliceSkill = append(tempSliceSkill, tempString)
					}

				default: // 2,NO
					tempSliceSkill = append(tempSliceSkill, "NONE")
				}

				color.Yellow("[?] Enter Student's Phone Number [Type : String] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				tempMapPriv8["phoneNumber"] = tempString

				color.Yellow("[?] Enter Student's Home Number [Type : String] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				tempMapPriv8["homeNumber"] = tempString

				color.Yellow("[?] Enter Student's ID Card Number [Type : String] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				tempMapPriv8["IDCardNumber"] = tempString

				color.Yellow("[?] Enter Location Of Student's Home [Type : String] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				tempMapPriv8["homeLocation"] = tempString

				s, err := student.Add(
					tempSliceString[0],
					tempSliceString[1],
					tempSliceInt[0],
					tempSliceString[2],
					tempMap,
					tempSliceSkill,
					tempMapPriv8,
				)

				if err != nil {
					config.SetLog("E", "student.Add() -> Invalid Argument")
				} else {
					config.SetLog("I", "student.Add() -> New Student Created -> ID : "+s.ID)
					config.Clear()
					banner.SchoolBanner()
					s.ShowNow()
					time.Sleep(time.Second * 5)
				}

			case "2": // Edit
				config.Clear()
				banner.SchoolBanner()

				color.Yellow("[?] Enter Student's ID")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&ID)

			edit_student:
				config.Clear()
				banner.SchoolBanner()
				banner.ChangeStudentInformationBanner()
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)

				config.Clear()
				banner.SchoolBanner()

				switch tempString {
				case "1": // First Name
					color.Yellow("[?] Enter New Student's First Name [Type:String]")
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempString)
					student.EditFirstName(ID, tempString)

				case "2": // Last Name
					color.Yellow("[?] Enter New Student's Last Name [Type:String]")
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempString)
					student.EditLastName(ID, tempString)

				case "3": // Age
					color.Yellow("[?] Enter New Student's Age")
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempString)

					tempInt, _ = strconv.Atoi(tempString)
					student.EditAge(ID, tempInt, tempString)

				case "4": // Gender
					banner.GenderBanner()
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempString)
					banner.GenderBanner()

					switch tempString {
					case "1":
						student.EditGender(ID, "Male")

					case "2":
						student.EditGender(ID, "Female")

					default:
						goto edit_student
					}

				case "5": // Skill
					banner.EditSkillBanner()
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempString)

					config.Clear()
					banner.SchoolBanner()

					switch tempString {
					case "1": // Add
						color.Yellow("[?] Enter New Student's Skill")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)
						student.EditSkill(ID, "1", tempString)

					case "2": // Remove
						config.Clear()
						banner.SchoolBanner()
						student.ShowSkill(ID)
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)
						student.EditSkill(ID, "2", tempString)

					default:
						goto edit_student
					}

				case "6": // Score
					banner.EditScoreBanner()
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempString)

					config.Clear()
					banner.SchoolBanner()

					switch tempString {
					case "1": // Add
						var score string = ""
						color.Yellow("[?] Enter New Student's Lesson")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)

						color.Yellow("[?] Enter New Student's Score")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&score)

						student.EditScore(ID, "1", tempString, score, "")

					case "2": // Edit
						var newScore string = ""
						student.ShowScore(ID)
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)

						config.Clear()
						banner.SchoolBanner()
						color.Yellow("[?] Enter New Score")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&newScore)

						student.EditScore(ID, "2", tempString, newScore, "")

					case "3": // Remove
						student.ShowScore(ID)
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)

						student.EditScore(ID, "3", "", "", tempString)

					default:
						goto edit_student
					}

				case "7": // Private Information
					banner.EditPrivateBanner()
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempString)

					config.Clear()
					banner.SchoolBanner()

					switch tempString {
					case "1": // Add
						var value string = ""
						color.Yellow("[?] Enter New Student's Private Information's Title")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)

						color.Yellow("[?] Enter New Student's Private Information's Value")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&value)

						student.EditPrivateInformation(ID, "1", tempString, value, "")

					case "2": // Edit
						var newValue string = ""
						student.ShowPrivateInformation(ID)
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)

						config.Clear()
						banner.SchoolBanner()
						color.Yellow("[?] Enter New Value")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&newValue)

						student.EditPrivateInformation(ID, "2", tempString, newValue, "")

					case "3": // Remove
						student.ShowPrivateInformation(ID)
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)

						student.EditPrivateInformation(ID, "3", "", "", tempString)

					default:
						goto edit_student
					}

				case "0": // BACK To Student Menu
					goto student

				default: // Back To Edit Student's Information
					goto edit_student
				}

			case "3": // Remove
				config.Clear()
				banner.SchoolBanner()

				color.Yellow("[?] Enter Student's ID")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&ID)

				student.Remove(ID)

			case "4": // Show One
				config.Clear()
				banner.SchoolBanner()

				color.Yellow("[?] Enter Student's ID")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&ID)

				config.Clear()
				banner.SchoolBanner()
				student.Show(ID)
				time.Sleep(time.Second * 10)

			case "5": // Show All
				config.Clear()
				banner.SchoolBanner()
				student.ShowAll()

			case "6": // JSON
				student.JsonOutput()

			case "0": // Main Menu
				goto main_menu

			default: // Back To Student Menu
				goto student
			}
			goto student

		case "2": // Employee
		employee:
			config.Clear()
			banner.SchoolBanner()
			banner.EmployeeBanner()
			fmt.Print("\t\t\t[?] -> ")
			fmt.Scan(&tempString)
			config.Clear()
			banner.SchoolBanner()

			switch tempString {
			case "1": // Add
				var employeeRank string = ""
				color.Yellow("[?] Enter Employee's First Name [Type : String] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				tempSliceString = append(tempSliceString, tempString)

				color.Yellow("[?] Enter Employee's Last Name [Type : String] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				tempSliceString = append(tempSliceString, tempString)

				color.Yellow("[?] Enter Employee's Age [Type : Int] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempInt)
				tempSliceInt = append(tempSliceInt, tempInt)

				banner.GenderBanner()
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				switch tempString {
				case "1": // Male
					tempSliceString = append(tempSliceString, "Male")

				case "2": // Female
					tempSliceString = append(tempSliceString, "Female")

				default: // None Of Them
					tempSliceString = append(tempSliceString, "None")
				}

				color.Yellow("[?] Enter Employee's Rank")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&employeeRank)

				banner.SkillBanner()
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)

				switch tempString {
				case "1": // YES
					color.Yellow("[?] How Many Skills Do U Want To Enter [Type : Int] ")
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempInt)

					for i := 0; i < tempInt; i++ {
						color.Yellow("[?] Enter Skill Name [Type : String] ")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)

						tempSliceSkill = append(tempSliceSkill, tempString)
					}

				default: // 2,NO
					tempSliceSkill = append(tempSliceSkill, "NONE")
				}

				color.Yellow("[?] Enter Employee's Phone Number [Type : String] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				tempMapPriv8["phoneNumber"] = tempString

				color.Yellow("[?] Enter Employee's Home Number [Type : String] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				tempMapPriv8["homeNumber"] = tempString

				color.Yellow("[?] Enter Employee's ID Card Number [Type : String] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				tempMapPriv8["IDCardNumber"] = tempString

				color.Yellow("[?] Enter Location Of Employee's Home [Type : String] ")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)
				tempMapPriv8["homeLocation"] = tempString

				e, err := employee.Add(
					tempSliceString[0],
					tempSliceString[1],
					tempSliceInt[0],
					tempSliceString[2],
					employeeRank,
					tempSliceSkill,
					tempMapPriv8,
				)

				if err != nil {
					config.SetLog("E", "employee.Add() -> Invalid Argument")
				} else {
					config.SetLog("I", "employee.Add() -> New Student Created -> ID : "+e.ID)
					config.Clear()
					banner.SchoolBanner()
					e.ShowNow()
					time.Sleep(time.Second * 5)
				}

			case "2": // Edit
				config.Clear()
				banner.SchoolBanner()

				color.Yellow("[?] Enter Employee's ID")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&ID)

			edit_employee:
				config.Clear()
				banner.SchoolBanner()
				banner.ChangeEmployeeInformationBanner()
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&tempString)

				config.Clear()
				banner.SchoolBanner()

				switch tempString {
				case "1": // First Name
					color.Yellow("[?] Enter New Employee's First Name [Type:String]")
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempString)
					employee.EditFirstName(ID, tempString)

				case "2": // Last Name
					color.Yellow("[?] Enter New Employee's Last Name [Type:String]")
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempString)
					employee.EditLastName(ID, tempString)

				case "3": // Age
					color.Yellow("[?] Enter New Employee's Age")
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempString)

					tempInt, _ = strconv.Atoi(tempString)
					employee.EditAge(ID, tempInt, tempString)

				case "4": // Gender
					banner.GenderBanner()
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempString)
					banner.GenderBanner()

					switch tempString {
					case "1":
						employee.EditGender(ID, "Male")

					case "2":
						employee.EditGender(ID, "Female")

					default:
						goto edit_employee
					}

				case "5": // Skill
					banner.EditSkillBanner()
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempString)

					config.Clear()
					banner.SchoolBanner()

					switch tempString {
					case "1": // Add
						color.Yellow("[?] Enter New Employee's Skill")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)
						employee.EditSkill(ID, "1", tempString)

					case "2": // Remove
						config.Clear()
						banner.SchoolBanner()
						employee.ShowSkill(ID)
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)
						employee.EditSkill(ID, "2", tempString)

					default:
						goto edit_employee
					}

				case "6": // Rank
					color.Yellow("[?] Enter New Employee's Rank")
					fmt.Scan(&tempString)
					employee.EditRank(ID, tempString)

				case "7": // Private Information
					banner.EditPrivateBanner()
					fmt.Print("\t\t[?] -> ")
					fmt.Scan(&tempString)

					config.Clear()
					banner.SchoolBanner()

					switch tempString {
					case "1": // Add
						var value string = ""
						color.Yellow("[?] Enter New Employee's Private Information's Title")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)

						color.Yellow("[?] Enter New Employee's Private Information's Value")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&value)

						employee.EditPrivateInformation(ID, "1", tempString, value, "")

					case "2": // Edit
						var newValue string = ""
						employee.ShowPrivateInformation(ID)
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)

						config.Clear()
						banner.SchoolBanner()
						color.Yellow("[?] Enter New Value")
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&newValue)

						employee.EditPrivateInformation(ID, "2", tempString, newValue, "")

					case "3": // Remove
						employee.ShowPrivateInformation(ID)
						fmt.Print("\t\t[?] -> ")
						fmt.Scan(&tempString)

						employee.EditPrivateInformation(ID, "3", "", "", tempString)

					default:
						goto edit_employee
					}

				case "0": // BACK To Student Menu
					goto employee

				default: // Back To Edit Student's Information
					goto edit_employee
				}

			case "3": // Remove
				config.Clear()
				banner.SchoolBanner()

				color.Yellow("[?] Enter Employee's ID")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&ID)

				employee.Remove(ID)

			case "4": // Show One
				config.Clear()
				banner.SchoolBanner()

				color.Yellow("[?] Enter Employee's ID")
				fmt.Print("\t\t[?] -> ")
				fmt.Scan(&ID)

				config.Clear()
				banner.SchoolBanner()
				employee.Show(ID)
				time.Sleep(time.Second * 10)

			case "5": // Show All
				config.Clear()
				banner.SchoolBanner()
				employee.ShowAll()

			case "6": // JSON
				employee.JsonOutput()

			case "0": // Main Menu
				goto main_menu

			default: // Back To Student Menu
				goto employee
			}
			goto employee

		case "0": // EXIT
			config.Clear()
			config.SetLog("W", "App Closed")
			os.Exit(1)

		default: // Back To Main Menu
			goto main_menu
		}
	}
}
