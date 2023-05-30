package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	transcodingVideo("./input.MOV", "../web-client/output")
}

func transcodingVideo(inputFile string, outputDirectory string) {
	segmentDuration := 10
	bitrates := []int{500, 1000, 2000}

	err := os.MkdirAll(outputDirectory, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating output directory: %s", err)
	}

	variantPlaylists := []string{}

	for _, bitrate := range bitrates {
		segmentOutputFile := fmt.Sprintf("%s/output_%dk%%03d.ts", outputDirectory, bitrate)
		playlistOutputFile := fmt.Sprintf("%s/output_%dk.m3u8", outputDirectory, bitrate)
		variantPlaylists = append(variantPlaylists, playlistOutputFile)

		cmd := exec.Command(
			"./ffmpeg",
			"-i", inputFile,
			"-c:v", "libx264",
			"-b:v", fmt.Sprintf("%dk", bitrate),
			"-c:a", "aac",
			"-map", "0",
			"-f", "segment",
			"-segment_list", playlistOutputFile,
			"-segment_time", fmt.Sprintf("%d", segmentDuration),
			segmentOutputFile,
		)
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Error transcoding video into segments for bitrate %dk: %s", bitrate, err)
		}
	}

	// Create a master playlist file
	masterPlaylist := fmt.Sprintf("%s/master.m3u8", outputDirectory)
	file, err := os.Create(masterPlaylist)
	if err != nil {
		log.Fatalf("Error creating master playlist file: %s", err)
	}
	defer file.Close()

	file.WriteString("#EXTM3U\n")
	for _, bitrate := range bitrates {
		file.WriteString(fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d000\n", bitrate))
		file.WriteString(fmt.Sprintf("output_%dk.m3u8\n", bitrate))
	}

	fmt.Println("Adaptive bitrate transcoding complete!")
}

func transcodingVideoOld(inputFile string) {
	outputDirectory := "output"        // Output directory for the transcoded files
	segmentDuration := 10              // Segment duration in seconds
	bitrates := []int{500, 1000, 2000} // Desired bitrates in kilobits per second

	// Create output directory if it doesn't exist
	err := os.MkdirAll(outputDirectory, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating output directory: %s", err)
	}

	// Generate segments for each bitrate
	for _, bitrate := range bitrates {
		outputFile := fmt.Sprintf("%s/output_%dk%%03d.ts", outputDirectory, bitrate)

		cmd := exec.Command("./ffmpeg", "-i", inputFile, "-c:v", "libx264", "-b:v", fmt.Sprintf("%dk", bitrate), "-c:a", "copy", "-map", "0", "-segment_time", fmt.Sprintf("%d", segmentDuration), "-f", "segment", outputFile)
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Error transcoding video into segments for bitrate %dk: %s", bitrate, err)
		}
	}

	// Generate HLS playlist (manifest) file
	playlistFile := fmt.Sprintf("%s/playlist.m3u8", outputDirectory)
	cmd := exec.Command("./ffmpeg", "-i", inputFile, "-c:v", "libx264", "-c:a", "aac", "-f", "hls", "-hls_time", fmt.Sprintf("%d", segmentDuration), "-hls_segment_filename", fmt.Sprintf("%s/output_%%d.ts", outputDirectory), playlistFile)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error generating HLS playlist file: %s", err)
	}

	fmt.Println("Adaptive bitrate transcoding complete!")
}
