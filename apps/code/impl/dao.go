package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) delete(ins *code.Code) error {
	result, err := s.col.DeleteOne(context.TODO(), bson.M{"_id": ins.Id})
	if err != nil {
		return exception.NewInternalServerError("delete verify code(%s) error, %s", ins.Code, err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("delete verify code %s not found", ins.Code)
	}

	return nil
}
