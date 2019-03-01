package winserv

import (
	"log"
	"os"
	"path/filepath"
)

// RedirectLog helps with logging into file output while debugging Windows Service.
func RedirectLog(file string) (*os.File, error) {
	f, err := os.OpenFile(filepath.Join(file), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	log.SetOutput(f)
	return f, nil
}
