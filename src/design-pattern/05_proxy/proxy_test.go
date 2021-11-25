package proxy

import "testing"

func TestStaticProxy(t *testing.T) {
	url := "http://baidu.com"
	download := &ConcreteDownload{URL: url}
	proxy := &DownloadProxy{
		URL:        url,
		downloader: download,
	}
	proxy.Download()
}
