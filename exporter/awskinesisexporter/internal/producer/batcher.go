// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package producer // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter/internal/producer"

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter/internal/batch"
)

type batcher struct {
	stream *string

	client Kinesis
	log    *zap.Logger
}

var (
	_ Batcher = (*batcher)(nil)
)

var (
	permanentErrResourceNotFound = new(*types.ResourceNotFoundException)
	permanentErrInvalidArgument  = new(*types.InvalidArgumentException)
)

func NewBatcher(kinesisAPI Kinesis, stream string, opts ...BatcherOptions) (Batcher, error) {
	be := &batcher{
		stream: aws.String(stream),
		client: kinesisAPI,
		log:    zap.NewNop(),
	}
	for _, opt := range opts {
		if err := opt(be); err != nil {
			return nil, err
		}
	}
	return be, nil
}

func (b *batcher) Put(ctx context.Context, bt *batch.Batch) error {
	for _, records := range bt.Chunk() {
		out, err := b.client.PutRecords(ctx, &kinesis.PutRecordsInput{
			StreamName: b.stream,
			Records:    records,
		})

		if err != nil {
			if errors.As(err, permanentErrResourceNotFound) || errors.As(err, permanentErrInvalidArgument) {
				err = consumererror.NewPermanent(err)
			}
			fields := []zap.Field{
				zap.Error(err),
			}
			if out != nil {
				fields = append(fields, zap.Int32p("failed-records", out.FailedRecordCount))
			}
			b.log.Error("Failed to write records to kinesis", fields...)
			return err
		}

		b.log.Debug("Successfully wrote batch to kinesis", zap.Stringp("stream", b.stream))
	}
	return nil
}

func (b *batcher) Ready(ctx context.Context) error {
	_, err := b.client.DescribeStream(ctx, &kinesis.DescribeStreamInput{
		StreamName: b.stream,
	})
	return err
}
