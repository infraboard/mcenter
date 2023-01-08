package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/scm"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) HandleGitlabEvent(r *restful.Request, w *restful.Response) {
	req := scm.NewDefaultWebHookEvent()
	set, err := h.scm.HandleGitlabEvent(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
