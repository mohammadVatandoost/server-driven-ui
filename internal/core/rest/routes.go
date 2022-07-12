package rest

const (
	versionPrefix      = "/v1"
	pagePrefix     = "/page"
)

func (s *Server) Routes() {
	router := s.Engine
	groupV1 := router.Group(versionPrefix)

	// authentication
	pagesGroup := groupV1.Group(pagePrefix)
	pagesGroup.GET("/demo", s.GetDemoPage)
}

