package instagot

import (
	"net/http"
	"strings"
	"log"
	"strconv"
	"bytes"
	"image"
	"image/jpeg"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func GetImage(url string) image.Image{
	// request and parse the front page
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	var AllBody string

	// Search for the Body
	Body, ok := scrape.Find(root, scrape.ByTag(atom.Body))
	if ok {	    
	    AllBody = scrape.Text(Body)
	}

	IndexOfDisplayUrl := strings.Index(AllBody, "display_url")
	IndexOfDisplayUrl += 12

	AfterSignDisplayUrl := AllBody[IndexOfDisplayUrl:]
	IndexOfLinkStart := strings.Index(AfterSignDisplayUrl, `"`)
	AfterSignDisplayUrl = AfterSignDisplayUrl[IndexOfLinkStart:]

	IndexOfLinkEnd := strings.Index(AfterSignDisplayUrl[1:], `"`)
	DisplayUrl := AfterSignDisplayUrl[1:IndexOfLinkEnd+1]

	response, e := http.Get(DisplayUrl)
    if e != nil {
        log.Fatal(e)
    }

    defer response.Body.Close()

	ig_img,err := jpeg.Decode(response.Body)
	if err != nil {
		return nil
	}
	return ig_img
}

func WriteImageToResponseWriter(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}