package notify

import (
	"encoding/json"

	"github.com/infraboard/mcenter/common/format"
	"github.com/infraboard/mcenter/common/meta"
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
		*meta.Meta
		*SendNotifyRequest
		Response []*SendResponse
	}{r.Meta, r.Request, r.Response})
}

func (r *Record) ToJson() string {
	return format.Prettify(r)
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