package main

import (
	"syscall"
	"time"
	"unsafe"
)

func main() {
	ticker := time.Tick(time.Second) // 每一秒都会触发。
	for i := range ticker {
		msgBox(i.Format("2006-01-02 15:04:05"))
	}
}
func msgBox(timeStr string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")
	hwnd := 0 // 0表示将弹窗放在桌面的中心位置
	title := "Hello"
	text := timeStr
	flags := 0x00000000 | 0x00000040 // 0x00000000表示弹出消息框并且默认按钮为OK，0x00000040表示消息框的图标为信息图标
	messageBox.Call(uintptr(hwnd), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), uintptr(flags))
}
