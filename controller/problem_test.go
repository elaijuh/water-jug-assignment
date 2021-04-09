package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var h ProblemHandler

func TestProblem(t *testing.T) {
	t.Run("test ResolveProblem", func(t *testing.T) {
		body := `
		{
			"x": 2,
			"y": 10,
			"z": 4
		}
		`
		req := httptest.NewRequest(http.MethodPost, "/problem", strings.NewReader(body))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		c := e.NewContext(req, res)
		err := h.ResolveProblem(c)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.Code)

		var answer Answer
		err = json.Unmarshal(res.Body.Bytes(), &answer)
		assert.Nil(t, err)
		t.Logf("answer: %+v\n", answer)
		assert.True(t, answer.Resolvable)
		assert.Equal(t, 4, len(answer.Actions))
	})
}
