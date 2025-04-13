// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: calculator.proto

package calculatorv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	calculatorv1 "github.com/Wuei-Rain/calculator/server/calculatorv1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// CalculatorServiceName is the fully-qualified name of the CalculatorService service.
	CalculatorServiceName = "calculator.v1.CalculatorService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CalculatorServiceCalculateProcedure is the fully-qualified name of the CalculatorService's
	// Calculate RPC.
	CalculatorServiceCalculateProcedure = "/calculator.v1.CalculatorService/Calculate"
)

// CalculatorServiceClient is a client for the calculator.v1.CalculatorService service.
type CalculatorServiceClient interface {
	Calculate(context.Context, *connect_go.Request[calculatorv1.CalculatorRequest]) (*connect_go.Response[calculatorv1.CalculatorResponse], error)
}

// NewCalculatorServiceClient constructs a client for the calculator.v1.CalculatorService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCalculatorServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CalculatorServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &calculatorServiceClient{
		calculate: connect_go.NewClient[calculatorv1.CalculatorRequest, calculatorv1.CalculatorResponse](
			httpClient,
			baseURL+CalculatorServiceCalculateProcedure,
			opts...,
		),
	}
}

// calculatorServiceClient implements CalculatorServiceClient.
type calculatorServiceClient struct {
	calculate *connect_go.Client[calculatorv1.CalculatorRequest, calculatorv1.CalculatorResponse]
}

// Calculate calls calculator.v1.CalculatorService.Calculate.
func (c *calculatorServiceClient) Calculate(ctx context.Context, req *connect_go.Request[calculatorv1.CalculatorRequest]) (*connect_go.Response[calculatorv1.CalculatorResponse], error) {
	return c.calculate.CallUnary(ctx, req)
}

// CalculatorServiceHandler is an implementation of the calculator.v1.CalculatorService service.
type CalculatorServiceHandler interface {
	Calculate(context.Context, *connect_go.Request[calculatorv1.CalculatorRequest]) (*connect_go.Response[calculatorv1.CalculatorResponse], error)
}

// NewCalculatorServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCalculatorServiceHandler(svc CalculatorServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	calculatorServiceCalculateHandler := connect_go.NewUnaryHandler(
		CalculatorServiceCalculateProcedure,
		svc.Calculate,
		opts...,
	)
	return "/calculator.v1.CalculatorService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CalculatorServiceCalculateProcedure:
			calculatorServiceCalculateHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCalculatorServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCalculatorServiceHandler struct{}

func (UnimplementedCalculatorServiceHandler) Calculate(context.Context, *connect_go.Request[calculatorv1.CalculatorRequest]) (*connect_go.Response[calculatorv1.CalculatorResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("calculator.v1.CalculatorService.Calculate is not implemented"))
}
