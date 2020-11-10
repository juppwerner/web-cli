package main

import (
	"github.com/juppwerner/web-cli/cmd"
)

func main() {
	/*
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")
		for i := 0; i < 5; i++ {
			fmt.Println(feed.Items[i].Title)
		}
	*/
	cmd.Exec()
}
