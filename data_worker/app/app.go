package app

import (
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

var Log *Logger
var Client *mongo.Client


var TimerChan chan bool

type Logger struct {
	*zerolog.Logger
}
