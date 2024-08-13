cmake_minimum_required(VERSION 3.5.0)
project(urlcatcher LANGUAGES C)

set(CMAKE_C_COMPILER clang)
set(CMAKE_C_FLAGS_DEBUG "${CMAKE_C_FLAGS_DEBUG} -g")

#add_library(urlcatcher urlcatcher.c)
add_executable(urlcatcher urlcatcher.c)

target_link_libraries(urlcatcher PRIVATE "-framework CoreFoundation")
target_link_libraries(urlcatcher PRIVATE "-framework Carbon")
target_link_libraries(urlcatcher PRIVATE "-framework ApplicationServices")
target_link_libraries(urlcatcher PRIVATE "-framework Foundation")







