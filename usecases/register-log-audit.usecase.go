package usecases

import (
	"errors"
	"github/hallex-abreu/lambda-chapter-products/entities"
	"github/hallex-abreu/lambda-chapter-products/interfaces"
)

func RegisterLogAudit(logAudit entities.LogAuditEntity, noSqlInterface interfaces.NoSqlInterface) error {
	err := noSqlInterface.RegisterLogAudit(logAudit)

	if err != nil {
		return errors.New("Error in register log audit")
	}
	return nil
}
