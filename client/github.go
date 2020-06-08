package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/xyclos/github-gpg-keys/model"
)

type GithubUser string

const (
	BaseURL              string        = "https://api.github.com"
	DefaultClientTimeout time.Duration = 30 * time.Second
)

type GithubClient struct {
	client  *http.Client
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

func (gc *GithubClient) Fetch(user GithubUser, email *string) (model.GithubPGPKeys, error) {
	resp, err := gc.client.Get(gc.buildURL(user))
	if err != nil {
		return model.GithubPGPKeys{}, err
	}
	defer resp.Body.Close()
	var githubResp model.GithubPGPKeys
	if err := json.NewDecoder(resp.Body).Decode(&githubResp); err != nil {
		return model.GithubPGPKeys{}, err
	}
	return githubResp, nil
}

func (gc *GithubClient) SaveToDisk(keyId string, rawKey string, savePath string) (*string, error) {
	absSavePath, _ := filepath.Abs(savePath)
	filePath := fmt.Sprintf("%s/%s.gpg", absSavePath, keyId)

	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = io.WriteString(file, rawKey)
	if err != nil {
		return nil, err
	}
	return &filePath, nil
}

func (gc *GithubClient) buildURL(u GithubUser) string {
	return fmt.Sprintf("%s/users/%s/gpg_keys", gc.baseURL, u)
}
