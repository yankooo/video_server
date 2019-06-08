package ossops

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"yankooo/config"
)

var (
	AK string
	SK string
	EP string
)

func init() {
	AK = "FckdpsCxW8dtsIgmzvA0MDZdF8BQ0U"
	SK = ""
	EP = config.GetOssAddr()
}
func UploadToOss(filename string, path string, bn string) bool {

	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss client with error : %s", err)
		return false
	}

	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Fatalf("Get bucket with error : %s", err)
		return false
	}

	err = bucket.UploadFile(filename, path, 500*1024, oss.Routines(3))
	if err != nil {
		log.Printf("Uploading object error: %s", err)
		return false
	}
	return true
}

func DeleteObject(filename string, bn string) bool {
	client, err := oss.New(EP, AK, SK)
	if err != nil {
		log.Printf("Init oss client with error : %s", err)
		return false
	}

	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Fatalf("Get bucket with error : %s", err)
		return false
	}

	if err := bucket.DeleteObject(filename); err != nil {
		log.Printf("Deleting Object error:%v", err)
		return false
	}
	return true
}
