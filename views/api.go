package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.4" )

func U8KMdgKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Db5S7UDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRRGGQL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtlqI4EmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQEcB3z4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L1SyAVweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6Rz1VUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nATuiHwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCRTOPzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4kkI1qkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func apEJB7wdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcfWRwUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSOTF0TyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYh76Y9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjBSiBSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VIiQofwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9OYXwy5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbCcV8DlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9gFLOwCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUpHjka6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kDKM5JMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aTrcE8qKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRl7tarQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0TAZCoqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1PcSoUETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxzVnhUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjCA2l3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyK86mIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0CSsitpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tUYKLzgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqKHKPBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func viOJJW6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XcoldIA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQakNbxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EX4FuztIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fqQk1VyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FAnfSHOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdVksjucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9wHa0qiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UDEz9o6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func voVRzAciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oaul80W9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIucGJjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhZ1eRMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIwW1EyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WCYkUlN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1OUDWlYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FrBEyycRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hDrNcKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHQd2hkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3YeNPCamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptpCEKztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d60TK5IsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQ05cTZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N75nfBHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rf14E5THWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ieV4FhnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WkPqjgQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTxT03q8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NtegzoBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cagw3ZzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFwJAmeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KbZLOUYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UHXYOBvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xs05tAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITmmd0pAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5WlrmaaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HgJz26FfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odKNyZYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRFPFFHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VRtXR1QfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1VmGeypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBkFdralWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ysuZ8Mr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHY3lhTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jbv837e1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AZ4CMcOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0vwN7NvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yxncd0nXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Ls9x7UkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V1bUVWPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D5BcoEo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zyrReDfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgZ2cf1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HyPZ4UUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3RvqqeeMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tM3i6yzLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxmkK4uMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IoDUD19yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDmFOrn3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlNQldgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vTYS6hQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFOEqHw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYjzLI7eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMqyAUFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4WB4dyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vG8zYeF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qAokOJASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kD0fE2bNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8U8nge9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mrGOGhyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WcXH1abOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3aP3BVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RNrRYA0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N481GWvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YqcdAblkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TsJuN0xKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nyf2NnyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYJZafkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9pyk5iuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pz0Kv6waWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjXreM8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UbovyhVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFy2zzh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASaGpOR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVVDmfhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8cz9ufyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NoTndiMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lr6qVTfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TjdfjPJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlrdAhJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JaqN4QzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDLgWVNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZsJGWUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTs5OsIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n3pybvBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozNlQ4yyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwGpvsUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Rw009UXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qijkFo4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwXNE5GgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8HkXOygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgdJrs54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9hhFGYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3KEk9iNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CfU5JzhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5amrRua4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q03kGabTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0cdXqHT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1Nkg7hUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KiVKCq3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbIeTPWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func COySCzMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgVg4jLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqKVPPp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISb2HUrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rbQFDKAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MW2aYPcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IsBcnZ0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gx70Fj4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxmZgi2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxQ6iDh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQGniBOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZmnLDaf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89ce5YRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jofhdYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9mnR4MTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4GnHBgfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pwWrAQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFLwpDdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b7RjqolWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jfoUFmsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1uBcvYIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cMS76zYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dftTwEbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhhd7bsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z2AgtFPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5t9rLUMdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qLNYJmoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUUHfVgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KYSD5HxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjU37CqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0SpyU0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FYuyMDJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qGi2UrydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKW8WgeSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GmVxNKoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QB6xK39OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6c7tw3ZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJOIgjmhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0hFDf3icWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mwnjr9F9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ha3jWqBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3H9PVztcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hyKpBPGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SAOIM79FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fh1pcGcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gT0QuScZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XZs8RTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MnKwDnTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V1FtuFXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROJ3h8eZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTsOJO8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLttLbE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ywo3W1rjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCIU2IxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4rkyYu8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OvnUvpgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyP5axcpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hkix6G0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WiSDFbiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItVS9ZeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BVpE1yDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QfawrKbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkFX1hSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqfsQQCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkzjxYuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EREr303aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4QzsUji7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3DQVcpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZI2fNkdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQFd2irJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5I7OeeGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTyv0EVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLPvlN2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVUBmegAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qRgMBRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cERQWD7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CkR3DAlRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5os42odyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hofneoDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LxU03fJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WheamUnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AXX73t4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qAhxaCs7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYgEH83jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ty74E8OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yvYDfEUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUZ4YgTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nlfny9gjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KofBcZTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGn3DXvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bL1mC2jhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func txqf2n6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jTuLsYzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oeUZkIztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mEC8qDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQNBCdDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FvDMiqnaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JtDKZU9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KjSGZxQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2E4G61J3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2DunraXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tm9DB1NHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nr1i3V1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HVtXfvjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5zkryyCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JBwMPi7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVBwSKjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a9j1H8UCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sf1OLIl3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3XPvLfkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aNKkDQ48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HvjSphdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B7nfYrN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdpqyKJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxYr7FClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cVg7GiHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mvAlJVmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g9HxUnKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TEhUB3ORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vjzsX9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sru55oHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pc2DC6GzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d7cAmCWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJmvGpxVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F71XqsmqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REChONteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWkl313KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 87HSPuEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUgzOWuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ooPdShxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dmgSR0aJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKrJiFcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o5NOXHGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PoVF9mBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwaJezDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CeMqgZxOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V53EAMrPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUJXKiWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubP31k1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2SUkxduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlfonhNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWdh2N87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bphWuNJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGnbc8PLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qubo6Ek4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElbLA8ZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EUdXjJ2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wrS6OwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQ0VsNkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tvka3Sq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxOeB4dbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMGzFm5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToU7czQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G02bt3zhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6nPN1OdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9iqV7JLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EUmTGumAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6bGnAW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12vm9hbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mx80l8gpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6idpi5oiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

