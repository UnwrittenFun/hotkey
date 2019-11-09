package hotkey

import (
	"errors"
	"time"
)

// Modifiers
const (
	ModAlt = 1 << iota
	ModCtrl
	ModShift
	ModWin
)

// Errors
var (
	ErrRegisterFailed = errors.New("Failed to register hotkey")
)

// A Listener is used to register hotkeys with the system
type Listener struct {
	lastID int

	PollRate time.Duration
	Hotkeys  map[int]Hotkey
}

// NewListener sets up and returns a new Listener
func NewListener() *Listener {
	return &Listener{
		PollRate: 50 * time.Millisecond,
		Hotkeys:  make(map[int]Hotkey),
	}
}

// Hotkey stores relevant information about a created hotkey
type Hotkey struct {
	ID        int
	Modifiers uint
	KeyCode   uint
	Handler   func()
}

// CreateHotkey creates a hotkey with the next available id
func (hl *Listener) CreateHotkey(modifiers, keycode int, handler func()) Hotkey {
	return Hotkey{hl.nextID(), uint(modifiers), uint(keycode), handler}
}

// RegisterHotkey registers a hotkey with the system
func (hl *Listener) RegisterHotkey(hotkey Hotkey) error {
	if err := registerHotkey(hotkey); err != nil {
		return err
	}

	hl.Hotkeys[hotkey.ID] = hotkey
	return nil
}

// CreateAndRegisterHotkey is a shortcut for calling CreateHotkey then RegisterHotkey
func (hl *Listener) CreateAndRegisterHotkey(modifiers, keycode int, handler func()) (hotkey Hotkey, err error) {
	hotkey = hl.CreateHotkey(modifiers, keycode, handler)
	err = hl.RegisterHotkey(hotkey)
	return
}

// Listen starts polling the system for hotkey presses
func (hl *Listener) Listen() {
	listen(hl)
}

func (hl *Listener) nextID() int {
	hl.lastID = hl.lastID + 1
	return hl.lastID
}
