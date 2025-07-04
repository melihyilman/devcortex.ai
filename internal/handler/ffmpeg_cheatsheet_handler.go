package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func FFmpegCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"Basic Conversion": {
			{Command: "ffmpeg -i input.mp4 output.avi", Description: "Convert a video from one format to another."},
			{Command: "ffmpeg -i input.mp4 output.gif", Description: "Convert a video into an animated GIF."},
			{Command: "ffmpeg -i input.mp3 output.wav", Description: "Convert an audio file from one format to another."},
		},
		"Video Manipulation": {
			{Command: "ffmpeg -i input.mp4 -ss 00:01:30 -t 00:00:10 output.mp4", Description: "Cut a 10-second clip starting from 1 minute 30 seconds."},
			{Command: "ffmpeg -i input.mp4 -vf scale=640:480 output.mp4", Description: "Resize a video to 640x480 pixels."},
			{Command: "ffmpeg -i input.mp4 -c:v libx264 -crf 28 output.mp4", Description: "Change video quality (CRF). Lower is better quality (18-28 is a good range)."},
			{Command: "ffmpeg -i input.mp4 -an output.mp4", Description: "Remove the audio track from a video."},
		},
		"Audio Manipulation": {
			{Command: "ffmpeg -i input.mp4 -vn -ab 192k output.mp3", Description: "Extract the audio track from a video and save as MP3."},
			{Command: "ffmpeg -i input.mp3 -af 'volume=0.5' output.mp3", Description: "Change the volume of an audio file (0.5 is half, 2.0 is double)."},
		},
		"Subtitles": {
			{Command: "ffmpeg -i input.mp4 -i subtitles.srt -c copy -c:s mov_text output.mp4", Description: "Add a subtitle track to a video."},
			{Command: "ffmpeg -i input.mp4 -vf subtitles=subtitles.srt output.mp4", Description: "Burn subtitles directly onto the video (hardsubs)."},
		},
	}

	data := &view.PageData{
		Title:            "FFmpeg Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "ffmpeg-cheatsheet.html", data)
}
