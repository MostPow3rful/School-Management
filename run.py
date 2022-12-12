import os
import subprocess
import mysql.connector

from time import sleep
from sys import exit as error
from colorama import (
    Fore,
    init
)


class Color:
    RED = Fore.RED
    GREEN = Fore.GREEN
    CYAN = Fore.CYAN
    MAGENTA = Fore.MAGENTA
    WHITE = Fore.WHITE


class Config(Color):
    def __init__(self, _password: str = "", _distro: int = 0) -> None:
        self._clear()
        print(self._banner())
        sleep(1)
        self._USERNAME: str = "root"

        # Checking Password
        if _password is str():
            error(f"{Config.RED}[!] {Config.CYAN}Invalid Password For MySQL")
        self._PASSWORD: str = _password

        # Checking Distro
        if _distro not in range(1, 5):
            error(f"{Config.RED}[!] {Config.CYAN}Invalid Distro")
        self._DISTRO: str = _distro

        # Checking MySQL Package
        self._mysql_help() if self._check_mysql() is False else print(
            f"{Config.GREEN}[+] {Config.CYAN}MySQL Package Is Valid On Your System"
        )
        sleep(1)

        # Checking GoLang package
        self._golang_help if self._check_golang() is False else print(
            f"{Config.GREEN}[+] {Config.CYAN}GoLang Package Is Valid On Your System"
        )
        sleep(1)

        # Checking Directories And Files
        self._check_directory_file()

        # Configing MySQL
        self._config_mysql()

        # Creating Executable File Of main.go
        print(
            f"{Config.GREEN}[+] {Config.CYAN}Creating Executable File Of main.go"
        )
        os.system("go build ./main.go")

    def _banner(self) -> str:
        return f"""
{Config.MAGENTA}  ____       _                 _ 
{Config.MAGENTA} / ___|  ___| |__   ___   ___ | |  
{Config.MAGENTA} \___ \ / __| '_ \ / _ \ / _ \| |  
{Config.MAGENTA}  ___) | (__| | | | (_) | (_) | |  
{Config.MAGENTA} |____/ \___|_| |_|\___/ \___/|_|
        """

    def _check_mysql(self) -> bool:
        RESULT: str = subprocess.run(
            "command -v mysql", shell=True, capture_output=True
        ).stdout.decode()

        return False if RESULT is str() else True

    def _check_golang(self) -> bool:
        RESULT: str = subprocess.run(
            "command -v go", shell=True, capture_output=True
        ).stdout.decode()

        return False if RESULT is str() else True

    def _golang_help(self) -> None:
        match self._DISTRO:
            case 1:  # Debian
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install GoLang Package On Your System With -> [dpkg , apt , ...]"
                )

            case 2:  # Arch
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install GoLang Package On Your System With -> [pacman , yay , ...]"
                )

            case 3:  # Fedora
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install GoLang Package On Your System With -> [dnf , yum , ...]"
                )

            case 4:  # Another
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install GoLang Package On Your System With Your Package Manager"
                )

    def _mysql_help(self) -> None:
        match self._DISTRO:
            case 1:  # Debian
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install MySQL Package On Your System With -> [dpkg , apt , ...]"
                )

            case 2:  # Arch
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install MySQL Package On Your System With -> [pacman , yay , ...]"
                )

            case 3:  # Fedora
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install MySQL Package On Your System With -> [dnf , yum , ...]"
                )

            case 4:  # Another
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install MySQL Package On Your System With Your Package Manager"
                )

    def _check_directory_file(self) -> None:
        DIRS: tuple = (
            "assets",
            "json",
            "log",
            "source",
            "source/student",
            "source/employee",
            "source/config",
            "source/banner",
        )

        FILES: tuple = (
            "go.mod",
            "go.sum",
            "main.go",
            "log/log.log",
            "json/Employees.json",
            "json/Secret.json",
            "json/Students.json",
            "requirements.txt",
            "source/banner/function.go",
            "source/config/function.go",
            "source/employee/function.go",
            "source/employee/struct.go",
            "source/student/function.go",
            "source/student/struct.go",
        )

        for directory in DIRS:
            error(
                f"{Config.RED}[-] {Config.CYAN}Invalid Directory ({directory})"
            ) if os.path.exists(directory) is False else print(
                f"{Config.GREEN}[+] {Config.CYAN}Valid Directory ({directory})"
            )
            sleep(1)

        for file in FILES:
            error(
                f"{Config.RED}[-] {Config.CYAN}Invalid File ({file})"
            ) if os.path.exists(file) is False else print(
                f"{Config.GREEN}[+] {Config.CYAN}Valid File ({file})"
            )
            sleep(1)

    def _config_mysql(self) -> None:
        database = mysql.connector.connect(
            host="localhost",
            user=self._USERNAME,
            password=self._PASSWORD
        )
        cursor = database.cursor()

        cursor.execute(
            f"CREATE database IF NOT EXISTS SchoolManagement"
        )
        database.commit()
        [... for _ in cursor]

        cursor.close()

        database = mysql.connector.connect(
            host="localhost",
            user="root",
            password=self._PASSWORD,
            database="SchoolManagement"
        )
        cursor = database.cursor()

        cursor.execute(
            "CREATE TABLE if not exists MYSQL_USER_PASS (username VARCHAR(300), password VARCHAR(300))"
        )

        cursor.execute(
            "INSERT INTO MYSQL_USER_PASS (username,password) VALUES (%s, %s)",
            (
                'root',
                self._PASSWORD
            )
        )
        database.commit()
        [... for _ in cursor]

        cursor.execute(
            "CREATE TABLE IF NOT EXISTS Students (FirstName VARCHAR(300), LastName VARCHAR(300), Age INT, Gender VARCHAR(6), Skill VARCHAR(1000), Score VARCHAR(1000), Private VARCHAR(1000), ID VARCHAR(10))"
        )
        database.commit()
        [... for _ in cursor]

        cursor.execute(
            "CREATE TABLE IF NOT EXISTS Employees (FirstName VARCHAR(300), LastName VARCHAR(300), Age INT, Gender VARCHAR(6), Skill VARCHAR(1000), Rank VARCHAR(100), Private VARCHAR(1000), ID VARCHAR(10))"
        )
        database.commit()
        [... for _ in cursor]

        cursor.execute(
            "SELECT username, password FROM MYSQL_USER_PASS GROUP BY username, password"
        )
        RESULT = cursor.fetchall()

        os.system('echo " " > log/log.log')
        with open(file="./json/Secret.json", mode="a") as file:
            file.write(
                '{\n\t"username":"'+RESULT[0][0] +
                '",\n\t"password":"'+RESULT[0][1]+'"\n}\n'
            )

    def _clear(self) -> None:
        os.system("clear")


def main() -> None:
    PASSWORD: str = input(
        f"""
        {Color.RED}[?] {Color.CYAN} Please Enter Your MySQL Password
        {Color.RED}[{Color.GREEN}Default=System Password{Color.RED}] {Color.WHITE}- {Color.RED}[{Color.GREEN}Need For Config{Color.RED}]
                    {Color.CYAN}-> """
    )

    try:
        KERNEL: int = int(
            input(f"""
        {Color.RED}[1] {Color.CYAN} Debian
        {Color.RED}[2] {Color.CYAN} Arch
        {Color.RED}[3] {Color.CYAN} Fedora
        {Color.RED}[4] {Color.CYAN} Another . . .
            
                    -> """)
        )
    except ValueError:
        error(
            f"{Color.RED}[!] {Color.CYAN}ValueError : U Must Enter Number In Range 1-4 In Kernel Option"
        )

    Config(
        PASSWORD,
        KERNEL
    )


if __name__ == "__main__":
    init(autoreset=True)
    main()


