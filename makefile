cmd := go run .

test:
	$(cmd) register kahya
	$(cmd) addfeed Boot.dev https://blog.boot.dev/index.xml
	$(cmd) feeds
	$(cmd) following
	$(cmd) register holgith
	$(cmd) addfeed Hacker_news https://news.ycombinator.com/rss
	$(cmd) follow https://blog.boot.dev/index.xml
	$(cmd) feeds
	$(cmd) following

scrape_test:
	$(cmd) register kahya
	$(cmd) addfeed Boot.dev https://blog.boot.dev/index.xml
	$(cmd) feeds
	$(cmd) following


scrape_test2:
	$(cmd) register kahya
	$(cmd) addfeed Mono29 https://mono29.com/livetv/feed
	$(cmd) addfeed WorkPoint https://www.workpointtv.com/feed
	$(cmd) addfeed CatDumb https://www.catdumb.com/feed
	$(cmd) feeds
	$(cmd) following

reset:
	$(cmd) reset
