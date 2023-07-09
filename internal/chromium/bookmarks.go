package chromium

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type BookmarksFile struct {
	Checksum string `josn:"checksum"`
	Version  int    `json:"version"`
	Roots    map[string]BookmarkRoot
}

type BookmarkRoot struct {
	Children     []Bookmark `json"children"`
	DateAdded    string     `json:"date_added"`
	DateLastUsed string     `json:"date_last_used"`
	DateModified string     `json:"date_modified"`
	Guid         string     `json:"guid"`
	Id           string     `json:"id"`
	Type         string     `json:"type"`
	Url          string     `json:"url"`
}

type Bookmark struct {
	DateAdded    string           `json:"date_added"`
	DateLastUsed string           `json:"date_last_used"`
	Guid         string           `json:"guid"`
	Id           string           `json:"id"`
	MetaInfo     BookmarkMetaInfo `json:"meta_info"`
	Name         string           `json:"name"`
	Type         string           `json:"type"`
	Url          string           `json:"url"`
}

type BookmarkMetaInfo struct {
	PowerBookmarkMeta string `json:"power_bookmark_meta"`
}

func GetBookmarks(path ...string) *BookmarksFile {
	fullPath := ""
	for _, p := range path {
		fullPath += p
	}
	
	bookmarksFileContent, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	var bookmarksFile BookmarksFile
	err = json.Unmarshal(bookmarksFileContent, &bookmarksFile)
	if err != nil {
		log.Fatal(err)
	}

	return &bookmarksFile
}
