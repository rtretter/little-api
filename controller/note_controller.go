package controllers

import (
	"little-api/model"
	"little-api/repo"
    "little-api/repo/impl"
	"little-api/usecase"

    "fmt"
	"github.com/labstack/echo"
)

type NoteController struct {
    Interactor usecase.NoteInteractor
}

func NewNoteController(d repo.Db) *NoteController {
    return &NoteController {
        Interactor: usecase.NoteInteractor {
            NoteRepository: &impl.NoteRepository {
                Db: d,
            },
        },
    }
}

func (controller *NoteController) Create(c echo.Context) error {
    u := model.Note{}
    c.Bind(&u)
    key := c.Request().Header.Get("Authorization")
    err := u.Encrypt(key)
    if err != nil {
        return c.NoContent(500)
    }
    controller.Interactor.Add(u)
    return c.NoContent(201)
}

func (controller *NoteController) Update(c echo.Context) error {
    key := c.Request().Header.Get("Authorization")
    u := model.Note{}
    c.Bind(&u)
    res := controller.Interactor.GetInfoById(fmt.Sprintf("%d", u.ID))
    if res == nil {
        return c.NoContent(404)
    } 
    if u.Notebook != res.Notebook {
        return c.NoContent(401)
    }
    err := res.Decrypt(key)
    if err != nil {
        return c.NoContent(401)
    }
    err = u.Encrypt(key)
    if err != nil {
        return c.NoContent(500)
    }
    controller.Interactor.Update(u)
    return c.NoContent(200)
}

func (controller *NoteController) GetNotes(c echo.Context, notebook string) [] model.Note {
    res := controller.Interactor.GetInfo()
    key := c.Request().Header.Get("Authorization")
    result := make([]model.Note, 0)
    for _, elm := range res {
        if elm.Notebook != notebook {
            continue
        }
        err := elm.Decrypt(key)
        if err == nil {
            result = append(result, elm)
        }
    }
    return result
}

func (controller * NoteController) Delete(id string) int {
    return controller.Interactor.Delete(id)
}
