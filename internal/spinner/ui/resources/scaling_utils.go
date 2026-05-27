package res

import "github.com/hajimehoshi/ebiten/v2"

func applyIntDeviceScale(toScale int) int {
	return int(float64(toScale) * ebiten.Monitor().DeviceScaleFactor())
}

func applyF64DeviceScale(toScale float64) float64 {
	return toScale * ebiten.Monitor().DeviceScaleFactor()
}

func resizeImage(img *ebiten.Image, targetWidth, targetHeight int) *ebiten.Image {
	scaleX := float64(targetWidth) / float64(img.Bounds().Dx())
	scaleY := float64(targetHeight) / float64(img.Bounds().Dy())
	scaledImg := ebiten.NewImage(targetWidth, targetHeight)
	scaledImgOpts := &ebiten.DrawImageOptions{}
	scaledImgOpts.GeoM.Scale(scaleX, scaleY)
	scaledImgOpts.Filter = ebiten.FilterLinear

	scaledImg.DrawImage(img, scaledImgOpts)
	return scaledImg
}
