package toakbut

import (
	"context"
	"toakbut/cmd/toakbut/config"
	atdSvc "toakbut/pkg/attendance/svc"
	breaksSvc "toakbut/pkg/breaks/svc"
	"toakbut/pkg/common/db"

	commonMdw "toakbut/pkg/common/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"

	discordBotSvc "toakbut/pkg/bot/svc"
)

func InitServer(cfg *config.Config) *fiber.App {
	db := db.NewPostgresDatabase(cfg.DB)
	_ = db
	_, err := db.Exec("SET TIME ZONE 'Asia/Bangkok'", nil)
	if err != nil {
		panic(err)
	}
	app := fiber.New()

	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceQuote:      true,
		DisableQuote:    true,
		ForceColors:     true,
	})
	log.SetLevel(logrus.StandardLogger().Level)
	log.SetReportCaller(true)

	commonMdw.Init()
	app.Use(commonMdw.RequestLogger)
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	_attendance := &atdSvc.Attendance{
		Ctx: context.Background(),
		Db:  db,
		Log: log,
	}
	_breaks := &breaksSvc.Breaks{
		Ctx: context.Background(),
		Db:  db,
		Log: log,
	}
	discordBotSvc.Init(cfg.Token, discordBotSvc.DiscordClient{
		Ctx:        context.Background(),
		Db:         db,
		Log:        log,
		ChannelID:  cfg.ChannelID,
		Attendance: _attendance,
		Breaks:     _breaks,
	})
	return app
}
