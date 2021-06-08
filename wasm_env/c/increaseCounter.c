#include <json.h>
// https://github.com/json-c/json-c , the emcc compilation only worked when the build directory (following the make instructions) is INSIDE the main one and not alongside it.
// json-c was built using emconfigure and emmake.
#include <emscripten/emscripten.h>
#include <stdio.h>
//#include <time.h>

// emcc increaseCounter.c /Users/snufon/c/json-c/*.c -o increaseCounterC.js -I/Users/snufon/c/json-c -I/Users/snufon/c/json-c/json-c-build -s EXPORTED_FUNCTIONS='["_malloc", "_free"]' -s EXPORTED_RUNTIME_METHODS='["allocate", "UTF8ToString", "intArrayFromString", "ALLOC_NORMAL"]' -s MODULARIZE -s ALLOW_MEMORY_GROWTH=1
#ifdef __cplusplus
extern "C"
{
#endif

    EMSCRIPTEN_KEEPALIVE
    const char *increaseCounter(const char *str)
    {
        struct json_object *counter;
        struct json_object *jsonObj = json_tokener_parse(str);
        json_object_object_get_ex(jsonObj, "counter", &counter);
        int num = json_object_get_int(counter) + 1;
        int length = snprintf(NULL, 0, "%d", num);
        char *value = malloc(length + 1);
        snprintf(value, length + 1, "%d", num);
        json_object_object_add(jsonObj, "result", json_object_new_string(value));
        json_object_object_add(jsonObj, "Accepted", json_object_new_string("true"));
        /* int a;
        //srand(time(NULL)); // Initialization, should only be called once.
        for (int i = 1; i < 1000000; ++i)
        {
            a = rand();
            a *= rand();
            printf("%d", a);
        } */
        return json_object_to_json_string(jsonObj);
    }

    /*EMSCRIPTEN_KEEPALIVE
    const char *increaseCounterTest()
    {
        const char *str = increaseCounter("{ \"counter\" : 4, \"contractName\" : \"increaseCounter\", \"contractLanguage\" : \"go\",}");
        struct json_object *jsonObj = json_tokener_parse(str);
        struct json_object *counter;
        json_object_object_get_ex(jsonObj, "counter", &counter);
        return json_object_to_json_string(jsonObj);
    }*/

#ifdef __cplusplus
}
#endif

/*int main()
{
    const char *str = "{ \"counter\" : 4, \"contractName\" : \"increaseCounter\", \"contractLanguage\" : \"go\",}";
    printf("%s\n", increaseCounter(str));
}*/