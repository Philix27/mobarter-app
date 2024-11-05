package workspaceEpics

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

type iRoutes interface {
	manager(router fiber.Router)
	create(c *fiber.Ctx) error
	update(c *fiber.Ctx) error
	getAll(c *fiber.Ctx) error
	getOne(c *fiber.Ctx) error
	deleteOne(c *fiber.Ctx) error
}

type Routes struct {
	repository iRepository
	logger     *slog.Logger
}

func NewRoutes(repo iRepository, logger *slog.Logger, logGroupKey string) iRoutes {
	return &Routes{
		repository: repo,
		logger:     logger.WithGroup(logGroupKey),
	}
}

func (r *Routes) manager(route fiber.Router) {
	route.Post("/", r.create).Name("WorkspaceEpicCreate")
	route.Put("/", r.update).Name("WorkspaceEpicUpdate")
	route.Delete("/:id", r.deleteOne).Name("WorkspaceEpicDelete")
	route.Get("/:id", r.getOne).Name("WorkspaceEpicGetOne")
	route.Get("/", r.getAll).Name("WorkspaceEpicGetAll")
}

func (r *Routes) create(c *fiber.Ctx) error {
	input := &WorkspaceEpicCreateInput{}

	if err := c.BodyParser(input); err != nil {
		r.logger.Error("Error passing body")
		return err
	}

	if err := r.repository.Create(*input); err != nil {
		return c.SendString("Error in request")
	}

	r.logger.Info("CREATED_" + ModuleName)
	return c.SendString("Created successfully")
}

func (r *Routes) update(c *fiber.Ctx) error {
	var input = &WorkspaceEpicUpdateInput{}

	if err := c.BodyParser(input); err != nil {
		r.logger.Error("Error passing body")
		return err
	}

	if err := r.repository.Update(*input); err != nil {
		return err
	}

	r.logger.Info("UPDATED_" + ModuleName)
	return c.JSON(WorkspaceEpicUpdateResponse{
		Msg: "success",
	})
}

func (r *Routes) getAll(c *fiber.Ctx) error {
	// var input = &WorkspaceEpicGetAllInput{}
	workspaceId := c.Query("workspace_id")
	limit := c.Query("limit")

	workspaceIdValue, err := strconv.Atoi(workspaceId)
	limitValue, limitErr := strconv.Atoi(limit)

	if err != nil {
		r.logger.Error("Cannot parse workspaceId")
		return err
	}

	if limitErr != nil {
		r.logger.Error("Cannot parse limit")
		return err
	}

	obj, err := r.repository.FindWorkspaceId(workspaceIdValue, limitValue)

	if err != nil {
		return err
	}

	r.logger.Info("GET_ALL_" + ModuleName)

	return c.JSON(WorkspaceEpicGetAllResponse{
		Data: obj,
	})
}

func (r *Routes) getOne(c *fiber.Ctx) error {
	id := c.Params("id")

	if intValue, err := strconv.Atoi(id); err != nil {
		r.logger.Error("Error passing id")
		return err

	} else {
		obj, err := r.repository.FindById(intValue)
		if err != nil {
			return err
		}
		r.logger.Info("GET_" + ModuleName)
		return c.JSON(obj)
	}
}

func (r *Routes) deleteOne(c *fiber.Ctx) error {
	id := c.Params("id")

	intValue, err := strconv.Atoi(id)
	if err != nil {
		r.logger.Error("Error passing id")
		return err
	}

	if err := r.repository.Delete(intValue); err != nil {
		return err
	}

	r.logger.Info("DELETED_" + ModuleName)
	return c.JSON(WorkspaceEpicDeleteResponse{
		Msg: "Deleted successfully",
	})
}
