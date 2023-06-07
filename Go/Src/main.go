package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Url struct {
	web string
}

func SCRAP(data *http.Response) (string, error) {
	doc, err := goquery.NewDocumentFromReader(data.Body)
	if err != nil {
		return "", err
	}

	links := doc.Find("a").Text()
	return links, nil
}

func GET(url *Url) (string, error) {
	result, err := http.Get(url.web)
	if err != nil {
		return "", errors.New("FATAL ERROR!!!!")
	}

	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return "", errors.New("You've got some troubles!")
	}

	file, err := os.Create("data.html")

	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		return "", errors.New("Shit")
	}

	return string(body), nil
}

type Love struct {
	heart string
	kiss  string
}

func choke(love *Love) (string, string) {
	return love.heart, love.kiss
}

func main() {
	var (
		linux string
		err   error
		link  string = "https://go.dev/"
	)

	fmt.Scanln(&linux)

	nums := []int{1, 2, 3, 4}
	nums = append(nums, 5, 6)
	fmt.Println(nums)

	strs := []string{"Wesley", "Mwah", "Hug", "Kiss"}
	strs = append(strs, "Cutie")
  fmt.Println(strs)

	cute := &Love{"wesley", "mwah"}
	content := &Url{"https://go.dev/"}

	/* scraping, err := GET(content)
	if err != nil {
		log.Fatal(err)
	} */

  resp, err := http.Get(content.web)
  if err != nil {
  	log.Fatal(err)
  }

	scrapingResult, err := SCRAP(resp)
	if err != nil {
		log.Fatal(err)
	}
  
  const limit = 17

  for num := 0; num < limit; num++ {
  	fmt.Println(num)
  }


	fmt.Println(scrapingResult)
	fmt.Println(choke(cute))
	fmt.Println("Wesley! :}", link, linux)

	if err != nil {
		log.Fatal(errors.New("Sheet we got an error lmfao"))
	}
}
