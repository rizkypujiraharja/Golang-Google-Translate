package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func main() {
	sl := flag.String("sl", "en", "Source Language")
	tl := flag.String("tl", "id", "Target Language")
	text := flag.String("text", "Hallo world!", "Text to be translated")
	flag.Parse()

	apiUrl := "https://translate.google.com/translate_a/single?client=at&dt=t&dt=ld&dt=qca&dt=rm&dt=bd&dj=1&ie=UTF-8&oe=UTF-8&inputm=2&otf=2&iid=1dd3b944-fa62-4b55-b330-74909a99969e"
	data := url.Values{}
	data.Set("sl", *sl)
	data.Set("tl", *tl)
	data.Set("q", *text)

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	httpResponse, _ := client.Do(r)

	if httpResponse.StatusCode != 200 {
		log.Fatalf("An Error Occured %v", httpResponse.Status)
	}
	defer httpResponse.Body.Close()

	bodyBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	fmt.Println("Source Language :", *sl)
	fmt.Println("Target Language :", *tl)
	fmt.Println("Result\t\t:", bodyString)
}
