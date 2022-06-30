package router

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"todoapp-in-go/usecase"
)

type requestTask struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type responseTask struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// TaskHandler task hadlerのinterface
type TaskHandler interface {
	PostTask(echo echo.Context) error
	GetTasks(echo echo.Context) error
	GetTask(echo echo.Context) error
	PutTask(echo echo.Context) error
	DeleteTask(echo echo.Context) error
}

type taskHandler struct {
	taskUsecase usecase.TaskUsecase
}

// NewTaskHandler task handlerのインスタンスを生成する
func NewTaskHandler(taskUsecase usecase.TaskUsecase) TaskHandler {
	return &taskHandler{taskUsecase}
}

// GetTasks taskを全件取得する
func (h *taskHandler)GetTasks(c echo.Context) error {
	tasks, err := h.taskUsecase.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	responseTasks := make([]responseTask, 0)
	for _, task := range tasks {
		responseTasks = append(responseTasks, responseTask{
			ID:      task.ID,
			Title:   task.Title,
			Content: task.Content,
		})
	}

	return c.JSON(http.StatusOK, responseTasks)
}

// PostTask taskを生成してDBに保存する
func (h *taskHandler)PostTask(c echo.Context) error {
	var req requestTask
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdTask, err := h.taskUsecase.Create(req.Title, req.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := responseTask{
		ID:      createdTask.ID,
		Title:   createdTask.Title,
		Content: createdTask.Content,
	}

	return c.JSON(http.StatusCreated, res)
}

// GetTask taskをIDで検索する
func (h *taskHandler)GetTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	task, err := h.taskUsecase.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res := responseTask{
		ID:      task.ID,
		Title:   task.Title,
		Content: task.Content,
	}

	return c.JSON(http.StatusOK, res)
}

// PutTask taskを更新
func (h *taskHandler)PutTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var req requestTask
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	updatedTask, err := h.taskUsecase.Update(id, req.Title, req.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := responseTask{
		ID:      updatedTask.ID,
		Title:   updatedTask.Title,
		Content: updatedTask.Content,
	}

	return c.JSON(http.StatusOK, res)
}

// DeleteTask taskを削除
func (h *taskHandler)DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.taskUsecase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
