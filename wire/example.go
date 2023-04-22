package wire

// repo

// IPostRepo IPostRepo
type IPostRepo interface{}

// NewPostRepo NewPostRepo
func NewPostRepo() IPostRepo {
	return new(IPostRepo)
}

// usecase

// IPostUsecase IPostUsecase
type IPostUsecase interface{}
type postUsecase struct {
	repo IPostRepo
}

// NewPostUsecase NewPostUsecase
func NewPostUsecase(repo IPostRepo) IPostUsecase {
	return postUsecase{repo: repo}
}

// service service

// PostService PostService
type PostService struct {
	usecase IPostUsecase
}

// NewPostService NewPostService
func NewPostService(u IPostUsecase) *PostService {
	return &PostService{usecase: u}
}
