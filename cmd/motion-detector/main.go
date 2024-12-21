// cmd/motion-detector/main.go
package main

import (
	"log"
	"time"

	"github.com/IvanAndreevichPle/smart-home-security/pkg/camera"
	"github.com/IvanAndreevichPle/smart-home-security/pkg/detection"
	"gocv.io/x/gocv"
)

func main() {
	config := detection.DetectorConfig{
		MinArea:       500,
		Sensitivity:   25,
		CheckInterval: 100,
	}

	cam := camera.NewRTSPCamera("rtsp://username:password@ip:port/stream")
	if err := cam.Connect(); err != nil {
		log.Fatal(err)
	}
	defer cam.VideoCapture.Close()

	detector := detection.NewMotionDetector(config, "./motion_frames")
	window := gocv.NewWindow("Motion Detection")
	defer window.Close()

	for {
		frame, err := cam.ReadFrame()
		if err != nil {
			log.Printf("Error reading frame: %v", err)
			continue
		}

		motionDetected, resultFrame := detector.DetectMotion(*frame)
		if motionDetected {
			log.Println("Движение обнаружено!")
			if resultFrame != nil {
				window.IMShow(*resultFrame)
				resultFrame.Close()
			}
		}

		if window.WaitKey(1) == 27 { // ESC
			break
		}

		time.Sleep(time.Duration(config.CheckInterval) * time.Millisecond)
	}
}
