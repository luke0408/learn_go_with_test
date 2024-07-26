package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) { // goroutine
			// 채널에 결과를 보냄
			resultsChannel <- result{u, wc(u)} // channel
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// 채널에서 결과를 받음
		r := <-resultsChannel // channel
		results[r.string] = r.bool
	}

	return results
}
