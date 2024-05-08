package repo

import "little-api/model"

type NoteRepository interface {
    Store(model.Note)
    Select() []model.Note
    SelectById(id string) *model.Note
    Modify(model.Note)
    Delete(id string) int
}
