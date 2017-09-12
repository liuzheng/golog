package golog

import (
    "testing"
)

func TestNormalPrint(t *testing.T) {
    Debug("aa", "%v", "laskdfas")
    Info("aa", "%v, %v", "laskdfas","ccc")
    Warn("aa", "%v", "laskdfas")
    Error("aa", "%v", "laskdfas")
    Critical("aa", "%v", "laskdfas")
}
func TestResetLogLevel(t *testing.T) {
    Logs("", "INFO", "DEBUG")

    Debug("aa", "%v", "laskdfas")
    Info("aa", "%v", "laskdfas")
    Warn("aa", "%v", "laskdfas")
    Error("aa", "%v", "laskdfas")
    Critical("aa", "%v", "laskdfas")
}
