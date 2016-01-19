package net

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	gio "github.com/nzai/go-utility/io"
)

//	发送GET请求并返回字符串
func DownloadString(url string) (string, error) {
	return DownloadStringReferer(url, "")
}

//	发送GET请求并返回字符串(带Referer)
func DownloadStringReferer(url, referer string) (string, error) {
	return DownloadStringRefererRetry(url, referer, 1, 0)
}

//	访问网址并返回字符串
func DownloadStringRetry(url string, retryTimes, intervalSeconds int) (string, error) {
	return DownloadStringRefererRetry(url, "", retryTimes, intervalSeconds)
}

//	访问网址并返回字符串
func DownloadStringRefererRetry(url, referer string, retryTimes, intervalSeconds int) (string, error) {
	buffer, err := DownloadBufferRefererRetry(url, referer, retryTimes, intervalSeconds)

	return string(buffer), err
}

//	发送GET请求并返回缓冲区
func DownloadBuffer(url string) ([]byte, error) {
	return DownloadBufferReferer(url, "")
}

//	发送GET请求并返回缓冲区(带Referer)
func DownloadBufferReferer(url, referer string) ([]byte, error) {
	return DownloadBufferRefererRetry(url, referer, 1, 0)
}

//	访问网址并返回缓冲区
func DownloadBufferRetry(url string, retryTimes, intervalSeconds int) ([]byte, error) {
	return DownloadBufferRefererRetry(url, "", retryTimes, intervalSeconds)
}

//	访问网址并返回缓冲区
func DownloadBufferRefererRetry(url, referer string, retryTimes, intervalSeconds int) ([]byte, error) {
	var err error
	client := &http.Client{}

	for times := retryTimes - 1; times >= 0; times-- {
		//	构造请求
		request, err := http.NewRequest("GET", url, nil)
		if err == nil {
			//	引用页
			if referer != "" {
				request.Header.Set("Referer", referer)
			}

			//	发送请求
			response, err := client.Do(request)
			if err == nil {
				defer response.Body.Close()

				//	读取结果
				buffer, err := ioutil.ReadAll(response.Body)
				if err == nil {
					return buffer, nil
				}
			}
		}

		if times > 0 {
			if err != nil {
				log.Printf("访问%s出错，还有%d次重试机会，%d秒后重试:%s", url, times, intervalSeconds, err.Error())
			}

			//	延时
			time.Sleep(time.Duration(intervalSeconds) * time.Second)
		}
	}

	return nil, fmt.Errorf("访问%s出错，已重试%d次，不再重试:%s", url, retryTimes, err.Error())
}

//	下载文件
func DownloadFile(url, path string) error {
	return DownloadFileReferer(url, "", path)
}

//	下载文件(带Referer)
func DownloadFileReferer(url, referer, path string) error {
	return DownloadFileRefererRetry(url, referer, path, 1, 0)
}

//	下载文件
func DownloadFileRetry(url, path string, retryTimes, intervalSeconds int) error {
	return DownloadFileRefererRetry(url, "", path, retryTimes, intervalSeconds)
}

//	下载文件
func DownloadFileRefererRetry(url, referer, path string, retryTimes, intervalSeconds int) error {
	var err error
	client := &http.Client{}

	for times := retryTimes - 1; times >= 0; times-- {
		//	构造请求
		request, err := http.NewRequest("GET", url, nil)
		if err == nil {
			//	引用页
			if referer != "" {
				request.Header.Set("Referer", referer)
			}

			//	发送请求
			response, err := client.Do(request)
			if err == nil {
				defer response.Body.Close()

				//	打开文件
				file, err := gio.OpenForWrite(path)
				if err == nil {
					//	写文件
					_, err := io.Copy(file, response.Body)
					return err
				}

			}
		}

		if times > 0 {
			if err != nil {
				log.Printf("访问%s出错，还有%d次重试机会，%d秒后重试:%s", url, times, intervalSeconds, err.Error())
			}

			//	延时
			time.Sleep(time.Duration(intervalSeconds) * time.Second)
		}
	}

	return fmt.Errorf("访问%s出错，已重试%d次，不再重试:%s", url, retryTimes, err.Error())
}