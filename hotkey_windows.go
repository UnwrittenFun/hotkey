package hotkey

import (
	"syscall"
	"time"

	"github.com/lxn/win"
)

// KeyCodes
const (
	KeyEscape      = win.VK_ESCAPE
	KeyPrintScreen = win.VK_SNAPSHOT
)

var funcRegisterHotKey uintptr

func init() {
	libuser32 := win.MustLoadLibrary("user32.dll")
	funcRegisterHotKey = win.MustGetProcAddress(libuser32, "RegisterHotKey")
}

func registerHotkey(hotkey Hotkey) error {
	if !winRegisterHotKey(0, hotkey.ID, hotkey.Modifiers, hotkey.KeyCode) {
		return ErrRegisterFailed
	}

	return nil
}

func listen(hl *Listener) {
	for {
		msg := win.MSG{}
		for win.GetMessage(&msg, 0, 0, 0) > 0 {
			if msg.WParam != 0 {
				if key, ok := hl.Hotkeys[int(msg.WParam)]; ok {
					go key.Handler()
				}
			}
		}

		time.Sleep(hl.PollRate)
	}
}

func winRegisterHotKey(hwnd win.HWND, id int, fsModifiers, vk uint) bool {
	ret, _, _ := syscall.Syscall6(funcRegisterHotKey, 4,
		uintptr(hwnd),
		uintptr(id),
		uintptr(fsModifiers),
		uintptr(vk),
		0,
		0)
	return ret != 0
}
