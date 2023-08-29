package usecases

import (
	"errors"
	"gitlab/assai-clientes/lambda-chapter-products/entities"
	"gitlab/assai-clientes/lambda-chapter-products/interfaces"
)

func RegisterLogAudit(logAudit entities.LogAuditEntity, noSqlInterface interfaces.NoSqlInterface) error {
	err := noSqlInterface.RegisterLogAudit(logAudit)

	if err != nil {
		return errors.New("Error in register log audit")
	}
	return nil
}
