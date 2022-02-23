workspace "ccallgo"
    configurations { "Debug", "Release" }

project "ccallgo"
    kind "ConsoleApp"
    language "C"
    targetdir "bin/%{cfg.buildcfg}"

    files { "**.c"}

    libdirs { "libs" }
    links { "goforc" }

    filter "configurations:Debug"
        defines { "DEBUG" }
        symbols "On"

    filter "configurations:Release"
        defines { "NODEBUG" }
        optimize "On"

