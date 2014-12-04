package feed

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"diektronics.com/carter/dl/types"
)

type data struct {
	DateStamp string `xml:"channel>lastBuildDate"`
	ItemList  []item `xml:"channel>item"`
}

type item struct {
	Title   string `xml:"title"`
	Content string `xml:"encoded"`
}

func Link(linkRegexp string, show *types.Show) string {
	titleEp := fmt.Sprintf("%s\\.%s\\.720p.*\\.mkv",
		strings.ToLower(strings.Replace(deparenthesize(show.Name), " ", "\\.", -1)),
		strings.ToLower(show.Eps))
	reStr := "(?i)(?P<link>" + linkRegexp + titleEp + ")"
	ret, err := match(reStr, show.Blob)
	if err != nil {
		return ""
	}
	return ret["link"]
}

func ScrapeShows(url string) ([]*types.Show, time.Time, error) {
	var timestamp time.Time
	stuff, err := http.Get(url)
	if err != nil {
		return nil, timestamp, err
	}
	defer stuff.Body.Close()

	body, err := ioutil.ReadAll(stuff.Body)
	if err != nil {
		return nil, timestamp, err
	}

	var d *data
	err = xml.Unmarshal([]byte(string(body)), &d)
	if err != nil {
		return nil, timestamp, err
	}

	timestamp, err = date(d.DateStamp)
	if err != nil {
		return nil, timestamp, err
	}
	shows := []*types.Show{}
	for _, entry := range d.ItemList {
		title, eps := tokenize(entry.Title)
		title = parenthesize(title)
		shows = append(shows, &types.Show{
			Name: title,
			Eps:  eps,
			Blob: entry.Content})
	}
	return shows, timestamp, nil
}

func Season(ep string) (string, error) {
	reStr := `S(?P<season>\d+)E\d+`
	ret, err := match(reStr, ep)
	if err != nil {
		return ep, fmt.Errorf("Bad episode format: %v", ep)
	}
	season, err := strconv.Atoi(ret["season"])
	if err != nil {
		return ep, err
	}

	return fmt.Sprintf("Season%d", season), nil
}

// func (f Feed) SetLinks(shows []*common.Episode) ([]*common.Episode, error) {}
// 	// We range the array in reverse because episodes are added on the top of the feed,
// 	// and when a show has two episodes back to back, we will first find the newest one.
// 	for i := len(shows[name]) - 1; i >= 0; i-- {
// 		s := shows[name][i]
// 		if latest_ep < s.eps {
// 			log.Printf("title: %q episode: %q latest_ep: %q\n", name, s.eps, latest_ep)
// 			link := s.it.Link(d.linkRegexp)

// 			if len(link) != 0 {
// 				log.Printf("link: %q\n", link)
// 				log.Println("update latest_ep in DB")
//
// 				log.Println("download the thing")

// 			}
// 		}
// 	}
// }

func match(reStr string, s string) (map[string]string, error) {
	re := regexp.MustCompile(reStr)
	matches := re.FindStringSubmatch(s)
	if len(matches) == 0 {
		return nil, errors.New("no matches found")
	}
	ret := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if len(name) == 0 {
			continue
		}
		ret[name] = matches[i]
	}

	return ret, nil
}

func date(timestamp string) (time.Time, error) {
	format := "Mon, 02 Jan 2006 15:04:05 -0700"
	if timestamp == "" {
		timestamp = format
	}
	return time.Parse(format, timestamp)
}

func parenthesize(str string) string {
	// RlsBB doesn't use parenthesis when a Series name has a year attached to it,
	// eg. Castle (2009), but the DB has them.
	// So, if "title" ends with four digits, we are going to add
	// parenthesis around it.
	stuff := `\d{4}$`
	epsRegexp := regexp.MustCompile(stuff)
	return epsRegexp.ReplaceAllString(str, "($0)")
}

func deparenthesize(str string) string {
	stuff := `[\(|\)]`
	epsRegexp := regexp.MustCompile(stuff)
	return epsRegexp.ReplaceAllString(str, "")
}

func tokenize(title string) (string, string) {
	reStr := `(?P<name>.*)\s+(?P<eps>S\d{2}E\d{2})`
	ret, err := match(reStr, title)
	if err != nil {
		return title, ""
	}
	return ret["name"], ret["eps"]
}
