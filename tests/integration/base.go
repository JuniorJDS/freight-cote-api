package integration_test

import (
	"fmt"
	"io"
	"net/http"
	"os"

	api "freight-cote-api/api/app"

	"github.com/stretchr/testify/suite"
)

type BaseTest struct {
	suite.Suite
}

func NewBaseTest() *BaseTest {
	return &BaseTest{}
}

func (b *BaseTest) appClient(
	verb string, pathEndpoint string, body io.Reader,
) (*http.Response, error) {
	baseURL := os.Getenv("APP_BASE_URL")
	url := baseURL + pathEndpoint

	req, err := http.NewRequest(verb, url, body)
	if err != nil {
		fmt.Println("Erro to make request: ", err.Error())
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	app := api.MakeApp()
	resp, err := app.Test(req, -1)

	return resp, err
}
