package logging

import (
	"fmt"
	log_ "log"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	"golang.org/x/sys/unix"
)

type LogLevel int

func (level LogLevel) Name() string        { return LogLevelName(level) }
func (level LogLevel) Color() aurora.Color { return LogLevelColors[level] }
func (level LogLevel) Colored(s string, au aurora.Aurora) string {
	return au.Colorize(s, level.Color()).String()
}

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

var LogLevelColors = map[LogLevel]aurora.Color{
	FATAL:   aurora.WhiteFg | aurora.BrightFg | aurora.RedBg,
	ERROR:   aurora.RedFg | aurora.BrightFg,
	WARNING: aurora.YellowFg | aurora.BrightFg,
	INFO:    aurora.CyanFg | aurora.BrightFg,
	DEBUG:   aurora.WhiteFg,
}

func (self LogLevel) String() string { return self.Name() }

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
		Root.Fatal("Wrong level name: %+q", name)
	}
	if Root.level != lvl {
		Root.Say("Log level %s -> %s (%s)", Root.level, lvl, name)
		Root.SetLevel(lvl)
	}
}

type Logger struct {
	log    *log_.Logger
	level  LogLevel
	panics bool
	color  aurora.Aurora
}

func newLogger(level LogLevel, color bool) *Logger {
	var l = &Logger{log: log_.New(os.Stderr, "", log_.LstdFlags)}
	if !ValidLogLevel(level) {
		panic(fmt.Errorf("FATAL: Invalid level %#v (%s)", level, level))
	}
	l.level = level
	l.color = aurora.NewAurora(color)
	return l
}

func isAtty(f *os.File) bool {
	var fd uintptr = f.Fd()
	_, err := unix.IoctlGetTermios(int(fd), unix.TCGETS)
	return err == nil
}

func NewLogger(level LogLevel) *Logger          { return newLogger(level, false) }
func NewColorLogger(level LogLevel) *Logger     { return newLogger(level, true) }
func NewAutoColorLogger(level LogLevel) *Logger { return newLogger(level, isAtty(os.Stdout)) }

func (self *Logger) UsePanic(use bool) {
	self.panics = use
}

func (self *Logger) SetLevel(level LogLevel) {
	if !ValidLogLevel(level) {
		self.Fatal("Invalid level %#v (%s)", level, level)
	}
	if self.level != level {
		self.Say("Log level %s -> %s", self.level, level)
		self.level = level
	}
}

// unconditionally write to Logger with "NOTE:" prefix
func (self *Logger) Say(message string, args ...interface{}) {
	self.log.Printf("NOTE: "+message, args...)
}

func (self *Logger) colored(level LogLevel, message string) string {
	return level.Colored(level.Name()+":"+message, self.color)
}

func (self *Logger) Log(level LogLevel, message string, args ...interface{}) {
	if level <= self.level {
		self.log.Printf(self.colored(level, message), args...)
	}
}

func (self *Logger) Fatal(message string, args ...interface{}) {
	out := self.colored(FATAL, fmt.Sprintf(message, args...))
	f := self.log.Fatal
	if self.panics {
		f = self.log.Panic
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

var Root = NewAutoColorLogger(INFO)

/* EOF */
