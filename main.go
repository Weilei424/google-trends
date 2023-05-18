package google_trends

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
