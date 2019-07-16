package finding

// Finding represents actions for all types of findings
type Finding interface {
	Investigate() []Finding
	GetUID() string
	GetType() string
}

const (
	// FindingTypePII is the finding type for findings related to
	// user personally identifiable info such as email, phone number,
	// home address, etc
	FindingTypePII = "PersonallyIdentifiableInformation"

	// FindingTypeDigitalAlias is the finding type for findings related to
	// social media accounts and online profiles
	FindingTypeDigitalAlias = "DigitalAlias"

	// FindingTypeImage is the finding type for relevant image URLs
	FindingTypeImage = "Event"

	// FindingTypeEventDetails is the finding type for dated/timestamped
	// events such as a public payment on Venmo
	FindingTypeEventDetails = "Event"

	// FindingTypeGeolocation is the finding type for events whose main
	// feature is geophysical coordinates or other location types
	FindingTypeGeolocation = "Geolocation"

	// FindingTypeGenericData is the finding type for findings which do not
	// fall into a more appropriate category
	FindingTypeGenericData = "GenericData"
)
