package inject

import (
	log "github.com/Sirupsen/logrus"
	"github.com/google/uuid"
	"github.com/howardplus/lirest/job"
	"time"
)

// Job is an object that gets recorded
// when a injection is complete
type Job struct {
	Id      uuid.UUID `json:"id"`
	Execute time.Time `json:"executed"`
	Old     string    `json:"old"`
	Data    string    `json:"new"`
	Name    string    `json:"writeTo"`
}

// Revert implements the Job interface
func (j *Job) Revert() error {
	return nil
}

// RecordJob creates a new completed job object
func RecordJob(inj Injector, old string, data string) {
	newJobChan <- &Job{
		Execute: time.Now(),
		Id:      uuid.New(),
		Name:    inj.Name(),
		Old:     old,
		Data:    data,
	}
}

var newJobChan chan job.Job
var jobReqChan chan int
var jobRespChan chan []job.Job

const (
	jobInitSize  = 100
	jobStaleTime = 24 // hours
)

// JobTracker stores the jobs and respond to query
func JobTracker() {

	newJobChan = make(chan job.Job, jobInitSize)
	jobReqChan = make(chan int, 1)
	jobRespChan = make(chan []job.Job, 1)

	jobs := make([]job.Job, 0, jobInitSize)

	for {
		select {
		case job := <-newJobChan:
			jobs = append(jobs, job)
		case n := <-jobReqChan:

			log.WithFields(log.Fields{
				"n": n,
			}).Debug("Request jobs")

			if n != 0 && n < len(jobs) {
				jobRespChan <- jobs[len(jobs)-n:]
			} else {
				jobRespChan <- jobs
			}
		case <-time.After(time.Second):
			// delete stale jobs
			now := time.Now()
			for i, j := range jobs {
				job := j.(*Job)
				if now.Unix() > job.Execute.Add(time.Duration(jobStaleTime)*time.Hour).Unix() {
					jobs = append(jobs[0:i-1], jobs[i+1:]...)
				}
			}
		}
	}
}

// RequestJobs request last n jobs from the tracker
func RequestJobs(n int) []job.Job {
	jobReqChan <- n
	return <-jobRespChan
}
