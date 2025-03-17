package golang_reverso_api

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type reverso struct {
	proxy       string
	scraper     colly.Collector
	sl_language string
	tl_language string
}

func New_Reverso_Requester() reverso {
	r := reverso{
		proxy:       "",
		scraper:     *colly.NewCollector(),
		sl_language: "",
		tl_language: "",
	}

	return r
}

/*
translation requests on reverso are written in the url with the starting
& finishing language & the sentence to translate written w/ %2520 inbetween the words in the url
for reference

	https://www.reverso.net/text-translation#sl=spa&tl=eng&text=te%2520quiero
*/
func (r reverso) get_translation(prompt string) string {

	full_url := fmt.Sprintf("https://www.reverso.net/text-translation#sl=%s&tl=%s&text=", r.langstring_langabrev(), r.langstring_langabrev())
	query_string := ""
	split_prompt := strings.Split(prompt, " ")

	for i, v := range split_prompt {
		if i == len(split_prompt)-1 {
			query_string += fmt.Sprintf("%s", v)
		} else {
			query_string += fmt.Sprintf("%s%%2520", v)
		}
	}
	full_url = fmt.Sprintf("%s%s", full_url, query_string)
	r.scraper.Visit(full_url)
	return ""
}

/*
 */
func (r reverso) get_grammar_check(prompt string) {

}

func (r reverso) get_context(prompt string) {

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
