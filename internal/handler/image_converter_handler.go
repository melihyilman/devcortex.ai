package handler

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"strings"

	"devcortex.ai/internal/view"
	// Anonymous import for image decoding
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func ImageConverterTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Image Converter",
	}

	if r.Method == http.MethodPost {
		file, header, err := r.FormFile("image_file")
		if err != nil {
			data.ToolSpecificData = map[string]interface{}{"Error": "Could not read uploaded file."}
			view.Render(w, r, "image-converter.html", data)
			return
		}
		defer file.Close()

		img, originalFormat, err := image.Decode(file)
		if err != nil {
			data.ToolSpecificData = map[string]interface{}{"Error": "Could not decode image. Please upload a valid PNG, JPG, or GIF file."}
			view.Render(w, r, "image-converter.html", data)
			return
		}

		targetFormat := r.FormValue("target_format")
		var outputBuffer bytes.Buffer
		var contentType string
		var newFileName string

		originalFileName := strings.TrimSuffix(header.Filename, "."+originalFormat)

		switch targetFormat {
		case "png":
			contentType = "image/png"
			newFileName = fmt.Sprintf("%s.png", originalFileName)
			err = png.Encode(&outputBuffer, img)
		case "jpeg":
			contentType = "image/jpeg"
			newFileName = fmt.Sprintf("%s.jpg", originalFileName)
			err = jpeg.Encode(&outputBuffer, img, &jpeg.Options{Quality: 90})
		case "gif":
			contentType = "image/gif"
			newFileName = fmt.Sprintf("%s.gif", originalFileName)
			err = gif.Encode(&outputBuffer, img, &gif.Options{})
		default:
			data.ToolSpecificData = map[string]interface{}{"Error": "Invalid target format selected."}
			view.Render(w, r, "image-converter.html", data)
			return
		}

		if err != nil {
			data.ToolSpecificData = map[string]interface{}{"Error": "Failed to encode image to the target format."}
			view.Render(w, r, "image-converter.html", data)
			return
		}

		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", newFileName))
		w.Write(outputBuffer.Bytes())
		return // Return after writing the file
	}

	view.Render(w, r, "image-converter.html", data)
}
