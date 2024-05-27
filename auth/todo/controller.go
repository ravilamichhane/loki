package todo

import (
	"auth/todo/dtos"
	"loki/common"
	"loki/thor"
)

type TodoController struct {
	TodoService TodoService
}

func NewTodoController(TodoService TodoService) *TodoController {
	return &TodoController{
		TodoService: TodoService,
	}
}

func (u *TodoController) Routes() []common.Route {
	return []common.Route{
		common.GET("/", u.FindAll),
		common.GET("/{id}", u.FindOne),
		common.POST("", u.Create),
		common.PUT("/{id}", u.Update),
		common.DELETE("/{id}", u.Delete),
	}
}

func (u *TodoController) FindAll(ctx common.HttpContext) error {
	ctx.SetStatusCode(200)

	todos, err := u.TodoService.FindAll()

	if err != nil {
		return err
	}

	ctx.JSON(200, todos)
	return nil
}

func (u *TodoController) FindOne(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}

	todo, err := u.TodoService.FindOne(uid)

	if err != nil {
		return err
	}

	ctx.SetStatusCode(200)

	ctx.JSON(200, todo)
	return nil
}

func (u *TodoController) Create(ctx common.HttpContext) error {
	var todo dtos.CreateTodo

	if err := ctx.Decode(&todo); err != nil {

		return err
	}

	newTodo := todo.ToTodo()

	if err := u.TodoService.Create(newTodo); err != nil {
		return err
	}

	ctx.JSON(200, newTodo)
	return nil

}

func (u *TodoController) Update(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}

	todo, err := u.TodoService.FindOne(uid)

	if err != nil {
		return err
	}

	var updateTodo dtos.UpdateTodo

	if err := ctx.Decode(&updateTodo); err != nil {
		return err
	}
	updateTodo.Decode(todo)

	if err = u.TodoService.Update(todo); err != nil {
		return err
	}

	ctx.JSON(200, todo)
	return nil
}

func (u *TodoController) Delete(ctx common.HttpContext) error {
	uid, err := ctx.GetUUIDParam("id")

	if err != nil {
		return err
	}
	err = u.TodoService.Delete(uid)

	if err == nil {
		ctx.JSON(200, struct {
			Message string `json:"message"`
		}{
			Message: "Todo Deleted",
		})
	}

	return err

}

func (u *TodoController) Prefix() string {
	return "/api/v1/todo"
}

func (u *TodoController) Middlewares() []common.MiddleWare {
	return []common.MiddleWare{
		thor.LoggingMiddleware,
		thor.ErrorMiddleware,
		thor.PanicMiddleWare,
	}
}
