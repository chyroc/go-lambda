package internal

var intSize = 32 << (^uint(0) >> 63) // 32 or 64
