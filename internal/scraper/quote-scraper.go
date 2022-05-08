package scraper

import (
	"github.com/gocolly/colly/v2"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//goland:noinspection SpellCheckingInspection
var quotesHardcoded = []string{"Jangan Pernah Lelah Mencintai Negeri Ini.",
	"Kurang cerdas bisa diperbaiki dengan belajar. Kurang cakap dapat dihilangkan dengan pengalaman. Namun tidak jujur itu sulit diperbaiki.",
	"Setiap orang menjadi guru, setiap rumah menjadi sekolah.",
	"Bermimpilah setinggi langit. Jika kamu jatuh, kamu akan jatuh di antara bintang-bintang.",
	"Kita tunjukan bahwa kita adalah benar-benar orang yang ingin merdeka, lebih baik kita hancur lebur daripada tidak merdeka.",
	"Bangsa yang tidak percaya kepada kekuatan dirinya sebagai suatu bangsa, tidak dapat berdiri sebagai suatu bangsa yang merdeka.",
	"Pahlawan yang setia itu berkorban bukan buat dikenal namanya, tetapi semata-mata membela cita-cita.",
	"Negeri ini, Republik Indonesia, bukanlah milik suatu golongan, bukan milik suatu agama, bukan milik suatu kelompok etnis, bukan juga milik suatu adat-istiadat tertentu, tapi milik kita semua dari Sabang sampai Merauke!",
	"Kadang kita terlalu sibuk memikirkan kesulitan-kesulitan sehingga kita tidak punya waktu untuk mensyukuri rahmat Tuhan.",
	"Banyak hal yang bisa menjatuhkanmu. Tapi, satu-satunya hal yang benar-benar dapat menjatuhkanmu adalah sikapmu sendiri.",
	"Kejahatan akan menang bila orang yang benar tidak melakukan apa-apa.",
	"Indonesia merdeka bukan tujuan akhir kita. Indonesia merdeka hanya syarat untuk bisa mencapai kebahagiaan dan kemakmuran rakyat.",
	"Kalau sistem itu tak bisa diperiksa kebenarannya dan tak bisa dikritik, maka matilah ilmu pasti itu.",
	"Jangan pernah berhenti dan jangan pernah lelah untuk memberikan yang terbaik dan melakukan inovasi serta kreativitas untuk mendukung tugas-tugas DJBC."}
var authorsHardcoded = []string{"Sri Mulyani", "Mohammad Hatta", "Ki Hadjar Dewantara", "Soekarno", "Bung Tomo", "Soekarno", "Mohammad Hatta", "Soekarno", "Soedirman", "R.A. Kartini", "Soedirman", "Mohammad Hatta", "Tan Malaka", "Askolani"}
var uniqueTimeNow = rand.NewSource(time.Now().UnixNano())

// ScrapeQuoteWithTimeLimiter limit how long QuoteScraper can run, otherwise return default value Bea Cukai Makin Baik
func ScrapeQuoteWithTimeLimiter(timeLimitInMs float64) (quote string, author string) {
	timeLimit := time.Duration(timeLimitInMs) * time.Millisecond
	done := make(chan struct{})
	go func() {
		quote, author = QuoteScraper()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(timeLimit): //Not finished scraping before timeLimitInMs
		quote, author = getHardcodedQuoteAndAuthor()
	}
	return quote, author
}

//QuoteScraper scrape quotes from quotes.toscrape.com on random page (1-10) and random n-th of quote found on the page
func QuoteScraper() (string, string) {
	c := colly.NewCollector()
	var webQuotes, webAuthors []string
	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		webQuotes = append(webQuotes, e.ChildText(".text"))
		webAuthors = append(webAuthors, e.ChildText(".author"))
	})
	err := c.Visit(getRandomPage())
	if err != nil {
		return getHardcodedQuoteAndAuthor()
	}
	if len(webQuotes) > 0 && len(webQuotes) == len(webAuthors) { //Total webQuotes must be the same with total of webAuthors (pairing to each other)
		return getWebQuoteAndAuthor(webQuotes, webAuthors)
	} else {
		return getHardcodedQuoteAndAuthor()
	}
}

func getHardcodedQuoteAndAuthor() (quote string, author string) {
	randomForHardcoded := rand.New(uniqueTimeNow).Intn(len(authorsHardcoded) - 1)
	return quotesHardcoded[randomForHardcoded], authorsHardcoded[randomForHardcoded]
}

func getRandomPage() (page string) {
	pageRandom := rand.New(uniqueTimeNow).Intn(9) + 1 //will result of random number from 1 to 10
	return "https://quotes.toscrape.com/page/" + strconv.Itoa(pageRandom)
}

func getWebQuoteAndAuthor(quotes, authors []string) (quote string, author string) {
	quoteRandom := rand.New(uniqueTimeNow).Intn(len(quotes) - 1)
	return formatScrapedQuote(quotes[quoteRandom]), authors[quoteRandom]
}

func formatScrapedQuote(quote string) (formattedQuote string) {
	return strings.ReplaceAll(strings.ReplaceAll(quote, "“", ""), "”", "")
}
