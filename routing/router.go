package routing

import (
    controller "little-api/controller"
    "little-api/repo/impl"
    "net/http"

    "github.com/labstack/echo"
)

func Init(e *echo.Echo) {
    noteController := controller.NewNoteController(impl.CreateDB())

    e.GET("/notes/:notebook", func(ctx echo.Context) error {
        notebook := ctx.Param("notebook")
        notes := noteController.GetNotes(ctx, notebook)
        ctx.Bind(&notes)
        return ctx.JSON(http.StatusOK, notes)
    })

    e.POST("/notes", func(c echo.Context) error {
        return noteController.Create(c)
    })

    e.PUT("/notes", func(c echo.Context) error {
        return noteController.Update(c)
    })

    e.DELETE("/notes/:id", func(c echo.Context) error {
        id := c.Param("id")
        count := noteController.Delete(id)
        if count > 0 {
            return c.NoContent(http.StatusOK)
        } 
        return c.NoContent(http.StatusNoContent)
    })
}
