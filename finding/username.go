package finding

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const usernameType = "GenericUsername"

var siteUsernameTemplates = map[string]string{
	"facebook": "https://facebook.com/{}",
	"venmo":    "https://venmo.com/{}",
	"github":   "https://github.com/{}",
	"twitter":  "https://twitter.com/{}",
	// add sites username templates here
}

var siteUsernameFoundHandler = map[string]func(string) Finding{
	"facebook": func(n string) Finding { return &FacebookUser{name: n} },
	"venmo":    func(n string) Finding { return &VenmoUser{name: n} },
	"github":   func(n string) Finding { return &GithubUser{name: n} },
	"twitter":  func(n string) Finding { return &TwitterUser{name: n} },
	// add site username found handler functions here
}

// Username is the finding type for a generic username search
type Username struct {
	name string
}

// GetID is equal for two Username's with the same name for the same site
func (u *Username) GetID() string {
	return fmt.Sprintf("%s-%s", usernameType, u.name)
}

// Investigate attempts to find existing users in several social media platforms
func (u *Username) Investigate() []Finding {
	found := []Finding{}

	for site, tmpl := range siteUsernameTemplates {
		url := strings.Replace(tmpl, "{}", u.name, 1)
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
			found = append(found, siteUsernameFoundHandler[site](u.name))
		}
	}

	return found
}
