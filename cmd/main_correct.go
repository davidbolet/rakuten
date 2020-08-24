package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"cursmedia.com/rakuten/database"
)

func readDictionary() *map[string]string {
	file, err := os.Open("./resources/dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var diccionario = make(map[string]string)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		diccionario[arr[0]] = arr[1]
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return &diccionario
}

// Correct corrects the postTitle and PostExcerpt against a dictionary
func Correct() {
	dictionary := readDictionary()
	connector = database.Connect()
	defer connector.Close()
	posts := connector.GetPosts(300)
	for i, post := range *posts {
		fmt.Println(fmt.Sprintf("Corrigiendo post %d", i))
		post.PostExcerpt = replaceOnText(post.PostExcerpt, dictionary)
		post.PostTitle = replaceOnText(post.PostTitle, dictionary)
		connector.DB.Save(post)
	}

}

func replaceOnText(text string, dictionary *map[string]string) string {
	corrected := ""
	dict := *dictionary
	for _, word := range strings.Fields(text) {
		if val, ok := dict[strings.ToLower(word)]; ok {
			corrected = corrected + " " + val
		} else {
			corrected = corrected + " " + word
		}
	}
	return corrected
}
