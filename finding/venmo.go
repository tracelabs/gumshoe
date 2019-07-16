package finding

// VenmoAccount is the finding implementation for a venmo account
type VenmoAccount struct {
	Username string
	URL      string
}

// GetUID returns the unique url of the venmo account profile
func (v *VenmoAccount) GetUID() string {
	return v.URL
}

// GetType returns the correct finding type
func (v *VenmoAccount) GetType() string {
	return FindingTypeDigitalAlias
}

// Investigate runs an investigation on the venmo account
func (v *VenmoAccount) Investigate() []Finding {
	// TODO
	return nil
}
