package logging

import log_ "log"
import "fmt"
import "strings"

type LogLevel int

const (
	FATAL = LogLevel(0) + iota
	ERROR
	WARNING
	INFO
	DEBUG

	WARN     = WARNING
	CRITICAL = FATAL
	CRYTICAL = FATAL
	INVALID  = LogLevel(-1)
)

var LogLevelNames = map[LogLevel]string{
	FATAL:   "FATAL",
	ERROR:   "ERROR",
	WARNING: "WARNING",
	INFO:    "INFO",
	DEBUG:   "DEBUG",
}

func CanBeLogLevel(level int) bool {
	_, ok := LogLevelNames[LogLevel(level)]
	return ok
}

func LogLevelName(level LogLevel) string {
	name, ok := LogLevelNames[level]
	if ok {
		return name
	}
	return fmt.Sprintf("<Unknown log level #%d>", level)
}

var LogLevelValues = map[string]LogLevel{
	"FATAL":   FATAL,
	"ERROR":   ERROR,
	"WARNING": WARNING,
	"INFO":    INFO,
	"DEBUG":   DEBUG,

	"WARN":     WARNING,
	"CRITICAL": FATAL,
	"CRIT":     FATAL,
	"CRYTICAL": FATAL,
	"CRYT":     FATAL,
}

func (self LogLevel) String() string {
	return LogLevelName(self)
}

func ValidLogLevel(level LogLevel) bool {
	return level >= FATAL && level <= DEBUG
}

func ValidLogLevelName(level string) bool {
	_, ok := LogLevelValues[strings.ToUpper(level)]
	return ok
}

func LogLevelByName(level string) LogLevel {
	lvl, ok := LogLevelValues[strings.ToUpper(level)]
	if !ok {
		return INVALID
	}
	return lvl
}

func adjustLogLevel(name string) {
	lvl, ok := LogLevelValues[strings.ToUpper(name)]
	if !ok {
		log.Fatal("Wrong level name: %+q", name)
	}
	if log.Level != lvl {
		log.Say("Log level %s -> %s (%s)", log.Level, lvl, name)
		log.SetLevel(lvl)
	}
}

type Logger struct {
	Level LogLevel
	Panic bool
}

func NewLogger(level LogLevel) *Logger {
	var l = &Logger{}
	if !ValidLogLevel(level) {
		panic(fmt.Errorf("FATAL: Invalid level %#v (%s)", level, level))
	}
	l.Level = level
	return l
}

func (self *Logger) UsePanic(use bool) {
	self.Panic = use
}

func (self *Logger) SetLevel(level LogLevel) {
	if !ValidLogLevel(level) {
		self.Fatal("Invalid level %#v (%s)", level, level)
	}
	if self.Level != level {
		self.Say("Log level %s -> %s", self.Level, level)
		self.Level = level
	}
}

// unconditionally write to Logger with "NOTE:" prefix
func (self *Logger) Say(message string, args ...interface{}) {
	log_.Printf("NOTE: "+message, args...)
}

func (self *Logger) Log(level LogLevel, message string, args ...interface{}) {
	if level <= self.Level {
		log_.Printf(level.String()+": "+message, args...)
	}
}

func (self *Logger) Fatal(message string, args ...interface{}) {
	out := fmt.Sprintf("FATAL: "+message, args...)
	f := func (x string) { log_.Fatal(x) }
	if self.Panic {
		f = func (x string) { panic(x) }
	}
	f(out)
}

func (self *Logger) Error(message string, args ...interface{}) {
	self.Log(ERROR, message, args...)
}

func (self *Logger) Warning(message string, args ...interface{}) {
	self.Log(WARNING, message, args...)
}

func (self *Logger) Warn(message string, args ...interface{}) {
	self.Warning(message, args...)
}

func (self *Logger) Info(message string, args ...interface{}) {
	self.Log(INFO, message, args...)
}

func (self *Logger) Debug(message string, args ...interface{}) {
	self.Log(DEBUG, message, args...)
}

var log = NewLogger(INFO)

/* EOF */
