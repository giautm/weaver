// go:build !ignoreWeaverGen

package productcatalogservice

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.ReportVersion(0, 9, 0)

	codegen.Register(codegen.Registration{
		Name:        "github.com/ServiceWeaver/weaver/examples/onlineboutique/productcatalogservice/T",
		Iface:       reflect.TypeOf((*T)(nil)).Elem(),
		Impl:        reflect.TypeOf(impl{}),
		LocalStubFn: func(impl any, tracer trace.Tracer) any { return t_local_stub{impl: impl.(T), tracer: tracer} },
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return t_client_stub{stub: stub, getProductMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/productcatalogservice/T", Method: "GetProduct"}), listProductsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/productcatalogservice/T", Method: "ListProducts"}), searchProductsMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/examples/onlineboutique/productcatalogservice/T", Method: "SearchProducts"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return t_server_stub{impl: impl.(T), addLoad: addLoad}
		},
		RefData: "",
	})
}

// Local stub implementations.

type t_local_stub struct {
	impl   T
	tracer trace.Tracer
}

func (s t_local_stub) GetProduct(ctx context.Context, a0 string) (r0 Product, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "productcatalogservice.T.GetProduct", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetProduct(ctx, a0)
}

func (s t_local_stub) ListProducts(ctx context.Context) (r0 []Product, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "productcatalogservice.T.ListProducts", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.ListProducts(ctx)
}

func (s t_local_stub) SearchProducts(ctx context.Context, a0 string) (r0 []Product, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "productcatalogservice.T.SearchProducts", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.SearchProducts(ctx, a0)
}

// Client stub implementations.

type t_client_stub struct {
	stub                  codegen.Stub
	getProductMetrics     *codegen.MethodMetrics
	listProductsMetrics   *codegen.MethodMetrics
	searchProductsMetrics *codegen.MethodMetrics
}

func (s t_client_stub) GetProduct(ctx context.Context, a0 string) (r0 Product, err error) {
	// Update metrics.
	start := time.Now()
	s.getProductMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "productcatalogservice.T.GetProduct", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.getProductMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.getProductMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.getProductMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.getProductMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

func (s t_client_stub) ListProducts(ctx context.Context) (r0 []Product, err error) {
	// Update metrics.
	start := time.Now()
	s.listProductsMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "productcatalogservice.T.ListProducts", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.listProductsMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.listProductsMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	var shardKey uint64

	// Call the remote method.
	s.listProductsMetrics.BytesRequest.Put(0)
	var results []byte
	results, err = s.stub.Run(ctx, 1, nil, shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.listProductsMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_Product_3e9d9e07(dec)
	err = dec.Error()
	return
}

func (s t_client_stub) SearchProducts(ctx context.Context, a0 string) (r0 []Product, err error) {
	// Update metrics.
	start := time.Now()
	s.searchProductsMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "productcatalogservice.T.SearchProducts", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.searchProductsMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.searchProductsMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	s.searchProductsMetrics.BytesRequest.Put(float64(len(enc.Data())))
	var results []byte
	results, err = s.stub.Run(ctx, 2, enc.Data(), shardKey)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}
	s.searchProductsMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_Product_3e9d9e07(dec)
	err = dec.Error()
	return
}

// Server stub implementations.

type t_server_stub struct {
	impl    T
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s t_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GetProduct":
		return s.getProduct
	case "ListProducts":
		return s.listProducts
	case "SearchProducts":
		return s.searchProducts
	default:
		return nil
	}
}

func (s t_server_stub) getProduct(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetProduct(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) listProducts(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.ListProducts(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_Product_3e9d9e07(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s t_server_stub) searchProducts(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.SearchProducts(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_Product_3e9d9e07(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = &Product{}

func (x *Product) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Product.WeaverMarshal: nil receiver"))
	}
	enc.String(x.ID)
	enc.String(x.Name)
	enc.String(x.Description)
	enc.String(x.Picture)
	(x.PriceUSD).WeaverMarshal(enc)
	serviceweaver_enc_slice_string_4af10117(enc, x.Categories)
}

func (x *Product) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Product.WeaverUnmarshal: nil receiver"))
	}
	x.ID = dec.String()
	x.Name = dec.String()
	x.Description = dec.String()
	x.Picture = dec.String()
	(&x.PriceUSD).WeaverUnmarshal(dec)
	x.Categories = serviceweaver_dec_slice_string_4af10117(dec)
}

func serviceweaver_enc_slice_string_4af10117(enc *codegen.Encoder, arg []string) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.String(arg[i])
	}
}

func serviceweaver_dec_slice_string_4af10117(dec *codegen.Decoder) []string {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = dec.String()
	}
	return res
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_Product_3e9d9e07(enc *codegen.Encoder, arg []Product) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		(arg[i]).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_slice_Product_3e9d9e07(dec *codegen.Decoder) []Product {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]Product, n)
	for i := 0; i < n; i++ {
		(&res[i]).WeaverUnmarshal(dec)
	}
	return res
}
