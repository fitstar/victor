package google

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
)

type ImageResult struct {
	UnescapedUrl string
}

type ImageResults struct {
	Results []ImageResult
}

type ImageResponseDate struct {
	ResponseData ImageResults
}

func ImageSearch(term string, gifOnly bool) (string, error) {
	search, err := url.Parse("http://ajax.googleapis.com/ajax/services/search/images")

	if err != nil {
		return "", err
	}

	q := search.Query()
	q.Add("v", "1.0")
	q.Add("rsz", "8")
	q.Add("q", term)
	q.Add("safe", "active")

	if gifOnly {
		q.Add("as_filetype", "gif")
	}

	search.RawQuery = q.Encode()

	resp, err := http.Get(search.String())

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Bad status from google: %v", resp.Status)
	}
	if err != nil {
		return "", err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	var result ImageResponseDate

	err = json.Unmarshal(buf, &result)

	if err != nil {
		return "", err
	}

	images := result.ResponseData.Results

	if len(images) > 0 {
		return images[rand.Intn(len(images))].UnescapedUrl, nil
	}

	return "", nil
}
