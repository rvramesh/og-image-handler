package function

import (
	"bytes"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"

	"github.com/fogleman/gg"
	handler "github.com/openfaas/templates-sdk/go-http"
	"github.com/pasztorpisti/qs"
)

type Query struct {
	title string
	// sig   string
}

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {

	q := parseQueryString(req)
	log.Print(req)
	log.Println("Req printed")
	//	h := hmac.New(sha512.New, []byte("2r5u8x/A?D*G-KaPdSgVkYp3s6v9y$B&E)H+MbQeThWmZq4t7w!z%C*F-JaNcRfU"))

	// Write Data to it
	//	h.Write([]byte(q.title))

	// titleHash := base64.StdEncoding.EncodeToString(h.Sum(nil))
	// if titleHash != q.sig {
	// 	return handler.Response{
	// 		Header: map[string][]string{
	// 			"X-Served-By": {"My Awesome Function"},
	// 		},
	// 		StatusCode: http.StatusBadRequest,
	// 	}, nil
	// }

	buffer := getOgImage("template2.jpg", "Hello World").Bytes()
	log.Println("Done getting buffer")
	return handler.Response{
		Body: buffer,
		Header: map[string][]string{
			"X-Served-By":    {"pico3"},
			"Content-Type":   {"image/jpeg"},
			"Content-Length": {strconv.Itoa(len(buffer))},
		},
		StatusCode: http.StatusOK,
	}, nil
}

func parseQueryString(req handler.Request) Query {
	var q Query
	log.Println(req.QueryString)
	err := qs.Unmarshal(&q, req.QueryString)
	if err != nil {
		panic(err)
	}
	return q
}

func getOgImage(templateLocation string, title string) *bytes.Buffer {

	log.Println("Starting image processing")
	backgroundImage, err := gg.LoadImage(templateLocation)
	if err != nil {
		panic(err)
	}
	log.Println("Loaded Image")
	dc := gg.NewContextForImage(backgroundImage)
	log.Println("Drew Image")
	if err := dc.LoadFontFace("SourceSans3-Regular.ttf", 48); err != nil {
		panic(err)
	}
	log.Println("Loaded Font")
	dc.SetRGB(1, 1, 1)

	dc.DrawStringWrapped(title, 80, 493, 0, 1, 1026, 1.25, gg.AlignLeft)
	log.Println("Wrote String")
	buffer := new(bytes.Buffer)
	img := dc.Image()
	log.Println("Converted to image")
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		panic(err)
	}
	log.Println("Done image processing")
	return buffer
}
