package workspace

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

type iRoutes interface {
	manager(router fiber.Router)
	create(c *fiber.Ctx) error
	update(c *fiber.Ctx) error
	getAllByUserId(c *fiber.Ctx) error
	getOne(c *fiber.Ctx) error
	deleteOne(c *fiber.Ctx) error
}

var ModuleName = "WORKSPACE"

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
	route.Post("/", r.create).Name("WorkspaceCreate")
	route.Put("/", r.update).Name("WorkspaceUpdate")
	route.Delete("/:id", r.deleteOne).Name("WorkspaceDelete")
	route.Get("/:id", r.getOne).Name("WorkspaceGetOne")
	route.Get("/", r.getAllByUserId).Name("WorkspaceGetAll")
}

func (r *Routes) create(c *fiber.Ctx) error {

	input := &WorkspaceCreateInput{}

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

	var input = &WorkspaceUpdateInput{}

	if err := c.BodyParser(input); err != nil {
		r.logger.Error("Error passing body")
		return err
	}

	if err := r.repository.Update(*input); err != nil {
		return err
	}

	r.logger.Info("UPDATED_" + ModuleName)
	return c.JSON(WorkspaceUpdateResponse{
		Msg: "success",
	})

}

func (r *Routes) getAllByUserId(c *fiber.Ctx) error {

	var input = &WorkspaceGetAllInput{}

	if err := c.BodyParser(input); err != nil {
		r.logger.Error("Error passing body")
		return err
	}

	obj, err := r.repository.FindAllByUserId(*input)

	if err != nil {
		return err
	}

	r.logger.Info("GET_ALL_"+ModuleName, obj)
	return c.JSON(WorkspaceGetAllResponse{
		Data: obj,
	})
}

func (r *Routes) getOne(c *fiber.Ctx) error {
	id := c.Params("id")

	if intValue, err := strconv.Atoi(id); err != nil {
		r.logger.Error("Error passing id")
		return err

	} else {
		obj, err := r.repository.FindByUserId(intValue)
		if err != nil {
			return err
		}
		r.logger.Info("GET_" + ModuleName)
		return c.JSON(obj)
	}

}

func (r *Routes) deleteOne(c *fiber.Ctx) error {
	// input := &WorkspaceDeleteInput{}
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
	return c.JSON(WorkspaceDeleteResponse{
		Msg: "Deleted successfully",
	})
}
