package texec

import "fmt"

func  Trune() {
    words := []rune{
        0x4e0d, 0x7ba1, 0x4ec0,
        0x4e48, 0x6218, 0x4e89,
        0xff0c, 0x90fd, 0x4e0d,
        0x5e94, 0x8be5, 0x60e9,
        0x6212, 0x6216, 0x5236,
        0x88c1, 0x65e0, 0x8f9c,
        0x5e73, 0x6c11, 0x3002,
        0x8f6f, 0x4ef6, 0x5f00,
        0x53d1, 0x6280, 0x672f,
        0x4e5f, 0x5982, 0x6b64,
        0x3002}
    fmt.Println(string(words))

    // string --> rune --> rune
    raw_en := "Randal"
    raw_cn := "开发，测试。" // 实际长度 6

    fmt.Printf("%v len %v\n", raw_en, len(raw_en))
    for i := 0; i < len(raw_en); i++ {
        fmt.Printf("%x, %c\n", raw_en[i], raw_en[i])
    }
    fmt.Println()

    fmt.Printf("raw_cn %v len %v\n", raw_cn, len(raw_cn))
    for i := 0; i < len(raw_cn); i++ {
        fmt.Printf("%x, %c\n", raw_cn[i], raw_cn[i])
    }

    var world_chinese []rune
    world_chinese = []rune(raw_cn)
    world_chinese_2 := []byte(raw_cn)

    fmt.Println(string(world_chinese))
    fmt.Println(string(world_chinese_2))

    fmt.Println(world_chinese)
    fmt.Println(world_chinese_2)
}
