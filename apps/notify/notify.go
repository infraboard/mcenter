package notify

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/infraboard/mcube/v2/pb/resource"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

func NewRecordSet() *RecordSet {
	return &RecordSet{
		Items: []*Record{},
	}
}

func (s *RecordSet) Add(items ...*Record) {
	s.Items = append(s.Items, items...)
}

func NewDefaultRecord() *Record {
	return NewRecord(NewSendMailRequest("", ""))
}

func (r *Record) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*resource.Meta
		*SendNotifyRequest
		Response []*SendResponse
	}{r.Meta, r.Request, r.Response})
}

func (r *Record) ToJson() string {
	return pretty.ToJSON(r)
}

func (r *Record) FailedResponse() (items []*SendResponse) {
	if r.Response == nil {
		return
	}

	for i := range r.Response {
		resp := r.Response[i]
		if !resp.Success {
			items = append(items, resp)
		}
	}

	return
}

func (r *Record) FailedResponseToMessage() string {
	resps := r.FailedResponse()
	if len(resps) == 0 {
		return ""
	}

	errors := []string{}
	for i := range resps {
		resp := resps[i]
		errors = append(errors, resp.DetailMessage())
	}

	return strings.Join(errors, ",")
}

func (r *SendResponse) DetailMessage() string {
	if r.Message == "" {
		return ""
	}
	return fmt.Sprintf("send message to user %s[%s] error, %s", r.User, r.Target, r.Message)
}
