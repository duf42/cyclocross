
#
# Integration : 1 step
#
set(TEST_NAME onestep)

add_executable(${TEST_NAME} ${TEST_NAME}.c)
target_link_libraries(${TEST_NAME} PRIVATE ${PROJECT_NAME})

add_test(NAME ${TEST_NAME} COMMAND ${TEST_NAME})
set_tests_properties(${TEST_NAME} PROPERTIES PASS_REGULAR_EXPRESSION "SUCCESS")
