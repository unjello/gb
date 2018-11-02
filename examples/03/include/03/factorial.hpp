#pragma once

template <typename T> T factorial(T number) {
  return number <= 1 ? number : factorial(number - 1) * number;
}
