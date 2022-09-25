package handler

type AdminHandler struct {
	Log    *AdminLogHandler
	Search *AdminSearchHandler
	User   *AdminUserHandler
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{
		Log:    NewAdminLogHandler(),
		Search: NewAdminSearchHandler(),
		User:   NewAdminUserHandler(),
	}
}
