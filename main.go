/* Author: xpl0ited1 (Bastian Muhlhauser)
   Date: July 16th, 2021
*/

package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"strconv"
)

const (
	GITHUB_URL           = "https://api.github.com/"
	SEARCH_CODE_ENDPOINT = "search/code"
)

type result struct {
	Total_Count        int32
	Incomplete_Results bool
	Items              []item
}

type item struct {
	Name       string
	Path       string
	Sha        string
	Url        string
	Git_Url    string
	Html_Url   string
	Repository repository
	Score      float32
}

type repository struct {
	Id                int64
	Node_Id           string
	Name              string
	Full_Name         string
	Private           bool
	Owner             owner
	Html_Url          string
	Description       string
	Fork              bool
	Url               string
	Forks_Url         string
	Keys_Url          string
	Collaborators_Url string
	Teams_Url         string
	Hooks_Url         string
	Issue_Events_Url  string
	Events_Url        string
	Assignees_Url     string
	Branches_Url      string
	Tags_Url          string
	Blobs_Url         string
	Git_Tags_Url      string
	Git_Refs_Url      string
	Trees_Url         string
	Statutes_Url      string
	Languages_Url     string
	Stargazers_Url    string
	Contributors_Url  string
	Subscribers_Url   string
	Subscription_Url  string
	Commits_Url       string
	GitCommits_Url    string
	Comments_Url      string
	IssueComment_Url  string
	Contents_Url      string
	Compare_Url       string
	Merges_Url        string
	Archive_Url       string
	Downloads_Url     string
	Issues_Url        string
	Pulls_Url         string
	Milestones_Url    string
	Notifications_Url string
	Labels_Url        string
	Releases_Url      string
	Deployments_Url   string
}

type owner struct {
	Login               string
	Id                  int64
	Node_Id             string
	Avatar_Url          string
	Gravatar_Id         string
	Url                 string
	Html_Url            string
	Followers_Url       string
	Following_Url       string
	Gists_Url           string
	Starred_Url         string
	Subscriptions_Url   string
	Organizations_Url   string
	Repos_Url           string
	Events_Url          string
	Received_Events_Url string
	Type                string
	Site_Admin          bool
}

type content struct {
	Name         string
	Path         string
	Sha          string
	Size         int32
	Url          string
	Html_Url     string
	Git_Url      string
	Download_Url string
	Type         string
	Content      string
	Encoding     string
	_links       links
}

type links struct {
	Self string
	Git  string
	Html string
}

func main() {

	search := flag.String("search", "", "code to search")
	org := flag.String("org", "", "organization to look at")
	lang := flag.String("lang", "", "programming language")
	showContent := flag.Bool("content", false, "display content of code")
	page := flag.Int("page", 1, "page number, only if results are more than 100")
	user := flag.String("user", "", "user to look at")
	repo := flag.String("repo", "", "repo to look at")

	flag.Parse()

	query := ""
	lookedAt := ""
	if *lang != "" {
		if *org != "" {
			query = *search + " language:" + *lang + " org:" + *org
			lookedAt = *org
		}

		if *user != "" {
			query = *search + " language:" + *lang + " user:" + *user
			lookedAt = *user
		}

		if *repo != "" {
			query = *search + " language:" + *lang + " repo:" + *repo
			lookedAt = *repo
		}
	} else {
		if *org != "" {
			query = *search + " org:" + *org
			lookedAt = *org
		}

		if *user != "" {
			query = *search + " user:" + *user
			lookedAt = *user
		}

		if *repo != "" {
			query = *search + " repo:" + *repo
			lookedAt = *repo
		}
	}
	url := GITHUB_URL + SEARCH_CODE_ENDPOINT + "?q=" + url2.QueryEscape(query) + "&per_page=100&page=" + strconv.Itoa(*page)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var data result
	json.Unmarshal(body, &data)

	print_formatted_results(&data, search, lookedAt, lang, showContent)
}

func print_formatted_results(result *result, search *string, lookedAt string, lang *string, showContent *bool) {
	fmt.Printf("Search: %s\n", *search)
	fmt.Printf("Looked at: %s\n", lookedAt)

	if *lang != "" {
		fmt.Printf("Language: %s\n", *lang)
	}

	fmt.Printf("Results: %d\n", result.Total_Count)
	for _, item := range result.Items {
		fmt.Printf("URL: %s\n", item.Html_Url)
		if *showContent {
			content, err := getContent(item.Git_Url)

			if err != nil {
				fmt.Println("Eror decoding content: ", err)
			}

			fmt.Printf("%s\n", content)
		}
	}
}

func getContent(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var data content
	json.Unmarshal(body, &data)

	if data.Encoding == "base64" {
		decodedData, err := base64.StdEncoding.DecodeString(data.Content)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s", decodedData), nil
	}

	return "", nil
}
