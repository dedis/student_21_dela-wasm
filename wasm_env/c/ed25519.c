// emcc ed25519.c /Users/snufon/c/json-c/*.c /Users/snufon/libsodium/randombytes/*.c /Users/snufon/libsodium/crypto_scalarmult/*.c /Users/snufon/libsodium/crypto_scalarmult/curve25519/*.c /Users/snufon/libsodium/crypto_scalarmult/curve25519/ref10/*.c /Users/snufon/libsodium/crypto_scalarmult/*.c -o ed25519.js -I/Users/snufon/libsodium/include -I/Users/snufon/libsodium/include/sodium -I/Users/snufon/libsodium/include/sodium/private -I/Users/snufon/c/json-c -I/Users/snufon/libsodium/include/sodium/private -I/Users/snufon/c/json-c/json-c-build -I/Users/snufon/libsodium/crypto_core/ed25519/ref10/fe_25_5 -I/Users/snufon/libsodium/crypto_core/ed25519/ref10/fe_51 -I/Users/snufon/libsodium/crypto_scalarmult/curve25519/ref10 -s EXPORTED_FUNCTIONS='["_malloc", "_free"]' -s EXPORTED_RUNTIME_METHODS='["allocate", "UTF8ToString", "intArrayFromString", "ALLOC_NORMAL"]' -s MODULARIZE

// emcc ed25519.c /Users/snufon/deps/b64/*.c /Users/snufon/c/json-c/*.c /Users/snufon/libsodium/sodium/utils.c /Users/snufon/libsodium/randombytes/*.c /Users/snufon/libsodium/crypto_scalarmult/curve25519/*.c /Users/snufon/libsodium/crypto_scalarmult/curve25519/ref10/*.c /Users/snufon/libsodium/crypto_core/ed25519/ref10/*.c /Users/snufon/libsodium/crypto_core/ed25519/*.c  -o ed25519.js -I/Users/snufon/deps/b64 -I/Users/snufon/libsodium/include -I/Users/snufon/libsodium/include/sodium -I/Users/snufon/libsodium/include/sodium/private -I/Users/snufon/c/json-c -I/Users/snufon/libsodium/include/sodium/private -I/Users/snufon/c/json-c/json-c-build -I/Users/snufon/libsodium/crypto_core/ed25519/ref10/fe_25_5 -I/Users/snufon/libsodium/crypto_core/ed25519/ref10/fe_51 -I/Users/snufon/libsodium/crypto_scalarmult/curve25519/ref10 -s EXPORTED_FUNCTIONS='["_malloc", "_free"]' -s EXPORTED_RUNTIME_METHODS='["allocate", "UTF8ToString", "intArrayFromString", "ALLOC_NORMAL"]' -s MODULARIZE

#include <json.h>
#include <sodium.h>
// https://github.com/json-c/json-c , the emcc compilation only worked when the build directory (following the make instructions) is INSIDE the main one and not alongside it.
// json-c was built using emconfigure and emmake.
#include <emscripten/emscripten.h>
#include <stdio.h>
#include "b64.h"
#ifdef __cplusplus
extern "C"
{
#endif

    EMSCRIPTEN_KEEPALIVE
    const char *cryptoOp(const char *str)
    {
        struct json_object *point1;
        struct json_object *point2;
        struct json_object *scalar;
        struct json_object *jsonObj = json_tokener_parse(str);
        json_object_object_get_ex(jsonObj, "point1", &point1);
        json_object_object_get_ex(jsonObj, "point2", &point2);
        json_object_object_get_ex(jsonObj, "scalar", &scalar);
        const char *point1 = json_object_get_string(point1);
        json_object_object_add(jsonObj, "Accepted", json_object_new_string("true"));

        int a;
        for (int i = 1; i < 1; ++i)
        {
            a = 1 + 1;
        }
        unsigned char client_publickey[crypto_box_PUBLICKEYBYTES];
        unsigned char server_publickey[crypto_box_PUBLICKEYBYTES];
        unsigned char server_secretkey[crypto_box_SECRETKEYBYTES];
        unsigned char scalarmult_q_by_client[crypto_scalarmult_BYTES];
        unsigned char scalarmult_q_by_server[crypto_scalarmult_BYTES];
        unsigned char sharedkey_by_client[crypto_generichash_BYTES];
        unsigned char sharedkey_by_server[crypto_generichash_BYTES];
        crypto_generichash_state h;

        unsigned char *client_secretkey = malloc(crypto_box_SECRETKEYBYTES);
        unsigned char bytes[] = {0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30};

        a = crypto_scalarmult_curve25519(bytes, bytes, bytes);
        int b = crypto_core_ed25519_add(bytes, bytes, bytes);

        for (size_t i = 0; i < crypto_box_SECRETKEYBYTES; i++)
        {
            client_secretkey[i] = 0x0D;
        }
        char buffer[50];
        sprintf(buffer, "%d", a);
        const char *stringue = buffer;
        return stringue;
        //return json_object_to_json_string(jsonObj);
    }

#ifdef __cplusplus
}
#endif