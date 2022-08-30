package main

import (
	"encoding/json"
	"fmt"
	"ihojose.com/wikivoyage-stats/model"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/**
 * wikivoyage-stats-bot was developed by Jose Buelvas (ihojose)
 *
 * @author        @ihojose
 * @author_url    dev.ihojose.com
 * @licence       Apache Licence v2.0
 * @year          2022
 * @donations     buymeacoff.ee/ihojose
 */
func main() {

	var result = ""

	// Do API Requests
	for _, lang := range languages {
		log.Println("Getting Stats from ", lang, ".wikivoyage.org...")

		// Start api call
		res, err := http.Get("https://" + lang + ".wikivoyage.org/w/api.php?action=query&meta=siteinfo&formatversion=2&format=json&siprop=statistics")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Get Request Result
		data, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err.Error())
		}

		// Add to WikiVoyage Data list
		var item model.SiteInfo
		err = json.Unmarshal(data, &item)
		if err != nil {
			log.Fatalln(err)
		}

		// Now create wiki template
		var stat = model.LangVariable
		var prof = float64(item.Query.Statistics.Edits/item.Query.Statistics.Pages) * math.Pow(float64(item.Query.Statistics.Pages-item.Query.Statistics.Articles)/float64(item.Query.Statistics.Articles), 2)

		// Set Data
		stat = strings.Replace(stat, "{LANG}", lang, 1)
		stat = strings.Replace(stat, "{ARTICLES}", strconv.FormatInt(item.Query.Statistics.Articles, 10), 1)
		stat = strings.Replace(stat, "{FILES}", strconv.FormatInt(item.Query.Statistics.Images, 10), 1)
		stat = strings.Replace(stat, "{PAGES}", strconv.FormatInt(item.Query.Statistics.Pages, 10), 1)
		stat = strings.Replace(stat, "{USERS}", strconv.FormatInt(item.Query.Statistics.Users, 10), 1)
		stat = strings.Replace(stat, "{ACTIVEUSERS}", strconv.FormatInt(item.Query.Statistics.ActiveUsers, 10), 1)
		stat = strings.Replace(stat, "{ADMINS}", strconv.FormatInt(item.Query.Statistics.Admins, 10), 1)
		stat = strings.Replace(stat, "{EDITS}", strconv.FormatInt(item.Query.Statistics.Edits, 10), 1)
		stat = strings.Replace(stat, "{DEPTH}", fmt.Sprintf("%2.2f", prof), 1)

		// Add to result
		result += stat
	}

	log.Println("Now setting wiki template for https://es.wikivoyage.org/wiki/Plantilla:Variables/datos")

	// Finish
	fmt.Println("RESULT:")
	fmt.Println(result)
}
