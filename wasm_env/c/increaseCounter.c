#include <json.h> // https://github.com/json-c/json-c
#include <emscripten/emscripten.h>
#include <stdio.h>

// emcc increaseCounter.c /Users/snufon/c/json-c/*.c -o increaseCounterC.js -I/Users/snufon/c/json-c -I/Users/snufon/c/json-c/json-c-build -s EXPORTED_RUNTIME_METHODS='["ccall","cwrap"]' -s MODULARIZE
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
        json_object_set_int(counter, json_object_get_int(counter) + 1);
        return json_object_to_json_string(jsonObj);
    }

    EMSCRIPTEN_KEEPALIVE
    int increaseCounterTest()
    {
        const char * str = increaseCounter("{ \"counter\" : 4, \"contractName\" : \"increaseCounter\", \"contractLanguage\" : \"go\",}");
        struct json_object *jsonObj = json_tokener_parse(str);
        struct json_object *counter;
        json_object_object_get_ex(jsonObj, "counter", &counter);
        return json_object_get_int(counter);
    }

#ifdef __cplusplus
}
#endif

/*int main()
{
    const char *str = "{ \"counter\" : 4, \"contractName\" : \"increaseCounter\", \"contractLanguage\" : \"go\",}";
    printf("%s\n", increaseCounter(str));
}*/