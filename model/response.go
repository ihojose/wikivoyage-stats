package model

/**
 * wikivoyage-stats-bot was developed by Jose Buelvas (ihojose)
 *
 * @author        @ihojose
 * @author_url    dev.ihojose.com
 * @licence       Apache Licence v2.0
 * @year          2022
 * @donations     buymeacoff.ee/ihojose
 */

type SiteInfo struct {
	BatchComplete bool  `json:"batchcomplete"`
	Query         Query `json:"query"`
}

type Query struct {
	Statistics Statistics `json:"statistics"`
}

type Statistics struct {
	Pages       int64 `json:"pages"`
	Articles    int64 `json:"articles"`
	Edits       int64 `json:"edits"`
	Images      int64 `json:"images"`
	Users       int64 `json:"users"`
	ActiveUsers int64 `json:"activeusers"`
	Admins      int64 `json:"admins"`
	Jobs        int64 `json:"jobs"`
}
