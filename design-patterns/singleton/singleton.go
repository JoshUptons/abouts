package singleton

import (
	"fmt"
	"time"
)

/*
the main issue that the singleton pattern solves is when you are trying to
pass data between different section of your application and require it to keep
a single instance of the data without getting overwritten or transformed
accidentally.

A scenario off the top of my head would be something like the store within
a React or Vue project in javascript.
*/

// if we take an object like MyLogger for instance

type Log struct {
	text string
}

func (l *Log) out() {
	fmt.Println(fmt.Sprintf("%v: %s", time.Now(), l.text))
}

type MyLogger struct {
	logs []Log
}

func (m *MyLogger) log(text string) {
	m.logs = append(m.logs, Log{text})
}

func (m *MyLogger) totalLogs() int {
	return len(m.logs)
}

func firstProcess() {
	l := MyLogger{}
	l.log("this is a log")
	l.log("this is a second log")
}

func secondProcess() {
	l := MyLogger{}
	l.log("this is my third log")
}

/*
if we were to execute these two functions, they would each instantiate a new
instance of MyLogger, which would each have an empty log set, thus, outputting
an incorrect count of logs
Contrived, I know, but still gets the idea.
*/

/*
Instead, we can create a singleton to handle this, that we pass between functions
to maintain the data on the object throughout subsequent function calls
*/

// the call of each function becomes
func firstProcessSingleton(l *MyLogger) *MyLogger {
	l.log("this is the first log")
	l.log("this is the second log")
	return l
}

func secondProcessSingleton(l *MyLogger) *MyLogger {
	l.log("this is the third log")
	return l
}

/*
now when we call the l.totalLogs, we will receive out all 3 logs in our main
process, because we have passed the reference to the singleton throughout the
processes, and returned back the singleton, maintianing the state
*/
