package finding

import "fmt"

const venmoType = "VenmoUser"

// VenmoUser is a user of Venmo
type VenmoUser struct {
	Name string
}

// GetID impl.
func (v *VenmoUser) GetID() string {
	return fmt.Sprintf("%s-%s", venmoType, v.Name)
}

// Investigate a Venmo user
func (v *VenmoUser) Investigate() []Finding {
	found := []Finding{}
	// TODO
	return found
}

// Display a Venmo user
func (v *VenmoUser) Display() {
	fmt.Printf("[venmo username] %s\n", v.Name)
}
