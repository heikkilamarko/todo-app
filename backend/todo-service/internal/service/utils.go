package service

import "os"

func (s *Service) logInfo(msg string, v ...interface{}) {
	s.logger.Info().Msgf(msg, v...)
}

func (s *Service) logFatal(err error) {
	s.logger.Fatal().Err(err).Send()
}

func env(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
