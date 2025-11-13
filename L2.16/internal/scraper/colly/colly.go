// Package colly implements a Scraper using the gocolly collector. It
// wires response/HTML handlers and downloads resources according to the
// provided configuration.
package colly

import (
	"fmt"
	"regexp"
	"strings"
	"sync/atomic"

	"time"

	"L2.16/internal/downloader"
	"L2.16/internal/input"
	"L2.16/internal/utils"
	c "github.com/gocolly/colly"
)

// Scraper is a colly-backed implementation of the scraping logic. It keeps
// stats such as total files and total size and retains the underlying
// collector instance.
type Scraper struct {
	startTime  time.Time
	collector  *c.Collector
	totalSize  atomic.Int64
	totalFiles atomic.Int64
	domain     string
}

// NewCollyScraper constructs a new Scraper configured according to the
// provided input.Data. It wires handlers based on enabled download types.
func NewCollyScraper(data *input.Data) *Scraper {

	collector := c.NewCollector(
		c.Async(data.Config.Async),
		c.AllowedDomains(data.Domain),
	)

	scraper := &Scraper{collector: collector, domain: data.Domain}

	if data.Flags.Flagr {
		collector.MaxDepth = int(data.Flags.Flagl + 1) // offset by 1 because Colly counts depth from 0
	} else {
		collector.MaxDepth = 1
	}

	collector.OnResponse(scraper.handleResponse())
	if data.Config.DownloadFonts {
		collector.OnResponse(scraper.handleFonts())
	}

	collector.OnError(scraper.handleError())
	collector.OnHTML(utils.BuildSelector(data.Config), scraper.handleHTML())

	collector.UserAgent = data.Config.UserAgent
	collector.IgnoreRobotsTxt = data.Config.IgnoreRobotsTxt
	collector.SetRequestTimeout(data.Config.RequestTimeout)

	if err := collector.Limit(&c.LimitRule{
		DomainGlob:  "*",
		Parallelism: data.Config.Parallelism,
		Delay:       data.Config.Delay,
		RandomDelay: data.Config.RandomDelay,
	}); err != nil {
		panic(fmt.Sprintf("wget: fatal: Colly failed to set collector limit: %v", err))
	}

	return scraper

}

// handleResponse returns a response handler that saves the response body
// to a localized path and updates scraper statistics.
func (s *Scraper) handleResponse() func(r *c.Response) {

	return func(r *c.Response) {

		localPath, err := utils.LocalizePath(r.Request.URL)
		if err != nil {
			utils.LogError(fmt.Sprintf("wget: failed to determine save path for ‘%s’", r.Request.URL.String()), err)
			return
		}

		savedFile := downloader.Save(localPath, r.Body)
		if savedFile.Err != nil {
			utils.LogError(fmt.Sprintf("wget: failed to save file ‘%s’: %v", localPath, savedFile.Err), savedFile.Err)
			return
		}

		s.totalFiles.Add(1)
		s.totalSize.Add(int64(savedFile.Size))

		utils.WgetHeader(r.Request.URL.String(), localPath, r.Headers.Get("Content-Type"), savedFile.Size)

	}

}

// handleFonts returns a response handler that scans CSS responses for
// font URLs and schedules them for download.
func (s *Scraper) handleFonts() func(r *c.Response) {

	return func(r *c.Response) {

		if !strings.HasSuffix(r.FileName(), ".css") {
			return
		}

		re := regexp.MustCompile(`url\(['"]?([^'")]+)['"]?\)`)
		matches := re.FindAllSubmatch(r.Body, -1)

		for _, match := range matches {
			if len(match) > 1 {
				absURL := r.Request.AbsoluteURL(string(match[1]))
				if absURL != "" {
					_ = r.Request.Visit(absURL) // Colly handles errors via callback
				}
			}
		}

	}

}

// handleError returns an error handler that logs host resolution and
// request-level errors.
func (s *Scraper) handleError() func(r *c.Response, err error) {
	return func(r *c.Response, err error) {
		if r == nil || r.Request == nil {
			utils.LogError("wget: request error (no response)", err)
			return
		}
		utils.LogError(fmt.Sprintf("wget: unable to resolve host address ‘%s’", r.Request.URL), err)
	}
}

// handleHTML returns an HTML element handler that extracts resource URLs
// from HTML elements and schedules them for visiting/downloading.
func (s *Scraper) handleHTML() func(*c.HTMLElement) {

	return func(e *c.HTMLElement) {

		var rawURL string

		switch e.Name {
		case "a", "link":
			rawURL = e.Attr("href")
		case "script", "video", "audio", "source", "iframe":
			rawURL = e.Attr("src")
		case "img":
			rawURL = e.Attr("src")
			if rawURL == "" {
				rawURL = e.Attr("data-src")
			}
		}

		if rawURL != "" {
			absURL := e.Request.AbsoluteURL(rawURL)
			if absURL != "" {
				_ = e.Request.Visit(absURL) // Colly handles errors via callback
			}
		}

	}

}

// Scrape starts the scraping process for the given URL. Although Colly's
// robot rules handling is configured during scraper initialization,
// it does not actually download the robots.txt file like wget does,
// so a manual visit to /robots.txt is performed here when not ignored.
func (s *Scraper) Scrape(URL string) {

	s.startTime = time.Now()

	if !s.collector.IgnoreRobotsTxt {
		_ = s.collector.Visit("http://" + s.domain + "/robots.txt")
	}

	_ = s.collector.Visit(URL) // Colly handles errors via callback

	s.collector.Wait()

}

// GetStats returns number of downloaded files, total size in bytes and the
// elapsed scraping duration.
func (s *Scraper) GetStats() (int, int64, time.Duration) {
	return int(s.totalFiles.Load()), s.totalSize.Load(), time.Since(s.startTime)
}
