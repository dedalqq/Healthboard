package main

import (
	"sync"
	"time"
)

const (
	statusOK = iota
	statusErr
	statusWaiting
	statusNotCheck
)

// ServiceList -
type ServiceList struct {
	mutex      sync.Mutex
	LastUpdate time.Time
	Services   []*service
}

type service struct {
	Name      string
	Resources []*resource
}

type resource struct {
	URL         string
	Status      uint8
	LastAttempt time.Time
}

func (r *resource) check(wg *sync.WaitGroup) {
	defer wg.Done()

	if statusWaiting == r.Status {
		return
	}

	r.Status = statusWaiting
	r.LastAttempt = time.Now()

	if checkHTTPStatus(r.URL) {
		r.Status = statusOK
		return
	}

	r.Status = statusErr
}

func (sl *ServiceList) runCheckServers() {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()

	var wg sync.WaitGroup

	for _, s := range sl.Services {
		for _, r := range s.Resources {
			wg.Add(1)
			go r.check(&wg)
		}
	}

	wg.Wait()
}

func (sl *ServiceList) updateConfig(fileName string) {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()

	config, err := parseConfig(fileName)

	if nil != err {
		println(err.Error())
		return
	}

	ser := make([]*service, 0, len(config))
	for _, c := range config {
		s := service{
			Name:      c.Name,
			Resources: make([]*resource, 0, len(c.URL)),
		}

		for _, u := range c.URL {
			r := resource{
				URL:    u,
				Status: statusNotCheck,
			}
			s.Resources = append(s.Resources, &r)
		}

		ser = append(ser, &s)
	}

	sl.Services = ser
}
