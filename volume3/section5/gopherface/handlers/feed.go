package handlers

import (
	"net/http"

	"github.com/ramanujadasu/gofullstack/volume3/section5/gopherface/common"
	"github.com/ramanujadasu/gofullstack/volume3/section5/gopherface/forms"

	"go.isomorphicgo.org/go/isokit"
)

func FeedHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		formParams := isokit.FormParams{ResponseWriter: w, Request: r}
		p := forms.NewSocialMediaPostForm(&formParams)
		p.PageTitle = "Feed"
		env.TemplateSet.Render("feed_page", &isokit.RenderParams{Writer: w, Data: p})
	})
}
