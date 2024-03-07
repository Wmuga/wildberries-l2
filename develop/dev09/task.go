package main

import (
	"flag"

	"github.com/Wmuga/wildberries-l2/develop/dev09/downloader"
	"github.com/Wmuga/wildberries-l2/develop/dev09/flags"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	flags, sites := flags.ParseArgs()
	if len(sites) == 0 {
		flag.Usage()
	}

	if flags.Recurse < 1 {
		flags.Recurse = 1
	}

	if flags.Wait < 1 {
		flags.Wait = 1
	}

	for i := range sites {
		downloader.DownloadPageToDirRecursive(sites[i], "/", flags.Prefix, flags.Wait, 0, flags.Recurse)
	}

}
