// Author stone-bird created on 2021/8/8 14:06.
// Email 1245863260@qq.com or g1245863260@gmail.com.
// Use of  
package logger

func ExampleNewLogger() {
	logName := "test.log"
	logPrefix := "[stone-bird]:"
	NewLogger(logName, logPrefix)
	defer GlobalLogger.LogFile.Close()
	GlobalLogger.Logger.Println(LogBegin)
	GlobalLogger.Logger.Println(LogEnd)
	// Output:
}