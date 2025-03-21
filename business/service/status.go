package service

import (
	"log"

	"github.com/dfcfw/rock-migrate/datalayer/repository"
)

func NewStatus(source, target repository.Status, log *log.Logger) *Status {
	return &Status{
		source: source,
		target: target,
		log:    log,
	}
}

type Status struct {
	source repository.Status
	target repository.Status
	log    *log.Logger
}
