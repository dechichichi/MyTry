package log

import (
	stlog "log"
	"os"
)

var log *stlog.Logger

type filelog string

func (f filelog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(f), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
}
