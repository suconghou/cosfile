package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
	"os"
	"net/http"
)

const (
	defaultSignExpireTime = 3600 * 24
)

type COS struct {
	AppID     string
	SecretID  string
	SecretKey string
}

func NewCosClient() *COS {
	return &COS{
		AppID:     Config.AppID,
		SecretID:  Config.SecretID,
		SecretKey: Config.SecretKey,
	}
}

func (c *COS) appSign(bucket, fileID string, expired uint64) string {
	now := time.Now()
	rdm := rand.New(rand.NewSource(now.UnixNano())).Intn(999999999)
	plainText := []byte(fmt.Sprintf("a=%s&k=%s&e=%d&t=%d&r=%d&f=%s&b=%s", c.AppID, c.SecretID, expired, now.Unix(), rdm, fileID, bucket))
	mac := hmac.New(sha1.New, []byte(c.SecretKey))
	mac.Write(plainText)
	signature := base64.StdEncoding.EncodeToString(append(mac.Sum(nil), plainText...))
	return signature
}

func (c *COS) SignOnce(bucket, fileID string) string {
	return c.appSign(bucket, fileID, 0)
}

func (c *COS) SignMore(bucket string, expired uint64) string {
	return c.appSign(bucket, "", expired)
}

func (c *COS) UploadFile(bucket, filePath, localFileName string) error {
	fileHandle, err := os.Open(localFileName)
	if err != nil {
		return err
	}
	defer fileHandle.Close()
	return nil

}

func (c *COS) DownloadFile(bucket, filePath string) error {
	return nil
}


func doHttpRequest(method, url, sign, contentType string, content []byte)(err error,) {
	req, err := http.NewRequest(method, url, bytes.NewReader(content))
	if err != nil {
		return fmt.Errorf("create request error: %v", err), nil
	}
	req.Header.Add("Authorization", sign)
	req.Header.Add("Content-Type", contentType)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("send request error: %v", err), nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response error: %v", err), nil
	}

	jsrResp, err = simplejson.NewJson(body)
	if err != nil {
		return fmt.Errorf("decode response error: %v, Body: %s", err, body), nil
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP StatusCode: %v, Body: %s", resp.StatusCode, body), jsrResp
	}

	return nil, jsrResp
}
