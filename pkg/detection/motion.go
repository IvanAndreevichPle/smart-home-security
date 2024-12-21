// pkg/detection/motion.go
package detection

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"gocv.io/x/gocv"
)

// MotionDetector структура для обнаружения движения в видеопотоке
type MotionDetector struct {
	firstFrame gocv.Mat
	config     DetectorConfig
	savePath   string
}

// NewMotionDetector создает новый экземпляр детектора движения
func NewMotionDetector(config DetectorConfig, savePath string) *MotionDetector {
	return &MotionDetector{
		config:   config,
		savePath: savePath,
	}
}

// DetectMotion анализирует кадр на наличие движения
func (md *MotionDetector) DetectMotion(frame gocv.Mat) (bool, *gocv.Mat) {
	// Конвертация в оттенки серого
	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(frame, &gray, gocv.ColorBGRToGray)

	// Размытие для уменьшения шума
	gocv.GaussianBlur(gray, &gray, image.Pt(21, 21), 0, 0, gocv.BorderDefault)

	// Сохранение первого кадра как эталонного
	if md.firstFrame.Empty() {
		gray.CopyTo(&md.firstFrame)
		return false, nil
	}

	// Поиск разницы между кадрами
	frameDelta := gocv.NewMat()
	defer frameDelta.Close()
	gocv.AbsDiff(md.firstFrame, gray, &frameDelta)

	// Бинаризация разницы
	thresh := gocv.NewMat()
	defer thresh.Close()
	gocv.Threshold(frameDelta, &thresh, float32(md.config.Sensitivity), 255, gocv.ThresholdBinary)

	// Морфологическая операция для улучшения контуров
	kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
	defer kernel.Close()
	gocv.Dilate(thresh, &thresh, kernel)

	// Поиск контуров движения
	contours := gocv.FindContours(thresh, gocv.RetrievalExternal, gocv.ChainApproxSimple)
	defer contours.Close()

	result := frame.Clone()
	motionDetected := false

	// Анализ найденных контуров
	for i := 0; i < contours.Size(); i++ {
		contour := contours.At(i)
		area := gocv.ContourArea(contour)

		// Пропуск маленьких контуров
		if area < md.config.MinArea {
			continue
		}

		// Отрисовка прямоугольника вокруг области движения
		rect := gocv.BoundingRect(contour)
		gocv.Rectangle(&result, rect, color.RGBA{
			R: 0,
			G: 255,
			B: 0,
			A: 0,
		}, 2)
		motionDetected = true
	}

	// В методе DetectMotion изменить вызов saveFrame
	if motionDetected {
		if err := md.saveFrame(&result); err != nil {
			// Обработка ошибки сохранения
			fmt.Printf("Ошибка сохранения кадра: %v\n", err)
		}
		return true, &result
	}

	return false, nil
}

// saveFrame сохраняет кадр в файл как метод структуры MotionDetector
func (md *MotionDetector) saveFrame(frame *gocv.Mat) error {
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s/motion_%s.jpg", md.savePath, timestamp)

	success := gocv.IMWrite(filename, *frame)
	if !success {
		return fmt.Errorf("не удалось сохранить изображение в %s", filename)
	}
	return nil
}
