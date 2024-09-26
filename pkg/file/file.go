package file

import (
	"context"
	"io"
)

// Downloader 提供下载文件功能
type Downloader interface {
	// Download ctx 提供超时控制，filename 为源文件路径
	Download(ctx context.Context, filename string) (io.Reader, error)
}

// Uploader 提供上传/覆盖文件功能
type Uploader interface {
	// Upload ctx 提供超时控制，filename 为文件上传路径，reader为上传的文件内容
	Upload(ctx context.Context, filename string, reader io.ReadCloser) error
}

// File 文件具备的两种功能
type File interface {
	Downloader
	Uploader
}
