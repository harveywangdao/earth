package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

func do1() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Print("this is zerolog")
	log.Print("this is zerolog2")
}

func do2() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Debug().Str("Scale", "833 cents").Float64("Interval", 833.09).Msg("Fibonacci is everywhere")
	log.Debug().Str("Name", "Tom").Send()
}

func do3() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("hello world")
}

func do4() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Debug().Msg("This message appears only when log level set to Debug")
	log.Info().Msg("This message appears when log level set to Debug or Info")

	if e := log.Debug(); e.Enabled() {
		fmt.Println("AAAAAAAAAAAAAAAAA")
		value := "bar"
		e.Str("foo", value).Msg("some debug message")
	}
}

func do5() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Log().
		Str("foo", "bar").
		Msg("")
}

func do6() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	err := errors.New("A repo man spends his life getting into tense situations")
	service := "myservice"

	log.Fatal().
		Err(err).
		Str("service", service).
		Msgf("Cannot start %s", service)
}

func do7() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.Info().Str("foo", "bar").Msg("hello world")

	sublogger := log.With().Str("component", "foo").Logger()
	sublogger.Info().Msg("hello world")

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Str("foo", "bar").Msg("Hello world")
}

func do8() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	log := zerolog.New(output).With().Timestamp().Logger()

	log.Info().Str("foo", "bar").Msg("Hello World")
}

func do9() {
	log.Info().
		Str("foo", "bar").
		Dict("dict", zerolog.Dict().
			Str("bar", "baz").
			Int("n", 1),
		).Msg("hello world")
}

func do10() {
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"

	log.Info().Msg("hello world")
}

func do11() {
	log.Logger = log.With().Caller().Logger()
	log.Info().Msg("hello world")
}

func do12() {
	wr := diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages", missed)
	})
	log := zerolog.New(wr)
	log.Print("test")

	time.Sleep(1 * time.Second)
}

func do13() {
	sampled := log.Sample(&zerolog.BasicSampler{N: 10})

	for i := 0; i < 20; i++ {
		sampled.Info().Msgf("will be logged every 10 messages %d", i+1)
	}
}

func do14() {
	sampled := log.Sample(zerolog.LevelSampler{
		DebugSampler: &zerolog.BurstSampler{
			Burst:       5,
			Period:      1 * time.Second,
			NextSampler: &zerolog.BasicSampler{N: 100},
		},
	})

	for i := 0; i < 20; i++ {
		sampled.Debug().Msgf("hello world %d", i+1)
		if i == 8 {
			time.Sleep(time.Second)
		}
	}
}

type SeverityHook struct{}

func (h SeverityHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level != zerolog.NoLevel {
		e.Str("severity", level.String())
	}
}

func do15() {
	hooked := log.Hook(SeverityHook{})
	hooked.Warn().Msg("")
}

func do16() {
	lo := log.With().Str("component", "module").Logger()
	ctx := lo.WithContext(context.Background())

	//ctx := log.With().Str("component", "module").Logger().WithContext(context.WithValue(context.Background(), "xiaoming", "xiaohong"))

	log.Ctx(ctx).Info().Msg("hello world")
}

func do17() {
	log := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("role", "my-service").
		Str("host", "host").
		Logger()

	c := alice.New()

	c = c.Append(hlog.NewHandler(log))

	c = c.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))
	c = c.Append(hlog.RemoteAddrHandler("ip"))
	c = c.Append(hlog.UserAgentHandler("user_agent"))
	c = c.Append(hlog.RefererHandler("referer"))
	c = c.Append(hlog.RequestIDHandler("req_id", "Request-Id"))

	h := c.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hlog.FromRequest(r).Info().
			Str("user", "current user").
			Str("status", "ok").
			Msg("Something happened")

		// Output: {"level":"info","time":"2001-02-03T04:05:06Z","role":"my-service","host":"local-hostname","req_id":"b4g0l5t6tfid6dtrapu0","user":"current user","status":"ok","message":"Something happened"}
	}))
	http.Handle("/", h)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal().Err(err).Msg("Startup failed")
	}
}

func do18() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	multi := zerolog.MultiLevelWriter(consoleWriter, os.Stdout)
	logger := zerolog.New(multi).With().Timestamp().Logger()
	logger.Info().Msg("Hello World!")
}

func main() {
	do18()
}
