package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"golang.org/x/net/html"
)

const targetButtonId = "buy-now"

// const targetButtonClass = "btn-add-cart btn-addcart"
const targetUrl = "https://www.advice.co.th/product/graphic-card-vga-/amd-radeon-rx-9000-series/vga-gigabyte-radeon-rx-9070xt-gaming-oc-16gb-gddr6"

// const targetUrl = "https://www.advice.co.th/product/graphic-card-vga-/amd-radeon-rx-7000-series/vga-power-color-radeon-rx-7800xt-16gb-gddr6"

const greenColor = "\033[32m"
const redColor = "\033[31m"
const defaultColor = "\033[0m"

// second
const timeInterval = 5

func main() {
	fmt.Println("Current listen on item  : " + targetUrl)

	for {
		statuscode, isFound, err := webListener(targetUrl)
		if !isFound && err != nil {
			fmt.Println(redColor+"Error:", err, defaultColor)
			time.Sleep(5 * time.Second)
			continue
		}
		if isFound {
			fmt.Printf("%s[%d]\tItem Instock right now !!! \n", greenColor, statuscode)
			//fmt.Println(statuscode+"Item Instock right now !!! " , )
			break
		} else {
			fmt.Printf("%s[%d]%s\tNo item in stock refetch again in %d sec.%s\n", greenColor, statuscode, redColor, timeInterval, defaultColor)
			//fmt.Println("No item in stock refetch again in 5 sec.")
			time.Sleep(timeInterval * time.Second)
		}

	}
}

func webListener(url string) (int, bool, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 500, false, err
	}
	request.Header.Set("Referer", "https://www.advice.co.th/")
	request.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	request.Header.Set("Pragma", "no-cache")
	request.Header.Set("Expires", "0")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookieJar}
	res, err := client.Do(request)
	if err != nil {
		return 500, false, err
	}
	defer res.Body.Close()
	// Check if response status is OK
	if res.StatusCode != http.StatusOK {
		return res.StatusCode, false, fmt.Errorf("received status code %d", res.StatusCode)
	}
	doc := html.NewTokenizer(res.Body)

	for {
		tokenType := doc.Next()
		switch tokenType {
		case html.ErrorToken:
			// End of the document
			return res.StatusCode, false, nil
		case html.StartTagToken:
			token := doc.Token()
			if token.Data == "div" {
				var idMatch, classMatch bool

				// Check attributes
				for _, attr := range token.Attr {
					if attr.Key == "id" && attr.Val == targetButtonId {
						idMatch = true
					}
					if attr.Key == "class" {
						classList := strings.Fields(attr.Val) // Split class names
						for _, class := range classList {
							if class == "btn-add-cart" || class == "btn-addcart" {
								classMatch = true
							}
						}
					}
				}

				if idMatch && classMatch {
					return res.StatusCode, true, nil
				}
			}
		}
	}

}
