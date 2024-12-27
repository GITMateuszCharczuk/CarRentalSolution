package factories

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/nfnt/resize"
)

const (
	ImageWidth  = 1920
	ImageHeight = 1080
)

var carImageURLs = []string{
	// Luxury Cars
	"https://images.unsplash.com/photo-1494976388531-d1058494cdd8",
	"https://images.unsplash.com/photo-1583121274602-3e2820c69888",
	"https://images.unsplash.com/photo-1503376780353-7e6692767b70",
	"https://images.unsplash.com/photo-1542282088-72c9c27ed0cd",
	"https://images.unsplash.com/photo-1553440569-bcc63803a83d",
	"https://images.unsplash.com/photo-1511919884226-fd3cad34687c",
	"https://images.unsplash.com/photo-1552519507-da3b142c6e3d",
	"https://images.unsplash.com/photo-1580273916550-e323be2ae537",
	"https://images.unsplash.com/photo-1514316454349-750a7fd3da3a",
	"https://images.unsplash.com/photo-1555353540-64580b51c258",
	// Sports Cars
	"https://images.unsplash.com/photo-1525609004556-c46c7d6cf023",
	"https://images.unsplash.com/photo-1544829099-b9a0c07fad1a",
	"https://images.unsplash.com/photo-1535732820275-9ffd998cac22",
	"https://images.unsplash.com/photo-1554744512-d6c603f27c54",
	"https://images.unsplash.com/photo-1588258524935-ab2f048a4c8c",
	// SUVs and Crossovers
	"https://images.unsplash.com/photo-1533473359331-0135ef1b58bf",
	"https://images.unsplash.com/photo-1549317661-bd32c8ce0db2",
	"https://images.unsplash.com/photo-1606016159991-dfe4f2746ad5",
	"https://images.unsplash.com/photo-1519641471654-76ce0107ad1b",
	"https://images.unsplash.com/photo-1533558701576-23c65e0272fb",
	// Classic Cars
	"https://images.unsplash.com/photo-1566008885218-90abf9200ddb",
	"https://images.unsplash.com/photo-1567343411792-32a88d5b6ee9",
	"https://images.unsplash.com/photo-1565043589221-1a6fd9ae45c7",
	"https://images.unsplash.com/photo-1562911791-c7a97b729ec5",
	"https://images.unsplash.com/photo-1567808291548-fc3ee04dbcf0",
	// Modern Cars
	"https://images.unsplash.com/photo-1618843479313-40f8afb4b4d8",
	"https://images.unsplash.com/photo-1617814076367-b759c7d7e738",
	"https://images.unsplash.com/photo-1619767886558-efdc259cde1a",
	"https://images.unsplash.com/photo-1620891549027-942fdc95d3f5",
	"https://images.unsplash.com/photo-1621007947382-bb3c3994e3fb",
}

var blogImageURLs = []string{
	// Workspace and Office
	"https://images.unsplash.com/photo-1499750310107-5fef28a66643",
	"https://images.unsplash.com/photo-1542435503-956c469947f6",
	"https://images.unsplash.com/photo-1488190211105-8b0e65b80b4e",
	"https://images.unsplash.com/photo-1486312338219-ce68d2c6f44d",
	"https://images.unsplash.com/photo-1516321318423-f06f85e504b3",
	"https://images.unsplash.com/photo-1515378960530-7c0da6231fb1",
	"https://images.unsplash.com/photo-1434030216411-0b793f4b4173",
	"https://images.unsplash.com/photo-1455390582262-044cdead277a",
	"https://images.unsplash.com/photo-1519337265831-281ec6cc8514",
	"https://images.unsplash.com/photo-1512486130939-2c4f79935e4f",
	// Technology and Devices
	"https://images.unsplash.com/photo-1496171367470-9ed9a91ea931",
	"https://images.unsplash.com/photo-1498050108023-c5249f4df085",
	"https://images.unsplash.com/photo-1461749280684-dccba630e2f6",
	"https://images.unsplash.com/photo-1504639725590-34d0984388bd",
	"https://images.unsplash.com/photo-1517694712202-14dd9538aa97",
	// Writing and Creativity
	"https://images.unsplash.com/photo-1455390582262-044cdead277a",
	"https://images.unsplash.com/photo-1506252374453-ef5237291d83",
	"https://images.unsplash.com/photo-1510442650500-93217e634e4c",
	"https://images.unsplash.com/photo-1518932945647-7a1c969f8be2",
	"https://images.unsplash.com/photo-1513542789411-b6a5d4f31634",
	// Coffee and Lifestyle
	"https://images.unsplash.com/photo-1495474472287-4d71bcdd2085",
	"https://images.unsplash.com/photo-1509042239860-f550ce710b93",
	"https://images.unsplash.com/photo-1511537190424-bbbab87ac5eb",
	"https://images.unsplash.com/photo-1459755486867-b55449bb39ff",
	"https://images.unsplash.com/photo-1495474472287-4d71bcdd2085",
	// Modern Office Culture
	"https://images.unsplash.com/photo-1497215728101-856f4ea42174",
	"https://images.unsplash.com/photo-1497366216548-37526070297c",
	"https://images.unsplash.com/photo-1497366811353-6870744d04b2",
	"https://images.unsplash.com/photo-1505330622279-bf7d7fc918f4",
	"https://images.unsplash.com/photo-1516387938699-a93567ec168e",
}

