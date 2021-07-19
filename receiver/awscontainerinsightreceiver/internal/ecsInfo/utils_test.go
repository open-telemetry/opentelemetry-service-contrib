// Copyright  OpenTelemetry Authors
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

package ecsinfo

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContainerInstanceIDFromArn(t *testing.T) {

	oldFormatARN := "arn:aws:ecs:region:aws_account_id:task/task-id"
	result, _ := GetContainerInstanceIDFromArn(oldFormatARN)
	assert.Equal(t, "task-id", result, "Expected to be equal")

	newFormatARN := "arn:aws:ecs:region:aws_account_id:task/cluster-name/task-id"
	result, _ = GetContainerInstanceIDFromArn(newFormatARN)
	assert.Equal(t, "task-id", result, "Expected to be equal")

	wrongFormatARN := "arn:aws:ecs:region:aws_account_id:task"
	_, err := GetContainerInstanceIDFromArn(wrongFormatARN)
	assert.NotNil(t, err)
}

func TestIsClosed(t *testing.T) {

	channel := make(chan bool)

	assert.Equal(t, false, IsClosed(channel))

	close(channel)

	assert.Equal(t, true, IsClosed(channel))

}

type fakeClient struct {
	response *http.Response
	err      error
}

func (f *fakeClient) Do(req *http.Request) (*http.Response, error) {
	return f.response, f.err
}

func TestRequestSuccessWithKnownLength(t *testing.T) {

	respBody := "body"
	response := &http.Response{
		StatusCode:    200,
		Body:          ioutil.NopCloser(bytes.NewBufferString(respBody)),
		Header:        make(http.Header),
		ContentLength: 5 * 1024,
	}

	fakeClient := &fakeClient{
		response: response,
		err:      nil,
	}

	ctx := context.Background()

	body, err := request(ctx, "0.0.0.0", fakeClient)

	assert.Nil(t, err)

	assert.NotNil(t, body)

}

func TestRequestSuccessWithUnknownLength(t *testing.T) {

	respBody := "body"
	response := &http.Response{
		StatusCode:    200,
		Body:          ioutil.NopCloser(bytes.NewBufferString(respBody)),
		Header:        make(http.Header),
		ContentLength: -1,
	}

	fakeClient := &fakeClient{
		response: response,
		err:      nil,
	}

	ctx := context.Background()

	body, err := request(ctx, "0.0.0.0", fakeClient)

	assert.Nil(t, err)

	assert.NotNil(t, body)

}

func TestRequestWithFailedStatus(t *testing.T) {

	respBody := "body"
	response := &http.Response{
		Status:        "Bad Request",
		StatusCode:    400,
		Body:          ioutil.NopCloser(bytes.NewBufferString(respBody)),
		Header:        make(http.Header),
		ContentLength: 5 * 1024,
	}

	fakeClient := &fakeClient{
		response: response,
		err:      errors.New(""),
	}

	ctx := context.Background()

	body, err := request(ctx, "0.0.0.0", fakeClient)

	assert.Nil(t, body)

	assert.NotNil(t, err)

}

func TestRequestWithLargeContentLength(t *testing.T) {

	respBody := "body"
	response := &http.Response{
		StatusCode:    200,
		Body:          ioutil.NopCloser(bytes.NewBufferString(respBody)),
		Header:        make(http.Header),
		ContentLength: 5 * 1024 * 1024,
	}

	fakeClient := &fakeClient{
		response: response,
		err:      nil,
	}

	ctx := context.Background()

	body, err := request(ctx, "0.0.0.0", fakeClient)

	assert.Nil(t, body)

	assert.NotNil(t, err)

}
