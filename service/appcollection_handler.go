package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func appCollectionHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, "foo")
	}
}