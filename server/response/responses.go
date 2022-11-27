package response

type Bookmark struct {
	BookmarkID uint64 `json:"bookmark_id"`
	SiteTitle  string `json:"site_title"`
	SiteURL    string `json:"site_url"`
	Note       string `json:"note"`
}
