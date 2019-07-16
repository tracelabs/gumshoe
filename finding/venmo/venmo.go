package venmo

import "github.com/tracelabs/gumshoe/finding"

// Account is the finding implementation for a venmo account
type Account struct {
	Username string
	URL      string
}

// GetUID returns the unique url of the account profile
func (a *Account) GetUID() string {
	return a.URL
}

// GetType returns FindingTypeDigitalAlias
func (a *Account) GetType() string {
	return finding.FindingTypeDigitalAlias
}

// Investigate runs an investigation on the venmo account
func (a *Account) Investigate() []finding.Finding {
	// TODO
	return nil
}
