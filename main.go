package google_trends

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type News struct {
	Headline     string `xml:"news_item_title"`
	HeadlineLink string `xml:"news_item_url"`
}

func main() {
	var rss RSS
	data := readGoogleTrends()
	err := xml.Unmarshal(data, &rss)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n Below are all the Google Search Trends in Canada For Today! ")
	fmt.Println("---------------------------------------------------------------")

	for i := range rss.Channel.ItemList {
		rank := i + 1
		fmt.Println("#", rank)
		fmt.Println("Search Term: ", rss.Channel.ItemList[i].Title)
		fmt.Println("Link to the Trend: ", rss.Channel.ItemList[i].Link)
		fmt.Println("Headline: ", rss.Channel.ItemList[i].NewsItems[0].Headline)
		fmt.Println("---------------------------------------------------------------")
	}
}

func readGoogleTrends() []byte {
	response := getGoogleTrends()
	data, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return data
}

// get http request
func getGoogleTrends() *http.Response {
	response, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=CA")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return response
}
