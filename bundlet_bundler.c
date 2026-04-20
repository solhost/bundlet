#include <stdlib.h>

int main() {
  system("gcc worker.c -o worker");
  system("go build bundlet.go");
  return 0;
}
