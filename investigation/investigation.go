package investigation

import (
	"sync"
	"time"

	"github.com/tracelabs/gumshoe/finding"
)

type investigation struct {
	sync.RWMutex
	ongoing   sync.WaitGroup
	processed map[string]bool
	findings  []finding.Finding
	startedAt time.Time
}

// Run begins the investigation with the initial finding(s) provided
func Run(fs ...finding.Finding) []finding.Finding {
	if fs == nil || len(fs) == 0 {
		return []finding.Finding{}
	}
	return investigate(fs...).findings
}

// investigate is a wrapper around the recursive processFindings func
func investigate(fs ...finding.Finding) *investigation {
	i := &investigation{
		ongoing:   sync.WaitGroup{},
		processed: make(map[string]bool),
		findings:  []finding.Finding{},
		startedAt: time.Now(),
	}
	i.processFindings(fs...)
	i.ongoing.Wait() // wait for all go routines to finish
	return i
}

func (i *investigation) processFindings(fs ...finding.Finding) {
	for _, f := range fs {
		if !i.isProcessed(f) {
			i.markProcessed(f)
			i.ongoing.Add(1)

			go func(ff finding.Finding) {
				defer i.ongoing.Done()
				i.processFindings(ff.Investigate()...)
			}(f)

		}
	}
}

func (i *investigation) markProcessed(f finding.Finding) {
	i.Lock()
	defer i.Unlock()
	i.findings = append(i.findings, f)
	i.processed[f.GetID()] = true
}

func (i *investigation) isProcessed(f finding.Finding) bool {
	i.RLock()
	defer i.RUnlock()
	_, ok := i.processed[f.GetID()]
	return ok
}
