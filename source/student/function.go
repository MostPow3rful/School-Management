package student

import (
	"github.com/JesusKian/School-Management/source/config"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"

	"os"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var err error = nil

func Counter() int {
	var count int = 0

	result, err := config.Database.Query("SELECT COUNT(*) FROM Students")

	if err != nil {
		config.SetLog("E", "student.Counter() -> Couldn't Count Students")
		config.SetLog("D", err.Error())
		return -1
	}

	for result.Next() {
		err = result.Scan(&count)

		if err != nil {
			config.SetLog("E", "student.Counter() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
			return -1
		}

	}

	return count
}

func Add(FN string, LN string, A int, G string, SC map[string]string, SK []string, Priv8 map[string]string) (*Student, error) {
	var (
		dbSC    string = ""
		dbSK    string = ""
		dbPriv8 string = ""
	)

	var hour, min, sec int = time.Now().Clock()

	if A == 0 {
		config.SetLog("E", "student.Add() -> Invalid Age")
		return nil, errors.New("Add() -> Invalid Age")
	}

	for lesson, score := range SC {
		dbSC += fmt.Sprintf("%s:%s ", lesson, score)
	}

	for _, skill := range SK {
		dbSK += fmt.Sprintf("%s ", skill)
	}

	for question, answer := range Priv8 {
		dbPriv8 += fmt.Sprintf("%s:%s ", question, answer)
	}

	_, err = config.Database.Query(`
	INSERT INTO Students 
	(FirstName, LastName, Age, Gender, Skill, Score, Private, ID) 
	VALUES 
	(?, ?, ?, ?, ?, ?, ?, ?)`, FN, LN, A, G, dbSK, dbSC, dbPriv8, string(FN[0])+strconv.Itoa(hour)+strconv.Itoa(min)+strconv.Itoa(sec))

	if err != nil {
		config.SetLog("E", "student.Add() -> Couldn't Add Student's Data In Database")
		config.SetLog("D", err.Error())
	}

	return &Student{
		FirstName: FN,
		LastName:  LN,
		Age:       A,
		Gender:    G,
		Skill:     SK,
		Score:     SC,
		Private:   Priv8,
		ID:        string(FN[0]) + strconv.Itoa(hour) + strconv.Itoa(min) + strconv.Itoa(sec),
	}, nil
}

func (s *Student) ShowNow() {
	color.Green("\t\t[!] Student Secceful Added !")
	color.Cyan("\t\t[!] First Name -> " + s.FirstName)
	color.Cyan("\t\t[!] Last Name -> " + s.LastName)
	color.Cyan("\t\t[!] Age -> " + strconv.Itoa(s.Age))
	color.Cyan("\t\t[!] Gender -> " + s.Gender)
	color.Cyan("\t\t[!] ID -> " + s.ID)

	for _, skill := range s.Skill {
		if skill == "NONE" {
			break
		}
		color.Magenta("\t\t[!] Skill -> " + skill)
	}

	for lesson, score := range s.Score {
		if lesson == "NONE" {
			break
		}
		color.Green("\t\t[!] " + lesson + " -> " + score)
	}

	for question, answer := range s.Private {
		color.Red("\t\t[!] " + question + " -> " + answer)
	}

}

func ShowSkill(ID string) {
	var (
		skillValue string   = ""
		tempSlice  []string = []string{}
	)

	result, err := config.Database.Query("SELECT Skill FROM Students WHERE ID=?", ID)

	if err != nil {
		config.SetLog("E", "student.ShowSkill() -> Couldn't Get Students's Skill")
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&skillValue)

		if err != nil {
			config.SetLog("E", "student.ShowSkill() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}
	}

	tempSlice = strings.Split(skillValue, " ")

	for i := 0; i < len(tempSlice)-1; i++ {
		color.Cyan("\t" + "[" + strconv.Itoa(i+1) + "] -> " + tempSlice[i])
	}
}

func ShowScore(ID string) {
	var (
		scoreValue string   = ""
		tempArray  []string = []string{}
	)

	result, err := config.Database.Query("SELECT Score FROM Students WHERE ID=?", ID)

	if err != nil {
		config.SetLog("E", "student.ShowScore() -> Couldn't Get Students's Score With ID="+ID)
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&scoreValue)

		if err != nil {
			config.SetLog("E", "student.ShowScore() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}
	}

	tempArray = strings.Split(scoreValue, " ")

	for i := 0; i < len(tempArray)-1; i++ {
		color.Cyan("\t" + "[" + strconv.Itoa(i+1) + "] -> " + tempArray[i])
	}
}

func EditFirstName(ID string, value string) {
	_, err = config.Database.Query("UPDATE Students SET FirstName=? WHERE ID=?", value, ID)

	if err != nil {
		config.SetLog("E", "student.EditFirstName() -> Couldn't Edit Student's First Name With ID="+ID)
		config.SetLog("D", err.Error())
	} else {
		config.SetLog("E", "student.EditFirstName() -> Student With ID="+ID+" Changed First Name To "+value)
	}
}

func EditLastName(ID string, value string) {
	_, err = config.Database.Query("UPDATE Students SET LastName=? WHERE ID=?", value, ID)

	if err != nil {
		config.SetLog("E", "student.EditLastName() -> Couldn't Edit Student's Last Name With ID="+ID)
		config.SetLog("D", err.Error())
	} else {
		config.SetLog("E", "student.EditLastName() -> Student With ID="+ID+" Changed Last Name To "+value)
	}
}

func EditAge(ID string, value int, stringValue string) {
	_, err = config.Database.Query("UPDATE Students SET Age=? WHERE ID=?", value, ID)

	if err != nil {
		config.SetLog("E", "student.EditAge() -> Couldn't Edit Student's Age")
		config.SetLog("D", err.Error())
	} else {
		config.SetLog("E", "student.EditAge() -> Student With ID="+ID+" Changed Age To "+stringValue)
	}
}

func EditGender(ID string, value string) {
	_, err = config.Database.Query("UPDATE Students SET Gender=? WHERE ID=?", value, ID)

	if err != nil {
		config.SetLog("E", "student.EditGender() -> Couldn't Edit Student's Gender With ID="+ID)
		config.SetLog("D", err.Error())
	} else {
		config.SetLog("E", "student.EditGender() -> Student With ID="+ID+" Changed Gender To "+value)
	}
}

func EditSkill(ID string, option string, value string) {
	var (
		skillsLen int      = 0
		skills    string   = ""
		tempSlice []string = []string{}
	)

	result, err := config.Database.Query("SELECT Skill FROM Students WHERE ID=?", ID)

	if err != nil {
		config.SetLog("E", "student.EditSkill() -> Couldn't Get Students's Skill With ID="+ID)
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&skills)

		if err != nil {
			config.SetLog("E", "student.EditSkill() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}
	}

	tempSlice = strings.Split(skills, " ")

	switch option {
	case "1": // Add
		result, err := config.Database.Query("SELECT COUNT(Skill) FROM Students WHERE ID=?", ID)
		if err != nil {
			config.SetLog("E", "student.EditSkill() -> Couldn't Count Student's Skill With ID="+ID)
			config.SetLog("D", err.Error())
		}

		for result.Next() {
			_, err = fmt.Scan(&skillsLen)
			if err != nil {
				config.SetLog("E", "student.EditSkill() -> Couldn't Get Length Of Student's Skill With ID="+ID)
				config.SetLog("D", err.Error())
			}
		}

		skills += strconv.Itoa(skillsLen) + ":" + value + " "
		result, err = config.Database.Query("UPDATE Students SET Skill=? WHERE ID=?", skills, ID)

		if err != nil {
			config.SetLog("E", "student.EditSkill() -> Couldn't Add New Student's Skill With ID="+ID)
			config.SetLog("D", err.Error())
		}

	case "2": // Remove
		skills = ""

		for i := 0; i < len(tempSlice)-1; i++ {
			if strconv.Itoa(i+1) != value {
				skills += string(tempSlice[i] + " ")
			}
		}

		result, err = config.Database.Query("UPDATE Students SET Skill=? WHERE ID=?", skills, ID)

		if err != nil {
			config.SetLog("E", "student.EditSkill() -> Couldn't Remove Student's Skill With ID="+ID)
			config.SetLog("D", err.Error())
		}
	}
}

func EditScore(ID string, option string, key string, value string, Remove string) {
	var (
		scores    string   = ""
		tempSlice []string = []string{}
	)

	result, err := config.Database.Query("SELECT Score From Students Where ID=?", ID)

	if err != nil {
		config.SetLog("E", "student.EditScore() -> Couldn't Get Student's Score With ID="+ID)
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&scores)

		if err != nil {
			config.SetLog("E", "student.EditScore() -> Couldn't Scan Student's Score With ID="+ID)
			config.SetLog("D", err.Error())
		}
	}

	tempSlice = strings.Split(scores, " ")

	switch option {
	case "1": // Add
		scores += key + ":" + value + " "
		result, err = config.Database.Query("UPDATE Students SET Score=? WHERE ID=?", scores, ID)

		if err != nil {
			config.SetLog("E", "student.EditScore() -> Couldn't Add New Score With ID="+ID)
			config.SetLog("D", err.Error())
		} else {
			config.SetLog("E", "student.EditScore() -> Student With ID="+ID+" Added New Score "+key+":"+value)
		}

	case "2": // Edit
		scores = ""

		index, err := strconv.Atoi(key)

		for i := 0; i < len(tempSlice)-1; i++ {
			if strings.Split(tempSlice[i], ":")[0] == strings.Split(tempSlice[index-1], ":")[0] {
				scores += strings.Split(tempSlice[index-1], ":")[0] + ":" + value + " "
				continue
			}
			scores += tempSlice[i] + " "
		}

		result, err = config.Database.Query("UPDATE Students SET Score=? WHERE ID=?", scores, ID)

		if err != nil {
			config.SetLog("E", "student.EditScore() -> Couldn't Edit Student's Score With ID="+ID)
			config.SetLog("D", err.Error())
		} else {
			config.SetLog("E", "student.EditScore() -> Student With ID="+ID+" Edited Score Of "+strings.Split(tempSlice[index-1], ":")[0]+" Lesson")
		}

	case "3": // Remove
		scores = ""

		for i := 0; i < len(tempSlice)-1; i++ {
			if strconv.Itoa(i) != Remove {
				scores += tempSlice[i] + " "
			}
		}

		result, err = config.Database.Query("UPDATE Students SET Score=? WHERE ID=?", scores, ID)

		if err != nil {
			config.SetLog("E", "student.EditScore() -> Couldn't Remove Student's Score With ID="+ID)
			config.SetLog("D", err.Error())
		} else {
			config.SetLog("E", "student.EditScore() -> Student With ID="+ID+" Removed Score")
		}

	}
}

func ShowPrivateInformation(ID string) {
	var (
		answers   string   = ""
		tempArray []string = []string{}
	)

	result, err := config.Database.Query("SELECT Private FROM Students WHERE ID=?", ID)

	if err != nil {
		config.SetLog("E", "student.ShowPrivateInformation() -> Couldn't Get Students's Private Information With ID="+ID)
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&answers)

		if err != nil {
			config.SetLog("E", "student.ShowPrivateInformation() -> Couldn't Scan Result Of Query")
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

	result, err := config.Database.Query("SELECT Private From Students Where ID=?", ID)

	if err != nil {
		config.SetLog("E", "student.EditPrivateInformation() -> Couldn't Get Student's Private Information With ID="+ID)
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&priv8)

		if err != nil {
			config.SetLog("E", "student.EditPrivateInformation() -> Couldn't Scan Student's Private Information With ID="+ID)
			config.SetLog("D", err.Error())
		}
	}

	tempSlice = strings.Split(priv8, " ")

	switch option {
	case "1": // Add
		priv8 += key + ":" + value + " "
		result, err = config.Database.Query("UPDATE Students SET Private=? WHERE ID=?", priv8, ID)

		if err != nil {
			config.SetLog("E", "student.EditPrivateInformation() -> Couldn't Add New Private Information With ID="+ID)
			config.SetLog("D", err.Error())
		} else {
			config.SetLog("E", "student.EditPrivateInformation() -> Student With ID="+ID+" Added New Private Information "+key+":"+value)
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

		result, err = config.Database.Query("UPDATE Students SET Private=? WHERE ID=?", priv8, ID)

		if err != nil {
			config.SetLog("E", "student.EditPrivateInformation() -> Couldn't Edit Student's Private Information With ID="+ID)
			config.SetLog("D", err.Error())
		} else {
			config.SetLog("E", "student.EditPrivateInformation() -> Student With ID="+ID+" Edited Private Information Of "+strings.Split(tempSlice[index-1], ":")[0])
		}

	case "3": // Remove
		priv8 = ""

		for i := 0; i < len(tempSlice)-1; i++ {
			if strconv.Itoa(i) != Remove {
				priv8 += tempSlice[i] + " "
			}
		}

		result, err = config.Database.Query("UPDATE Students SET Private=? WHERE ID=?", priv8, ID)

		if err != nil {
			config.SetLog("E", "student.EditPrivateInformation() -> Couldn't Remove Student's Private Information With ID="+ID)
			config.SetLog("D", err.Error())
		} else {
			config.SetLog("E", "student.EditPrivateInformation() -> Student With ID="+ID+" Removed Private Information")
		}

	}
}

func Remove(ID string) {
	_, err := config.Database.Query("DELETE FROM Students WHERE ID=?", ID)
	if err != nil {
		config.SetLog("E", "student.Remove() -> Couldn't Remove Student With ID="+ID)
		config.SetLog("D", err.Error())
	} else {
		config.SetLog("E", "student.Remove() -> Student With ID="+ID+" Removed")
	}
}

func Show(ID string) {
	var (
		FN        string = ""
		LN        string = ""
		A         int    = 0
		G         string = ""
		SK        string = ""
		SC        string = ""
		Priv8     string = ""
		StudentID string = ""
	)

	result, err := config.Database.Query("SELECT * FROM Students WHERE ID=?", ID)

	if err != nil {
		config.SetLog("E", "student.Show() -> Couldn't Get Student's Info With ID="+ID)
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&FN, &LN, &A, &G, &SK, &SC, &Priv8, &StudentID)

		if err != nil {
			config.SetLog("E", "student.Show() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}

		color.Cyan("\t\t[!] First Name -> " + FN)
		color.Cyan("\t\t[!] Last Name -> " + LN)
		color.Cyan("\t\t[!] Age -> " + strconv.Itoa(A))
		color.Cyan("\t\t[!] Gender -> " + G)
		color.Cyan("\t\t[!] ID -> " + StudentID)

		for _, skill := range strings.Split(SK, " ")[0 : len(strings.Split(SK, " "))-1] {
			if skill == "NONE" {
				break
			}

			color.Magenta("\t\t[!] Skill -> " + skill)
		}

		for _, data := range strings.Split(SC, " ")[0 : len(strings.Split(SC, " "))-1] {
			color.Green("\t\t[!] " + strings.Split(data, ":")[0] + " -> " + strings.Split(data, ":")[1])
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
		SC        string = ""
		Priv8     string = ""
		StudentID string = ""
	)
	result, err := config.Database.Query("SELECT * FROM Students")

	if err != nil {
		config.SetLog("E", "student.ShowAll() -> Couldn't Execute Query")
		config.SetLog("D", err.Error())
	}

	for result.Next() {
		err = result.Scan(&FN, &LN, &A, &G, &SK, &SC, &Priv8, &StudentID)

		if err != nil {
			config.SetLog("E", "student.ShowAll() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}

		color.Cyan("\t\t[!] First Name -> " + FN)
		color.Cyan("\t\t[!] Last Name -> " + LN)
		color.Cyan("\t\t[!] Age -> " + strconv.Itoa(A))
		color.Cyan("\t\t[!] Gender -> " + G)
		color.Cyan("\t\t[!] ID -> " + StudentID)

		for _, skill := range strings.Split(SK, " ")[0 : len(strings.Split(SK, " "))-1] {
			if skill == "NONE" {
				break
			}

			color.Magenta("\t\t[!] Skill -> " + skill)
		}

		for _, data := range strings.Split(SC, " ")[0 : len(strings.Split(SC, " "))-1] {
			color.Green("\t\t[!] " + strings.Split(data, ":")[0] + " -> " + strings.Split(data, ":")[1])
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
		FN        string  = ""
		LN        string  = ""
		A         int     = 0
		G         string  = ""
		SK        string  = ""
		SC        string  = ""
		Priv8     string  = ""
		StudentID string  = ""
		data      Student = Student{}
		Counter   int     = Counter() - 1
	)
	result, _ := config.Database.Query("SELECT * FROM Students")

	if err != nil {
		config.SetLog("E", "student.JsonOutput() -> Couldn't Scan Result Of Query")
		config.SetLog("D", err.Error())
	}

	err = os.WriteFile("json/Students.json", []byte{'\n'}, 0644)
	if err != nil {
		config.SetLog("E", "student.JsonOutput() -> Couldn't Clear File")
		config.SetLog("D", err.Error())
	}

	file, err := os.OpenFile("json/Students.json", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		config.SetLog("E", "student.JsonOutput() -> Couldn't Open File"+config.PWD)
		config.SetLog("D", err.Error())
	}

	_, err = fmt.Fprintln(file, "[")

	for result.Next() {
		var (
			temp      []string          = []string{}
			tempSK    []string          = []string{}
			tempSC    map[string]string = map[string]string{}
			tempPriv8 map[string]string = map[string]string{}
		)
		err = result.Scan(&FN, &LN, &A, &G, &SK, &SC, &Priv8, &StudentID)

		if err != nil {
			config.SetLog("E", "student.JsonOutput() -> Couldn't Scan Result Of Query")
			config.SetLog("D", err.Error())
		}

		tempSK = strings.Split(SK, " ")

		temp = strings.Split(SC, " ")
		for _, data := range temp[0 : len(strings.Split(SC, " "))-1] {
			tempSC[strings.Split(data, ":")[0]] = strings.Split(data, ":")[1]
		}

		temp = strings.Split(Priv8, " ")
		for _, data := range temp[0 : len(strings.Split(Priv8, " "))-1] {
			tempPriv8[strings.Split(data, ":")[0]] = strings.Split(data, ":")[1]
		}

		data = Student{
			FirstName: FN,
			LastName:  LN,
			Age:       A,
			Gender:    G,
			Skill:     tempSK[0 : len(tempSK)-1],
			Score:     tempSC,
			Private:   tempPriv8,
			ID:        StudentID,
		}

		res, err := json.Marshal(data)
		if err != nil {
			config.SetLog("E", "student.JsonOutput() -> Couldn't Convert Data To Json")
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
