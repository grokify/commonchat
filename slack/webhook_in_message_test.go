package slack

import (
	"encoding/json"
	"testing"

	"github.com/grokify/mogo/encoding/jsonutil"
)

var MessageTests = []struct {
	v                     string
	attachmentCount       int
	attachment0FieldCount int
}{
	{
		`{"attachments":[{"color":"#D63232","fallback":"[Alerting] TEST","fields":[{"short":true,"title":"server $tag_name replication_lag","value":0},{"short":true,"title":"server $tag_name replication_lag","value":0},{"short":true,"title":"server $tag_name replication_lag","value":0},{"short":true,"title":"server $tag_name replication_lag","value":0},{"short":true,"title":"server $tag_name replication_lag","value":0},{"short":true,"title":"server $tag_name replication_lag","value":0}],"footer":"Grafana v8","footer_icon":"https://grafana.com/assets/img/fav32.png","text":"","title":"[Alerting] TEST","title_link":"https://grafana/","ts":1643390857}],"channel":""}`,
		1, 6,
	},
}

func TestMessageUnmarshal(t *testing.T) {
	for _, tt := range MessageTests {
		m := &Message{}
		err := json.Unmarshal([]byte(tt.v), m)
		if err != nil {
			t.Errorf("json.Unmarshal(%s): err (%s)", tt.v, err.Error())
			continue
		}
		if tt.attachmentCount != len(m.Attachments) {
			t.Errorf("json.Unmarshal(...) attachmentCount: mismatch want (%v), got (%v)", tt.attachmentCount, len(m.Attachments))
			continue
		}
		if tt.attachment0FieldCount != len(m.Attachments[0].Fields) {
			t.Errorf("json.Unmarshal(...) attachment0FieldCount: mismatch want (%v), got (%v)", tt.attachment0FieldCount, len(m.Attachments[0].Fields))
		}
	}
}

var FieldTests = []struct {
	v    string
	want jsonutil.String
	json string
}{
	{`{"value":"mystring"}`, "mystring", `{"value":"mystring"}`},
	{`{"value":1}`, "1", `{"value":"1"}`},
	{`{"value":false}`, "false", `{"value":"false"}`},
}

func TestFieldUnmarshal(t *testing.T) {
	for _, tt := range FieldTests {
		f := &Field{}
		err := json.Unmarshal([]byte(tt.v), f)
		if err != nil {
			t.Errorf("json.Unmarshal(%s): err (%s)", tt.v, err.Error())
			continue
		}
		if f.Value != tt.want {
			t.Errorf("json.Unmarshal(%s): mismatch want (%v), got (%v)", tt.v, tt.want, f.Value)
			continue
		}
		m, err := json.Marshal(f)
		if err != nil {
			panic(err) // should not happen
		}
		if string(m) != tt.json {
			t.Errorf("json.Marshal(%v): mismatch want (%v), got (%v)", f, tt.json, string(m))
		}
	}
}
