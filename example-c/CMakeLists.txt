
cmake_minimum_required(VERSION 3.18)

project(
  GoIntegrationC
  VERSION 1.0
  LANGUAGES C)


include_directories(sources)
link_directories(sources)

# Library
add_subdirectory(sources)

set_target_properties(${PROJECT_NAME} PROPERTIES PUBLIC_HEADER sources/simple_lib.h)
install(TARGETS ${PROJECT_NAME} PUBLIC_HEADER DESTINATION lib)
SET(CMAKE_INSTALL_PREFIX .)

# Tests
enable_testing()
add_subdirectory(tests)