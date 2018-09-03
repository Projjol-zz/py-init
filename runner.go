package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/logrusorgru/aurora"
)

func reEnterDetails(errorText string, cliMessage string, scanner *bufio.Scanner) string {
	var au aurora.Aurora
	au = aurora.NewAurora(true)
	fmt.Print(au.Red(errorText))
	var option string
	if errorText != ErrorName {
		scanner.Scan()
		option = scanner.Text()
	}

	if option == "y" || option == "Y" || errorText == ErrorName {
		scanner.Scan()
		return scanner.Text()
	}
	return ""
}

func sanitizeInput(input string) string {
	// regex to accept only
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal("Error is :", err)
	}
	return reg.ReplaceAllString(input, " ")
}

func confirmInput(details map[string]string, scanner *bufio.Scanner) {

	for k, v := range details {
		fmt.Printf("%s : %s \n", sanitizeInput(k), v)
	}
	fmt.Printf(ConfirmationMessage)
	scanner.Scan()
	if scanner.Text() == "n" {
		runner()
	} else {
		ex, err := os.Executable()
		if err != nil {
			log.Fatal("Error is :", err)
		}
		currPath := filepath.Dir(ex)
		nestedPath := currPath + "/" + details["name"] + "/" + details["name"]
		// fmt.Printf(nestedPath)
		os.MkdirAll(nestedPath, os.ModePerm)
		filePath := nestedPath + "/" + "__init__.py"
		_, err = os.Create(filePath)
		if err != nil {
			log.Fatal("Init creation error: ", err)
		}
		outerPath := currPath + "/" + details["name"] + "/"

		// setup.py
		setupPath := outerPath + "setup.py"
		setupFile, err := os.Create(setupPath)
		if err != nil {
			log.Fatal("Setup creation error: ", err)
		}
		defer setupFile.Close()

		// README
		readmePath := outerPath + "README"
		_, err = os.Create(readmePath)
		if err != nil {
			log.Fatal("README creation error: ", err)
		}

		// License
		licensePath := outerPath + "LICENSE"
		_, err = os.Create(licensePath)
		if err != nil {
			log.Fatal("License creation error: ", err)
		}

		// write to setup.py
		writer := bufio.NewWriter(setupFile)
		fmt.Fprintf(writer, setupText,
			details["name"],
			details["version"],
			details["author"],
			details["author_email"],
			details["short_description"],
			details["long_description"],
			details["url"])

		writer.Flush()
	}

}

func runner() {
	fmt.Println(Opening)
	fmt.Printf(PackageName)

	packageDetails := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	packageDetails["name"] = scanner.Text()

	if len(packageDetails["name"]) == 0 {
		packageDetails["name"] = reEnterDetails(ErrorName, "package name: ", scanner)
	}

	fmt.Printf("version number (default 1.0.0) : ")
	scanner.Scan()
	packageDetails["version"] = scanner.Text()

	if len(packageDetails["version"]) == 0 {
		packageDetails["version"] = "1.0.0"
	}

	fmt.Printf("Author: ")
	scanner.Scan()
	packageDetails["author"] = scanner.Text()

	if len(packageDetails["author"]) == 0 {
		packageDetails["author"] = reEnterDetails(ErrorAuthor, "Author: ", scanner)
	}

	fmt.Printf("Author Email: ")
	scanner.Scan()
	packageDetails["author_email"] = scanner.Text()

	if len(packageDetails["author_email"]) == 0 {
		packageDetails["author_email"] = reEnterDetails(ErrorAuthorEmail, "Author Email: ", scanner)
	}

	fmt.Printf("Short Description: ")
	scanner.Scan()
	packageDetails["short_description"] = scanner.Text()

	fmt.Printf("Long Description: ")
	scanner.Scan()
	packageDetails["long_description"] = scanner.Text()

	fmt.Printf("Url: ")
	scanner.Scan()
	packageDetails["url"] = scanner.Text()

	fmt.Printf("\n\n")
	confirmInput(packageDetails, scanner)
}
