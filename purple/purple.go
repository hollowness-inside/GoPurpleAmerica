package purple

import "archive/zip"

type Purple struct {
	counties []string
	year     int
	colors   [3][3]int

	dataArchive    *zip.ReadCloser
	regionsArchive *zip.ReadCloser
}
