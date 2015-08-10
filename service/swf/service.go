// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package swf

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go/internal/signer/v4"
)

// The Amazon Simple Workflow Service (Amazon SWF) makes it easy to build applications
// that use Amazon's cloud to coordinate work across distributed components.
// In Amazon SWF, a task represents a logical unit of work that is performed
// by a component of your workflow. Coordinating tasks in a workflow involves
// managing intertask dependencies, scheduling, and concurrency in accordance
// with the logical flow of the application.
//
// Amazon SWF gives you full control over implementing tasks and coordinating
// them without worrying about underlying complexities such as tracking their
// progress and maintaining their state.
//
// This documentation serves as reference only. For a broader overview of the
// Amazon SWF programming model, see the Amazon SWF Developer Guide (http://docs.aws.amazon.com/amazonswf/latest/developerguide/).
type SWF struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*service.Request)

// New returns a new SWF client.
func New(config *aws.Config) *SWF {
	service := &service.Service{
		Config:       defaults.DefaultConfig.Merge(config),
		ServiceName:  "swf",
		APIVersion:   "2012-01-25",
		JSONVersion:  "1.0",
		TargetPrefix: "SimpleWorkflowService",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &SWF{service}
}

// newRequest creates a new request for a SWF operation and runs any
// custom request initialization.
func (c *SWF) newRequest(op *service.Operation, params, data interface{}) *service.Request {
	req := service.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
