package handlers

import (
	"context"
	"gitlab.com/tokend/blobs/internal/data"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	blobsCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}


func CtxBlobsQ(q data.Blobs) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, blobsCtxKey, q)
	}
}

func BlobsQ(r *http.Request) data.Blobs {
	return r.Context().Value(blobsCtxKey).(data.Blobs).New()
}
