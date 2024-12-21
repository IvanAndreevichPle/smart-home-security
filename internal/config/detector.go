// internal/config/detector.go
package config

type DetectorConfig struct {
	MinArea       float64 `yaml:"minArea"`
	Sensitivity   float64 `yaml:"sensitivity"`
	CheckInterval int     `yaml:"checkInterval"`
}
