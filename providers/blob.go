package providers

import (
	"context"
	"fmt"
	"github.com/Azure/azure-storage-file-go/azfile"
	"github.com/gourds/upload2remote/config"
	"github.com/wonderivan/logger"
	"log"
	"net/url"
	"os"
)

func (c *BLOB) Auth() (SessionType, error) {
	client, err := azfile.NewSharedKeyCredential(c.AccessKeyID, c.AccessKeySecret)
	if err != nil {
		logger.Fatal(err)
	}
	return SessionType{blob: client}, nil
}

func (c *BLOB) UploadFile(objName string, filePath string, client SessionType, wg *config.Multi) (err error) {
	defer wg.WG.Done()
	file, err := os.Open(filePath)
	if err != nil {
		logger.Fatal(err)
	}
	defer file.Close()
	fileSize, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	u, _ := url.Parse(fmt.Sprintf("https://%s.file.core.windows.net/myshare/BigFile.bin", c.AccessKeyID))
	fileURL := azfile.NewFileURL(*u, azfile.NewPipeline(client.blob, azfile.PipelineOptions{}))
	ctx := context.Background()
	err = azfile.UploadFileToAzureFile(ctx, file, fileURL,
		azfile.UploadToAzureFileOptions{
			Parallelism: 3,
			FileHTTPHeaders: azfile.FileHTTPHeaders{
				CacheControl: "no-transform",
			},
			Metadata: azfile.Metadata{
				"createdby": "Jeffrey&Jiachen",
			},
			// If Progress is non-nil, this function is called periodically as bytes are uploaded.
			Progress: func(bytesTransferred int64) {
				fmt.Printf("Uploaded %d of %d bytes.\n", bytesTransferred, fileSize.Size())
			},
		})
	if err != nil {
		log.Fatal(err)
	}
	return
}
