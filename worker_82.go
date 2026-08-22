package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.8" )

func ppMMM3upWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qRzfjBMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6gJMpeHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wc42TjDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C7reTEv4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rjeU7axgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0LWEYlOhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ctKT1aZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CF5t9J5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yB1lMEqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pl4LB3SAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wRhLrtnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O1jm3iATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QjdQ3RN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A7K0tj4mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9AcvZIT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9u4tRGx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TXizj3c0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hIW0z8SDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OR194VmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JFbpkqlUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 64nqc6D3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func we5E8KoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XYZO3EbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func klhLbsOtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IxXCRMLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XRLf7MgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 99bzR04SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TIbS495JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iyfl9etBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sDbzdbU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5uTcbl1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func asPlDecnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a7HmNbWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CU4Q1HYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CUmfPBZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nmw472zCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fsa3o2pDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jSVc2G31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BNnUoS7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lKvaYlyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RP2IkePXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mG3iWIKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iCIQiINoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pnzOkOTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QhFmFD7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oFd7fPXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LJrm1RSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GjCztgnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DIvoOxi7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cM3JZPgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DlKc0HwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GT5j4cvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yMsIURILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mALX9jcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w4vsGTuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wP98uTasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3SZMxXLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F89XMsfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2MR01o5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N1rEOOjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hXQ0t8qNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jyAo77sLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mo8qGip5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Q8FcR9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jxM3wOLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wuHG16u5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pix61AX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pnbNBRv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dBA6sBJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rzjFChwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W1cqRP2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m8N0b2XuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kLKDu4FeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L3geAT0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KOjXtFLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z57XfvmjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C1nvmjWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mczJ81zhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bn67WpVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FnE7Kdt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tjvQy8tWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BRw4HQmoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bVpBwWE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pVMTyLCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZKP3ghidWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mDqmPzaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 98FU7jlHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ywxgDIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RTMPAI4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M0JMfZPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6wmC7CLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TRpAKunVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tdBLYIvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C1yecpyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DUDieJezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jysp6v0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uuXzUVIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QLkhxcBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pWCTiWeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j6GzIExDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6tJwejPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vu554mGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FcrDe6MjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2NBwazsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yx6JpKD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QFkoa7L6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7w6INlepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E5sJMiLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sXEgbQqWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OWw2T5seWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fpg2tiwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PElwqEL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CqijQpeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t4Ssu42eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HZRZSYJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FFbeYuvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Xsxp5VdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6wWWZ2wOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lzvl87HbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vo5YIJzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CA1MvR1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q52M7yG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rD76E2NmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 12bvdHXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kNg7PKbBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z7t2qNDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FAtvdlkHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TNkDncbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Toq3OH2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5vVus08jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CkIUWVfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dGszl8pHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5GnFEXcVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func skhS9YFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OFKujMgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lUoIdzWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I2aHneboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uqKfbMyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mfEl9ugbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rXmUOuaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LhVmkqwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WLFemLKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mu1rnIVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V9sFMqYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hR69CBvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dHfVNE2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M5Aof5vSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 65kU8fFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ru2fmJlQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wIwGUHrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E1wxytwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mVVVaIIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DbBtXG65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PZ4yXjphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uXwknIjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Imi7xZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B7oSLA1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kqxC5lI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N9t4tbscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qa80Sq5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q7YNLCw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GuxFqaLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i7ikEKwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func slKdNdYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E2zTWn3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YGuj0gMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6wyqBWZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8iJV5WP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BbB78PcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C0JlWd1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zwNuYy6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wtdPYEiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 85iuNljWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UnGS2FRiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UVvFZ1HOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YE7M4paiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YkTi28xrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GN41u0J9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JvmONFTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a0JVV0hFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jk2E9GLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JgtDzhPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NZ0K2eq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JKrj3obnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VeHk2Tx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5LSTJkM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7NsXZQnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FMDtgMB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DVkHpwvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 94W3VHQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Euu2Xzo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7ttaiLWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B44dWxhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y8hxM1m9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N0dVFTooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xdy9MK6HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0McSd4qvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QIUdBRzhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fyaxIh0rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8npTDhSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TbsmKoeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q6W2IFfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JTRaC0MFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GHxPwWVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EB9GR7mqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i1Y6jfo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y34Hvy4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sArNi0k0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eIRzJ9t4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QjwRJoANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IHlx8tZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8EITpruuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sS1sUENyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xD7gFbALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PDMLpDzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BcgwUI6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kKeWJqFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yGlUdK6LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ayuhkGNqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rHJFXZBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OCpAP1GrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WBqwm4lNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HRkz2rYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CZePxQbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LFCR7UseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KHl6IgTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PEPrGQ8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pZaycwLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DhkeuFykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P14IAPhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3LQSHh3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func deFjBWOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L8PxJJ4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lr6B5QKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k6Svc5y0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E5ElSpOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RvG4Jz5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w1PACJPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oQSsbOkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c9vXj0h3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ur1MFDDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b5yag3QxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F8trEv8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b7LLtZU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N7EkjLidWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KpSXfLKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ScCnOq8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E2vSymucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func elJQ6fv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QaLFV3fJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hVh0j0mEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OFewx9qfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rYBK1TwZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o1zPsLnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tzo3YqeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K9UGTFbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fmXZEP12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yzgANFDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qdZgowIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FZJcTZ0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1kVowMFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 30JiFWceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zUSRWgLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ncOEHm6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KkQ8jrnMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LCbPcN8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9bIMZZPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O8Q3HkebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aW2ZHHbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hNTpuRnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S1klOgkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EcJNDbxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ho7ivl54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RkWIQaXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jAJH5zZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QFpSLs2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JWRpQUj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8C4cAgJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IwYwAeY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MWrptAktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3aaLmqWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aP463csJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vyKuiH6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0kxgRzg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TFeG3N1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0BjDF7UoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OccTfBsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yFHRlvbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func waCkvSdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F4nPuxffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SkVY3VjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hrxE7tEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UWssimn7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NeJNGxqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m3eIGHXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C2oYvwoTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p1B7tKd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0hdlLeOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UPofD0WDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YG3QhAE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iY9kUYdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JjHXTRaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tefmjWTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2oENy26EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BdM3zH8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7tEWpdV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VsicVG2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PaJCF1SvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YNAeNl6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W1Egx9eOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XeokEbM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func va1QmaxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3MvciH5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mipxqV68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WiQWyltdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z15QaL4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kiOV2wPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yurk1hFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V0NzRDaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kD6Sq5tSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5cBeFPqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E6dEVzbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sn08A6InWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SZE0BUmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lxm88EbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mmMs0CGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kteNcpTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aycw4mtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VrsbTic5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XTgdMDnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kQII6NilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y4mHEnVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yX0Wb6oZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nJjlMhaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bcwZDCKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yhJH5GvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func syQOdpk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xs75u18WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P4m3POgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ztQHT85OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qwjxccV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k8sACVarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KS1lBkk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HLKwhIoLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8agsMBcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 21YBtCTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uO7lxCrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8N3xmLKlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eprkXWdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qlaZ7a8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3qk0zNIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AflvPzCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9OizuphFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qtb6qcOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func igbo2en4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oy1Zehv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jVU326PwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YFZUVOSUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RMHUquy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3nFDZ8UtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3jYjyRUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y3azQBnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JFQrvkwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 99LOqb4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aAdFenVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B8z1PTdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YBKXVX82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rekwabJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WZ8FNfJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q8v65ZlUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ugRBzE9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tnfpcpOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c5FbLgSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D12vAF9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func neVIXiKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 71c13viOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2TPV2DTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e85lELxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 30FmlM0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tuo6ROhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func foXIiZYJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5C6TdaKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SenPywfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KP9Tc0bdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JKS5Qx9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8aWOTIpLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0KrloKh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GI2mjpBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IOKOv9rCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xM9Bz4uOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UjHASeQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2nUO2wKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5zyucdHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KpNVY1YdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rT1CI3sZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RM9apbaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BLEpvxzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yFvMpqT4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RvihbkKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ICksNlzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ivA6mbC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jiaGk75xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Y1eMgtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uAzrQX8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AOG6nAQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VMoOKU0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JO7SZro8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func buVRg1iwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W12XcefnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IPoAHAT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1rmSoVT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZVre2YtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5cfmBuX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AIG0xvhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U1gsheM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EfUBz6vMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pqoPrnQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u1Su1h75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func huOB212DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x0iOzdXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ux8YrW1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wjq7W2B4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JD7XCw6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zVDcwc7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func soW2IttlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GEAcMfl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oJ246jgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VaOYbZWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oACujKHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wfAJa4gXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wm4s51gcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TwuMvXuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BaYQBd0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mL1RbuMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZMcJFf1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yXGRghmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ovcZqaDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func exAGGXJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iLHnB7bqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nrx6rhqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PuM2G34cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vExGGQYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Ugc3oWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X5zh8sT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LOmJ2lFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CdOwovrBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nidArteDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jFqTFlAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gJQM6aEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aOiVOcicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DDW3ahXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l4arS6AXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BOO59VAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fj5XHwbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2SlFdB1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GjsOa2rRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kj5mXysFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lDainLAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l8tQjKpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qn1KizJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4x1WFUdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func seKt2rizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mRqi749rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vembGRupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func prMJIH7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CRfGkOrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BPNU2xAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YCU6L7MLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 49qbIpmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dwkovz7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sc4E0ZMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ivKvxfj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B09uqbqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KGQn0QP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zzzsyWwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func le3tE3GjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func owPMy4svWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func djYFxjFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WTor3B3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func njHwZMbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ebSIRbCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BEKEvvoaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1qNwsk1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sla36Q8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VMJ6OiivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gen6x6JoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nw6PCZIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p6qmfKyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NBgoOZ9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qmJLGQMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FYzf2C1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hENj4oI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hFEYBshkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e0RCFhHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SzePNt6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UsM40QlPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1MufbQwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 37Y7MbtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4upv0VZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nNELXQ1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dj39yPVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UF76fpKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iYvwId7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vK2vxd2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wxxhys3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XIXX1Ny6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ucTSpr4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I6NRDYWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 04gDo74TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j9g2tJolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lpqGyMIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TjB8fH6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ju0erQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func limYUGK1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EvzYwuU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AtbLPSUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FPK1dhSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oRu23VgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C2rbBRLHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wnOlu0p0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zdniTQuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a3UbyLdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mwSQpsdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L07IqMiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3eBYosVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TLUrz0JyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tMPYiuBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BBQfhpynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mcO0AspSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f4vyEQVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3A4h9QB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J7NRskeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fPO8hHLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XMM2Qx2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2BNqY4ZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J68COBaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iLa2OlkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sa6NoDwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EzRVt3xqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eQTIt4wFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dU704OKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GEgwtgp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TkWDsMtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OzKo4mjPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gIK6en0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dYzKmrZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Lk2OykWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ASaH8gCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EIzEQRgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dydo86GlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fLO8HuETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qpiqdhSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HE4J7JCaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func swjJhvGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3UO4gdeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func guav9QJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u4GvFYowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ivSZuvSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u0dCsEtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2IpbYNXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func utFd2nsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a1xNeYYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wWbsC1aBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func peRMnZAGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TCzKhrOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aSOhhqx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T6RiXXgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oddf0ANIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xp0m24tEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9K55MBwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NCzU4XWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CZot7tBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tSckylw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rRa0XRQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NuqQSumqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KROBmVSUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K4JDLMEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fPYQpDPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func boZ1P9OGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func igao7GRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 58aQ6SK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9SEo8CS0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fS5rHm2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jzDQ59wqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6wdSQVFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NJ9mxt9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FOgVFTKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lgtjS6cpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iqm6tfTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y1TKhP5uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func twN9AdLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F2lO4UEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8aMFwRjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JP8tEfhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V1UaNCqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yE2wYuCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RRamJiAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rI4OJtlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yo3sbocVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uBiynw0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AMI1TEheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ukfZsueYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qBGKLw6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JKWOwUvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sjetcYTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JXgzprC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cD5WQgTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oOjGVRdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func unaOn3M6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bjZUNci4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WGpWNcWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rmy2lBIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aB0eTn61Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WIOx5CYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KMVUIXoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GUj3I2wdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hCwsy3Q6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xIWIGQSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VWFFWkLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GMhqAVS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6K5v8RIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yg7OH1PsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fMysD0egWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PnvripwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3uzfjtSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ttoPsOU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TOG1TRETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zABraNNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func daffaVbWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7UeXdPHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5WWFQVsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zmvdlMNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U8Snu05bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JFjy3BvPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CQsHU529Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3DsEdCuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rAA4N1PhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TdAb8K1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nYSlUGBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DBzUz2yFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1g7TxKivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J6OIAhUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M0KncSADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hGNMkrb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h88riyNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func isOTsdhAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1lyCTQswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 52Yr2mnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TSn5w0zAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uptEV4y9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7TpNYSexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mji8xrsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h9NXFF7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g9aIYiApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o0qQmJy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kP7Cro6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZZIJN0a6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rkcL2hINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qfIIQHHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k4x9jJnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ENX8Qb1eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7PpC5nzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LnjmzZ7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BUsK5LatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e5jGrW2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fNLNiZ2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8JjcDt2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PoE9iyfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Paj3kKOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QE4juLbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sjd1iiHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wcs6vWgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A95ntPqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m5BseaoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PmJxz4XfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0sZkgns3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e2vrOA4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rw1ZtAF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q6eb858VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GOv3HFucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AiC4ZamGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D1xj31uEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EEePZOgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cCD6XK0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func prvFfXOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SdCREOYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Do15BwohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U5fLpekaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6OT3VOhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func faLqTJRHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cL91iukdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W1DYqgsIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EvzLHbuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UMF1W7ZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XoIUDP5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qhjYJVVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M0YHMcT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LDhPJOxxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N9YnyGtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t6Nk1nOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7FxdqMsYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JfJQjAdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vO7xID6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aWkfj8bKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xJVBLystWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vIijCXLmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func js2aeuriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KQfpNAokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func daaautFYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tsdzh8UvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OpCInBrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PeG7fbChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CEyBuAYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8wMtYLajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func phlMWed1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Cv9yiZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zlRT3UJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vD85MTOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HzmxgyNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sfQom8mpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BuoOyuzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FbA3bHYJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 787Xo7YDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8XF8qPeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jNhmkx6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lUCeztfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6SZDRQcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wmmsXqKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DHavNWCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wc3GC9jCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6PXMg7PDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 74YXwWtbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MQShxSQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1eEHOJZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UCNC0MoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 39icOeSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NcrBP5yNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AxX6aNxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Im4f8rroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9KU1fdxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x1qgNRRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jfFDhXxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YvRNcWg7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lF6PuOv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hVN7ODoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TF8eftAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FO6d7uW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1OwznSIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VMzroI3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LoywDqbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hXUTprQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KIoEM5HUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HJT8PO1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PAkKB1pMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VOax9xZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z7rrgcmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t9V4ON5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QJxeq6fzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6eDzZBgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iFc1tzzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func navs5lY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qlJm9CqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oWk6St4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bVkuD9q3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lz9nDHUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iJQCWLGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zv6d9FnfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8cY3L6B0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HmxVHDCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v7CxFBUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P7pYsNqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qg2hWsgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cx9ek0eHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KibuMuOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func owqNep9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O5rQNNOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9VIcnc3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dl3nhI42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IbrfCoKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QN3Vu5enWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OuTP96a4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZQMPWWnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aVuAyLIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lgI6gbFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jxqj335OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fvEmwGaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OeC9kaqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lUGHhnKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hGdHPJ3WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func scND5WhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qh9dO7vlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0AU7vfQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9mFkRrYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m1B9sdrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jUHLajV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A3iJOSgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZTPXgMVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JtCVRUMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WmGh3ZgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NSzQlhsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SdqsGUk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c6dMucneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f8IFdfF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gBiUdlAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uI3iX23SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HwEo8pHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nosuHyIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lvmKcHYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cUVsehAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mgBaksBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HnY15sZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gVTrq7LtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n32TbU6LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3qUhzlLmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B399gra6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LqmuJz2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hdV12NW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xXN5bpmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SGSzXgBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WyNox0e1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ga4fRNZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3N7Y6YQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E6x5YNzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jk12ACU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dDnnTnCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d8azlKaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NJjT3K2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mTakrOBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1XYejqR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tO12Rt3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pnXAjY9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rbqv0GeSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ZY7Qwv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oRA3l7DnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EtbmZqjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qSdeEwXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uMAAKH49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OsZaJdyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nNhStOZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9C5J2ZpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8FKvPPz3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1hZ9FXPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 27RBgqRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kDULNDAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FAyqQZqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nuOI5fKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SNDhjeOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hwW38m1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6nmHg3DAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XdYJLPiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vHe3FhlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qt93xv01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4OiATCbFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cGDQ9BRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bImVgEUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aWtMGXeNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I3CcUcztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TVfeWceZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GsHay0DMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r83PMtHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ralMpjQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BgTQcfklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TvX3NHMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eJ07hqFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JGZvD4A7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vvaf8rpXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mp6xsvekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1MQAuinWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qof3qTx8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zfDaqHoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VFnsgKE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ECx8oN3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gUumJcyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func McjjZahqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wLHo9U1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nckuya0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XyVpjGzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W7DSAJmIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NEAbM688Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J3WF2DtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QZnjliw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4SnK4IAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wtNcKPY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ziTtjyCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func drzyVj6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SzkQKSXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gokaclj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AXjCK1doWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9xuTizTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3x96qswuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LK20xKGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SQAO73QoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yEuXN2F5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g2kwwUEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kk1teKIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h0TgSTt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XbhNNcS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CwRAO92tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YwJokPr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DhTTixIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vUkSmxaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S8kqcdbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func neJAtQ4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2FAvCLH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zcEJXfjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iufrfyi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j7WKdlXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CbvJnl8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wZS9zUgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x7rH07tTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xPQTrlJFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XPSmJsTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xk6WmGqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NhPvJNw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IMbClC3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IGhTUz16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oNIc4PfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tTlirKznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W3PlPsFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g1b5gO7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yzAYHRgIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func es19q8GJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nE4f3OHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Su6uzZ8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func szqFMWYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RSIr89mtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kCWScM7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dQfVaXrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cblbj8RpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d2Z7MDzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8rz3mOOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ammZDaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1EzkAkWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SmdtfWQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pbV9ElJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GNozV7tTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DT4dmyT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6j6OByqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i5Ud9gozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c1MXJBlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xTQi2ubuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WH6Ajo66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xvinQrQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GqlvUfFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CIMZUaEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ku176dk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p6S2OHvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6hSBgV5yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dZ7SX648Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RpQnHdYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7byXboFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pahNJMK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I4gSsomxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h68oh0biWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0pa8x43JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8H8EJuC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z8dWnKUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pcgp0TKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SP3J4ZMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9lHSa9kjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7bV8VgeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func imvfY4TkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eZK5N8TpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2nu7RoCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6xy37onoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t8Hxv5YYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dlkPlZtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func opSiD3l8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Pa1ETz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nL4d5wgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5x49kiOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4e8CLUbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FlS2IgpLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DruzDeOOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BCvnlgh7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LF0xxlbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iAMX3jEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RI06B3VaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 03e7Y7g1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func urBoChIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CBXSdMcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func trNSyTc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xTTMk84QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RaUbS1LjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pvvBRXRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ePGrOO0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ncsDZ5rlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0FUcmBLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WlWay8cnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QFW1NLWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dXg6FkbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JZVaFSk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1puY2UWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XqsTyKkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 73Uo8ZoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GPPROl5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wq60kupWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p1yoGspOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yxf1bZiYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZehWL4MNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b2Nlu4fHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ql06v9ZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 24dPu3PGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func shM8eB96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PeyRcqlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fxf6bH0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AKTUq73SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JJe1DBR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jSptNOqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IUiOzgcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nyXc8MWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0hOPBwiZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bzrvKSQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kpIBu9rLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fEmLfRoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EGOtTaLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4cOfoprEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func taOm3sbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0wZLVsweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GbzX99hvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QsTbqAsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ySkICu5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6NgwnWk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RPNJJN5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CGnfvy1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E6y0lbUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d21fAVjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Eg1rtxavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s6Fv4v2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3lYM91ecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B9xaGkMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eWKtFJBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hMr6k5FJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PhJL6dmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LdueSZkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TcXitNJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SiuXf4mUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QmlwlxUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func naMj0fpqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DR1lOOVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yRkXLvArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oSsomhxYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RZ9F08NuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hy4cnvA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func shz14Wo2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RyZSI45IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cwBR8SXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xOxiSk81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kMW28mhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cY7AISX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4tsBNiigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BicGYAlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6BNdLXSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5UZpz57nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IW86hoxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pgmbWA6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4FLgZNjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KOsv8C5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MTVxoFHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CTUOPC0sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G7pODtsCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FaeKBHx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kjEAiFUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func isOrCfs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3jXjuCd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ipPaJEtCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nUrtjAFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V88L6idGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oP57OchoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func loLjfNyjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bzp2Wi9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLHFWTWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CmBfuoc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K8SjKYbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YwlPmEpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g47iNu3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IFhy4JGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jLy55AaxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Leb5kmI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q1enkHFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QtBUCnL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PL0nxX19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func waXWSetEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RpxOE5TqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CkRFQzHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pi5AWXHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dw3Z5uQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eWaKAtpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HfC4MtPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ngJrO6IWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nr5YLuXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZaveiHwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3CNj9cKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PTm5pcw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func daNcDVWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P6kIlGq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9XmhuhNZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tQVWC2a9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pn9htlbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pQwvMrsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t0D85bEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lqn688qaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ndrk3KhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7ZHdQAyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l3U6CeOhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EeB4pGPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fK36akeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BBNFKOu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P2i7IuBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B2Kn0HM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f6x3bQdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OiONouk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sUgpVlv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bgIRLCNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xL86j9D8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OoKDD4rIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jooLVAe6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X1x1kN3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kb64rtK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vd6og3H4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rulBefy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Im1AUQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j0I105HcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5BYZUujjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W2PYLcaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nN8HyraNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WCLntQJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zUqwIgetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z7UIKn7MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Abhwfy7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2umuX596Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hi9Qg6sdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xyuOSPVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d9FJNM0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NjEWDbhsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MbXRrJiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ZWeJGB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1T279nvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lru0H5DRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yPYgqM1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ebyYO0rEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lbpxVpxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3CAqr382Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2hEMtzQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Yu5RUxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func my9d12iWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c8shV4UzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7xQYoHcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func phr15dedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PnezHcsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KEGOpZgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j7F1abzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dINRRFe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zq3Hns24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NFUmVvbxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aMEeTqgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i5Ak0XlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PBjCEaRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z1ESd6I4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func forsKVJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q8HxnNaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a0ku0BZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i7jsSdtnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KByhVWjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ArqerTOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XJPbGjYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tLTKHcZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gybccxwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qnUPGWzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SLGNEqw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dmKghHngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S0DbSm3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZlW6GPyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ICm5jt9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lfCN12LSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0oxwuy1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QOgSqZMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nHDWuQRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KjkAj4UHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PaZJRHdgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sx74FvxcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cO5jh3H3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uyz8dX4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OvJff0qcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GMAZyfklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dNRT43oaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y0DTp9tDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func stA0f3LiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OD864tR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tcrOPnAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uGkDuXvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WLxHMrzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DjUfqrs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nmIHekaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y7Gx2C4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wr39ZTQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L6Xedey8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8HBF9I7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7l27F3vuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6OdqIIJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hcRUrWYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QBOV9ieLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rOkQYlkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VprAI82qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8BUmtbhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L49QUkg7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QHMsvazLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IP9g6ojZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Spl3Dy02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func clVKMAV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YDeqr6UgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5rC7RBPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nD886dJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VRSeomHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9BsSJVP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JOQgvNL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ep56rh0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DrjZFegxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JkJfC4ujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gmy7F3kDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DfELBOXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HTa9vZsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wcngTojKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R7jYM5sCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TwVHqAQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z7EFlRi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dtwTBfQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PAQqyjQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eDUDqLG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vy4c9M9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TuBjZn3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KRxmqOZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func thckDBTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 63yQFkNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uHJvPshFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X7IEYg0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4v91Xdr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0iIeME4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mFKWpvZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wTtQsDrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hn40zNdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N5KkWqnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func liv3XTxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Pp11bS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i3xXe2W1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EUuiV9ptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QV3DqOizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RDrCddXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ewtiVYA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1mzXZ4u7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cTJWB1YEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9euyASG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wy5kvGrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LIV13Ms1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UB3qWfw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1NdVXcDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J1EQSPWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VvXqqIr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RU6cxhfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zVdUIXQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J6ZkDIWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E0Tv2XFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ehcmbg1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EIPXUNyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tVPMlsa5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yv6kBUiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wRP9FmoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PS5W9qY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KCJ6ClX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5LzRHvHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lv9ahsEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func omXn4Vr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZsqNPqQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q1TaS6heWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ty4irzo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aQMvQ1jrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5WuP8YAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x7ydrMCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bvmi4Ek5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JrQhxE4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0CexZrhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oL9O1EiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hWtVbf0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jFX09uB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OUh0TS0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zMP11M9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2BDyTTwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mnsf7AXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gvM1RQmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cvMKevJNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SaShWnZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1rqnK7hAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lKS5LpSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q9tauEiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1WUtDTfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u5Lxu5nIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hjnt5UoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oGIbgIznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d1t7v5nTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hmOD9XEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zZmN3TuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h1KU0K8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EIJkQKHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dBGJqgiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fOfF8xaHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kHhDOKCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RRXiTqFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X73mrqgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JN3cxOxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func otwxu9muWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fikI9nxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SC85unX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CMogovXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RPqstn2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nJPhcLt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ofv7wxTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jNgmp2O7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oipsCBQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BIpPNjPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nticydd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y6nmlqRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SBM9uhp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ecZLpo03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ineFxx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l8AqEAPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ixZzvuTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jSxYbTSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UQjSOAhBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L59Ny59zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oRLsmSYPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OPWZ0YjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CXqJc95UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MYtjiaOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1l9xMWaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oJ7o23J9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FNfEMVIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DHy1huE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jQ8wjRRdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p90xAIpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func caSxhOzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xELdGGYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KnsliAmAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RVvOn9IaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9x4fhtiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SeaRRCTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gvr9dCOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vBzbjTKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0ebxPrtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2KCL3AbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X6U7PNbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dstPpIyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tQwFX58SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zk7LJBudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0T9SzijrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eDQdWTEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZnsNTi9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GbF48GESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZZIEkA5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 17BaNfD8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V3yzOWcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GQaMRd77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ka20IxrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NXr9LI2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UKEhY0kZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P1zrL7ysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NbVdTHTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rptHfHwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lnn9Jf43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cg5HucGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1H58RK8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hb6CgvqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tlU0OHVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cr1Vh3PBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wdtgCbzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jv7IHIyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4xA35iYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VLdZBwdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nVcD4ko0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bYhRZxDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T10BU1VMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ywdqZOZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QPFInkSZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qxGEPbRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FSChz0fQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I3MSaV2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 88nG4zf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rJ0eQNjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BfgrMf8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AF7DZrgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func neHcCRAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U7485eTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9cTFZYjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fkcNh0WqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iFlWTVDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4qDb8NcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2e1ID2SvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j08EyN1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UPugk7TEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0o1hEVvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AJd0BG0wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d6IKqlrlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8xoeGgi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CuGCBJugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aEFNv5u3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AQ6oNMfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lHTJM8QlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 05K49uAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8PiHkbFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func voih273iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tB5lif90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AjVsmD5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xe3ua6nVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rAwDHTjBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func INIC2cwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JCKkU2HNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zf8dWiIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U5QQnADZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GeyHYxUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4RHpaa6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KAMgL1ftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GN0nFaCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I0u5NQc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MfXw5MQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7xOVOs6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TPQPsOtmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yR9qUd4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PbTIeuRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VcmVLi3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cbGEtpSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KLuFe96uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d9uymtAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0BSMMYJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G1ssI74KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MBGOw7w6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AyeaUASJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pnh8pPhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FSVoNYbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wv9E4IEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NTalzFH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OMsFLufJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yAHMjh7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6OiqIZuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ek61jUo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7STBvMJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BDsKc9ODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zXTOMqYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ExMyhPN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0stbH2JDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FJtaLXDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P5U3yXMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AbrMUN2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9sD5snsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cvq6AHeVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SHv4EV8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jcazPHJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hpu3jGmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8JSP0AHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tWdLr5WRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aVg686JPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iXkXibCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mfTCgAUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Ans9vq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A48wG5ByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W5BirE8RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9HtwPTrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5L7ThbCYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A0SAlEe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DkdbE5QtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func quNudQBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lv94LsaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DP37FXaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HFqI6cJNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 99lolWv7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ovg2nokhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3Q17zIwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6XHi2alPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XDsR6Dc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7RNdcAMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L7N5PxyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 98SMC0CLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vedfaBZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j97pYhaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ai70OsSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R3mEunOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F7QrtMh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7QKHX0osWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wFmlcSswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wvfn43h5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wLNy4xr2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f4vXXXeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OFHbFcV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oYfRiEoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1j66TGV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZhDxGPgpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Kbb1ynSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6EA5yIhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A3fjmRYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KYpJNSaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FzEza0GbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mrev5ZgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VhZRiNFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CruV8X0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CneivmpQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3BHClNkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cILlcXONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yhrGfWFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cM5nPgjPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GfiAjWvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pn5eOpvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Ta5Jt2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v2C6NZY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4EQ92hvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JBFUv6QzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c28O2SdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q0AZ7jSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dghK5gspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mAH3EJq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lq5wvCwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HbxZaGadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4YLOAjwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fgRYJR2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z1gVdFHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Nv6vhCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9JY0V8LZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sUtZDMnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GmsbyOwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9HadFsR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ptUXqqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d37crNumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UAc5VlnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kIRNJqz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wz7HYYA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XgeVPnK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OJcBXsV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v0mWIDvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TLFgxzt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WhcCimQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zr1gvr8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yma1be3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bVLLQmMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kl3s3EIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FtxmWGbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jFyO8hw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rpkWQQOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3W4DbbJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f46PajL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5RRbdezrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y0Q0ytQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EPKryxgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5mId0rYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 51DxlNoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Gt6Hv0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IjNrrOqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W71cOSMdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xLi5v02lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func beMkc3HFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hMkDMuGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aqvnDxOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oHHLBpJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oRuslCG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fPk6zWqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nv5cFDjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LR0AUNl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2nY0WINGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AGMDnqZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dM7qB69UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ui9g4oceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fPBIQa4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2EpuJYLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yvHgb5wrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GngFseiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qi0GiYqhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bvIn2Q0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5GnXAXvWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xMR4UAFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tl8kdBCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ivaZnmG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gytd5WCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y3uVEKJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PpKg27zoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RLueDg8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CrYvEqOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pjkNXeh7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 16hRKjcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1ifzSLYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3hJtDNsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IMDMXZpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CMux0ZgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jqBU3Ke6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P75DlF7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e0ibw3YfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FtuawtWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4cA3BnX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SrkNr3XVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Th0OxvPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9UK7ZszKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mbIeus4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wQfZcqAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4O2hTXtVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 72VHuRc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LbWJ94DsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8TawJYWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R09y8or7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vUGGq2HlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jh0Ot31dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PaWzCbJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pRmJcdKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z9PQxosPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8A7nEPNLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bDUIiwWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jJqVRHxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fHBKGBmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tQf5Np3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VVbTN9ZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lpscwhkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ORUf7eGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O6ABzOoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GcfDmUykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1BiZ5lWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OLLpDOCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u5ZP3xLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uusyDdTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C5H9N5M0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VcF6aOquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mpsNwEajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AfhyXwjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hBl7L5yPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hXhXcnzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fCebaZAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2E9aBnSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sk78iWSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hrBObmKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ReTNu0EaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cPHjgbxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ug3MPA2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3nnJUVAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pm9seIByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 88PEZiUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D1WGszC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OhcLBOLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YXaJ3z9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bCmaYRQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PFSex7rGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LZ7vQAPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xVgqyxmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rKfw59BSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8IzHthyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cBzxecc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PuRFi9inWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I97ufeaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mFUKM0VVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DnzPgmuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5iuSL3jsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LVsWk6lJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LTWrJqgqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iW7ogMkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TE2jwV4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 74GpEF3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w3TTDhrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NJ6kovrgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UgrxsB1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dIB4IuYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qwf8rzsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VALKwZ9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pvun5QbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nl6Mf6VSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AAdHYthEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 17LmyUQUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jcriKVYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tLcBVkvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xEjcLJg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 840Uc81eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TocwSBnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8NjTwCJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EOcjK7ILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func huBOOc5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ESiLYp48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3w6xzstuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iKXFz95AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PQ9nvExtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GxCAd5VGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5kHM81DVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PXFOS0sKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eeASfvmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sqnVORd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6k0ECcqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P7odDthZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hMgWiZrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pRtBrr9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WPyqycWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xWQTc7i7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WUJKhFFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N0dtMqQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ByHJd5uvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8elMg3GyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f6MY3GZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5O3FeorLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qsPuDPZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NGFrpo3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nfkap9leWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5QxBBwdAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PquyUr0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8HMJ7kcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 28QdCkGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CWlruK3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a8pUDZ9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vrJ9uGPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TBof0co8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tj2ItvsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ozhCq02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wIXcqfUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e1TBTD4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G5hrmtfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zfj5KKYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oTtIi8KtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v46MMxyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zJgKx7JFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nVqIC580Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lYP7OgIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FJKpCH32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fQ3l50MKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QI8QYwbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KxaYqTBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bf9SXgECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gh5gYFDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WcOnQyE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func biFr2F1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2WfJcFn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CQsyafnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OtcIwsrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ky8xsobzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oi8GhM6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eH1iKLrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5bdArOjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ce2DbJBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 43bwtbh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VFkyNqRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OiLkQ2mSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DrO1bUFQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func giVfBzQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func azCRIQp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YiB3mARsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jo5GETsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zzRMru2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gdSbqPA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dxq6jWlIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sT6xrfPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZHJZhp9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PFtkt2t3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4PTupQq4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hMAhZ4h7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7ffusbYgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o6KuddERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3sly7dsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tDKqSQvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iArKee92Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MWezPlZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dzBKaOCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1nJhPSYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IgRUxY5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LBuLd8aoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eLJQnahUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6wIykxgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FpXEWwvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u9hOTmlQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VWYdeiULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 85fdOSAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kjKtLkoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tUZEzpGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LYPDyRxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IBOYNe07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bOxLm5HEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ez7DxedbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P93yHouQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yy3wCpZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RBP5w5pvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sJJ4qgzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fdjzcR5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uVhUVi9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yhpg6TesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qSdX7Gy4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tt1wWmMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 19KeKpIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KrCxfNwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1zU9EvYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d8H5me6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DiJaHGPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oJs4KxmrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sA9JXxV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nOKiJcnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qN6R5lEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YbCrUeTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JSJ6bQwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eo9fpYePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WFqqXkJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GAWSpqjaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ygOCpTqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L7Zn45JXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VsxN7h0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q9w15MGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OBifdgQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eStdb1sDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YJrBcfzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DgZUVY6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rPjl0aQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vjjtPyGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d1U4BA1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NIZA06t9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bArEzAMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FjHuxWmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HUE8tCouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P7d4ysqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kXOe0kJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xQchNVB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EOi6LanyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5CLM5Jq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nKcPoiNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nJH5bMpUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YEjZZP6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mOfIFReHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IwIDpPptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8MtyY8VJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CalE126TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jzaI7ZB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JRGV5V8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BuglhqmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 33E9zzD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZHtQSP2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yqY5fshcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lra3lNswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zClOLIIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qgCxqlU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4vVbi39FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qGBnDYsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aP87BuLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 63OdstFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h7DEIq81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KocA5dNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jd0hsjHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FLSD05gfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oNcBKiNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HWiUAmHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yyzssJZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XGhXKDl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CYiyoL1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cmPN2LlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UWH3o4wVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sspqWKdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UMNGY8Z0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func olccAEePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jCFjJEKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UHmUz2IWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U9YYueQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KhjInJHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7g6eaYAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k9xSWG0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qvZ2562VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WRdYsRWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W16hItbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 58QGeXyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sku3Ivd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4aAzWeGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 09MB1jfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zLesPOG3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XEqD5coCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4XEqRk2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qEpNwMcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vwwCrxmBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SXVKmZs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uENDswboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d2LdkbCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func feRonTf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o5MZgSvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func owo4qP5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ARWVFqssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UkXumPUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ObkPFy5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6mQeQktDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GVoSZEdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UwhPFbmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ea9dQvYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZuWbF9VeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XxNKhKCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NjyT5EkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func URGfC2qrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 96TW5m2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5q8JvYujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1cjEg4FgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vQ8OyxPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oK2pKJInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yx1bNqDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UIPpjFzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IXwr5OkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VlEOo7JvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2XALOQ2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZRcL9wm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ISpTah8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aCUBHMTjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XzZYEnmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vZ0WIiPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ostfAldMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mdqd2ASgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zZv9FW3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jMxElLZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nDdOSnbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bVKDMkvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9uq11Z8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oyNqh8q4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CHWLdp1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5F6XPUpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ih4B0bykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tbakl9rLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TX8bXxdHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func psxO8079Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jwBIcMfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LIgveMCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bBPF1fVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pTFI9khiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HTaIL6MmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CIaxRvrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EVtOlRXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tCLw2m5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q8Rt4Df1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QE2Kdek4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ra9RXINKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func idKSRBCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PvVYu6EtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func clLLQKUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6iPN67CVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6PMcngPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rugZQHPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yv7coiGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KdTa9d5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F3SVWwAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kegSy4ONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o8bYLqPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ILLHENt4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7GqAYlbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func leUFzS7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kzbJfAh7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f3YKQdxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KkmoTB1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 16wIgiC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cYeeNu77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rX5hS4G1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bg19WGhAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FFwIID6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YnsXwrPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ODwYRvXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x6HJFrLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zea1ByecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z1mo2bR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Es6mmw1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BUAxl3aLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WNogf0sgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eISRqxvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SkbJRhIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G6zR3FBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FirpApZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ylOWsD53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c8YlqrJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 73BBh3ouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gqfd4DotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func StJzWg3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AOL7anZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Apq2sDyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CRAvqKoXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b8YWxSanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MdRLjpkHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m69y5HSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ueVJzIx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JperuWuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z0hepy5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7GLtWrucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F0jGkGL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w02RWuBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vJ5Y3DxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eAs6UJv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kn1vObxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M2inQJQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r3Xzc9HzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Eg3nbkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SsURtXmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LdefZi3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hc1zLkH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NaXAzrDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UWSrFIfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WzpQRfLsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gf9VT7elWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nKxmp179Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9UICy7djWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SL7OeFMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YrOwA6MBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ONmKTZKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NJV4o03HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dE9YAcGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AJQ9kB3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 26ENQVpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qqodyfwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5bfRsCcYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aMIiqFXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 49vdOhcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8SWd8U5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yrUQ0UN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4cITIqkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WoBfXaV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JHa4IUiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZCD3EtiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P73BYU3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PoeF5mU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 13CjARdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ie1PDGNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xt1yhYLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PKla2SH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ABCRd790Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ny1kjbTNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NizA5I7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eVKyxYf1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x7Hd1Q2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8uxIPNq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WTGVOcJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I0XGPwd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Q2qUv6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x0Nqy2cjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rMhXSQdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hmA6mI5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FjVUDRLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9dTid5zpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RJTPahiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZEbGu8w8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DgpOLbeYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tSTAqWfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7771eznOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MYh5ihFJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X5Kt5YlnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jnOsS3I5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KN59oiSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fvbnHgdLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b2nLWYmMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S879Yqc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qURXM97nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AomzkwiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FCCmf8XdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7FBmRPVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mh2vRXjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Kj7LJSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ywQviUG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PlMtZta3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y0118oXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IluJf2FmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CmqxbiJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 52TeHvszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hR82BlQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rnAxNa3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F5WNo1cWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHyE0hFJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ftReje94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f64YHxloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tv5Lbga5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X0vPxMO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8hlbduYgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func la0Zx9pvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gQmTDabvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iLt6EBYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 51KJzmbUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7M1x772HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tsrgaeKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tt68ibl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FWma6bE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AYtI6LA4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zdnb9RIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iBSKezhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VDagJncdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Li99GTVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dIHUDlmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vIsOKY50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pP1fUQKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6nrE6nR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yc4mM2PiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y92VbbCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7XIeHRYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S4h9hPRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QnmDMf5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RxEcneYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7gX4kU1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BfrSgDCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i52xVAINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I7wNwjx3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BM90KLrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3u7rhP9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3zv0NvOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YQGwb48QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z1wEKwYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0pTYtOfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func axdVXbAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SCFZf6K0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SzPONWNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PdvzHtkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4N0MXfymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9e8g4BvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4zbZp3zwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yj1eGJJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eVA0OtLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ii8JPrghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BUNsL5YJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RHlr6vbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J0VDpOlMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xo1X1e8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hwyKfTGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qkqpI9ZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BMON7NbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vjSeBKCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6wkixMnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0naC8716Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ehFyjj6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2zPPARbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Aqog585BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yvxU71EaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8v9TGLCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tjgq9fOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ewyoL0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tuzonoo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fmmtZRnjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FgoGDtqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fXxah1dVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wcn8Zh7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EofIPpfsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 03rH0Uq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func umEeuRxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A9Tw72XvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Buz8MrvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mzm7NAFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ezzCOXIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cF52u6elWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wyDO0oV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 82zSMYujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m6yHblg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZxLBGGIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ogXBJPBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9EtNtknIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V8QUj2jOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 11avdeh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o3EeKWemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xMMFqP9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ToUgVFQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZvZljdIIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U4hwRSLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pjfj8GrlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7xpFGRt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tm7eYPoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HFS4tCWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z5lyPFymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lgrOvNxDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cDBi9PTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M9K8XXzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mC8LoTaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vBcwuKetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func McH4E2NTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GEU0hHWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2L06EPdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JeYyWi3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GukhpDmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ArhaoSO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xR40HJQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u9EycHGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LSPZtOduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bzfMNDGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sb06ZCQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AGPVvIhjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M4qT3As5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jClquahQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XlWxmVxcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hqtu1Nj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aLkVxTdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iny6z27EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EUS2hW81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s9qqk0ZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IL4us86LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ggKirDWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ElxWPZs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XFiN8GW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OoCt3dCaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d5njFR34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func udKL1VHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1CMk0Z2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GyeIuSJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l2Od1HJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s1fhgHbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RkSbxiX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y0KvEjNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wo83elKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yJqpiwFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n5O2X2liWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DwJk2TjBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FNWBkNzDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wvqkv0ElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hsnOVO7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gGUtQvhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iE1VyrewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UnfHlA2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7iTaYzdgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func afzYaCbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q2g8JF5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CzjotNd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 98dfQgWvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DsVB3G0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dGkXvs7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZCBv6QpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func du00IQudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D3eycFf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IMIsJqiiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func htqpiP5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wrikTQoCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 41UzoBbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rhoqNVbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func soj06FneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fGtUJhAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2XOiuCr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AXMiihWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v6Mar1a2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5o49TZp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8IHWdsCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BqlJtTGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nuG1IgcZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w4RVJVNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YywqinXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YLWP7lb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LDYqHb8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NxqaW5udWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z8vkboagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CN5ORpNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nN37FFutWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lLN2CjY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3QK7fLIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1LGSCtSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VTrPd8eJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FTMQo29hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iXdAh3YNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CakqTAcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EkfhjVD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1wUt81xcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dEigrN7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i5VIe7XhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PQaKGMadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VYd4HE5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XzoRQWFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4qLcVoGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1AFR7FSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 896nJ1ukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZPST4EbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QBUJOIZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lpCGOBSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mr4U6IdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OI8bjRKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nNHkCIaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VNXOh8dgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sMW6CRSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dN97FuOGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Jz1ANeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P4HOuNeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BJJSw3QUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y8wDnRomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WvvY9DL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5xqoSeCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UNHT88ZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YfX9Yc0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pqh5WiZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C8hYO4AwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func beMZ5MK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ynUM0UdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func urBchxyFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MzJfKCuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rpFv3HCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1HqEQu2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U4ZjMflxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fshBKRwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z5CWfQ9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JFIoHeCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Q6hopEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0gABItNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U9mF6ydPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 90rNo7VPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RbMDQD6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wf9ooqs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GVNMFa4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VVw5GjWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u5y8BaViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZaVGvAppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B5Jdx8T1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9nbpQWRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lng3tDbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dBfKS3F8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iig41KV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 27xOsFECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MN9Tn3HJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ugC1nRYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PuhsqWGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WCOmdJQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zYbQjelxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n5zMh999Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oIPhITGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dsoz7IsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MtqSaxUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m0PSTeVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GZNDi8d1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FcuPtqY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func URJxJruDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yKJ1rQQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QnLiRCThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TrO7wreDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pD3zxzb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yg9UsbTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0xIPECZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func syPQR32UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZpnVztOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ALgJerd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6eUn5v1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xe3IV6W5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iklTLJyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func saZyi9ZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FnvmlYacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f1dpfV1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FfdhmpzhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IYcehijuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 28NNTcHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OG2gyOpuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YC50gmuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O8OZGSFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lj1QVGCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cuH9DPYdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mc2LanxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8qA76wWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K3yYlNlKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S2ZLihRHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kWCI1qpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tyjLsw5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func khcGqYbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8HBzkKudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func adW2h7oCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1CQ2F9U9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4YVDe3phWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EZeYTABAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kC5X1XKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YXkFkU6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vVk64lFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4RKCsznTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fKlhAQPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OOVrZzffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wIvdoGTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MCdNM943Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G5qeSaeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OKFVK8SuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RAmNKaXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mys0b6sxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RNbA5khyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KIfNHujLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 72bu21a0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x6CHM1yJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z6xooUX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kaTB17oPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9pf9y5awWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CzNmFbrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func APa265nLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PdiskloCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cbfHAq2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HeSGZrIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T2VBD45pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jyGB33FwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func szmH0vdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jU1PdQNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jSCtrkX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2LwnehGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZsmomCDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func COBmHD0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7604PxE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GU5WCOVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CtTRPXCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ncfN8ffVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VlrmNItOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DqawaIJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2kEezHSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SSy5i2CQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Y8Ro9MWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XkN6w9AHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vu1CYpFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T51GBKPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ijeu84EcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xkLqYSDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XQ4tXsZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dEdEGZi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cStRFa4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BYuxVBaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y9iUIXMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uZYcERnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e151xMDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ytUnL66UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 88Qt4lWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yh9Fls6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r7XGFNl9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tRKsej7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yH8R3w0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iWkPKRDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ykbKr1QSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eHhTqc44Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 93mcK5ctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UVbgx5QVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s5suApXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M6KoYMaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wxPioQL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Icn2lkySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2hxu79kYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nq76e5xKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8RFYeEJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mlKSuau1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RunrELUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func asKdXdPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wJiBhlMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B38hjzxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TB91n9LiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V4d5BYaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1qw8j91kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7vD43YvfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CPcrdmhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WdrSHXS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QS2Q2lBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HqRN95f1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ogZtlHIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6wjycI2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IyOfLRnsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aL7R0FitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2XJwABbXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DjdBFeRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0mJa0Op1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WckkkhP0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q7yKYPItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zZY6XLO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7PoUgOEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R3qO6YFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F4PPPSzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VWR71KBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F4Vv7gZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func COzLDPBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hbI2nIE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mxbb90twWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AJHdYmEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sXdNtCCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ozCWXMPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TQEnCQboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GJxlzH58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KX7j2BmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TIh9MFhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1Ewrw31IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VjHghR6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func chftkLLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c2jLZRVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XDTJGW9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VVmWj5lUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eEEWIQkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sZvSZ1iXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4DL5K5MvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jFsybuoEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HIwVzvznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qrMxV2UBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iZap59i8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6HFRTwrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g68s6tsLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tBx5DvwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mNsudgR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pGB67dj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bj6ShQ1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3cy3LLeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H5Da9anNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P4TvIi6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oqOyGWeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tfZoBxCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lLkdWeu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0agpv7ucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TiCOfPUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 27qDhV6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YbiyAa9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L7CRcUJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FCKNHIhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QcBenn7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TTciPmR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GdNWon7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SqWIj3hqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GcVxkLTTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aRDcOUS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p7AdXD9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qzfNc92OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZYtqwHwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GPxPfpaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gE4M1o1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GBgQqWwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LRG6qlyOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iliqpGYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wxL98UhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5sIXoRWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B9fbq8nJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vpVUY6vjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func scb9nKUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cgkqQu9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vEBKS0GSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xC2XRkc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jKcyOByRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tlc6MnIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uJHBgrEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G1HvnuUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aoFshY59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T1khR5IBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tMkOROYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V62MGCaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OabUIpVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hmBji1GTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PXKTWpNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z3XNbjEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D1wlcUkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UzkMqVvdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XDNguCcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WYakpHy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A805LPcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D7w4PLZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9oKlnJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ftahHB1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4K327K1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PGvv4un8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WBAcUPn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q8MEJrH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gYK4qhLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PfHsFMwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YOosRmP0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NOR5AXTTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3mbC1LAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1gD2Jlv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a2jSN2HgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5iIZuxaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wJ264J6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WWbzSIquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sYcIC8TvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EjxAmKitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9GPYzgG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KLduA8LjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ElMF7dhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W1zuhSiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MTGor2VJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GOG1lAubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func noMwLhwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N9sypPTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j6xwqsxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NAj8mX3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qwjZDqsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F7EeyQqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pYXVzsbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P27dg2YBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L1HoSaBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func REWcAqkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NKsKl6e7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qVhDcoC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lq0ceD59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o7gCmMnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NsjDWTL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8gZXNHVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cpfKsLV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TktCuhCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4SDoHzaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CstNoiHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YEeJhz6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qTmbyLCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c2oNiSiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VGixfA1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func icHFlsZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func coQpL1RXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IaOsIh6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 172rbjjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KeKCRnj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v1bCzoDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fbfntSAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func crze5qCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kCdfv7JgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TiEq05MEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uWt8XFT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hp4GYUwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2is021IpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H4CWKdcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nsLRSC9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bzdcqKVxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yf2BfqH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tgMbKBZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9lSjTQfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AvqvbvmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jQihY0zyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N8ajhrWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vOL60oSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ray4nhfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PZQ8N5zyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kSeVdjTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3nfiyhUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZRkiUdsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7nfXqn7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hYxw3GnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eR4mNijeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 66WFwAMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YZjtUuvMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KzsM8dbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JsF4LNKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PLMEIkveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jxylRlqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FxhDEpCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W904y538Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zIwWToX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8OdEouDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GW9YE229Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2b41tbhIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YpkSPVTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5VCfvSyjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JCA5oec5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iCeiuGnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B7lCXKeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XowmT0P5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qDtqM5lJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 74WT44eTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VhYNTCoCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2QLDT01lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zwp2y8dGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W58LDh4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b91egsq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oYlaaMQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JUJqmxFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kgIR7jsQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M7U3pC2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ztxcH7bUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OGn85tgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lfsHCWFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DyMkCviNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wMyi8nB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kGB7xp8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XUD0s09yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hXGB30xIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pXUmCSrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZJvpn9iFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fr3KXUuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NneEyQCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9lkpcgIIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FnPYCbZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GET8eegmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rNsoGax9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CNtL6rWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IcDic47OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7zEjoZBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mEe7NuaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OVyh8MFYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 36NposhPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TqvdMayUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TeUkJccIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7kvTZfKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hpcF2A40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1L9lZMC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oOQSpNWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vzDxGDYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HUlaBJDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 97qO5XmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dZIpUvDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MgwDc0IIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yn2Cp8b9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MIq5iCLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AqV0zhR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zn5Q2eYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func agtGz2eSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QFc9YFmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BKf9fHSGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KtD3t4bTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9NIQGeeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MG5rpVHpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 53hJKCSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5VyIvKCeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y1LOKZyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UnfMd5ceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6dCl8EnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ScBteLyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oSM5gO7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TdWEgnF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tSSKGkSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9FLkRweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ffOP3FIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0seMwDpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C4RZEpGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t2SOZae0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tvthAVSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tRPuESacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xCHQEakGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GrJeEKW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pAOFVBK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vlzPc7UOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uirFzNelWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BI9KYJZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1WiiXhReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q5c8j2eTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZUCnXMEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nz1sLi75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dPJQAx8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0vWbF0qnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6OwwzJQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bzOupUG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x45Nd9CXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CXdCqe0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YuxWNQlzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GGalttVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y0LerWp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q0c2GRaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IIzvXQybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9HQg7a22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T6KkyvhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WtLDeKPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9NHj4QRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func szXFlpzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BZVyndGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6HelWbhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V3WCllmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lU4s5nMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t3h9y9LkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LCvjtcFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c3EJL5ObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1lKDn9rMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iG8DexTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GmdlPrNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i314LifXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kBiTZEHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pp2IH0HPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 32T6DnbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zfSmHC7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0RYYFY4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7cv62OIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8GishjweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uFI6PxxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8YRnwV6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A6Ar4b9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dsEpOFlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oglLj23EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RWQPqmSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Smfr3GgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KvHGKCuRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CPy3bUxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0MUpQrKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2MVc9eNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jpdemg2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qbJLMlbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yeE0un36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MVdWltW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6GhcPgjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NClouiOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lfFrszs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3pIfLbNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 91e3wJbUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DqvdKenVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QpCl6veuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MsZ1c4XhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3mlJDeXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j8NAFtmyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cDUyfiHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v70f79LhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zRMqICvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q0Lqh4XLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ni6cNgp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GPrdjapeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pn369wFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O9bFwzzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hpK3E72jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ki0DyG1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vIHxfiDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MFZHFMQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K2HYulZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CbXuLe1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CEHuNfecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B3ffrC7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Bz7RYZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eLxiE0BYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p1XqYwb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uVZAdqagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 05v2ZhUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ns3FJQivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ybaAqRM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bL12kU76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uM4BibvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QsPZY6LtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VAhKuD96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gMDyCf1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q6u7tV4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KZVCnGlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NiQTVMvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func furMYostWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ehrQB0kbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cv07ERYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gpnPuisOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mXHk0sMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dUYP0xbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wL048Zw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fyjCrumGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vhJmB3CoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZvJdfB6kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kR2BKvHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0HNhxquyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w0RSKNFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mhg52IAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1UDT1UZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kkvWtrplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sBPCqEfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0JFmx5omWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b69yLlmHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NwfNsplKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HRox6YOlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5KnSoK9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QMyE3NhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yiDoDUHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GhU8sQBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xjuzZYgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1TDnHiaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MnVOir5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VfUggZ43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vW2fy1oZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wisz09eIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jmHEOFrJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RSXwNE1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fTF63nVgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ozh8hKfsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ozcmpoU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BSMbZpzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 00nDUh8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H1rEy2TtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LrLdoVK7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oDi4zdXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DYFxjlbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DaB8UQzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BtzPahhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8yeuOWYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Ij9kyLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IjInglG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QOYQQhbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5fNS65vdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bu9U1iTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xeqTPtPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S5652mP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zqhIidTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mTLzLXKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func doaJAA0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BSRI3udSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7WNF1UZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t1BtrvzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qtjVHRVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fVBpc5t7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xekAW6jjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func klJ2WplwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EZeZ56QEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KlGoNF3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dl8wzgAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JQi2CEb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0MymVeBXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YknoOmL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J4WlvqPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HKHO6ayyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cJXXHJQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wmU4Y2ScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func djNfOJ3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1xtjxlT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i0vxBGToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KKdTRBoTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yNCAGJynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PLc2la7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func crbesveUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func knueRi9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jc0H9sZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I3xwY8y7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9dVNqzsYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3E9sFOdgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CsXyZBvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N1WXQMAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zLBIpcq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LgfzOZ3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wLNhLLWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 14cEDYzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yWXMJdehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7fScWVVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fp4BJprTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tXK3RJNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UvakUZwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g4XuszhUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KFrqs3VQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vMW2Wn76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MEpWaMM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C3B2kChEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CDYy4JlmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1f1pCC33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HOVQkdtCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qp8yFqsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ArscRaYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rDMOIooJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6gqotUZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 526kOreMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oxDqR6P3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iuTA7Q4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a2qELAVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hHOnmFYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KRHN4cS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q12uA3h2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yAvEOoDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kv6cVf6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RNvTwnahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func higtH8PJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QNcNxA4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J1O4FHlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oaKT9qm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pK1GThqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yvJZwSGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nF4WpO3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DcPWpXFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CPeYtvJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hfjxwwxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mjMQT4t4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func im2gpQHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D8UmX3BLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O7c4PorBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9f61cD1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fb0xUnTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UInIpm7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dA2z1ZROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hNqYuAkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f38ZPFMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f288gxDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6pZ2AkeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oe1KzDsiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y9CyE5KFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rRaiGvMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func it4gyLEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i20TvXMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qYOSkM3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SfsVbAdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fs0WJ6M7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PWJGc6IPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0GGpBqIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func isrergz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qdCNJtjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AalDBX5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func brI42Wm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qb83LFMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QIteItwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0sEMq8AfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FDFjH4wCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TUynjDs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E5BVDxCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vmsrSd7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1iRgGngZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZnLWXJrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YanWqOobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0h8YRCuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JVrX7SOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u0FsZgEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jmhV0azPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j6lbnQqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u0Q7GjOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

