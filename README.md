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

	gotrends "github.com/kahalKevin/instagot"
)

func main() {
	http.HandleFunc("/image", GetImage)
	http.ListenAndServe(":8001", nil)

	// fmt.Println(instagot.GetImage("https://www.instagram.com/p/BUeej53hENz/"))
}

func GetImage(w http.ResponseWriter, r *http.Request){
	ig_url := r.URL.Query().Get("ig_url")
	img := instagot.GetImage(ig_url)
	instagot.WriteImageToResponseWriter(w, &img)
}
```

## Output
Check it at
```console
http://127.0.0.1:8001/image?ig_url=https://www.instagram.com/p/BUeej53hENz/

```
