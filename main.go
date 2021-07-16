package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
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

func main() {
	query := "test language:HTML+ERB org:Github"
	url := GITHUB_URL + SEARCH_CODE_ENDPOINT + "?q=" + url2.QueryEscape(query)
	resp, err := http.Get(url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/vnd.github.v3+json")

	if err != nil {
		// handle error
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//sb := string(body)
	var data result
	json.Unmarshal(body, &data)

	fmt.Printf("Results: %v\n", data)

	//log.Printf(sb)

}
