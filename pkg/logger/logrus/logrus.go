// Copyright 2017 Kubernetes Community Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logrus

import (
	"github.com/k8s-community/codegen/pkg/logger"
	"github.com/sirupsen/logrus"
)

// New creates "github.com/sirupsen/logrus" logger
func New(config *logger.Config) logger.Logger {
	logger := logrus.New()
	logger.Level = logrusLevelConverter(config.Level)
	logger.WithFields(logrus.Fields(config.Fields))
	return logger
}

func logrusLevelConverter(level logger.Level) logrus.Level {
	switch level {
	case logger.LevelDebug:
		return logrus.DebugLevel
	case logger.LevelInfo:
		return logrus.InfoLevel
	case logger.LevelWarn:
		return logrus.WarnLevel
	case logger.LevelError:
		return logrus.ErrorLevel
	case logger.LevelFatal:
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
