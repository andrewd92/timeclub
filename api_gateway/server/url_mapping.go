package server

func mapUrls() {
	router.Any("/*proxyPath", ReverseProxy)
}
