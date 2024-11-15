package waitlist

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

type iRoutes interface {
	manager(router fiber.Router)
	create(c *fiber.Ctx) error
	getAll(c *fiber.Ctx) error
	getOne(c *fiber.Ctx) error
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
	route.Post("/", r.create).Name("WaitlistCreate")
	route.Get("/:id", r.getOne).Name("WaitlistGetOne")
	route.Get("/", r.getAll).Name("WaitlistGetAll")
}

func (r *Routes) create(c *fiber.Ctx) error {

	input := &CreateInput{}

	if err := c.BodyParser(input); err != nil {
		r.logger.Error("Error passing body")
		return err
	}

	if err := r.repository.Create(CreateInput{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}); err != nil {
		return c.SendString("Error in request")
	}

	r.logger.Info("CREATED_" + ModuleName)
	return c.SendString("Created successfully")
}

func (r *Routes) getAll(c *fiber.Ctx) error {

	// var input = &GetAllInput{}

	// if err := c.BodyParser(*input); err != nil {
	// 	r.logger.Error("Error passing body")
	// 	return err
	// }

	obj, err := r.repository.FindAll()

	if err != nil {
		return err
	}

	r.logger.Info("GET_ALL_"+ModuleName, obj)
	return c.JSON(GetAllResponse{
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
