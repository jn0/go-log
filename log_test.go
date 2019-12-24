package logging

import "testing"
import "strings"
import "github.com/stretchr/testify/assert"

func TestNames(t *testing.T) {
	for lvl, lnm := range LogLevelNames {
		assert.True(t,
			CanBeLogLevel(int(lvl)),
			"!CanBeLogLevel(%d) (%s)", int(lvl), lnm)
	}
}

func TestValues(t *testing.T) {
	for lvl := int(FATAL)-5; lvl < int(DEBUG)+5; lvl++ {
		lnm, ok := LogLevelNames[LogLevel(lvl)]
		assert.Equal(t,
			ValidLogLevel(LogLevel(lvl)), ok,
			"ValidLogLevel(%d) == %v != %v",
			lvl, ValidLogLevel(LogLevel(lvl)), ok)
		xlv := LogLevelByName(lnm)
		if ok {
			assert.Equal(t,
				int(xlv), lvl,
				"LogLevelByName(%q) == %d != %d",
				lnm, xlv, lvl)
			assert.True(t,
				CanBeLogLevel(lvl),
				"!CanBeLogLevel(%d) (%s)", lvl, lnm)
			assert.True(t,
				ValidLogLevelName(lnm),
				"!ValidLogLevelName(%q)", lnm)
		} else {
			assert.False(t,
				lvl != int(INVALID) && xlv != INVALID,
				"LogLevelByName(%q) == %d != %d",
				lnm, xlv, INVALID)
			assert.False(t,
				CanBeLogLevel(int(lvl)),
				"CanBeLogLevel(%d)", lvl)
			assert.False(t,
				ValidLogLevelName(lnm),
				"ValidLogLevelName(%q)", lnm)
		}
		lnm = LogLevelName(LogLevel(lvl))
		t.Logf("level=%d name=%q ok=%v", lvl, lnm, ok)
	}
	goodLevel := LogLevel(1)
	t.Logf("goodLevel=%d=%q", goodLevel, goodLevel)
	assert.False(t,
		strings.HasPrefix(goodLevel.String(), "<Unknown log level #"),
		"Bad value of %q for %d", goodLevel, goodLevel)
	badLevel := LogLevel(-9)
	t.Logf("badLevel=%d=%q", badLevel, badLevel)
	assert.True(t,
		strings.HasPrefix(badLevel.String(), "<Unknown log level #"),
		"Bad value of %q for %d", badLevel, badLevel)
}

func TestRoot(t *testing.T) {
	assert.NotPanics(t, func() { adjustLogLevel("debug") },
		`adjustLogLevel("debug")`)
	assert.NotPanics(t, func() { Root.UsePanic(true) }, `UsePanic(true)`)
	assert.Panics(t,
		func() { adjustLogLevel("no-level") },
		`adjustLogLevel("no-level")`)

	assert.NotPanics(t, func() { Root.SetLevel(DEBUG) }, `SetLevel(DEBUG)`)
	assert.NotPanics(t, func() { Root.Log(-1, "Log: hello %v", "world") },
		`Log(-1, "Log: hello %v", "world")`)
	assert.NotPanics(t, func() { Root.Say("Say: hello %v", "world") },
		`Say("Say: hello %v", "world")`)
	assert.NotPanics(t, func() { Root.Error("Error: hello %v", "world") },
		`Error("Error: hello %v", "world")`)
	assert.NotPanics(t, func() { Root.Warning("Warning: hello %v", "world") },
		`Warning("Warning: hello %v", "world")`)
	assert.NotPanics(t, func() { Root.Warn("Warn: hello %v", "world") },
		`Warn("Warn: hello %v", "world")`)
	assert.NotPanics(t, func() { Root.Info("Info: hello %v", "world") },
		`Info("Info: hello %v", "world")`)
	assert.NotPanics(t, func() { Root.Debug("Debug: hello %v", "world") },
		`Debug("Debug: hello %v", "world")`)
	assert.Panics(t,
		func() { Root.SetLevel(-99) },
		`Root.SetLevel(-99)`)
	assert.Panics(t,
		func() { Root.Fatal("Fatal: hello %v", "world") },
		`Root.Fatal("Fatal: hello %v", "world")`)
}

func TestNew(t *testing.T) {
	assert.Panics(t, func() { NewLogger(-100) }, `NewLogger(-100)`)
	var log2 *Logger
	assert.NotPanics(t, func() { log2 = NewLogger(DEBUG) },
		`log2 = NewLogger(DEBUG)`)
	assert.NotPanics(t, func() { log2.UsePanic(true) },
		`log2.UsePanic(true)`)
	assert.NotPanics(t, func() { log2.Log(-1, "Log: hello %v", "world") },
		`log2.Log(-1, "Log: hello %v", "world")`)
	assert.NotPanics(t, func() { log2.Say("Say: hello %v", "world") },
		`log2.Say("Say: hello %v", "world")`)
	assert.NotPanics(t, func() { log2.Error("Error: hello %v", "world") },
		`log2.Error("Error: hello %v", "world")`)
	assert.NotPanics(t, func() { log2.Warning("Warning: hello %v", "world") },
		`log2.Warning("Warning: hello %v", "world")`)
	assert.NotPanics(t, func() { log2.Warn("Warn: hello %v", "world") },
		`log2.Warn("Warn: hello %v", "world")`)
	assert.NotPanics(t, func() { log2.Info("Info: hello %v", "world") },
		`log2.Info("Info: hello %v", "world")`)
	assert.NotPanics(t, func() { log2.Debug("Debug: hello %v", "world") },
		`log2.Debug("Debug: hello %v", "world")`)
	assert.Panics(t,
		func() { log2.SetLevel(-99) },
		`log2.SetLevel(-99)`)
	assert.Panics(t,
		func() { log2.Fatal("Fatal: hello %v", "world") },
		`log2.Fatal("Fatal: hello %v", "world")`)
}

/* EOF */
