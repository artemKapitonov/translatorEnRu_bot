package translate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"translator/e"
	"unicode"
)

type TranslateAPI struct {
	Status string
	Data   Data
}

type Data struct {
	TranslatedText string
}

func Language(text string) (sourceLang string, targetLang string) {
	var count float64

	for _, r := range text {
		if r > unicode.MaxASCII {
			count += 1
		}
	}
	if count >= float64(len(text))*0.3 {
		return "ru", "en"
	}

	return "en", "ru"
}

func Translate(text string) (string, error) {

	sourceLang, targetLang := Language(text)

	translate := TranslateAPI{}

	url := "https://text-translator2.p.rapidapi.com/translate"

	payload := strings.NewReader(fmt.Sprintf("source_language=%s&target_language=%s&text=%s", sourceLang, targetLang, text))

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

	if err := json.Unmarshal(body, &translate); err != nil {
		return "", e.Wrap("body can't read", err)
	}

	/* 	fmt.Println(trnsl.Data.TranslatedText)

	   	fmt.Println(res)
	   	fmt.Println(string(body)) */
	fmt.Println(string(translate.Data.TranslatedText))

	return translate.Data.TranslatedText, nil
}
