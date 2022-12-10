package employee

import (
	"os"

	"github.com/JesusKian/School-Management/source/config"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"

	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	err error = nil
)

func Counter() int {
	var count int = 0

	result, err := config.Database.Query("SELECT COUNT(*) FROM Employees")

	if err != nil {
		config.SetLog("E", "employee.employee.Counter() -> Couldn't Count Employees")
		config.SetLog("D", err.Error())
		return -1
	}

	for result.Next() {
		err = result.Scan(&count)

		if err != nil {
			config.SetLog("E", "employee.employee.Counter() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
			return -1
		}

	}

	return count
}

func Add(FN string, LN string, A int, G string, R string, SK []string, Priv8 map[string]string) (*Employee, error) {
	var (
		dbSK    string = ""
		dbPriv8 string = ""
	)

	var hour, min, sec int = time.Now().Clock()

	if A == 0 {
		config.SetLog("E", "employee.employee.Add() -> Invalid Age")
		return nil, errors.New("Add() -> Invalid Age")
	}

	for _, skill := range SK {
		dbSK += fmt.Sprintf("%s ", skill)
	}

	for question, answer := range Priv8 {
		dbPriv8 += fmt.Sprintf("%s:%s ", question, answer)
	}

	_, err = config.Database.Query(`
	INSERT INTO Employees 
	(FirstName, LastName, Age, Gender, Skill, Rank, Private, ID) 
	VALUES 
	(?, ?, ?, ?, ?, ?, ?, ?)`, FN, LN, A, G, dbSK, R, dbPriv8, string(FN[0])+strconv.Itoa(hour)+strconv.Itoa(min)+strconv.Itoa(sec))

	if err != nil {
		config.SetLog("E", "employee.employee.Add() -> Couldn't Add Employee's Data In Database")
		config.SetLog("D", err.Error())
	}

	return &Employee{
		FirstName: FN,
		LastName:  LN,
		Age:       A,
		Gender:    G,
		Skill:     SK,
		Rank:      R,
		Private:   Priv8,
		ID:        string(FN[0]) + strconv.Itoa(hour) + strconv.Itoa(min) + strconv.Itoa(sec),
	}, nil
}

func (e *Employee) ShowNow() {
	color.Green("\t\t[!] Employee Secceful Added !")
	color.Cyan("\t\t[!] First Name -> " + e.FirstName)
	color.Cyan("\t\t[!] Last Name -> " + e.LastName)
	color.Cyan("\t\t[!] Age -> " + strconv.Itoa(e.Age))
	color.Cyan("\t\t[!] Gender -> " + e.Gender)
	color.Cyan("\t\t[!] Rank -> " + e.Rank)
	color.Cyan("\t\t[!] ID -> " + e.ID)

	for _, skill := range e.Skill {
		if skill == "NONE" {
			break
		}
		color.Magenta("\t\t[!] Skill -> " + skill)
	}

	for question, answer := range e.Private {
		color.Red("\t\t[!] " + question + " -> " + answer)
	}

}

func ShowSkill(ID string) {
	var (
		skillValue string   = ""
		tempSlice  []string = []string{}
	)

	result, err := config.Database.Query("SELECT Skill FROM Employees WHERE ID=?", ID)

	if err != nil {
		config.SetLog("E", "employee.employee.ShowSkill() -> Couldn't Get Employee's Skill")
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&skillValue)

		if err != nil {
			config.SetLog("E", "employee.employee.ShowSkill() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}
	}

	tempSlice = strings.Split(skillValue, " ")

	for i := 0; i < len(tempSlice)-1; i++ {
		color.Cyan("\t" + "[" + strconv.Itoa(i+1) + "] -> " + tempSlice[i])
	}
}

func EditFirstName(ID string, value string) {
	_, err = config.Database.Query("UPDATE Employees SET FirstName=? WHERE ID=?", value, ID)

	if err != nil {
		config.SetLog("E", "employee.employee.EditFirstName() -> Couldn't Edit Employee's First Name With ID="+ID)
		config.SetLog("D", err.Error())
	} else {
		config.SetLog("E", "employee.employee.EditFirstName() -> Employee With ID="+ID+" Changed First Name To "+value)
	}
}

func EditLastName(ID string, value string) {
	_, err = config.Database.Query("UPDATE Employees SET LastName=? WHERE ID=?", value, ID)

	if err != nil {
		config.SetLog("E", "employee.employee.EditFirstName() -> Couldn't Edit Employee's Last Name With ID="+ID)
		config.SetLog("D", err.Error())
	} else {
		config.SetLog("E", "employee.employee.EditFirstName() -> Employee With ID="+ID+" Changed Last Name To "+value)
	}
}

func EditAge(ID string, value int, stringValue string) {
	_, err = config.Database.Query("UPDATE Employees SET Age=? WHERE ID=?", value, ID)

	if err != nil {
		config.SetLog("E", "employee.employee.EditAge() -> Couldn't Edit Employee's Age")
		config.SetLog("D", err.Error())
	} else {
		config.SetLog("E", "employee.employee.EditAge() -> Employee With ID="+ID+" Changed Age To "+stringValue)
	}
}

func EditGender(ID string, value string) {
	_, err = config.Database.Query("UPDATE Employees SET Gender=? WHERE ID=?", value, ID)

	if err != nil {
		config.SetLog("E", "employee.employee.EditGender() -> Couldn't Edit Employee's Gender With ID="+ID)
		config.SetLog("D", err.Error())
	} else {
		config.SetLog("E", "employee.employee.EditGender() -> Employee With ID="+ID+" Changed Gender To "+value)
	}
}

func EditSkill(ID string, option string, value string) {
	var (
		skillsLen int      = 0
		skills    string   = ""
		tempSlice []string = []string{}
	)

	result, err := config.Database.Query("SELECT Skill FROM Employees WHERE ID=?", ID)

	if err != nil {
		config.SetLog("E", "employee.employee.EditSkill() -> Couldn't Get Employee's Skill With ID="+ID)
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&skills)

		if err != nil {
			config.SetLog("E", "employee.employee.EditSkill() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}
	}

	tempSlice = strings.Split(skills, " ")

	switch option {
	case "1": // Add
		result, err := config.Database.Query("SELECT COUNT(Skill) FROM Employees WHERE ID=?", ID)
		if err != nil {
			config.SetLog("E", "employee.employee.EditSkill() -> Couldn't Count Employee's Skill With ID="+ID)
			config.SetLog("D", err.Error())
		}

		for result.Next() {
			_, err = fmt.Scan(&skillsLen)
			if err != nil {
				config.SetLog("E", "employee.employee.EditSkill() -> Couldn't Get Length Of Student's Skill With ID="+ID)
				config.SetLog("D", err.Error())
			}
		}

		skills += strconv.Itoa(skillsLen) + ":" + value + " "
		result, err = config.Database.Query("UPDATE Employees SET Skill=? WHERE ID=?", skills, ID)

		if err != nil {
			config.SetLog("E", "employee.employee.EditSkill() -> Couldn't Add New Employee's Skill With ID="+ID)
			config.SetLog("D", err.Error())
		}

	case "2": // Remove
		skills = ""

		for i := 0; i < len(tempSlice)-1; i++ {
			if strconv.Itoa(i+1) != value {
				skills += string(tempSlice[i] + " ")
			}
		}

		result, err = config.Database.Query("UPDATE Employees SET Skill=? WHERE ID=?", skills, ID)

		if err != nil {
			config.SetLog("E", "employee.employee.EditSkill() -> Couldn't Remove Employee's Skill With ID="+ID)
			config.SetLog("D", err.Error())
		}
	}
}

func EditRank(ID string, value string) {
	_, err = config.Database.Query("UPDATE Employees SET Rank=? WHERE ID=?", value, ID)

	if err != nil {
		config.SetLog("E", "employee.employee.EditRank() -> Couldn't Change Employee's Rank With ID="+ID)
		config.SetLog("D", err.Error())
	} else {
		config.SetLog("E", "employee.employee.EditRank() -> Employee With ID="+ID+" Changed Rank")
	}
}

func ShowPrivateInformation(ID string) {
	var (
		answers   string   = ""
		tempArray []string = []string{}
	)

	result, err := config.Database.Query("SELECT Private FROM Employees WHERE ID=?", ID)

	if err != nil {
		config.SetLog("E", "employee.employee.ShowPrivateInformation() -> Couldn't Get Employee's Private Information With ID="+ID)
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&answers)

		if err != nil {
			config.SetLog("E", "employee.employee.ShowPrivateInformation() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}
	}

	tempArray = strings.Split(answers, " ")

	for i := 0; i < len(tempArray)-1; i++ {
		color.Cyan("\t" + "[" + strconv.Itoa(i+1) + "] -> " + tempArray[i])
	}
}

func EditPrivateInformation(ID string, option string, key string, value string, Remove string) {
	var (
		priv8     string   = ""
		tempSlice []string = []string{}
	)

	result, err := config.Database.Query("SELECT Private From Employees Where ID=?", ID)

	if err != nil {
		config.SetLog("E", "employee.employee.EditPrivateInformation() -> Couldn't Get Employee's Private Information With ID="+ID)
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&priv8)

		if err != nil {
			config.SetLog("E", "employee.employee.EditPrivateInformation() -> Couldn't Scan Employee's Private Information With ID="+ID)
			config.SetLog("D", err.Error())
		}
	}

	tempSlice = strings.Split(priv8, " ")

	switch option {
	case "1": // Add
		priv8 += key + ":" + value + " "
		result, err = config.Database.Query("UPDATE Employees SET Private=? WHERE ID=?", priv8, ID)

		if err != nil {
			config.SetLog("E", "employee.employee.EditPrivateInformation() -> Couldn't Add New Private Information With ID="+ID)
			config.SetLog("D", err.Error())
		} else {
			config.SetLog("E", "employee.employee.EditPrivateInformation() -> Employee With ID="+ID+" Added New Private Information "+key+":"+value)
		}

	case "2": // Edit
		priv8 = ""
		index, err := strconv.Atoi(key)

		for i := 0; i < len(tempSlice)-1; i++ {
			if strings.Split(tempSlice[i], ":")[0] == strings.Split(tempSlice[index-1], ":")[0] {
				priv8 += strings.Split(tempSlice[index-1], ":")[0] + ":" + value + " "
				continue
			}
			priv8 += tempSlice[i] + " "
		}

		result, err = config.Database.Query("UPDATE Employees SET Private=? WHERE ID=?", priv8, ID)

		if err != nil {
			config.SetLog("E", "employee.employee.EditPrivateInformation() -> Couldn't Edit Employee's Private Information With ID="+ID)
			config.SetLog("D", err.Error())
		} else {
			config.SetLog("E", "employee.employee.EditPrivateInformation() -> Employee With ID="+ID+" Edited Private Information Of "+strings.Split(tempSlice[index-1], ":")[0])
		}

	case "3": // Remove
		priv8 = ""

		for i := 0; i < len(tempSlice)-1; i++ {
			if strconv.Itoa(i) != Remove {
				priv8 += tempSlice[i] + " "
			}
		}

		result, err = config.Database.Query("UPDATE Employees SET Private=? WHERE ID=?", priv8, ID)

		if err != nil {
			config.SetLog("E", "employee.employee.EditPrivateInformation() -> Couldn't Remove Employee's Private Information With ID="+ID)
			config.SetLog("D", err.Error())
		} else {
			config.SetLog("E", "employee.employee.EditPrivateInformation() -> Employee With ID="+ID+" Removed Private Information")
		}

	}
}

func Remove(ID string) {
	_, err := config.Database.Query("DELETE FROM Employees WHERE ID=?", ID)
	if err != nil {
		config.SetLog("E", "employee.employee.Remove() -> Couldn't Remove Employee With ID="+ID)
		config.SetLog("D", err.Error())
	} else {
		config.SetLog("E", "employee.employee.Remove() -> Employee With ID="+ID+" Removed")
	}
}

func Show(ID string) {
	var (
		FN        string = ""
		LN        string = ""
		A         int    = 0
		G         string = ""
		SK        string = ""
		R         string = ""
		Priv8     string = ""
		StudentID string = ""
	)

	result, err := config.Database.Query("SELECT * FROM Employees WHERE ID=?", ID)

	if err != nil {
		config.SetLog("E", "employee.employee.Show() -> Couldn't Get Employee's Info With ID="+ID)
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&FN, &LN, &A, &G, &SK, &R, &Priv8, &StudentID)

		if err != nil {
			config.SetLog("E", "employee.employee.Show() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}

		color.Cyan("\t\t[!] First Name -> " + FN)
		color.Cyan("\t\t[!] Last Name -> " + LN)
		color.Cyan("\t\t[!] Age -> " + strconv.Itoa(A))
		color.Cyan("\t\t[!] Gender -> " + G)
		color.Cyan("\t\t[!] Rank -> " + R)
		color.Cyan("\t\t[!] ID -> " + StudentID)

		for _, skill := range strings.Split(SK, " ")[0 : len(strings.Split(SK, " "))-1] {
			if skill == "NONE" {
				break
			}

			color.Magenta("\t\t[!] Skill -> " + skill)
		}

		for _, data := range strings.Split(Priv8, " ")[0 : len(strings.Split(Priv8, " "))-1] {
			color.Green("\t\t[!] " + strings.Split(data, ":")[0] + " -> " + strings.Split(data, ":")[1])
		}

	}
}

func ShowAll() {
	var (
		FN        string = ""
		LN        string = ""
		A         int    = 0
		G         string = ""
		SK        string = ""
		R         string = ""
		Priv8     string = ""
		StudentID string = ""
	)
	result, err := config.Database.Query("SELECT * FROM Employees")

	if err != nil {
		config.SetLog("E", "employee.employee.ShowAll() -> Couldn't Execute Query")
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&FN, &LN, &A, &G, &SK, &R, &Priv8, &StudentID)

		if err != nil {
			config.SetLog("E", "employee.employee.ShowAll() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}

		color.Cyan("\t\t[!] First Name -> " + FN)
		color.Cyan("\t\t[!] Last Name -> " + LN)
		color.Cyan("\t\t[!] Age -> " + strconv.Itoa(A))
		color.Cyan("\t\t[!] Gender -> " + G)
		color.Cyan("\t\t[!] Rank -> " + R)
		color.Cyan("\t\t[!] ID -> " + StudentID)

		for _, skill := range strings.Split(SK, " ")[0 : len(strings.Split(SK, " "))-1] {
			if skill == "NONE" {
				break
			}

			color.Magenta("\t\t[!] Skill -> " + skill)
		}

		for _, data := range strings.Split(Priv8, " ")[0 : len(strings.Split(Priv8, " "))-1] {
			color.Green("\t\t[!] " + strings.Split(data, ":")[0] + " -> " + strings.Split(data, ":")[1])
		}
		color.Red("\n\t===============================================\n\n")
		time.Sleep(time.Second * 1)
	}

	time.Sleep(time.Second * 5)

}

func JsonOutput() {
	var (
		FN         string   = ""
		LN         string   = ""
		A          int      = 0
		G          string   = ""
		SK         string   = ""
		R          string   = ""
		Priv8      string   = ""
		EmployeeID string   = ""
		data       Employee = Employee{}
		Counter    int      = Counter() - 1
	)
	result, _ := config.Database.Query("SELECT * FROM Employees")

	if err != nil {
		config.SetLog("E", "employee.employee.JsonOutput() -> Couldn't Scan Result Of Query")
		config.SetLog("D", err.Error())
	}

	err = os.WriteFile("json/Employees.json", []byte{'\n'}, 0644)
	if err != nil {
		config.SetLog("E", "employee.employee.JsonOutput() -> Couldn't Clear File")
		config.SetLog("D", err.Error())
	}

	file, err := os.OpenFile("json/Employees.json", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		config.SetLog("E", "employee.employee.JsonOutput() -> Couldn't Open File"+config.PWD)
		config.SetLog("D", err.Error())
	}

	_, err = fmt.Fprintln(file, "[")

	for result.Next() {
		var (
			temp      []string          = []string{}
			tempSK    []string          = []string{}
			tempPriv8 map[string]string = map[string]string{}
		)
		err = result.Scan(&FN, &LN, &A, &G, &SK, &R, &Priv8, &EmployeeID)

		if err != nil {
			config.SetLog("E", "employee.employee.JsonOutput() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}

		tempSK = strings.Split(SK, " ")

		temp = strings.Split(Priv8, " ")
		for _, data := range temp[0 : len(strings.Split(Priv8, " "))-1] {
			tempPriv8[strings.Split(data, ":")[0]] = strings.Split(data, ":")[1]
		}

		data = Employee{
			FirstName: FN,
			LastName:  LN,
			Age:       A,
			Gender:    G,
			Rank:      R,
			Skill:     tempSK[0 : len(tempSK)-1],
			Private:   tempPriv8,
			ID:        EmployeeID,
		}

		res, err := json.Marshal(data)
		if err != nil {
			config.SetLog("E", "employee.employee.JsonOutput() -> Couldn't Convert Data To Json")
			config.SetLog("D", err.Error())
		}
		if Counter == 0 {
			_, err = fmt.Fprintln(file, string(res))
		} else {
			Counter -= 1
			_, err = fmt.Fprintln(file, string(res)+",\n")
		}
	}
	_, err = fmt.Fprintln(file, "]")
}
