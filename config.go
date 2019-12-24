package logging

import "io"

func (self *Logger) Prefix() string {
	return self.log.Prefix()
}

func (self *Logger) SetPrefix(prefix string) {
	self.log.SetPrefix(prefix)
}

func (self *Logger) SetFlags(flags int) {
	self.log.SetFlags(flags)
}

func (self *Logger) Flags() int {
	return self.log.Flags()
}

func (self *Logger) SetOutput(writer io.Writer) {
	self.log.SetOutput(writer)
}

/* old log version has no this method *
func (self *Logger) Output() io.Writer {
	return self.log.Writer()
}
*/

/* EOF */
