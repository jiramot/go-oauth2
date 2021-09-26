package usecases

type OpenAppGalleryUseCase interface {
    OpenFormAppId(appId string) (string, error)
}
