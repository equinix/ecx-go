package ecx

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	baseURL = "http://localhost:8888"
)

func TestClientImplementation(t *testing.T) {
	//given
	cli := NewClient(context.Background(), baseURL, &http.Client{})
	//then
	assert.Implements(t, (*Client)(nil), cli, "Rest client implements Client interface")
}

func readJSONData(filePath string, target interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, target); err != nil {
		return err
	}
	return nil
}
