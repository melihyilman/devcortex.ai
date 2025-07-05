package handler

import (
	"archive/zip"
	"bytes"
	"fmt"
	"image"
	"image/png"
	"net/http"

	"devcortex.ai/internal/view"
	"github.com/nfnt/resize"

	// Anonymous import for image decoding
	_ "image/jpeg"
	_ "image/png"
)

func FaviconGeneratorTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Favicon Generator",
	}

	if r.Method == http.MethodPost {
		file, _, err := r.FormFile("image_file")
		if err != nil {
			data.ToolSpecificData = map[string]interface{}{"Error": "Could not read uploaded file."}
			view.Render(w, r, "favicon-generator.html", data)
			return
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			data.ToolSpecificData = map[string]interface{}{"Error": "Could not decode image. Please upload a valid PNG or JPG file."}
			view.Render(w, r, "favicon-generator.html", data)
			return
		}

		sizes := []uint{16, 32, 48, 180}
		var zipBuffer bytes.Buffer
		zipWriter := zip.NewWriter(&zipBuffer)

		for _, size := range sizes {
			resizedImg := resize.Resize(size, size, img, resize.Lanczos3)
			var fileName string
			if size == 180 {
				fileName = "apple-touch-icon.png"
			} else {
				fileName = fmt.Sprintf("favicon-%dx%d.png", size, size)
			}

			f, err := zipWriter.Create(fileName)
			if err != nil {
				// handle error
			}
			png.Encode(f, resizedImg)
		}
		// A proper .ico file contains multiple sizes. For simplicity, we'll just add the 32x32 as favicon.ico
		ico, _ := zipWriter.Create("favicon.ico")
		png.Encode(ico, resize.Resize(32, 32, img, resize.Lanczos3))

		zipWriter.Close()

		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", "attachment; filename=\"favicons.zip\"")
		w.Write(zipBuffer.Bytes())
		return // Return after writing the zip file
	}

	view.Render(w, r, "favicon-generator.html", data)
}
