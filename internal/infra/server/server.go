package server

type Server interface {
	Start()	
	GracefullyShutdown()
	MountHandlers()
}