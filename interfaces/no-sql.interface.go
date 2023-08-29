package interfaces

import "github/hallex-abreu/lambda-chapter-products/entities"

type NoSqlInterface interface {
	RegisterLogAudit(logAudit entities.LogAuditEntity) *error
}
