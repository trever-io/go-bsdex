package bsdex

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const AUTH_HEADER_FMT = "hmac username=\"%v\", algorithm=\"hmac-sha1\", headers=\"date\", signature=\"%v\""
const HMAC_SIGNATURE_FMT = "date: %v"

const BASE_URL = "https://api-public.bsdex.de/api/"

const DATE_HEADER = "Date"
const API_KEY_HEADER = "ApiKey"
const AUTH_HEADER = "Authorization"

type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error: Code(%d) %v", e.Code, e.Message)
}

type APIClient struct {
	key    string
	secret string
}

func NewClient(key string, secret string) *APIClient {
	return &APIClient{
		key:    key,
		secret: secret,
	}
}

func (a *APIClient) getAuthHeader(now time.Time) string {
	nowStr := now.Format(time.RFC1123)
	h := hmac.New(sha1.New, []byte(a.secret))
	h.Write([]byte(fmt.Sprintf(HMAC_SIGNATURE_FMT, nowStr)))
	sha := h.Sum(nil)

	sig := base64.StdEncoding.EncodeToString(sha)
	return fmt.Sprintf(AUTH_HEADER_FMT, a.key, sig)
}

func (a *APIClient) requestGET(endpoint string, query map[string]string) ([]byte, error) {
	return a.requestNoBody(endpoint, http.MethodGet, query)
}

func (a *APIClient) requestNoBody(endpoint string, method string, query map[string]string) ([]byte, error) {
	if a.key == "" || a.secret == "" {
		return nil, errors.New("missing credentials")
	}

	now := time.Now().UTC()
	authHeader := a.getAuthHeader(now)
	url := fmt.Sprintf("%v%v", BASE_URL, endpoint)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(DATE_HEADER, now.Format(time.RFC1123))
	req.Header.Add(API_KEY_HEADER, a.key)
	req.Header.Add(AUTH_HEADER, authHeader)

	if query != nil {
		q := req.URL.Query()
		for key, val := range query {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	return a.doRequest(req)
}

func (a *APIClient) requestPOST(endpoint string, request interface{}) ([]byte, error) {
	if a.key == "" || a.secret == "" {
		return nil, errors.New("missing credentials")
	}

	var body io.Reader
	if request != nil {
		b, err := json.Marshal(request)
		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(b)
	}

	now := time.Now().UTC()
	authHeader := a.getAuthHeader(now)
	url := fmt.Sprintf("%v%v", BASE_URL, endpoint)

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add(DATE_HEADER, now.Format(time.RFC1123))
	req.Header.Add(API_KEY_HEADER, a.key)
	req.Header.Add(AUTH_HEADER, authHeader)
	req.Header.Add("Content-Type", "application/json")

	return a.doRequest(req)
}

func (a *APIClient) doRequest(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, &APIError{
			Code:    resp.StatusCode,
			Message: string(b),
		}
	}

	return b, nil
}
