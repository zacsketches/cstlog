package cstlog

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/apex/log"
)

// Handler formats Apex log entries with CST/CDT-aware timestamps.
type Handler struct {
	Writer io.Writer
	Zone   *time.Location
}

// New returns a new CST/CDT-aware CLI log handler.
func New(w io.Writer) *Handler {
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		// fallback to UTC if loading fails
		loc = time.UTC
	}
	return &Handler{
		Writer: w,
		Zone:   loc,
	}
}

// HandleLog writes the log entry with a timestamp in CST/CDT.
func (h *Handler) HandleLog(e *log.Entry) error {
	t := e.Timestamp.In(h.Zone).Format("2006-01-02 15:04:05 MST")
	_, err := fmt.Fprintf(h.Writer, "%s %-5s %s\n", t, strings.ToUpper(e.Level.String()), e.Message)
	return err
}

