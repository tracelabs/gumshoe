package finding

import "fmt"

const venmoType = "VenmoUser"

// VenmoUser is a user of Venmo
type VenmoUser struct {
	name string
}

// GetID impl.
func (v *VenmoUser) GetID() string {
	return fmt.Sprintf("%s-%s", venmoType, v.name)
}

// Investigate a Venmo user
func (v *VenmoUser) Investigate() []Finding {
	found := []Finding{}
	// TODO
	return found
}
