package memory

import (
	"context"
	"errors"

	"github.com/hashicorp/go-memdb"

	"github.com/Permify/permify/internal/repositories"
	db "github.com/Permify/permify/pkg/database/memory"
	"github.com/Permify/permify/pkg/dsl/compiler"
	"github.com/Permify/permify/pkg/dsl/schema"
	base "github.com/Permify/permify/pkg/pb/base/v1"
)

type SchemaReader struct {
	database *db.Memory
}

// NewSchemaReader creates a new SchemaReader
func NewSchemaReader(database *db.Memory) *SchemaReader {
	return &SchemaReader{
		database: database,
	}
}

// ReadSchema -
func (r *SchemaReader) ReadSchema(ctx context.Context, version string) (schema *base.IndexedSchema, err error) {
	if version == "" {
		version, err = r.findLastVersion(ctx)
		if err != nil {
			return schema, err
		}
	}

	txn := r.database.DB.Txn(false)
	defer txn.Abort()
	var it memdb.ResultIterator
	it, err = txn.Get("entity_config", "version", version)
	if err != nil {
		return schema, errors.New(base.ErrorCode_ERROR_CODE_EXECUTION.String())
	}

	var definitions []string
	for obj := it.Next(); obj != nil; obj = it.Next() {
		definitions = append(definitions, obj.(repositories.SchemaDefinition).Serialized())
	}

	schema, err = compiler.NewSchema(definitions...)
	if err != nil {
		return nil, err
	}

	return schema, nil
}

// ReadSchemaDefinition -
func (r *SchemaReader) ReadSchemaDefinition(ctx context.Context, entityType string, version string) (definition *base.EntityDefinition, err error) {
	if version == "" {
		version, err = r.findLastVersion(ctx)
		if err != nil {
			return nil, err
		}
	}

	txn := r.database.DB.Txn(false)
	defer txn.Abort()
	var raw interface{}
	raw, err = txn.First(schemaDefinitionTable, "id", entityType, version)
	if err != nil {
		return nil, errors.New(base.ErrorCode_ERROR_CODE_EXECUTION.String())
	}

	def, ok := raw.(repositories.SchemaDefinition)
	if ok {
		var sch *base.IndexedSchema
		sch, err = compiler.NewSchemaWithoutReferenceValidation(def.Serialized())
		if err != nil {
			return nil, err
		}
		return schema.GetEntityByName(sch, entityType)
	}

	return nil, errors.New(base.ErrorCode_ERROR_CODE_SCHEMA_NOT_FOUND.String())
}

// findLastVersion -
func (r *SchemaReader) findLastVersion(ctx context.Context) (string, error) {
	var err error
	txn := r.database.DB.Txn(false)
	defer txn.Abort()
	var raw interface{}
	raw, err = txn.Last(schemaDefinitionTable, "version")
	if err != nil {
		return "", errors.New(base.ErrorCode_ERROR_CODE_EXECUTION.String())
	}
	if _, ok := raw.(repositories.SchemaDefinition); ok {
		return "", nil
	}
	return "", errors.New(base.ErrorCode_ERROR_CODE_SCHEMA_NOT_FOUND.String())
}
