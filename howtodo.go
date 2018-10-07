package main

import (
    "flag"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "io"
    "log"
    "net/http"
    "net/url"
    "os"
    "strings"
)

const appName = "Howtodo"
const version = "v1.0.0"
const maintainer = "S.Hwang <lotsofluck4m@gmail.com>"

const searchUrlTemplate = "https://google.com/search?q=%s"
const targetSite = "stackoverflow.com"
const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"

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
        // TODO: show code from the answer
        fmt.Println("default")
    }
}

func formatVersion() string {
    return fmt.Sprintf("%s %s by %s", appName, version, maintainer)
}

func normalizeQuery(rawQueries []string) string {
    query := fmt.Sprintf("site:%s %s", targetSite, strings.Join(rawQueries, " "))
    return url.QueryEscape(query)
}

func getSearchUrl(query string) string {
    return fmt.Sprintf(searchUrlTemplate, query)
}

func getLink(questions []string) string {
    const linkSelector = "div#search div.r a"

    searchUrl := getSearchUrl(normalizeQuery(questions))

    responseBody := requestSearch(searchUrl)
    defer responseBody.Close()

    searchResult, parseError := goquery.NewDocumentFromReader(responseBody)
    if parseError != nil {
        log.Fatal(parseError)
    }

    link, found := searchResult.Find(linkSelector).Attr("href")
    if !found {
        return "Sorry. Try again with other words."
    }

    return link
}

func requestSearch(url string) io.ReadCloser {
    request, requestError := http.NewRequest("GET", url, nil)
    if requestError != nil {
        log.Fatal(requestError)
    }
    request.Header.Add("User-Agent", userAgent)

    client := &http.Client{}
    response, responseError := client.Do(request)
    if responseError != nil {
        log.Fatal(responseError)
    }

    return response.Body
}
