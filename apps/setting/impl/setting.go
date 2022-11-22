package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/setting"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) GetSetting(ctx context.Context) (*setting.Setting, error) {
	conf := setting.NewDefaultSetting()
	if err := s.col.FindOne(ctx, bson.M{"_id": conf.Version}).Decode(conf); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("version: %s system config not found, please init first", conf.Version)
		}

		return nil, exception.NewInternalServerError("find system config %s error, %s", conf.Version, err)
	}

	return conf, nil
}

func (s *service) UpdateSetting(ctx context.Context, ins *setting.Setting) (*setting.Setting, error) {
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			_, err := s.col.UpdateOne(ctx, bson.M{"_id": ins.Version}, bson.M{"$set": ins})
			if err != nil {
				return nil, exception.NewInternalServerError("update config document error, %s", err)
			}
		} else {
			return nil, err
		}
	}

	return ins, nil
}
