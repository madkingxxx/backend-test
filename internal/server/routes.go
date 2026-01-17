package server

func (s *HTTPServer) SetUpRoutes() {
	s.echo.Use(RequestLogger())
	api := s.echo.Group("/api")
	// User routes
	api.GET("/user/:id", s.transportManager.User.GetUser)
	api.POST("/user/:id/top-up", s.transportManager.User.TopUp)
	api.POST("/purchase", s.transportManager.User.Purchase)

	// Skinport routes
	api.GET("/items", s.transportManager.Skinport.GetAllItems)
}
