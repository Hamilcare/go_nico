package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Printf("hello, world\n")
	_, err := os.Create("test")
	check(err)

	// fmt.Printf(readWholeFileAsString("./resources/schubert.txt"))
	// readFileLineByLine("./resources/schubert.txt")
	readSomeByte("./resources/schubert.txt", 14)
}

// if error != nil print error message
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// return whole file content as a string
func readWholeFileAsString(filename string) string {

	fmt.Printf("readWholeFileAsString\n")
	data, erreur := ioutil.ReadFile(filename)
	check(erreur)

	var contentAsString = string(data)
	return contentAsString
}

func readSomeByte(filename string, nbByte int) {
	file, err := os.Open(filename)
	defer file.Close()
	check(err)

	b1 := make([]byte, nbByte)
	n1, err := file.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
}

func readFileLineByLine(filename string) {
	fmt.Printf("readFileLineByLine\n")

	//1 Ouvrir le fichier
	file, err := os.Open(filename)
	defer file.Close()
	check(err)

	//2 créer un objet reader pour parcourir le fichier
	reader := bufio.NewReader(file)

	var line string
	var allLines []string

	//3 lire le fichier ligne par ligne jusqu'à la fin du fichier
	for {
		line, err = reader.ReadString('\n')
		fmt.Printf(" > Read %d characters\n", len(line))

		fmt.Printf(" > Line " + strings.ToUpper(line))

		if err != nil {
			allLines = append(allLines, line)
			break
		}
	}
	fmt.Printf("Nombre de lignes lues : %d \n", len(allLines))
}
