clear
read -p "Please Enter Your MySQL Password (Default=System Password)-(Need For Config) -> " PASSWORD

while [ true ]; do
    # Clear Terminal
    clear

    # Show Banner
    echo "$(tput setaf 5) ____       _                 _ $(tput sgr0)"
    sleep 0.1s
    echo "$(tput setaf 5)/ ___|  ___| |__   ___   ___ | |$(tput sgr0)"
    sleep 0.1s
    echo "$(tput setaf 5)\___ \ / __| |_ \ / _ \ / _ \| |$(tput sgr0)"
    sleep 0.1s
    echo "$(tput setaf 5) ___) | (__| | | | (_) | (_) | |$(tput sgr0)"
    sleep 0.1s
    echo "$(tput setaf 5)|____/ \___|_| |_|\___/ \___/|_|$(tput sgr0)"
    echo ""
    echo ""

    # Show Distro's Menu
    echo "$(tput setaf 2)[?] Select Your Distro$(tput sgr0)"
    sleep 0.1s
    echo "$(tput setaf 1)[1] $(tput sgr0)" "$(tput setaf 6)Debian$(tput sgr0)"
    sleep 0.1s
    echo "$(tput setaf 1)[2] $(tput sgr0)" "$(tput setaf 6)Arch$(tput sgr0)"
    sleep 0.1s
    echo "$(tput setaf 1)[3] $(tput sgr0)" "$(tput setaf 6)Fedora$(tput sgr0)"
    sleep 0.1s
    echo "$(tput setaf 1)[4] $(tput sgr0)" "$(tput setaf 6)Another . . .$(tput sgr0)"
    echo ""

    # Get Distro
    read -p ">>> " DISTRO
    echo ""

    # Checking Packages
    case "${DISTRO}" in
    1)
        if [ "$(command -v mysql)" ]; then
            echo "$(tput setaf 2)[+] $(tput sgr0)" "$(tput setaf 6)MySQL Is Available On Your System$(tput sgr0)"
            echo ""
            sleep 1s
            echo "$(tput setaf 2)[+] $(tput sgr0)" "$(tput setaf 6)Executing MySQL Query$(tput sgr0)"

            # Execute MySQL Command
            mysql -u root -p"$PASSWORD" -N -e "CREATE DATABASE IF NOT EXISTS SchoolManagement;USE SchoolManagement;CREATE TABLE if not exists MYSQL_USER_PASS(username VARCHAR(300), password VARCHAR(300));INSERT INTO MYSQL_USER_PASS (username,password) VALUES ('root', '$PASSWORD');"
            mysql -u root -p"$PASSWORD" -N -e "use SchoolManagement;CREATE TABLE IF NOT EXISTS Students (FirstName VARCHAR(300), LastName VARCHAR(300), Age INT, Gender VARCHAR(6), Skill VARCHAR(1000), Score VARCHAR(1000), Private VARCHAR(1000), ID VARCHAR(10));"
            mysql -u root -p"$PASSWORD" -N -e "use SchoolManagement;CREATE TABLE IF NOT EXISTS Employees (FirstName VARCHAR(300), LastName VARCHAR(300), Age INT, Gender VARCHAR(6), Skill VARCHAR(1000), Rank VARCHAR(100), Private VARCHAR(1000), ID VARCHAR(10));"
            Counter=0
            for data in $(echo $(mysql -s -u root -p01 SchoolManagement -N -e "SELECT username, password FROM MYSQL_USER_PASS GROUP BY username, password;") | tr " " "\n"); do
                if [[ $Counter == 1 ]]; then
                    echo -e "\t\"password\":\"$data\"\n}" >>./json/Secret.json
                    break
                fi

                echo -e "{\n\t\"username\":\"$data\"," >./json/Secret.json
                Counter=1

            done
            break

        else
            echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)Please Install MySQL Package On Your System With -> $(tput sgr0)""$(tput setaf 2)dpkg , apt , ...$(tput sgr0)"
            sleep 3s
            exit 1

        fi
        ;;

    2)
        if [ "$(command -v mysql)" ]; then
            echo "$(tput setaf 2)[+] $(tput sgr0)" "$(tput setaf 6)MySQL Is Available On Your System$(tput sgr0)"
            echo ""
            sleep 1s
            echo "$(tput setaf 2)[+] $(tput sgr0)" "$(tput setaf 6)Executing MySQL Query$(tput sgr0)"

            # Execute MySQL Command
            mysql -u root -p"$PASSWORD" -N -e "CREATE DATABASE IF NOT EXISTS SchoolManagement;USE SchoolManagement;CREATE TABLE if not exists MYSQL_USER_PASS(username VARCHAR(300), password VARCHAR(300));INSERT INTO MYSQL_USER_PASS (username,password) VALUES ('root', '$PASSWORD');"
            mysql -u root -p"$PASSWORD" -N -e "use SchoolManagement;CREATE TABLE IF NOT EXISTS Students (FirstName VARCHAR(300), LastName VARCHAR(300), Age INT, Gender VARCHAR(6), Skill VARCHAR(1000), Score VARCHAR(1000), Private VARCHAR(1000), ID VARCHAR(10));"
            mysql -u root -p"$PASSWORD" -N -e "use SchoolManagement;CREATE TABLE IF NOT EXISTS Employees (FirstName VARCHAR(300), LastName VARCHAR(300), Age INT, Gender VARCHAR(6), Skill VARCHAR(1000), Rank VARCHAR(100), Private VARCHAR(1000), ID VARCHAR(10));"
            Counter=0
            for data in $(echo $(mysql -s -u root -p01 SchoolManagement -N -e "SELECT username, password FROM MYSQL_USER_PASS GROUP BY username, password;") | tr " " "\n"); do
                if [[ $Counter == 1 ]]; then
                    echo -e "\t\"password\":\"$data\"\n}" >>./json/Secret.json
                    break
                fi

                echo -e "{\n\t\"username\":\"$data\"," >./json/Secret.json
                Counter=1

            done
            break

        else
            echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)Please Install MySQL Package On Your System With -> $(tput sgr0)""$(tput setaf 2)pacman , yay , ...$(tput sgr0)"
            sleep 3s
            exit 1
        fi

        ;;

    3)
        if [ "$(command -v mysql)" ]; then
            echo "$(tput setaf 2)[+] $(tput sgr0)" "$(tput setaf 6)MySQL Is Available On Your System$(tput sgr0)"
            echo ""
            sleep 1s
            echo "$(tput setaf 2)[+] $(tput sgr0)" "$(tput setaf 6)Executing MySQL Query$(tput sgr0)"

            # Execute MySQL Command
            mysql -u root -p"$PASSWORD" -N -e "CREATE DATABASE IF NOT EXISTS SchoolManagement;USE SchoolManagement;CREATE TABLE if not exists MYSQL_USER_PASS(username VARCHAR(300), password VARCHAR(300));INSERT INTO MYSQL_USER_PASS (username,password) VALUES ('root', '$PASSWORD');"
            mysql -u root -p"$PASSWORD" -N -e "use SchoolManagement;CREATE TABLE IF NOT EXISTS Students (FirstName VARCHAR(300), LastName VARCHAR(300), Age INT, Gender VARCHAR(6), Skill VARCHAR(1000), Score VARCHAR(1000), Private VARCHAR(1000), ID VARCHAR(10));"
            mysql -u root -p"$PASSWORD" -N -e "use SchoolManagement;CREATE TABLE IF NOT EXISTS Employees (FirstName VARCHAR(300), LastName VARCHAR(300), Age INT, Gender VARCHAR(6), Skill VARCHAR(1000), Rank VARCHAR(100), Private VARCHAR(1000), ID VARCHAR(10));"
            Counter=0
            for data in $(echo $(mysql -s -u root -p01 SchoolManagement -N -e "SELECT username, password FROM MYSQL_USER_PASS GROUP BY username, password;") | tr " " "\n"); do
                if [[ $Counter == 1 ]]; then
                    echo -e "\t\"password\":\"$data\"\n}" >>./json/Secret.json
                    break
                fi

                echo -e "{\n\t\"username\":\"$data\"," >./json/Secret.json
                Counter=1

            done
            break

        else
            echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)Please Install MySQL Package On Your System With -> $(tput sgr0)""$(tput setaf 2)dnf , yum , ...$(tput sgr0)"
            sleep 3s
            exit 1
        fi
        ;;

    4)
        echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)Please Install MySQL Package On Your System With Your Package Manager$(tput sgr0)"
        sleep 3s
        exit 1
        ;;

    *)
        continue
        ;;
    esac

