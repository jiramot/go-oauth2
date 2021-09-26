package services

import "github.com/jiramot/go-oauth2/internal/core/domains"

type appGalleryService struct {
    clients domains.Clients
}

func NewAppGalleryService(clients domains.Clients) *appGalleryService {
    return &appGalleryService{
        clients: clients,
    }
}

func (svc *appGalleryService) OpenFormAppId(appId string) (string, error) {
    client, err := svc.clients.FindClientByClientId(appId)
    if err != nil {
        return "", nil
    }
    return client.RedirectUrl, nil
}
