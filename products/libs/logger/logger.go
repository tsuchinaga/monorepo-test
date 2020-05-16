package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var (
	loggerMap    = make(map[string]Logger)
	loggersMutex sync.Mutex
)

// CloseAll - 全てのログファイルを閉じる
func CloseAll() {
	loggersMutex.Lock()
	defer loggersMutex.Unlock()

	for key, logger := range loggerMap {
		logger.Close()
		delete(loggerMap, key)
	}
}

// Get - Loggerの取得Logger
func Get(fileName string) Logger {
	loggersMutex.Lock()
	defer loggersMutex.Unlock()

	fileFullName := fileName + "-" + time.Now().Format("20060102")
	// 既に存在する場合は既存のログを返す
	if logger, ok := loggerMap[fileFullName]; ok {
		return logger
	}

	loggerMap[fileFullName] = newLogger(fileFullName)
	return loggerMap[fileFullName]
}

// newLogger - Loggerの生成
func newLogger(fileName string) Logger {
	path := fmt.Sprintf("logs/%s.log", fileName)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("ログファイル(fileName: %s)のオープンに失敗しました(err: %s)。標準出力に出力します。\n", fileName, err.Error())
		return &logger{
			fileName: fileName,
			logger:   log.New(os.Stdout, "", log.Llongfile|log.LstdFlags|log.Lmicroseconds),
		}
	}

	return &logger{
		fileName: fileName,
		path:     path,
		file:     file,
		logger:   log.New(file, "", log.Llongfile|log.LstdFlags|log.Lmicroseconds),
	}
}

type Logger interface {
	Println(v ...interface{})
	Close()
}

type logger struct {
	fileName string
	path     string
	file     *os.File
	logger   *log.Logger
	mutex    sync.Mutex
}

func (l *logger) Println(v ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	_ = l.logger.Output(2, fmt.Sprintln(v...))
}

func (l *logger) Close() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	_ = l.file.Close()
}
