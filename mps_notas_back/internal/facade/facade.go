package facade

import (
	"mps_notas_back/internal/buisness/service"
	"mps_notas_back/internal/facade/handler"
	"net/http"
)
// Facade é função
type Facade interface {
	// GetFacade 
	GetFacade(userService *service.UserService, taskService *service.TaskService) Facade

	//Task Methods

	GetAllTasks(w http.ResponseWriter, r *http.Request)
	GetTaskByID(w http.ResponseWriter, r *http.Request)
	CreateTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request) 

	//User methods

	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GenerateReport(w http.ResponseWriter, r *http.Request)
}
type FacadeImpl struct {
	Facade
	task handler.TaskHandler
	user handler.UserHandler
}
func (FacadeImpl) GetFacade(userService *service.UserService, taskService *service.TaskService) Facade {
	return &FacadeImpl{
		task: *handler.NewTaskHandler(taskService),
		 user: *handler.NewUserHandler(userService),
		}
}


func (f FacadeImpl) GetAllTasks(w http.ResponseWriter, r *http.Request) { f.task.GetAllTasks(w,r )}

func (f FacadeImpl) GetTaskByID(w http.ResponseWriter, r *http.Request) { f.task.GetTaskByID(w,r )}

func (f FacadeImpl) CreateTask(w http.ResponseWriter, r *http.Request) { f.task.CreateTask(w,r )}

func (f FacadeImpl) UpdateTask(w http.ResponseWriter, r *http.Request) { f.task.UpdateTask(w,r )}

func (f FacadeImpl) DeleteTask(w http.ResponseWriter, r *http.Request) { f.task.DeleteTask(w,r )}



func (f FacadeImpl) GetAllUsers(w http.ResponseWriter, r *http.Request) {f.user.GetAllUsers(w,r) }

func (f FacadeImpl) GetUserByID(w http.ResponseWriter, r *http.Request) {f.user.GetUserByID(w,r)}

func (f FacadeImpl) CreateUser(w http.ResponseWriter, r *http.Request) {f.user.CreateUser(w,r)}


func (f FacadeImpl) Update(w http.ResponseWriter, r *http.Request) {f.user.Update(w,r)}

func (f FacadeImpl) Delete(w http.ResponseWriter, r *http.Request) {f.user.Delete(w,r)}

func (f FacadeImpl) Login(w http.ResponseWriter, r *http.Request) {f.user.Login(w,r)}

func (f FacadeImpl) GenerateReport(w http.ResponseWriter, r *http.Request) {f.user.GenerateReport(w,r)}