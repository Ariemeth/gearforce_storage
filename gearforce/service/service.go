package service

type GearForceService interface {
	Hello(s string) (string, error)
}

type ServiceMiddleware func(GearForceService) GearForceService
