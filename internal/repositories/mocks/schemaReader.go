package mocks

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"

	base "github.com/Permify/permify/pkg/pb/base/v1"
)

// SchemaReader is an autogenerated mock type for the SchemaReader type
type SchemaReader struct {
	mock.Mock
}

// ReadSchema -
func (_m *SchemaReader) ReadSchema(ctx context.Context, version string) (schema *base.IndexedSchema, err error) {
	ret := _m.Called(version)

	var r0 *base.IndexedSchema
	if rf, ok := ret.Get(0).(func(context.Context, string) *base.IndexedSchema); ok {
		r0 = rf(ctx, version)
	} else {
		r0 = ret.Get(0).(*base.IndexedSchema)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, version)
	} else {
		if e, ok := ret.Get(1).(error); ok {
			r1 = e
		} else {
			r1 = nil
		}
	}

	return r0, r1
}

// ReadSchemaDefinition -
func (_m *SchemaReader) ReadSchemaDefinition(ctx context.Context, entityType string, version string) (definition *base.EntityDefinition, err error) {
	ret := _m.Called(entityType, version)

	var r0 *base.EntityDefinition
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *base.EntityDefinition); ok {
		r0 = rf(ctx, entityType, version)
	} else {
		r0 = ret.Get(0).(*base.EntityDefinition)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, entityType, version)
	} else {
		if e, ok := ret.Get(1).(error); ok {
			r1 = e
		} else {
			r1 = nil
		}
	}

	return r0, r1
}
