package gumshoe

import (
	"fmt"
	"sync"
	"time"
)

type investigation struct {
	sync.RWMutex
	ongoing   sync.WaitGroup
	processed map[string]bool
	findings  []Finding
	startedAt time.Time
}

// Run begins the investigation with the initial finding(s) provided
func Run(fs ...Finding) []Finding {
	if fs == nil || len(fs) == 0 {
		return []Finding{}
	}
	return investigate(fs...).findings
}

// investigate is a wrapper around the recursive processFindings func
func investigate(fs ...Finding) *investigation {
	i := &investigation{
		ongoing:   sync.WaitGroup{},
		processed: make(map[string]bool),
		findings:  []Finding{},
		startedAt: time.Now(),
	}
	i.processFindings(fs...)
	i.ongoing.Wait() // wait for all go routines to finish
	return i
}

func (i *investigation) processFindings(fs ...Finding) {
	for _, f := range fs {
		if !i.isProcessed(f) {
			i.markProcessed(f)
			i.ongoing.Add(1)

			go func(ff Finding) {
				defer i.ongoing.Done()
				i.processFindings(ff.Investigate()...)
			}(f)

		}
	}
}

func (i *investigation) markProcessed(f Finding) {
	i.Lock()
	defer i.Unlock()
	i.findings = append(i.findings, f)
	i.processed[fmt.Sprintf("%s-%s", f.GetTypeName(), f.GetID())] = true
}

func (i *investigation) isProcessed(f Finding) bool {
	i.RLock()
	defer i.RUnlock()
	_, ok := i.processed[fmt.Sprintf("%s-%s", f.GetTypeName(), f.GetID())]
	return ok
}
