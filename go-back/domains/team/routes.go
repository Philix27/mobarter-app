package team

import (
	"github.com/gofiber/fiber/v2"
)

func RoutesHandler(router fiber.Router) {
	r := router.Group("/")
	r.Post("/create_team", createTeam).Name("TeamCreate")
	r.Post("/get_teams", getAllTeams).Name("TeamGetAll")
	r.Post("/get_team_members", getTeamMembers).Name("TeamGetMembers")
	r.Post("/add_member", addMember).Name("TeamAdd")
	r.Post("/remove_member", removeMember).Name("TeamRemove")
	r.Post("/delete_team", deleteTeam).Name("TeamDelete")
	// return c.SendStatus(fiber.StatusOK)
}
func getAllTeams(c *fiber.Ctx) error {
	return c.SendString("getAllTeams")
}
func getTeamMembers(c *fiber.Ctx) error {
	return c.SendString("getTeamMembers")
}
func createTeam(c *fiber.Ctx) error {
	return c.SendString("createTeam")
}
func addMember(c *fiber.Ctx) error {
	return c.SendString("addMember")
}
func removeMember(c *fiber.Ctx) error {
	return c.SendString("removeMember")
}
func deleteTeam(c *fiber.Ctx) error {
	return c.SendString("deleteTeam")
}
