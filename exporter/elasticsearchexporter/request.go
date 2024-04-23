package elasticsearchexporter

import (
	"bytes"
	"context"
)

type Request struct {
	bulkIndexer *esBulkIndexerCurrent
	Items       []bulkIndexerItem
}

func newRequest(bulkIndexer *esBulkIndexerCurrent) *Request {
	return &Request{bulkIndexer: bulkIndexer}
}

func (r *Request) Export(ctx context.Context) error {
	batch := make([]esBulkIndexerItem, len(r.Items))
	for i, item := range r.Items {
		batch[i] = esBulkIndexerItem{
			Index: item.Index,
			Body:  bytes.NewReader(item.Body),
		}
	}
	return r.bulkIndexer.AddBatchAndFlush(ctx, batch)
}

func (r *Request) ItemsCount() int {
	return len(r.Items)
}

func (r *Request) Add(item bulkIndexerItem) {
	r.Items = append(r.Items, item)
}

type bulkIndexerItem struct {
	Index string
	Body  []byte
}