done

echo ""
sleep 1s

# /log Dir
if [ -d "$PWD/log" ]; then
    echo "$(tput setaf 2)[+] $(tput sgr0)""$(tput setaf 6)/log Directory Is Valid$(tput sgr0)"
    sleep 1s
    echo ""

    # /log/log.log File
    if [ -f "$PWD/log/log.log" ]; then
        echo "$(tput setaf 2)[+] $(tput sgr0)""$(tput setaf 6)/log/log.log File Is Valid$(tput sgr0)"
        sleep 1s
        echo ""
    else
        echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)/log/log.log DirectoryFile Isn't Valid !$(tput sgr0)"
        sleep 3s
        exit 1
    fi

else
    echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)/log Directory Isn't Valid !$(tput sgr0)"
    sleep 3s
    exit 1
fi

# /json Dir
if [ -d "$PWD/json" ]; then
    echo "$(tput setaf 2)[+] $(tput sgr0)""$(tput setaf 6)/json Directory Is Valid$(tput sgr0)"
    sleep 1s
    echo ""

    # /json/Students.json File
    if [ -f "$PWD/json/Students.json" ]; then
        echo "$(tput setaf 2)[+] $(tput sgr0)""$(tput setaf 6)/json/Students.json File Is Valid$(tput sgr0)"
        sleep 1s
        echo ""
    else
        echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)/json/Students.json File Isn't Valid !$(tput sgr0)"
        sleep 3s
        exit 1
    fi

    # /json/Employees.json File
    if [ -f "$PWD/json/Employees.json" ]; then
        echo "$(tput setaf 2)[+] $(tput sgr0)""$(tput setaf 6)/json/Employees.json File Is Valid$(tput sgr0)"
        sleep 1s
        echo ""
    else
        echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)/json/Employees.json File Isn't Valid !$(tput sgr0)"
        sleep 3s
        exit 1
    fi

    # /json/Secret.json File
    if [ -f "$PWD/json/Secret.json" ]; then
        echo "$(tput setaf 2)[+] $(tput sgr0)""$(tput setaf 6)/json/Secret.json File Is Valid$(tput sgr0)"
        sleep 1s
        echo ""
    else
        echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)/json/Secret.json File Isn't Valid !$(tput sgr0)"
        sleep 3s
        exit 1
    fi

