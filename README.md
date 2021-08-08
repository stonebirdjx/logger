# readfile
A logger file package for Go
# Install
```
  go get -u github.com/stonebirdjx/logger
```
# Example

```
  logName := "test.log"
  logPrefix := "[stone-bird]:"
  NewLogger(logName, logPrefix)
  defer GlobalLogger.LogFile.Close()
  GlobalLogger.Logger.Println(LogBegin)
  GlobalLogger.Logger.Println(LogEnd)
  // Output:
```