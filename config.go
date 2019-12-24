package logging

import "io"

func (self *Logger) Prefix() string {
	return self.log.Prefix()
}

func (self *Logger) SetPrefix(prefix string) *Logger {
	self.log.SetPrefix(prefix)
	return self
}

func (self *Logger) SetFlags(flags int) *Logger {
	self.log.SetFlags(flags)
	return self
}

func (self *Logger) Flags() int {
	return self.log.Flags()
}

func (self *Logger) SetOutput(writer io.Writer) *Logger {
	self.log.SetOutput(writer)
	return self
}

/* old log version has no this method *
func (self *Logger) Output() io.Writer {
	return self.log.Writer()
}
*/

/* EOF */