// inMemoryFile implements multipart.File interface
type inMemoryFile struct {
	*bytes.Reader
}

func (f *inMemoryFile) Close() error {
	return nil
}

type FileFactory struct {
	client *http.Client
}

func NewFileFactory() *FileFactory {
	return &FileFactory{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// downloadAndResizeImage downloads an image from URL and resizes it to 1920x1080
func (f *FileFactory) downloadAndResizeImage(url string) (*bytes.Buffer, error) {
	// Add quality parameters to Unsplash URL
	url = url + "?q=85&w=1920&h=1080&fit=crop"

	// Download the image
	resp, err := f.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	// Read the image
	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read image data: %v", err)
	}

	// Decode the image
	img, err := jpeg.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	// Resize the image to exactly 1920x1080
	resized := resize.Resize(ImageWidth, ImageHeight, img, resize.Lanczos3)

	// Encode to JPEG
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, resized, &jpeg.Options{Quality: 90}); err != nil {
		return nil, fmt.Errorf("failed to encode JPEG: %v", err)
	}

	return buf, nil
}

func (f *FileFactory) createFileHeader(buf *bytes.Buffer, filename string) (*multipart.FileHeader, error) {
	// Create a file header
	fh := &multipart.FileHeader{
		Filename: filename,
		Header:   make(map[string][]string),
		Size:     int64(buf.Len()),
	}

	// Store the buffer in the file header's private field
	fh.Header.Set("X-Content", buf.String())

	return fh, nil
}

// CreateCarImage downloads and processes a random car image
func (f *FileFactory) CreateCarImage() (*multipart.FileHeader, error) {
	// Try up to 3 different images in case of errors
	var lastErr error
	for attempts := 0; attempts < 3; attempts++ {
		// Get random car image URL
		url := carImageURLs[rand.Intn(len(carImageURLs))]

		buf, err := f.downloadAndResizeImage(url)
		if err != nil {
			lastErr = err
			log.Printf("Failed to process car image from %s: %v, trying next image...", url, err)
			continue
		}

		filename := fmt.Sprintf("car_image_%d.jpg", time.Now().UnixNano())
		return f.createFileHeader(buf, filename)
	}
	return nil, fmt.Errorf("failed to create car image after 3 attempts, last error: %v", lastErr)
}

// CreateBlogImage downloads and processes a random blog image
func (f *FileFactory) CreateBlogImage() (*multipart.FileHeader, error) {
	// Try up to 3 different images in case of errors
	var lastErr error
	for attempts := 0; attempts < 3; attempts++ {
		// Get random blog image URL
		url := blogImageURLs[rand.Intn(len(blogImageURLs))]

		buf, err := f.downloadAndResizeImage(url)
		if err != nil {
			lastErr = err
			log.Printf("Failed to process blog image from %s: %v, trying next image...", url, err)
			continue
		}

		filename := fmt.Sprintf("blog_image_%d.jpg", time.Now().UnixNano())
		return f.createFileHeader(buf, filename)
	}
	return nil, fmt.Errorf("failed to create blog image after 3 attempts, last error: %v", lastErr)
}
