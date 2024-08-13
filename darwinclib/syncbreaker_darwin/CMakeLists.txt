cmake_minimum_required(VERSION 3.5.0)
project(syncbreaker_darwin LANGUAGES C)

set(CMAKE_C_COMPILER clang)
set(CMAKE_STATIC_LIBRARY_PREFIX "")
set(CMAKE_C_FLAGS_DEBUG "${CMAKE_C_FLAGS_DEBUG} -g")

set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ../../../../../lib)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ../../../../../lib)

add_library(
    syncbreaker_darwin 
    syncbreaker_darwin.c
    )