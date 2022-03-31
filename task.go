package main

import "time"

type Task struct {
	do               func()
	runningCondition RunningCondition
	isFinished       bool
	lastChecked      time.Time // time when the task was last checked, is needed to calculate the perfect timing for a running condition like "every day"
}

func (t Task) Execute() {
	t.isFinished = false
	// check if the running condition is met
	if t.runningCondition.isMet(t) {
	}
	t.lastChecked = time.Now()
}

// check if running condition is met
func (r RunningCondition) isMet(t Task) bool {
	if r.usingDuration {
		// get difference between now and last checked as time.Duration
		diff := time.Since(t.lastChecked)
		if diff >= r.duration {
			return true
		}
		return false
	}
	return false
}

/*
	how you use tasks:

	you execute a task by calling the function task.Execute() and then, if will be executed,
	if the conditions, which are necessary are there / existing.
*/

type RunningCondition struct {
	duration      time.Duration
	other         string // for example "every monday" or "every monday morning"
	usingDuration bool
}
