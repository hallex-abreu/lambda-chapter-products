package interfaces

import "gitlab/assai-clientes/lambda-chapter-products/entities"

type NoSqlInterface interface {
	RegisterLogAudit(logAudit entities.LogAuditEntity) *error
}
