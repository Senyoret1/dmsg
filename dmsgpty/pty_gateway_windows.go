//+build windows

package dmsgpty

import (
	"errors"

	"golang.org/x/sys/windows"
)

// NewWinSize creates a new WinSize object
func NewWinSize(w *windows.Coord) (*WinSize, error) {
	if w == nil {
		return nil, errors.New("pty size is nil")
	}
	return &WinSize{
		X: w.X,
		Y: w.Y,
	}, nil
}

// PtySize returns *windows.Coord object
func (w *WinSize) PtySize() *windows.Coord {
	return &windows.Coord{
		X: w.X,
		Y: w.Y,
	}
}
