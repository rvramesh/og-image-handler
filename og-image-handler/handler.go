package function

import (
	"bytes"
	"image/jpeg"
	"net/http"
	"strconv"

	"crypto/hmac"
	"crypto/sha512"

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

	h := hmac.New(sha512.New, []byte("2r5u8x/A?D*G-KaPdSgVkYp3s6v9y$B&E)H+MbQeThWmZq4t7w!z%C*F-JaNcRfU"))

	// Write Data to it
	h.Write([]byte(q.title))

	// titleHash := base64.StdEncoding.EncodeToString(h.Sum(nil))
	// if titleHash != q.sig {
	// 	return handler.Response{
	// 		Header: map[string][]string{
	// 			"X-Served-By": {"My Awesome Function"},
	// 		},
	// 		StatusCode: http.StatusBadRequest,
	// 	}, nil
	// }

	buffer, err := getOgImage("../template2.jpg", q.title)

	return handler.Response{
		Body: buffer.Bytes(),
		Header: map[string][]string{
			"X-Served-By":    {"pico3"},
			"Content-Type":   {"image/jpeg"},
			"Content-Length": {strconv.Itoa(len(buffer.Bytes()))},
		},
		StatusCode: http.StatusOK,
	}, err
}

func parseQueryString(req handler.Request) Query {
	var q Query
	err := qs.Unmarshal(&q, req.QueryString)
	if err != nil {
		panic(err)
	}
	return q
}

func getOgImage(templateLocation string, title string) (*bytes.Buffer, error) {
	var err error

	backgroundImage, err := gg.LoadImage(templateLocation)
	if err != nil {
		panic(err)
	}

	maxBounds := backgroundImage.Bounds().Max
	dc := gg.NewContext(maxBounds.X, maxBounds.Y)
	dc.DrawImage(backgroundImage, 0, 0)
	if err := dc.LoadFontFace("../SourceSans3-Regular.ttf", 48); err != nil {
		panic(err)
	}
	dc.SetRGB(1, 1, 1)

	dc.DrawStringWrapped(title, 80, 493, 0, 1, 1026, 1.25, gg.AlignLeft)
	buffer := new(bytes.Buffer)
	img := dc.Image()
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		panic(err)
	}
	return buffer, nil
}
