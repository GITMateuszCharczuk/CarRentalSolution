package constants

type FileExtensions string

const (
	JPEG FileExtensions = ".jpeg"
	JPG  FileExtensions = ".jpg"
	PNG  FileExtensions = ".png"
	GIF  FileExtensions = ".gif"
	BMP  FileExtensions = ".bmp"
)

func IsAllowedExtension(ext string) bool {
	switch FileExtensions(ext) {
	case JPEG, JPG, PNG, GIF, BMP:
		return true
	default:
		return false
	}
}
