#include <json.h>
#include <emscripten/emscripten.h>
#include <string.h>
//#include <stdio.h>
//#include <stdlib.h>


//emcc increaseCounter.c -o increaseCounter.js -I/usr/local/include/json-c -s  EXPORTED_RUNTIME_METHODS=['ccall'] -s LINKABLE=1 -s EXPORT_ALL=1

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

#ifdef __cplusplus
}
#endif

/*int main()
{
    const char *str = "{ \"counter\" : 4, \"contractName\" : \"increaseCounter\", \"contractLanguage\" : \"go\",}";
    printf("%s\n", increaseCounter(str));
}*/