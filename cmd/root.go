package cmd

import (
	"downloader-go/downloader"
	"downloader-go/scrapper"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() {
	app := &cli.App{
		Name:  "downloader-go",
		Usage: "Downloader-go allows downloading files with specific extensions",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Usage:    "URL of the page to be analyzed",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "extension",
				Aliases: []string{"e"},
				Usage:   "Desired extension(s), separated by commas (e.g., .zip, .mp3, .mp4)",
				Value:   ".zip, .mp3, .mp4",
			},
		},
		Action: func(c *cli.Context) error {
			url := c.String("url")
			extensions := c.String("extension")

			links, err := scrapper.ExtractLinks(url, extensions)
			if err != nil {
				log.Println("Error extracting links:", err)
				return err
			}

			if len(links) == 0 {
				fmt.Println("No files with the specified extensions were found.")
				return nil
			}

			fmt.Printf("Found %d files for download.\n", len(links))

			for _, link := range links {
				err := downloader.DownloadFiles(link)
				if err != nil {
					fmt.Printf("Error downloading %s: %v\n", link, err)
				}
			}

			fmt.Println("Download completed!")
			return nil
		},
	}

	app.Run(os.Args)
}
