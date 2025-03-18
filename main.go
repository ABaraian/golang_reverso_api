package golang_reverso_api

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type reverso struct {
	scraper     colly.Collector
	sl_language string
	tl_language string
}

func New_Reverso_Requester() reverso {
	r := reverso{
		scraper:     *colly.NewCollector(),
		sl_language: "",
		tl_language: "",
	}

	return r
}
func (r reverso) Set_Proxy(proxy string) {
	r.scraper.SetProxy(proxy)
}

/*
translation requests on reverso are written in the url with the starting
& finishing language & the sentence to translate written w/ %2520 inbetween the words in the url
for reference

	https://www.reverso.net/text-translation#sl=spa&tl=eng&text=te%2520quiero

Translations stored in either a what looks like a text area or app-context-box

	If it looks like a text area
		Its actually a div w/ a textare & divs w/in it, one that stores the translation w/in a span and another div that holds the transliteration
		The div that holds the translation & transliteration is one wth a class : translation-input__main translation-input__result
		Translation is stored in a span with the class : sentence-wrapper_without-hover
		Transliteration is stored in a div w/ the class : transliteration
	If app-context-box
		There are divs w/ the class context-result that stores translation, transliteration & part of speech
		Translation stored in a span w/ the class : text__translation
		Part of speech stored in a span w/ the classes : part-of-speech_context ng-star-inserted
		Transliteration stored in a span w/ the classes : transliteration ng-star-inserted
*/
func (r reverso) get_translation(prompt string) string {

	full_url := fmt.Sprintf("https://www.reverso.net/text-translation#sl=%s&tl=%s&text=", r.langstring_langabrev(), r.langstring_langabrev())
	query_string := ""
	split_prompt := strings.Split(prompt, " ")

	for i, v := range split_prompt {
		if i == len(split_prompt)-1 {
			query_string += fmt.Sprintf("%s%%250A", v)
		} else {
			query_string += fmt.Sprintf("%s%%2520", v)
		}
	}

	full_url = fmt.Sprintf("%s%s", full_url, query_string)
	web_scraper := r.scraper.Clone()
	//text__translation is the class for the direct translation
	//transliteration ng-star-inserted is the classes for the transliteration
	web_scraper.OnHTML("div .translation-inputs", func(e *colly.HTMLElement) {
		e.ForEach()
	})

	err := r.scraper.Visit(full_url)
	if err != nil {
		return "Error"
	}

	return ""
}

/*
 */
func (r reverso) get_grammar_check(prompt string) {

}

/*
https://context.reverso.net/translation/english-french/my+name+is+john
*/
func (r reverso) get_context(prompt string) string {
	full_url := fmt.Sprintf("https://context.reverso.net/translation/%s-%s/", r.sl_language, r.tl_language)
	query_string := ""
	split_prompt := strings.Split(prompt, " ")

	for i, v := range split_prompt {
		if i == len(split_prompt)-1 {
			query_string += fmt.Sprintf("%s", v)
		} else {
			query_string += fmt.Sprintf("%s+", v)
		}
	}
	full_url = fmt.Sprintf("%s%s", full_url, query_string)
	r.scraper.Visit(full_url)

	return ""
}
func (r reverso) allowed(reverso_req_type string) error {
	switch reverso_req_type {
	case "grammar_check":
		abbr := r.langstring_langabrev()
		if abbr != "eng" && abbr != "fra" && abbr != "spa" && abbr != "ita" {
			return errors.New("Language not supported")
		}
	}
	return nil

}
func (r reverso) langstring_langabrev() string {
	abrr := ""
	switch r.sl_language {
	case "arabic":
		abrr = "ara"
	case "chinese":
		abrr = "chi"
	case "chzech":
		abrr = "cze"
	case "danish":
		abrr = "dan"
	case "dutch":
		abrr = "dut"
	case "english":
		abrr = "eng"
	case "french":
		abrr = "fra"
	case "german":
		abrr = "german"
	case "greek":
		abrr = "gre"
	case "hebrew":
		abrr = "heb"
	case "hindi":
		abrr = "hin"
	case "hungarian":
		abrr = "hun"
	case "italian":
		abrr = "ita"
	case "japanese":
		abrr = "jpn"
	case "korea":
		abrr = "kor"
	case "persian":
		abrr = "per"
	case "polish":
		abrr = "pol"
	case "portuguese":
		abrr = "por"
	case "romanian":
		abrr = "rum"
	case "russian":
		abrr = "rus"
	case "slovak":
		abrr = "slo"
	case "spanish":
		abrr = "spa"
	case "swedish":
		abrr = "swe"
	case "thai":
		abrr = "tha"
	case "turkish":
		abrr = "tur"
	case "ukranian":
		abrr = "ukr"
	}
	return abrr
}
