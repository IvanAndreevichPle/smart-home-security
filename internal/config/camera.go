// internal/config/camera.go
package config

import "fmt"

type CameraConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Channel  string
}

func (c *CameraConfig) GetRTSPURL() string {
	return fmt.Sprintf("rtsp://%s:%s@%s:%d/%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Channel)
}
