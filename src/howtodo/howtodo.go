package main

import (
    "bytes"
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)

const appName = "Howtodo"
const version = "v1.0.0"
const maintainer = "S.Hwang <lotsofluck4m@gmail.com>"

const searchUrlTemplate = "https://google.com/search?q=site:%s %s"
const targetSite = "stackoverflow.com"
const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"

func main() {
    linkOnly := flag.Bool("l", false, "Show only link to answer")
    showFullAnswer := flag.Bool("a", false, "Show full content of answer")
    showVersion := flag.Bool("v", false, "Show version")

    // Parse the flags to generate help text for -h flag
    flag.Parse()

    switch {
    case *linkOnly:
        // TODO: show link
        fmt.Println("Show link only")
    case *showFullAnswer:
        // TODO: show full answer
        fmt.Println("Show full answer")
    case *showVersion:
        fmt.Println(formatVersion())
    default:
        // TODO: show code from the answer
        questions := os.Args[1:]
        fmt.Println(getAnswer(questions))
    }
}

func formatVersion() string {
    return fmt.Sprintf("%s %s by %s", appName, version, maintainer)
}

func normalizeQuery(rawQueries [] string) (trimmedQuery string) {
    joinedQuery := strings.Join(rawQueries, " ")
    questionMarkRemovedQuery := strings.Replace(joinedQuery, "?", "", -1)
    trimmedQuery = strings.Trim(questionMarkRemovedQuery, " ")
    return
}

func getSearchUrl(query string) (searchUrl string) {
    searchUrl = fmt.Sprintf(searchUrlTemplate, targetSite, query)
    return
}

func getAnswer(questions [] string) string {
    searchUrl := getSearchUrl(normalizeQuery(questions))

    return httpGet(searchUrl)
}

func httpGet(url string) string {
    client := &http.Client{}

    request, requestError := http.NewRequest("GET", url, nil)
    if requestError != nil {
        log.Fatal(requestError)
    }
    request.Header.Add("User-Agent", userAgent)

    response, responseError := client.Do(request)
    if responseError != nil {
        log.Fatal(requestError)
    }

    responseBuffer := new(bytes.Buffer)
    responseBuffer.ReadFrom(response.Body)
    return responseBuffer.String()
}
