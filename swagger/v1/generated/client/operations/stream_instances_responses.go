// Copyright 2016-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StreamInstancesReader is a Reader for the StreamInstances structure.
type StreamInstancesReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *StreamInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStreamInstancesOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewStreamInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStreamInstancesOK creates a StreamInstancesOK with default headers values
func NewStreamInstancesOK(writer io.Writer) *StreamInstancesOK {
	return &StreamInstancesOK{
		Payload: writer,
	}
}

/*StreamInstancesOK handles this case with default header values.

Stream instances - success
*/
type StreamInstancesOK struct {
	Payload io.Writer
}

func (o *StreamInstancesOK) Error() string {
	return fmt.Sprintf("[GET /stream/instances][%d] streamInstancesOK  %+v", 200, o.Payload)
}

func (o *StreamInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStreamInstancesInternalServerError creates a StreamInstancesInternalServerError with default headers values
func NewStreamInstancesInternalServerError() *StreamInstancesInternalServerError {
	return &StreamInstancesInternalServerError{}
}

/*StreamInstancesInternalServerError handles this case with default header values.

Stream instances - unexpected error
*/
type StreamInstancesInternalServerError struct {
	Payload string
}

func (o *StreamInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /stream/instances][%d] streamInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *StreamInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
