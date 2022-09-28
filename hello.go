package main

import (
	"runtime"
    "github.com/rs/zerolog"
	"os"
	"fmt"
)


func gln() int  {
	_, _, line, _ := runtime.Caller(1)
	return line
}

func main() {

	fo, err := os.Create("logs.txt")
    if err != nil {
        panic(err)
    }

    // UNIX Time is faster and smaller than most timestamps
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	//logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger := zerolog.New(fo).With().Timestamp().Logger().Level(zerolog.ErrorLevel)
	//logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.WarnLevel)
	//logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.InfoLevel)
	//logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	//logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.TraceLevel)
	//logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.Disabled)
	logger_module1 := logger.With().Str("module", "module1").Logger().Level(zerolog.WarnLevel)
	logger_module2 := logger.With().Str("module", "module2").Logger().Level(zerolog.InfoLevel)

	fmt.Println("............")
	//logger.Panic().Int("ln",gln()).Msg("hello.go") // will cause app to
	//logger.Fatal().Int("ln",gln()).Msg("hello.go") // will cause app to crash
	logger.Error().Int("ln",gln()).Msg("hello.go")
	logger.Warn().Int("ln",gln()).Msg("hello.go")
	logger.Info().Int("ln",gln()).Msg("hello.go")
	logger.Debug().Int("ln",gln()).Msg("hello.go")
	logger.Trace().Int("ln",gln()).Msg("hello.go")
	logger.Log().Int("ln",gln()).Msg("hello.go") //no log level - always logged unless zerolog.Disabled

	fmt.Println("............")
	logger_module1.Error().Int("ln",gln()).Msg("hello.go")
	logger_module1.Warn().Int("ln",gln()).Msg("hello.go")
	logger_module1.Info().Int("ln",gln()).Msg("hello.go")
	logger_module1.Debug().Int("ln",gln()).Msg("hello.go")
	logger_module1.Trace().Int("ln",gln()).Msg("hello.go")
	logger_module1.Log().Int("ln",gln()).Msg("hello.go")

	fmt.Println("............")
	logger_module2.Error().Int("ln",gln()).Msg("hello.go")
	logger_module2.Warn().Int("ln",gln()).Msg("hello.go")
	logger_module2.Info().Int("ln",gln()).Msg("hello.go")
	logger_module2.Debug().Int("ln",gln()).Msg("hello.go")
	logger_module2.Trace().Int("ln",gln()).Msg("hello.go")
	logger_module2.Log().Int("ln",gln()).Msg("hello.go")

	fo.Close()
}
