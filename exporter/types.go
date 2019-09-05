package exporter

type State int

type Update struct {
	State    string
	Instance string
	Name     string
}

const (
	running State = iota + 1
	created
	paused
	restarting
	removing
	exited
	dead
)

var states = map[string]State{
	"running":    running,
	"created":    created,
	"paused":     paused,
	"restarting": restarting,
	"removing":   removing,
	"exited":     exited,
	"dead":       dead,
}

func StateValue(s string) State {
	return states[s]
}
