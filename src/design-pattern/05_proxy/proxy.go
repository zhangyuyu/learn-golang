package proxy

import "fmt"

type Downloader interface {
	Download()
}

type ConcreteDownload struct {
	URL string
}

func (c *ConcreteDownload) Download() {
	fmt.Println(fmt.Sprintf("%s 在下载中", c.URL))
}

type DownloadProxy struct {
	URL        string
	downloader Downloader
}

func (d *DownloadProxy) Download() {
	fmt.Println(fmt.Sprintf("[Start] 开始下载 %s", d.URL))
	d.downloader.Download()
	fmt.Println(fmt.Sprintf("[Done] 下载 %s 完成", d.URL))
}
