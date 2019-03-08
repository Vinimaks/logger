package logger

import (
	"github.com/fatih/color"
	"log"
	"time"
)

type mode struct {
	name  string
	color *color.Color
}

func (m *mode) printLn(a ...interface{}) {
	ctime := time.Now().Format("2006-1-2 15:4:5")
	data := append([]interface{}{"[" + ctime + "]", m.name + ":"}, a...)
	_, err := m.color.Println(data...)
	if err != nil {
		log.Fatalln(err)
	}
}

type logger struct {
	notice  *mode
	warning *mode
	error   *mode
	fatal   *mode
}

func (l *logger) Notice(a ...interface{}) {
	l.notice.printLn(a...)
}
func (l *logger) Warning(a ...interface{}) {
	l.warning.printLn(a...)
}
func (l *logger) Error(a ...interface{}) {
	l.error.printLn(a...)
}
func (l *logger) Fatal(a ...interface{}) {
	l.fatal.printLn(a...)
}

var DefaultLogger = &logger{
	notice:  defaultNotice,
	warning: defaultWarning,
	error:   defaultError,
	fatal:   defaultFatal,
}

var defaultNotice = &mode{
	name:  "Notice",
	color: color.New(color.Bold, color.FgGreen),
}

var defaultWarning = &mode{
	name:  "Warning",
	color: color.New(color.Bold, color.FgYellow),
}

var defaultError = &mode{
	name:  "Error",
	color: color.New(color.Bold, color.FgRed),
}

var defaultFatal = &mode{
	name:  "Fatal",
	color: color.New(color.Bold, color.BgRed),
}
