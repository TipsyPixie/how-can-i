package main

import (
    "errors"
    "flag"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "net/http"
    "net/url"
    "strings"
)

const appName = "Howtodo"
const version = "v1.0.0"
const maintainer = "S.Hwang <lotsofluck4m@gmail.com>"

var errorMessages = map[string]string{
    "RESULT_NOT_FOUND": "Sorry. Try again with other words. ",
    "REQUEST_ERROR": "Sorry. Unable to get network connection. ",
    "HTML_PARSE_ERROR": "Sorry. Unable to parse HTML document. ",
}

func main() {
    linkOnly := flag.Bool("l", false, "Show only link to answer")
    showFullAnswer := flag.Bool("a", false, "Show full content of answer")
    showVersion := flag.Bool("v", false, "Show version")

    // Parse the flags to generate help text for -h flag
    flag.Parse()

    var output string
    var err error = nil
    arguments := flag.Args()
    switch {
    case *linkOnly:
        output, err = getLink(arguments)
    case *showFullAnswer:
        output, err = getAnswer(arguments, true)
    case *showVersion:
        output = formatVersion()
    default:
        output, err = getAnswer(arguments, false)
    }

    if err != nil {
        fmt.Println(err)
        return
    }

    if exception := recover(); exception != nil {
        fmt.Printf("Unknown Error:\n%s", exception)
    }

    fmt.Println(output)
}

func formatVersion() string {
    return fmt.Sprintf("%s %s by %s", appName, version, maintainer)
}

func getSearchURL(queryWords []string) string {
    const targetSite = "stackoverflow.com"
    singleStringQuery := fmt.Sprintf("site:%s %s", targetSite, strings.Join(queryWords, " "))

    const searchUrlTemplate = "https://google.com/search?q=%s"
    return fmt.Sprintf(searchUrlTemplate, url.QueryEscape(singleStringQuery))
}

func parseSearchResult(resultDocument *goquery.Document) *goquery.Selection {
    const resultSelector = "div#search div.r a"
    return resultDocument.Find(resultSelector)
}

func getLink(query []string) (string, error) {
    searchResultPage, searchError := requestGet(getSearchURL(query))
    if searchError != nil {
        return "", searchError
    }

    QNAThreads := parseSearchResult(searchResultPage)
    linkToFirstThread, linkExist := QNAThreads.Attr("href")
    if !linkExist {
        return "", errors.New(errorMessages["RESULT_NOT_FOUND"])
    }

    return linkToFirstThread, nil
}

func parseQNAThread(QNAThread *goquery.Document) *goquery.Selection {
    const answersSelector = "div#answers div.answer"
    answers := QNAThread.Find(answersSelector)

    const acceptedAnswerSelector = "div.accepted-answer.answer"
    acceptedAnswer := answers.Find(acceptedAnswerSelector)

    if len(acceptedAnswer.Nodes) > 0 {
        return acceptedAnswer
    } else if len(answers.Nodes) > 0 {
        return answers.First()
    } else {
        return nil
    }
}

func extractCodeBlock(answer *goquery.Selection) string {
    var codeBlockBuilder strings.Builder

    const postSelector = "div.post-text"
    answer.Find(postSelector).Contents().Each(
        func(index int, selection *goquery.Selection) {
            if goquery.NodeName(selection) == "pre" || goquery.NodeName(selection) == "code" {
                codeBlockBuilder.WriteString(fmt.Sprintf("%s\n", selection.Text()))
            }
        },
    )
    return codeBlockBuilder.String()
}

func extractFullAnswer(answer *goquery.Selection) string {
    var fullAnswerBuilder strings.Builder

    const postSelector = "div.post-text"
    answer.Find(postSelector).Contents().Each(
        func(index int, selection *goquery.Selection) {
            if goquery.NodeName(selection) != "#text" {
                fullAnswerBuilder.WriteString(fmt.Sprintf("%s\n", selection.Text()))
            }
        },
    )
    return fullAnswerBuilder.String()
}

func getAnswer(query []string, needFull bool) (string, error) {
    link, linkError := getLink(query)
    if linkError != nil {
        return "", linkError
    }

    linkToQNAThread := fmt.Sprintf("%s?answertab=votes", link)
    QNAThread, threadLinkError := requestGet(linkToQNAThread)
    if threadLinkError != nil {
        return "", threadLinkError
    }

    answer := parseQNAThread(QNAThread)

    var outputBuilder strings.Builder
    outputBuilder.WriteString(fmt.Sprintf("%s\n\n", link))
    if needFull {
        outputBuilder.WriteString(extractCodeBlock(answer))
    } else {
        outputBuilder.WriteString(extractFullAnswer(answer))
    }
    return outputBuilder.String(), nil
}

func requestGet(url string) (*goquery.Document, error) {
    request, requestError := http.NewRequest("GET", url, nil)
    if requestError != nil {
        return nil, requestError
    }
    const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
    request.Header.Add("User-Agent", userAgent)

    client := &http.Client{}
    response, responseError := client.Do(request)
    if responseError != nil {
        return nil, errors.New(errorMessages["REQUEST_ERROR"])
    }

    responseDocument, parseError := goquery.NewDocumentFromReader(response.Body)
    if parseError != nil {
        return nil, errors.New(errorMessages["HTML_PARSE_ERROR"])
    }

    response.Body.Close()
    return responseDocument, nil
}
