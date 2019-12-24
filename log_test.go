package logging

import "testing"
import "strings"
import "github.com/stretchr/testify/assert"

func TestEmAll(t *testing.T) {
	for lvl, lnm := range LogLevelNames {
		if !CanBeLogLevel(int(lvl)) {
			t.Errorf("!CanBeLogLevel(%d) (%s)", int(lvl), lnm)
		}
	}
	for lvl := int(FATAL)-5; lvl < int(DEBUG)+5; lvl++ {
		lnm, ok := LogLevelNames[LogLevel(lvl)]
		if ValidLogLevel(LogLevel(lvl)) != ok {
			t.Errorf("ValidLogLevel(%d) == %v != %v",
				lvl, ValidLogLevel(LogLevel(lvl)), ok)
		}
		xlv := LogLevelByName(lnm)
		if ok {
			if int(xlv) != lvl {
				t.Errorf("LogLevelByName(%q) == %d != %d",
					lnm, xlv, lvl)
			}
			if !CanBeLogLevel(lvl) {
				t.Errorf("!CanBeLogLevel(%d) (%s)", lvl, lnm)
			}
			if !ValidLogLevelName(lnm) {
				t.Errorf("!ValidLogLevelName(%q)", lnm)
			}
		} else {
			if lvl != int(INVALID) && xlv != INVALID {
				t.Errorf("LogLevelByName(%q) == %d != %d",
					lnm, xlv, INVALID)
			}
			if CanBeLogLevel(int(lvl)) {
				t.Errorf("CanBeLogLevel(%d)", lvl)
			}
			if ValidLogLevelName(lnm) {
				t.Errorf("ValidLogLevelName(%q)", lnm)
			}
		}
		lnm = LogLevelName(LogLevel(lvl))
		t.Logf("level=%d name=%q ok=%v", lvl, lnm, ok)
	}
	goodLevel := LogLevel(1)
	t.Logf("goodLevel=%d=%q", goodLevel, goodLevel)
	if strings.HasPrefix(goodLevel.String(), "<Unknown log level #") {
		t.Errorf("Bad value of %q for %d", goodLevel, goodLevel)
	}
	badLevel := LogLevel(-9)
	t.Logf("badLevel=%d=%q", badLevel, badLevel)
	if !strings.HasPrefix(badLevel.String(), "<Unknown log level #") {
		t.Errorf("Bad value of %q for %d", badLevel, badLevel)
	}

	adjustLogLevel("debug")
	log.UsePanic(true)
	assert.Panics(t,
		func() { adjustLogLevel("no-level") },
		`adjustLogLevel("no-level")`)

	log.SetLevel(DEBUG)
	log.Log(-1, "Log: hello %v", "world")
	log.Say("Say: hello %v", "world")
	log.Error("Error: hello %v", "world")
	log.Warning("Warning: hello %v", "world")
	log.Warn("Warn: hello %v", "world")
	log.Info("Info: hello %v", "world")
	log.Debug("Debug: hello %v", "world")
	assert.Panics(t,
		func() { log.Fatal("Fatal: hello %v", "world") },
		`log.Fatal("Fatal: hello %v", "world")`)

	assert.Panics(t, func() { NewLogger(-100) }, `NewLogger(-100)`)
	log2 := NewLogger(DEBUG)
	log2.UsePanic(true)
	log2.Log(-1, "Log: hello %v", "world")
	log2.Say("Say: hello %v", "world")
	log2.Error("Error: hello %v", "world")
	log2.Warning("Warning: hello %v", "world")
	log2.Warn("Warn: hello %v", "world")
	log2.Info("Info: hello %v", "world")
	log2.Debug("Debug: hello %v", "world")
	assert.Panics(t,
		func() { log2.SetLevel(-99) },
		`log2.SetLevel(-99)`)
	assert.Panics(t,
		func() { log2.Fatal("Fatal: hello %v", "world") },
		`log2.Fatal("Fatal: hello %v", "world")`)
}

/* EOF */
