go run . 執行setTmp.go, setTables.go 或 setRtp.go內的main函示 (可用//把不需執行的部分註解調)

1. 修改.env檔中的redis位置,GameCode以及檔案路徑
```env
//本地
# REDIS_HOST= 127.0.0.1
# REDIS_PASSWORD= 1
# REDIS_PORT= 6379
//alpha
# REDIS_HOST= 8.222.248.207
# REDIS_PASSWORD= "Taijc@888"
# REDIS_PORT=6384

GAMECODE="302"
# 逗號間不能有空格 否則路徑會錯
# FILES = 302/alpha/DSG,302/alpha/Fusion,302/alpha/NB
FILES = 302/prod/brazil/NB,302/prod/singapore/DSG,302/prod/singapore/Fusion
# FILES = 302/alpha/Fusion
```

2.go run . 執行main函式,等待顯示訊息:"傳了N個模板"後即完成上傳(訊息數=上傳資料夾數)

(Note:內部執行的Goroutine可能由於主函式中止而停止,導致模板上傳不完全, 設定wait或chan不確定效果是否較好,故先使用訊息確認)
```cmd
PS D:\Golang\src\github.com\adimax2953\bftrtpmodel\template\ConfigToRedis> go run .
上傳"302/prod/brazil/NB"
上傳"302/prod/singapore/DSG"
上傳"302/prod/singapore/Fusion"
2024-06-03 17:38:56 [Info] [main.go:61] 經過時間 => 0
等待顯示訊息:"傳了N個模板"後(訊息數=上傳資料夾數), 再按enter後繼續...
2024-06-03 17:39:03 [Info] [updateByFile.go:133]  => 上傳了29個模板
2024-06-03 17:39:03 [Info] [updateByFile.go:133]  => 上傳了57個模板
2024-06-03 17:39:04 [Info] [updateByFile.go:133]  => 上傳了57個模板
```