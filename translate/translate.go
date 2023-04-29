package translate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"translator/e"
)

type Translate struct {
	Status string
	Data   Data
}

type Data struct {
	TranslatedText string
}

func TranslateRuToEn(text string) (string, error) {
	trnsl := Translate{}

	url := "https://text-translator2.p.rapidapi.com/translate"

	payload := strings.NewReader(fmt.Sprintf("source_language=ru&target_language=en&text=%s", text))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("X-RapidAPI-Key", "7ded77776bmsh4bbc55fabcdcc1ep149dd9jsnce59b8ca9e82")
	req.Header.Add("X-RapidAPI-Host", "text-translator2.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", e.Wrap("body can't read", err)
	}

	if err := json.Unmarshal(body, &trnsl); err != nil {
		return "", e.Wrap("body can't read", err)
	}

	/* 	fmt.Println(trnsl.Data.TranslatedText)

	   	fmt.Println(res)
	   	fmt.Println(string(body)) */
	fmt.Println(string(trnsl.Data.TranslatedText))

	return trnsl.Data.TranslatedText, nil
}
