package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello, world\n")

	fmt.Printf(readWholeFileAsString("./resources/schubert.txt"))
	readSomeChars("./resources/schubert.txt", 14)
	readFileLineByLine("./resources/schubert.txt")

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
	var contentAsString string

	//TODO Lire le fichier d'une seule tratie à l'aide de ioutil.ReadFile

	return contentAsString
}

func readSomeChars(filename string, nbCharacters int) {
	fmt.Printf("about to read %d characters", nbCharacters)

	//Etape 1 ouvrir le fichier à l'aide de os.Open

	//Etape 2 préparer un slice de longeur nbCharacters à l'aide de make

	//Etape 3 Lire les nbCharacters premiers caractères du fichier

	//Etape 4 afficher les caractères lus
}

func readFileLineByLine(filename string) {
	fmt.Printf("readFileLineByLine\n")

	//1 Ouvrir le fichier

	//2 créer un objet reader pour parcourir le fichier (cf bufio.NewReader)

	//3 lire le fichier ligne par ligne jusqu'à la fin du fichier à l'aide d'une boucle

	//4 afficher le nombre de lignes lues
}
