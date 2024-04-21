package impl

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// 参考: https://www.mongodb.com/docs/manual/changeStreams/
// 参考: https://www.cnblogs.com/chnmig/p/16744242.html
func (i *impl) Watch(ctx context.Context) {
	i.log.Debug().Msg("start watch service instance change")

	stream, err := i.col.Watch(ctx, mongo.Pipeline{})
	if err != nil {
		i.log.Error().Msgf("watch service instance error, %s", err)
		return
	}

	for stream.Next(ctx) {
		change := stream.Current
		fmt.Println(string(change))
	}
}
