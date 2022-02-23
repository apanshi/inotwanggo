workspace "ccallm"
    configurations { "Debug", "Release" }

project "ccallc"
    kind "SharedLib"
    files {"./libc/**.c", "./libc/**.h"}

project "ccallm"
    kind "ConsoleApp"
    language "C"
    targetdir "bin/%{cfg.buildcfg}"

    files { "ccallcm.c" }

    links { "ccallc" }

    filter "configurations:Debug"
        defines { "DEBUG" }
        symbols "On"

    filter "configurations:Release"
        defines { "NODEBUG" }
        optimize "On"

