// pkg/camera/rtsp.go
package camera

import (
	"fmt"

	"gocv.io/x/gocv"
)

type RTSPCamera struct {
	URL          string
	VideoCapture *gocv.VideoCapture
}

func NewRTSPCamera(url string) *RTSPCamera {
	return &RTSPCamera{
		URL: url,
	}
}

func (c *RTSPCamera) Connect() error {
	capture, err := gocv.OpenVideoCapture(c.URL)
	if err != nil {
		return err
	}
	c.VideoCapture = capture
	return nil
}

func (c *RTSPCamera) ReadFrame() (*gocv.Mat, error) {
	frame := gocv.NewMat()
	if ok := c.VideoCapture.Read(&frame); !ok {
		return nil, fmt.Errorf("cannot read frame")
	}
	return &frame, nil
}
