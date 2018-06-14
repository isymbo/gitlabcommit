package gitlabcommit

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"gopkg.in/resty.v1"
)

type GitlabCommit struct {
	ID             string    `json:"id"`
	ShortID        string    `json:"short_id"`
	Title          string    `json:"title"`
	AuthorName     string    `json:"author_name"`
	AuthorEmail    string    `json:"author_email"`
	CommitterName  string    `json:"committer_name"`
	CommitterEmail string    `json:"committer_email"`
	CreatedAt      time.Time `json:"created_at"`
	Message        string    `json:"message"`
	CommittedDate  time.Time `json:"committed_date"`
	AuthoredDate   time.Time `json:"authored_date"`
	ParentIds      []string  `json:"parent_ids"`
	LastPipeline   struct {
		ID     int    `json:"id"`
		Ref    string `json:"ref"`
		Sha    string `json:"sha"`
		Status string `json:"status"`
	} `json:"last_pipeline"`
	Stats struct {
		Additions int `json:"additions"`
		Deletions int `json:"deletions"`
		Total     int `json:"total"`
	} `json:"stats"`
	Status string `json:"status"`
}

const (
	GITLABCOMMIT_TOKEN   string = ""
	GITLABCOMMIT_BASEURL string = ""
)

func GetGitlabCommit(projectID int, commitID string) (*GitlabCommit, error) {
	resp, err := resty.SetTimeout(time.Duration(1*time.Minute)).
		R().
		SetHeader("PRIVATE-TOKEN", GITLABCOMMIT_TOKEN).
		Get(GITLABCOMMIT_BASEURL + "/projects/" + strconv.Itoa(projectID) + "/repository/commits/" + commitID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	c := &GitlabCommit{}
	err = json.Unmarshal(resp.Body(), c)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return c, nil
}

