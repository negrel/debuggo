package main

import (
	"github.com/negrel/debuggo/pkg/log"
)

func main() {
	log.Traceln("[TRACE] - Trace log")
	log.Debugln("[DEBUG] - Debug log")
	log.Infoln("[INFO] - Info log")
	log.Warnln("[WARN] - Warning log")
	log.Errorln("[ERROR] - Error log")
	log.Fatalln("[FATAL] - Fatal log")
	// Will never be called because of log.Fatal
	log.Panicln("[PANIC] - Panic log")
}
