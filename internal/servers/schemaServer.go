package servers

import (
	"fmt"

	"google.golang.org/grpc/status"

	otelCodes "go.opentelemetry.io/otel/codes"
	"golang.org/x/net/context"

	"github.com/Permify/permify/internal/services"
	"github.com/Permify/permify/pkg/logger"
	v1 "github.com/Permify/permify/pkg/pb/base/v1"
)

// SchemaServer -
type SchemaServer struct {
	v1.UnimplementedSchemaServer

	schemaService services.ISchemaService
	l             logger.Interface
}

// NewSchemaServer -
func NewSchemaServer(s services.ISchemaService, l logger.Interface) *SchemaServer {
	return &SchemaServer{
		schemaService: s,
		l:             l,
	}
}

// Write -
func (r *SchemaServer) Write(ctx context.Context, request *v1.SchemaWriteRequest) (*v1.SchemaWriteResponse, error) {
	ctx, span := tracer.Start(ctx, "schemas.write")
	defer span.End()

	version, err := r.schemaService.WriteSchema(ctx, request.Schema)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		r.l.Error(fmt.Sprintf(err.Error()))
		return nil, status.Error(GetStatus(err), err.Error())
	}

	return &v1.SchemaWriteResponse{
		SchemaVersion: version,
	}, nil
}

// Write -
func (r *SchemaServer) Read(ctx context.Context, request *v1.SchemaReadRequest) (*v1.SchemaReadResponse, error) {
	ctx, span := tracer.Start(ctx, "schemas.write")
	defer span.End()

	var err error
	var response *v1.IndexedSchema
	response, err = r.schemaService.ReadSchema(ctx, request.SchemaVersion)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		r.l.Error(fmt.Sprintf(err.Error()))
		return nil, status.Error(GetStatus(err), err.Error())
	}

	return &v1.SchemaReadResponse{
		Schema: response,
	}, nil
}
