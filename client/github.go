package client

import (
	"encoding/json"
	"fmt"
	"github.com/xyclos/github-gpg-keys/model"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type GithubUser string

const (
	BaseURL string = "https://api.github.com"
	DefaultClientTimeout time.Duration = 30 * time.Second
)

type GithubClient struct {
	client *http.Client
	baseURL string
}

func NewGithubClient() *GithubClient {
	return &GithubClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		baseURL: BaseURL,
	}
}

func (gc *GithubClient) SetTimeout(d time.Duration) {
	gc.client.Timeout = d
}

func (gc *GithubClient) Fetch(u GithubUser, save bool) (model.GithubPGPKeys, error) {
	resp, err := gc.client.Get(gc.buildURL(u))
	if err != nil {
		return model.GithubPGPKeys{}, err
	}
	defer resp.Body.Close()
	var githubResp model.GithubPGPKeys
	if err := json.NewDecoder(resp.Body).Decode(&githubResp); err != nil {
		return model.GithubPGPKeys{}, err
	}

	if save {
		for _, key := range githubResp {
			if err := gc.SaveToDisk(key.KeyID, key.RawKey, "."); err != nil {
				fmt.Println("Failed to save key!")
			}
		}
	}

	return githubResp, nil
}

func (gc *GithubClient) SaveToDisk(keyId string, rawKey string, savePath string) error {
	absSavePath, _ := filepath.Abs(savePath)
	filePath := fmt.Sprintf("%s/%s.gpg", absSavePath, keyId)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, rawKey)
	if err != nil {
		return err
	}
	return nil
}

func (gc *GithubClient) buildURL(u GithubUser) string {
	return fmt.Sprintf("%s/users/%s/gpg_keys", gc.baseURL, u)
}
