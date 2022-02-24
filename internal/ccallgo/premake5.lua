workspace "ccallgo"
    configurations { "Debug", "Release" }

project "innnerlib"
    kind "SharedLib"
    files {"./libc/**.c", "./libc/**.h"}

project "ccallgo"
    kind "ConsoleApp"
    language "C"
    targetdir "bin/%{cfg.buildcfg}"

    files { "ccallgo.c" }

    libdirs { "libs" }
    links { "goforc", "innnerlib"}

    filter "configurations:Debug"
        defines { "DEBUG" }
        symbols "On"

    filter "configurations:Release"
        defines { "NODEBUG" }
        optimize "On"

