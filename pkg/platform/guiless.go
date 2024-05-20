//go:build !windows && !macos
// +build !windows,!macos

package platform

import (
	"github.com/lgnixai/sugua/pkg/server"
)

func Start(s *server.Server) {
	s.Start()
}
