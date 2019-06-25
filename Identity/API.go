package identity

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

//API ...
type API struct {
	IdentityAPIHost    string `json:"identity_api_host"`
	APIWorkers         int    `json:"api_workers"`
	MaxIdleConnections int    `json:"max_idle_connections"`
	MaxConnections     int    `json:"max_connections"`
	workChannel        chan *Job
	client             *http.Client
}

//Job ...
type Job struct {
	RequestName string
	Params      *url.Values
	Result      string
	Error       error
	Done        chan bool
}

//SetupAPI ...
func SetupAPI(IdentityAPIHost string, params ...int) *API {

	//apiWorkers, maxIdleConnections, maxConnections
	parameters := [3]int{1, 1, 1}

	for i, param := range params {
		if param > 0 {
			parameters[i] = param
		}
	}

	return &API{
		IdentityAPIHost:    IdentityAPIHost,
		APIWorkers:         parameters[0],
		MaxIdleConnections: parameters[1],
		MaxConnections:     parameters[2],
	}
}

//Init ...
func (a *API) Init() {
	tr := &http.Transport{
		MaxIdleConns:    a.MaxIdleConnections,
		MaxConnsPerHost: a.MaxConnections,
	}
	a.client = &http.Client{Transport: tr}
	a.StartWorkers()
}

//StartWorkers ...
func (a *API) StartWorkers() {
	a.workChannel = make(chan *Job, 512)

	for i := 0; i < a.APIWorkers; i++ {
		go a.worker()
	}
}

func (a *API) worker() {
	for {
		select {
		case job, ok := <-a.workChannel:
			if !ok {
				return
			}
			result, err := a.doJob(job)
			job.Result = result
			job.Error = err
			job.Done <- true
		}
	}
}

func (a *API) postRequest(params *url.Values, requestName string) (string, error) {

	job := &Job{
		RequestName: requestName,
		Params:      params,
		Done:        make(chan bool),
	}

	a.workChannel <- job
	<-job.Done

	return job.Result, job.Error
}

func (a *API) doJob(job *Job) (string, error) {

	params := job.Params
	endpointURL := fmt.Sprintf("%s/%s", a.IdentityAPIHost, job.RequestName)

	buffer := &bytes.Buffer{}
	buffer.Write([]byte(params.Encode()))

	resp, err := a.client.Post(endpointURL, "application/x-www-form-urlencoded", buffer)

	if err != nil {
		log.Println("Post request failed", err.Error())
		return "", err
	}

	defer resp.Body.Close()

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, resp.Body)

	if err != nil {
		log.Println("Copying response to buffer failed", err.Error())
		return "", err
	}

	return buf.String(), nil
}