else
    echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)/json Directory Isn't Valid !$(tput sgr0)"
    sleep 3s
    exit 1
fi

# /main.go File
if [ -f "$PWD/main.go" ]; then
    echo "$(tput setaf 2)[+] $(tput sgr0)""$(tput setaf 6)main.go File Is Valid$(tput sgr0)"
    sleep 1s
    echo ""
else
    echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)main.go File Isn't Valid !$(tput sgr0)"
    sleep 3s
    exit 1
fi

# /go.mod File
if [ -f "$PWD/go.mod" ]; then
    echo "$(tput setaf 2)[+] $(tput sgr0)""$(tput setaf 6)go.mod File Is Valid$(tput sgr0)"
    sleep 1s
    echo ""
else
    echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)go.mod File Isn't Valid !$(tput sgr0)"
    sleep 3s
    exit 1
fi

# /go.sum File
if [ -f "$PWD/go.sum" ]; then
    echo "$(tput setaf 2)[+] $(tput sgr0)""$(tput setaf 6)go.sum File Is Valid$(tput sgr0)"
    sleep 1s
    echo ""
else
    echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)go.sum File Isn't Valid !$(tput sgr0)"
    sleep 3s
    exit 1
fi

# requirements.txt
if [ -f "$PWD/requirements.txt" ]; then
    echo "$(tput setaf 2)[+] $(tput sgr0)""$(tput setaf 6)requirements.txt File Is Valid$(tput sgr0)"
    sleep 1s
    echo ""
else
    echo "$(tput setaf 1)[-] $(tput sgr0)" "$(tput setaf 6)requirements.txt File Isn't Valid !$(tput sgr0)"
    sleep 3s
    exit 1
fi

echo ""

if [ "$(command -v go)" ]; then
    echo "$(tput setaf 2)[+]$(tput sgr0)""$(tput setaf 6)Go Is Available On Your System$(tput sgr0)"
    sleep 1s
    echo ""
else
    echo "$(tput setaf 1)[-]$(tput sgr0)""$(tput setaf 6)XXXXXXXX$(tput sgr0)"
    sleep 3s
    exit 1
fi

echo "$(tput setaf 2)[+] $(tput sgr0)""$(tput setaf 6)Making Executable File Of main.go$(tput sgr0)"
go build ./main.go
echo ""
echo ""
echo "$(tput setaf 1)[*] $(tput sgr0)""$(tput setaf 1)Press Enter To Close$(tput sgr0)"
read -p "" x
clear
