package main

import (
    "flag"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "log"
    "net/http"
    "net/url"
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

    questions := flag.Args()
    switch {
    case *linkOnly:
        fmt.Println(getLink(questions))
    case *showFullAnswer:
        fmt.Println(getAnswer(questions, true))
    case *showVersion:
        fmt.Println(formatVersion())
    default:
        fmt.Println(getAnswer(questions, false))
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

func getAnswer(questions []string, needFull bool) string {
    link := getLink(questions)

    const answersSelector = "div#answers div.answer"
    answerDocument := getHTTP(fmt.Sprintf("%s?answertab=votes", link))
    answers := answerDocument.Find(answersSelector)

    var selectedAnswer *goquery.Selection
    const acceptedAnswerSelector = "div.accepted-answer.answer"
    if len(answers.Nodes) == 0 {
        log.Fatal(errorMessages["RESULT_NOT_FOUND"])
    } else if acceptedAnswer := answers.Find(acceptedAnswerSelector); len(acceptedAnswer.Nodes) > 0 {
        selectedAnswer = acceptedAnswer
    } else {
        selectedAnswer = answers.First()
    }

    var answerContentBuilder strings.Builder
    answerContentBuilder.WriteString(fmt.Sprintf("%s\n", link))

    const postSelector = "div.post-text"
    selectedAnswer.Find(postSelector).Contents().Each(
        func(index int, selection *goquery.Selection) {
            if (needFull && goquery.NodeName(selection) != "#text") ||
                (!needFull && (goquery.NodeName(selection) == "pre" || goquery.NodeName(selection) == "code")) {
                answerContentBuilder.WriteString(fmt.Sprintf("%s\n", selection.Text()))
            }
        },
    )

    return answerContentBuilder.String()
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
