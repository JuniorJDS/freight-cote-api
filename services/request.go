package services

import (
	"bytes"
	"fmt"
	"freight-cote-api/utils"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RequestsServices struct {
	baseURL       string
	errorsHandler utils.ErrorsHandler
}

func NewRequestsService(baseURL string) *RequestsServices {
	return &RequestsServices{
		baseURL:       baseURL,
		errorsHandler: *utils.NewErrorsHandler(),
	}
}

var (
	once   sync.Once
	Client HTTPClient
)

func init() {
	once.Do(func() {
		Client = &http.Client{
			Timeout: time.Second * 30,
		}
	})
}

func (r *RequestsServices) SendRequest(method string, body []byte) ([]byte, error) {
	requestBody := bytes.NewBuffer(body)

	req, err := http.NewRequest(method, r.baseURL, requestBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return respBody, nil
	case http.StatusBadRequest:
		return nil, fiber.NewError(resp.StatusCode, "Bad Request")
	case http.StatusUnauthorized:
		return nil, fiber.NewError(resp.StatusCode, "Unauthorized")
	case http.StatusForbidden:
		return nil, fiber.NewError(resp.StatusCode, "Forbidden")
	case http.StatusNotFound:
		return nil, fiber.NewError(resp.StatusCode, "Not Found")
	case http.StatusInternalServerError:
		return nil, fiber.NewError(resp.StatusCode, "Internal Server Error")
	default:
		return nil, fiber.NewError(resp.StatusCode, "Unexpected Status Code")
	}
}
