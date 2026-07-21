package jmap

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseUnmarshal(t *testing.T) {
	RegisterMethod("Test/method", newTest)
	assert := assert.New(t)
	data := []byte(`{"sessionState": "state","methodResponses":[["Test/method",{"Hello":"world"},"0"]]}`)

	resp := &Response{}
	err := json.Unmarshal(data, resp)
	assert.NoError(err)
	assert.Equal("state", resp.SessionState)
	assert.Equal(1, len(resp.Responses))

	inv := resp.Responses[0]
	assert.Equal("Test/method", inv.Name)
	assert.Equal(CallID("0"), inv.CallID)

	echo, ok := inv.Args.(*test)
	assert.Truef(ok, "invocation arguments are not type *Echo")
	assert.Equal("world", echo.Hello)
}

func TestResponseByCallID(t *testing.T) {
	desc := "account not found"
	resp := &Response{
		Responses: []*Invocation{
			{Name: "Test/method", Args: &test{Hello: "world"}, CallID: "0"},
			{Name: "error", Args: &MethodError{Type: "forbidden", Description: &desc}, CallID: "1"},
		},
	}

	t.Run("returns matching response", func(t *testing.T) {
		got, err := ResponseByCallID[*test](resp, "0")
		assert.NoError(t, err)
		assert.Equal(t, "world", got.Hello)
	})

	t.Run("returns method error under the call id", func(t *testing.T) {
		got, err := ResponseByCallID[*test](resp, "1")
		assert.Nil(t, got)
		var methodErr *MethodError
		assert.ErrorAs(t, err, &methodErr)
		assert.Equal(t, "forbidden", methodErr.Type)
		assert.EqualError(t, err, "forbidden: account not found")
	})

	t.Run("returns NotFoundError when call id absent", func(t *testing.T) {
		got, err := ResponseByCallID[*test](resp, "2")
		assert.Nil(t, got)
		var notFound NotFoundError[*test]
		assert.ErrorAs(t, err, &notFound)
	})
}

func TestResponseMarshal(t *testing.T) {
	assert := assert.New(t)
	resp := &Response{
		SessionState: "state",
		Responses: []*Invocation{
			{
				Name: "Test/method",
				Args: &struct {
					Hello string
				}{
					Hello: "world",
				},
				CallID: "0",
			},
		},
	}
	data, err := json.Marshal(resp)
	assert.NoError(err)
	expected := `{"methodResponses":[["Test/method",{"Hello":"world"},"0"]],"sessionState":"state"}`
	assert.Equal(expected, string(data))
}
