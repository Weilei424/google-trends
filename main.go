package google_trends

import "encoding/xml"

type RSS struct {
	XMLName xml.Name
	Channel *Channel
}

type Channel struct {
	Title    string
	ItemList []Item
}

type Item struct {
	Title     string
	Link      string
	Traffic   string
	NewsItems []News
}

type News struct {
	Headline     string
	HeadlineLink string
}

func main() {
	readGoogleTrends()
}

func readGoogleTrends() {
	getGoogleTrends()
}

// get http request
func getGoogleTrends() {

}
