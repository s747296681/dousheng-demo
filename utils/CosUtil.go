package utils

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	c    *cos.Client
	once sync.Once
)

const (
	bucktetUrl string = "https://dousheng-1307504639.cos.ap-beijing.myqcloud.com"
	SecretID   string = "找jj要"
	SecretKey  string = "找jj要"
)

func init() {
	GetCosClient()
}

func GetCosClient() *cos.Client {
	once.Do(func() {
		//change the cos url
		u, _ := url.Parse(bucktetUrl)
		b := &cos.BaseURL{BucketURL: u}
		client := cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				// 通过环境变量获取密钥
				// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
				SecretID: SecretID,
				// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
				SecretKey: SecretKey,
			},
		})
		c = client
	})
	return c
}

func Upload(key string, file multipart.File) error {
	// Case 3 上传 0 字节文件, 设置输入流长度为 0

	_, err := c.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetVideoUrl(key string) string {

	return c.Object.GetObjectURL(key).String()
}

func GenerateVideoKey(name string) string {
	split := strings.Split(name, ".")
	if split == nil || len(split) == 0 || len(split) == 1 {
		return name
	}
	return split[0] + strconv.Itoa(time.Now().Second()) + "." + split[1]
}
