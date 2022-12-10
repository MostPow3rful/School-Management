package config

import (
	_ "github.com/go-sql-driver/mysql"

	"bufio"
	"database/sql"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	err error
	PWD string = ""

	// Database Object
	Database *sql.DB
	// Username Of MySQL To Login
	dbUsername = ""
	// Password Of MySQL To Login
	dbPassword = ""

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

func Clear() {
	cmd := exec.Command("clear")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		SetLog("E", "Clear() Function -> Couldn't Run 'clear' Command")
		SetLog("D", err.Error())
	}
}

func SetLog(logType string, msg string) {
	PWD, err = os.Getwd()
	if err != nil {
		Clear()
		errorLog.Println("SetLog() -> Couldn't Get Output Of 'os.Getwd()'")
		os.Exit(1)
	}

	logFile, err := os.OpenFile(PWD+"/log/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Clear()
		errorLog.Println(`
		Invalid Directory -> /log
		Invalid File -> /log/log.log
		`)
		os.Exit(1)
	}

	switch logType {
	case "I":
		infoLog.SetOutput(logFile)
		infoLog.Println(msg)

	case "W":
		warnLog.SetOutput(logFile)
		warnLog.Println(msg)

	case "E":
		errorLog.SetOutput(logFile)
		errorLog.Println(msg)

	case "D":
		defaultLog.SetOutput(logFile)
		defaultLog.Println(msg)

	default:
		errorLog.SetOutput(logFile)
		errorLog.Printf("SetLog() -> Trying To Add Log Without Valid LogType '%s'", logType)
	}

	logFile.Close()
}

func Config() error {
	var (
		dbCounter int      = 0
		temp      []string = []string{}
	)

	log.SetFlags(log.Ldate | log.Ltime)
	infoLog.SetFlags(log.Ldate | log.Ltime)
	warnLog.SetFlags(log.Ldate | log.Ltime)
	errorLog.SetFlags(log.Ldate | log.Ltime)
	SetLog("I", "Config() -> Flags Setuped")

	secret, err := os.Open("./json/Secret.json")
	if err != nil {
		SetLog("E", "Config() -> Can't Open ./json/Secret.json")
		SetLog("D", err.Error())
		errorLog.Fatal(err)
	}
	defer secret.Close()

	scanner := bufio.NewScanner(secret)
	for scanner.Scan() {
		if (scanner.Text() != "{") && (scanner.Text() != "}") {
			temp = strings.Split(scanner.Text(), "\":\"")

			for i := 0; i < len(temp[1])-1; i++ {
				if dbCounter == 1 {
					dbPassword += string(temp[1][i])
				} else {
					if string(temp[1][i]) == "\"" {
						dbCounter += 1
						continue
					}

					dbUsername += string(temp[1][i])
				}
			}

		}
	}

	err = scanner.Err()
	if err != nil {
		SetLog("E", "Config() -> Unknow Error From bufio.Scanner()")
		SetLog("D", err.Error())
		errorLog.Fatal(err)
	}

	Database, err = sql.Open("mysql", dbUsername+":"+dbPassword+"@tcp(0.0.0.0:3306)/SchoolManagement")
	if err != nil {
		SetLog("E", "Config() -> Can't Open SchoolManagement Database")
		SetLog("D", err.Error())
		errorLog.Fatal(err)
	}

	err = Database.Ping()
	if err != nil {
		SetLog("E", "Config() -> MySQL Dosn't Response")
		SetLog("D", err.Error())
		errorLog.Fatal(err)
	}

	return nil
}
