//go:build wireinject
// +build wireinject

package wire

import "github.com/google/wire"

func GetPostService() *PostService {
	wire.Build(
		NewPostService,
		NewPostUsecase,
		NewPostRepo,
	)

	return nil
}
