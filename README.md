# instagot
Go wrapper to get an image from Post Url in instagram

## Using
### go get
You need `go` installed and `GOPATH` in your `PATH` , then run :
```shell
$ go get github.com/kahalKevin/instagot
```

## Usage (Library)
```go
package main

import (
	"net/http"
	"encoding/json"

	instagot "github.com/kahalKevin/instagot"
)

type IG_url struct{
	Image_url 		string  	`json:"image_url"`
}

func main() {
	http.HandleFunc("/image", GetImage)
	http.HandleFunc("/urlimage", GetUrlImage)	
	http.ListenAndServe(":8001", nil)

	// fmt.Println(instagot.GetImage("https://www.instagram.com/p/BUeej53hENz/"))
}

func GetImage(w http.ResponseWriter, r *http.Request){
	ig_url := r.URL.Query().Get("ig_url")
	img := instagot.GetImage(ig_url)
	instagot.WriteImageToResponseWriter(w, &img)
}

func GetUrlImage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	ig_url := r.URL.Query().Get("ig_url")
	img_url := instagot.GetUrlImage(ig_url)

	json.NewEncoder(w).Encode(
		IG_url{Image_url: 	img_url},
	)
}
```

## Output
Check it at
```console
http://127.0.0.1:8001/image?ig_url=https://www.instagram.com/p/BUeej53hENz/

```

```console
http://127.0.0.1:8001/urlimage?ig_url=https://www.instagram.com/p/BUeej53hENz/

```
