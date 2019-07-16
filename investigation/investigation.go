package investigation

import (
	"fmt"
	"github.com/tracelabs/gumshoe/finding"
	"sync"
	"time"
)

// investigation represents the currently ongoing investigation
type investigation struct {
	sync.RWMutex
	Findings  map[string]finding.Finding
	StartedAt time.Time
	Queue     []finding.Finding
}

// Run begins the investigation by populating the Investigation
// with the initial finding(s) provided
func Run(initFindings ...finding.Finding) error {
	if len(initFindings) == 0 {
		return fmt.Errorf("no initial finding provided")
	}
	i := &investigation{
		Findings:  make(map[string]finding.Finding),
		StartedAt: time.Now(),
		Queue:     initFindings,
	}
	return i.run()
}

func (i *investigation) run() error {
	// TODO
	return nil
}
