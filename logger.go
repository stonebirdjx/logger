// Author stone-bird created on 2021/8/8 14:01.
// Email 1245863260@qq.com or g1245863260@gmail.com.
// Use of program write log
package logger

import (
	"log"
	"os"
	"sync"
)

// log level
const (
	Info    string = "[INFO]"
	Warning string = "[WARNING]"
	Error   string = "[ERROR]"
	Notice  string = "[NOTICE]"
)

// head and tail
var LogBegin =  `

		                            _ooOoo_
		                           o8888888o
		                           88  .  88
		                           (| @_@ |)
		                           0\  =  /0
		                        ____/'---'\____
		                      .'  \\|     |//  '.
		                     /  \\|||  :  |||//  \
		                    /  _||||| -:- |||||_  \
		                    |   | \\\  -  /// |   |
		                    | \_|  ''\---/''  |   |
		                     \ .-\__  '-'  ___/-./
		                  ___'. .'  /--.--\  '. .'___
		               ."" '<  '.___\_<|>_/___.'  >'"".
		              | | :  '- \'.;'\ _ /';.'/ -'  : | |
		              \  \ '-.   \_ __\ /__ _/   .-' /  /
		         ======'-.____'-.___\_____/___.-'____.-'======
		                            '=---='
		        ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
		       ^^^^^^Buddha Bless hjx's programs never bugs !^^^^^
		      ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
		     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
	`
var LogEnd = "thank you for Buddha bless"

// log struck
type LogInfo struct {
	LogName   string      //log file name
	LogPrefix string      //log file prefix
	Logger    *log.Logger //logger Pointer
	LogFile   *os.File    //logger file Pointer
}

// global logger
var GlobalLogger LogInfo
var once sync.Once


// new logger Object
func NewLogger(logName, logPrefix string) {
	switch {
	case GlobalLogger.Logger != nil:
		return
	default:
		once.Do(func() {
			GlobalLogger.LogName = logName
			GlobalLogger.LogPrefix = logPrefix
			GlobalLogger.LogInit()
		})
	}
}

// logger init
func (li *LogInfo) LogInit() {
	var err error
	li.LogFile, err = os.OpenFile(li.LogName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("fail to create %s file!", li.LogFile)
	}
	li.Logger = log.New(li.LogFile, "", log.LstdFlags|log.Lshortfile)
	li.Logger.SetPrefix(li.LogPrefix)
}