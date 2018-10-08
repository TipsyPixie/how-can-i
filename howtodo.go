package main

import (
    "flag"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "log"
    "net/http"
    "net/url"
    "os"
    "strings"
)

const appName = "Howtodo"
const version = "v1.0.0"
const maintainer = "S.Hwang <lotsofluck4m@gmail.com>"

var errorMessages = map[string]string{
    "RESULT_NOT_FOUND": "Sorry. Try again with other words.",
}

func main() {
    linkOnly := flag.Bool("l", false, "Show only link to answer")
    showFullAnswer := flag.Bool("a", false, "Show full content of answer")
    showVersion := flag.Bool("v", false, "Show version")

    // Parse the flags to generate help text for -h flag
    flag.Parse()

    questions := os.Args[1:]
    switch {
    case *linkOnly:
        // TODO: show link
        fmt.Println(getLink(questions))
    case *showFullAnswer:
        // TODO: show full answer
        fmt.Println("Show full answer")
    case *showVersion:
        fmt.Println(formatVersion())
    default:
        fmt.Println(getAnswer(questions))
    }
}

func formatVersion() string {
    return fmt.Sprintf("%s %s by %s", appName, version, maintainer)
}

func normalizeQuery(rawQueries []string) string {
    const targetSite = "stackoverflow.com"
    query := fmt.Sprintf("site:%s %s", targetSite, strings.Join(rawQueries, " "))
    return url.QueryEscape(query)
}

func getSearchUrl(query string) string {
    const searchUrlTemplate = "https://google.com/search?q=%s"
    return fmt.Sprintf(searchUrlTemplate, query)
}

func parseSearchResults(document *goquery.Document) *goquery.Selection {
    const resultSelector = "div#search div.r a"
    return document.Find(resultSelector)
}

func getSearchResults(questions []string) *goquery.Selection {
    searchUrl := getSearchUrl(normalizeQuery(questions))

    searchResults := parseSearchResults(getHTTP(searchUrl))
    if len(searchResults.Nodes) == 0 {
        log.Fatal(errorMessages["RESULT_NOT_FOUND"])
    }

    return searchResults
}

func getLink(questions []string) string {
    link, linkExist := getSearchResults(questions).Attr("href")
    if !linkExist {
        log.Fatal(errorMessages["RESULT_NOT_FOUND"])
    }

    return link
}

func getAnswer(questions []string) string {
    answers := getSearchResults(questions)

    for index := 0; index < len(answers.Nodes); index++ {
        attributes := answers.Nodes[index].Attr
        for attrIndex := 0; attrIndex < len(attributes); attrIndex++ {
            if value := attributes[attrIndex].Val; attributes[attrIndex].Key == "href" {
                // TODO: check each answer
                fmt.Println(value)
            }
        }
    }

    return "TEST"
}

func getHTTP(url string) *goquery.Document {
    request, requestError := http.NewRequest("GET", url, nil)
    if requestError != nil {
        log.Fatal(requestError)
    }
    const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
    request.Header.Add("User-Agent", userAgent)

    client := &http.Client{}
    response, responseError := client.Do(request)
    defer response.Body.Close()
    if responseError != nil {
        log.Fatal(responseError)
    }

    responseDocument, parseError := goquery.NewDocumentFromReader(response.Body)
    if parseError != nil {
        log.Fatal(parseError)
    }

    return responseDocument
}
