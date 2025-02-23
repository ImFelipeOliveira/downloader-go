package downloader

import (
	"downloader-go/downloader/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/schollz/progressbar/v3"
)

func DownloadFiles(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error accessing %s, %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error %d accessing %s", resp.StatusCode, url)
	}
	filename := path.Base(url)

	downloadsPath, err := utils.SetDownloadPath()
	if err != nil {
		return err
	}

	filePath := path.Join(downloadsPath, filename)

	out, err := os.Create(filePath)

	if err != nil {
		return fmt.Errorf("error creating file %s: %v", filename, err)
	}
	defer out.Close()

	bar := progressbar.NewOptions(
		int(resp.ContentLength),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetWidth(40),
		progressbar.OptionSetDescription("Downloading file"),
		progressbar.OptionShowBytes(true),
	)

	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		return fmt.Errorf("error saving file %s: %v", filename, err)
	}

	fmt.Println("File saved:", filename)
	return nil
}
