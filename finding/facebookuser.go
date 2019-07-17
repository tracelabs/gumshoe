package finding

import "fmt"

const facebookType = "FacebookUser"

// FacebookUser is a user of Facebook
type FacebookUser struct {
	Name string
}

// GetID impl.
func (f *FacebookUser) GetID() string {
	return fmt.Sprintf("%s-%s", facebookType, f.Name)
}

// Investigate a Facebook user
func (f *FacebookUser) Investigate() []Finding {
	found := []Finding{}
	// TODO
	return found
}

// Display a Facebook user
func (f *FacebookUser) Display() {
	fmt.Printf("[facebook username] %s\n", f.Name)
}
