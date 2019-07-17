# gumshoe

Recursive OSINT investigation tool based on the [OSINT Framework](https://osintframework.com/) - a [TraceLabs](https://tracelabs.org) backed project

The idea is that the tool is fed an initial (set of) finding(s). These findings will be used to come up with more findings, which will in turn be used to come up with even more findings in a recursive manner... until no findings are generated -- duplicates are ignored.

### Some Concepts

- Investigation: the "current" ongoing investigation
- Finding: any data that is either fed into or produced by the investigation. These are many distinct kinds of finding types:
	- Personally Identifiable Information:
		- email
		- phone number
		- home address
	- Digital Aliases:
		- usernames
		- social media accounts
		- images (reverse image search)
	- Geolocation: 
		- physical coordinates of an event
	- Info:
		- school attended
		- family members
		- friends
		
### Vision:

- Ultimately, the [finding](https://github.com/tracelabs/gumshoe/tree/master/finding) package will contain a vast amount of implementations, each with their version of the Investigate() method, which produces 0 or more findings (of the same, or of other types). The produced findings are Investigate()d, and the pattern continues recursively.
 
### Proposed Design:

![](./docs/assets/initial_design.png)
