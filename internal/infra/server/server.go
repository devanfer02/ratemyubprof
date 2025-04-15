package server

type Server interface {
	// Start starts the server and listens for incoming requests
	Start()	

	// GracefullyShutdown gracefully shuts down the server on interrupt
	GracefullyShutdown()

	// MountHandlers mounts all the HTTP handlers and routes
	MountHandlers()

	// MountWorkers mounts all the workers or background jobs
	MountWorkers()
}