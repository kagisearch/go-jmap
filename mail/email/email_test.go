package email

import (
	"encoding/json"
	"testing"
	"time"

	"git.sr.ht/~rockorager/go-jmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmailUnmarshalDates(t *testing.T) {
	valid := time.Date(2025, 12, 4, 15, 30, 7, 0, time.UTC)

	tests := []struct {
		name string
		json string
		want *time.Time
	}{
		{name: "valid", json: `"2025-12-04T15:30:07Z"`, want: &valid},
		{name: "five-digit year", json: `"32548-12-04T15:30:07Z"`, want: nil},
		{name: "not a date", json: `"yesterday"`, want: nil},
		{name: "wrong type", json: `1764862207`, want: nil},
		{name: "null", json: `null`, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := `{"id":"a","subject":"hello","receivedAt":` + tt.json +
				`,"sentAt":` + tt.json + `,"smimeVerifiedAt":` + tt.json + `}`
			var e Email
			require.NoError(t, json.Unmarshal([]byte(data), &e))
			assert.Equal(t, jmap.ID("a"), e.ID)
			assert.Equal(t, "hello", e.Subject)
			assert.Equal(t, tt.want, e.ReceivedAt)
			assert.Equal(t, tt.want, e.SentAt)
			assert.Equal(t, tt.want, e.SMIMEVerifiedAt)
		})
	}

	t.Run("absent", func(t *testing.T) {
		var e Email
		require.NoError(t, json.Unmarshal([]byte(`{"id":"a"}`), &e))
		assert.Nil(t, e.ReceivedAt)
		assert.Nil(t, e.SentAt)
		assert.Nil(t, e.SMIMEVerifiedAt)
	})
}

func TestEmailGetResponseWithUnparseableDate(t *testing.T) {
	data := `{
		"methodResponses": [["Email/get", {
			"accountId": "acct",
			"state": "s1",
			"list": [
				{"id": "bad", "receivedAt": "30579-12-06T15:30:08Z", "sentAt": "32548-12-04T15:30:07Z"},
				{"id": "good", "receivedAt": "2025-12-06T15:30:08Z", "sentAt": "2025-12-06T15:30:08Z"}
			]
		}, "0"]],
		"sessionState": "x"
	}`

	var resp jmap.Response
	require.NoError(t, json.Unmarshal([]byte(data), &resp))
	getResp, err := jmap.ResponseByCallID[*GetResponse](&resp, "0")
	require.NoError(t, err)
	require.Len(t, getResp.List, 2)

	assert.Equal(t, jmap.ID("bad"), getResp.List[0].ID)
	assert.Nil(t, getResp.List[0].ReceivedAt)
	assert.Nil(t, getResp.List[0].SentAt)

	want := time.Date(2025, 12, 6, 15, 30, 8, 0, time.UTC)
	assert.Equal(t, jmap.ID("good"), getResp.List[1].ID)
	assert.Equal(t, &want, getResp.List[1].ReceivedAt)
	assert.Equal(t, &want, getResp.List[1].SentAt)
}
