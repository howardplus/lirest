package client

import (
	"encoding/json"
	"fmt"
	"github.com/howardplus/lirest/inject"
	"github.com/howardplus/lirest/util"
)

// JobListData contains the job list output
type JobListData struct {
	Name string       `json:"name"`
	Data []inject.Job `json:"data"`
}

// JobList returns the list of jobs where
// data has been successfully injected
func JobList() {
	r, err := util.Get("jobs")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	jobs := &JobListData{}
	if err := json.NewDecoder(r).Decode(&jobs); err != nil {
		fmt.Printf(err.Error())
		return
	}

	// format output
	if len(jobs.Data) == 0 {
		fmt.Printf("No jobs\n")
	} else {
		fmt.Printf("%-10s %-20s %-30s %-20s %-20s\n", "Id", "Execute Time", "Write To", "Write Data", "Old Data")
		for _, j := range jobs.Data {
			fmt.Printf("%-10s %-20s %-30s %-20s %-20s\n",
				j.Id.String()[:8],
				j.Execute.Format("2006-01-02 15:04:05"),
				j.Name,
				j.Data,
				j.Old)
		}
	}
}
