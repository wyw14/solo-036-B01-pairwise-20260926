package service

import (
	"github.com/wyw14/solo-036/internal/model"
)

func (r *Runtime) Equipment() []model.Equipment {
	return r.EquipmentStatus()
}
func (r *Runtime) Interlocks() []model.Interlock {
	return r.InterlockStatus()
}
func (r *Runtime) Incidents() []model.Incident {
	return r.Snapshot().Incidents
}
