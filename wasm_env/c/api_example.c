//api_example.c
#include <stdio.h>
#include <emscripten.h>
#include <json.h>

// emcc api_example.c -o api_example.js -s MODULARIZE -s EXPORTED_RUNTIME_METHODS=['ccall']

#ifdef __cplusplus
extern "C"
{
#endif

/*EMSCRIPTEN_KEEPALIVE
void sayHi() {
  printf("Hi!\n");
}

EMSCRIPTEN_KEEPALIVE
int daysInWeek() {
  return 7;
}*/

#ifdef __cplusplus
}
#endif