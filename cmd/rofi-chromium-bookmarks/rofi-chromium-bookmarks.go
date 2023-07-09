package main

import (
	"log"
	"os"
	"sort"

	"github.com/twoojoo/rofi-utils-go/internal/chromium"
	"github.com/twoojoo/rofi-utils-go/internal/rofi"
)

func main() {
	title := " Chromium Bookmarks:"

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	bookmarks := chromium.GetBookmarks(homeDir, "/.config/chromium/Default/Bookmarks")

	rofiItems := []rofi.RofiItem{}

	for _, bookmark := range bookmarks.Roots["other"].Children {
		rofiItems = append(rofiItems, rofi.RofiItem{
			Value: bookmark.Url,
			Label:  bookmark.Name,
		})
	}

	sort.Sort(rofi.ByLabel(rofiItems))

	rofi.LaunchMenu(rofi.RofiMenuOptions{
		Title: title,
		Items: rofiItems,
		Theme: homeDir + "/Projects/go/rofi-utils-go/configs/themes/theme.rasi",
		OnSelect: func(label string, value string) {
			chromium.OpenPage(value)
		},
	})
}
