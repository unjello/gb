#define DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN
#include <03/factorial.hpp>
#include <doctest.h>

TEST_CASE("testing the factorial function") {
  CHECK(factorial(1) == 1);
  CHECK(factorial(2) == 2);
  CHECK(factorial(3) == 6);
  CHECK(factorial(10) == 3628800);
}
