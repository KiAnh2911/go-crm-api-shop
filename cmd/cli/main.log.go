package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// logger dùng zap

	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d ", "Tip", 40) // like fmt.Printf(format, a)

	// logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "Tip"), zap.Int("age", 40))

	// 2.

	// Ex:
	// logger := zap.NewExample()
	// logger.info("Hello new exp")

	// Dev
	// logger, _ = zap.NewDevelopment()
	// logger.info("Hello new dev")

	// Production
	// logger, _ = zap.NewProduction()
	// logger.info("Hello new production")

	// 3. Custom logger zap
	encode := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encode, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("Info log", zap.Int("line", 1))
	logger.Info("Error log", zap.Int("line", 1))
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 177234832.84634 => 2024-05-26T16:16:07.877+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//ts => Time
	encodeConfig.TimeKey = "Time"
	// from info to INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// "caller":"cli/main.log.go:24"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}
func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./logs/log-dev.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
