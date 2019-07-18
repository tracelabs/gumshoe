package findings

import (
	"fmt"

	"github.com/tracelabs/gumshoe"
)

// VenmoType is the name of the venmo username finding struct
const VenmoType = "VenmoUser"

// VenmoUser is a user of Venmo
type VenmoUser struct {
	Name string
}

// Investigate a Venmo user
func (v *VenmoUser) Investigate() []gumshoe.Finding {
	found := []gumshoe.Finding{}
	// TODO
	return found
}

// GetTypeName impl.
func (v *VenmoUser) GetTypeName() string {
	return VenmoType
}

// GetID impl.
func (v *VenmoUser) GetID() string {
	return v.Name
}

// Display a Venmo user
func (v *VenmoUser) Display() {
	fmt.Printf("[venmo username] %s\n", v.Name)
}
