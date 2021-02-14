package main

import (
	"context"
	"fmt"
)

type jobDescription struct {
	Job      Job
	Criteria []func(user *User) bool
}

// ======================= Painter Position =======================

var painterPosition = jobDescription{Job: &Painter{},
	Criteria: []func(user *User) bool{
		func(user *User) bool {
			return user.Age > 15
		},
		func(user *User) bool {
			for _, l := range user.Languages {
				if l == "german" {
					return true
				}
			}
			return false
		},
	}}

// ======================= Taxi Driver Position ======================

var taxiDriverPosition = jobDescription{Job: &TaxiDriver{},
	Criteria: []func(user *User) bool{
		func(user *User) bool {
			return user.HigherEducation
		},
		func(user *User) bool {
			return user.IsAdult()
		},
	}}

// ======================= All Jobs =======================

var availableJobs = []jobDescription{painterPosition, taxiDriverPosition}





func (jd *jobDescription) SatisfiesCriteria(user *User, ctx context.Context, result chan<- Job) {
	fmt.Println(user.Name + ", "+jd.Job.GetDescription()+ ": enter SatisfiesCriteria")
	for _, criteria := range jd.Criteria {
		if criteria(user) == false {
			result <- nil
			return
		}

	}

	result <- jd.Job
}

func EvaluateApplicant(user *User) (Job, bool)  {
	ctx, cancel := context.WithCancel(context.Background())

	jobChan := make(chan Job)
	defer close(jobChan)

	for _, jd := range availableJobs {
		fmt.Println(user.Name + ", "+jd.Job.GetDescription()+ ": starting go rutine")
		go jd.SatisfiesCriteria(user, ctx, jobChan)
	}

	for result := range jobChan {
		fmt.Println(result)
		if result != nil{
			cancel()
			return result, true
		}
	}

	return nil, false

}