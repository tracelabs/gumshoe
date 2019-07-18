package findings

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/tracelabs/gumshoe"
)

// GenericUsernameType is the name of the generic username finding struct
const GenericUsernameType = "GenericUsername"

var siteUsernameTemplates = map[string]string{
	"facebook": "https://facebook.com/{}",
	"venmo":    "https://venmo.com/{}",
	"github":   "https://github.com/{}",
	"twitter":  "https://twitter.com/{}",
	// add sites username templates here
}

var siteUsernameFoundHandler = map[string]func(string) gumshoe.Finding{
	"facebook": func(n string) gumshoe.Finding { return &FacebookUser{Name: n} },
	"venmo":    func(n string) gumshoe.Finding { return &VenmoUser{Name: n} },
	"github":   func(n string) gumshoe.Finding { return &GithubUser{Name: n} },
	"twitter":  func(n string) gumshoe.Finding { return &TwitterUser{Name: n} },
	// add site username found handler functions here
}

// GenericUsername is the generic user of web services
type GenericUsername struct {
	Name string
}

// Investigate attempts to find existing users in several social media platforms
func (gu *GenericUsername) Investigate() []gumshoe.Finding {
	found := []gumshoe.Finding{}

	for site, tmpl := range siteUsernameTemplates {
		url := strings.Replace(tmpl, "{}", gu.Name, 1)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Println(err)
			continue
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}
		if resp.StatusCode == http.StatusOK {
			found = append(found, siteUsernameFoundHandler[site](gu.Name))
		}
	}

	return found
}

// GetTypeName impl.
func (gu *GenericUsername) GetTypeName() string {
	return GenericUsernameType
}

// GetID impl.
func (gu *GenericUsername) GetID() string {
	return gu.Name
}

// Display a generic user
func (gu *GenericUsername) Display() {
	fmt.Printf("[GENERIC USERNAME] %s\n", gu.Name)
}
