package controller

import (
	"net/http"

	"github.com/elaijuh/water-jug/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Problem struct {
	X int `json:"x" validate:"required,gt=0"`
	Y int `json:"y" validate:"required,gt=0"`
	Z int `json:"z" validate:"required,gt=0"`
}

type Answer struct {
	Resolvable bool     `json:"resolvable"`
	Actions    []string `json:"actions"`
}

type errorMessage struct {
	Message string `json:"message"`
}

type ProblemHandler struct {
}

func (h *ProblemHandler) ResolveProblem(c echo.Context) error {
	var problem Problem
	c.Echo().Validator = &ProblemValidator{validator: v}
	if err := c.Bind(&problem); err != nil {
		log.Errorf("Unable to bind : %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage{Message: "Unable to parse request payload"})
	}
	if err := c.Validate(problem); err != nil {
		log.Errorf("Invalid x/y/z %+v %v", problem, err)
		return c.JSON(http.StatusBadRequest, errorMessage{Message: "Invalid x/y/z"})
	}

	answer := &Answer{Actions: make([]string, 0)}
	s := service.NewProblemResolver(problem.X, problem.Y, problem.Z)
	if s.Resolvable() {
		answer.Resolvable = true
		answer.Actions = s.GetActions()
	}

	return c.JSON(http.StatusOK, answer)
}
