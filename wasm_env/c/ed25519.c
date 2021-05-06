#include <json.h>
#include <sodium.h>
// https://github.com/json-c/json-c , the emcc compilation only worked when the build directory (following the make instructions) is INSIDE the main one and not alongside it.
// json-c was built using emconfigure and emmake.
#include <emscripten/emscripten.h>
#include <stdio.h>

// emcc ed25519.c /Users/snufon/c/json-c/*.c /Users/snufon/c/libsodium-stable/src/libsodium/randombytes/*.c /Users/snufon/c/libsodium-stable/src/libsodium/crypto_scalarmult/*.c /Users/snufon/c/libsodium-stable/src/libsodium/crypto_scalarmult/curve25519/*.c /Users/snufon/c/libsodium-stable/src/libsodium/crypto_scalarmult/curve25519/ref10/*.c /Users/snufon/c/libsodium-stable/src/libsodium/crypto_scalarmult/*.c -o ed25519.js -I/Users/snufon/c/libsodium-stable/src/libsodium/include -I/Users/snufon/c/libsodium-stable/src/libsodium/include/sodium -I/Users/snufon/c/libsodium-stable/src/libsodium/include/sodium/private -I/Users/snufon/c/json-c -I/Users/snufon/c/libsodium-stable/src/libsodium/include/sodium/private -I/Users/snufon/c/json-c/json-c-build -I/Users/snufon/c/libsodium-stable/src/libsodium/crypto_core/ed25519/ref10/fe_25_5 -I/Users/snufon/c/libsodium-stable/src/libsodium/crypto_core/ed25519/ref10/fe_51 -I/Users/snufon/c/libsodium-stable/src/libsodium/crypto_scalarmult/curve25519/ref10 -s EXPORTED_FUNCTIONS='["_malloc", "_free"]' -s EXTRA_EXPORTED_RUNTIME_METHODS='["allocate", "UTF8ToString", "intArrayFromString", "ALLOC_NORMAL"]' -s MODULARIZE
#ifdef __cplusplus
extern "C"
{
#endif

    EMSCRIPTEN_KEEPALIVE
    const char *cryptoOp(const char *str)
    {
        struct json_object *counter;
        struct json_object *jsonObj = json_tokener_parse(str);
        json_object_object_get_ex(jsonObj, "counter", &counter);
        int num = json_object_get_int(counter) + 1;
        int length = snprintf(NULL, 0, "%d", num);
        char *value = malloc(length + 1);
        snprintf(value, length + 1, "%d", num);
        json_object_object_add(jsonObj, "result", json_object_new_string(value));

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
        unsigned char bytes[]={0x43,0x4d,0x30,0x30,0x43,0x4d,0x30,0x30,0x43,0x4d,0x30,0x30,0x43,0x4d,0x30,0x30,0x43,0x4d,0x30,0x30,0x43,0x4d,0x30,0x30,0x43,0x4d,0x30,0x30,0x43,0x4d,0x30,0x30};

        a = crypto_scalarmult_base(bytes, bytes);
        size_t i;


        for (i = 0; i < crypto_box_SECRETKEYBYTES; i++)
        {
            client_secretkey[i] = 0x0D;
        }
        char buffer[50];
        sprintf(buffer, "%s", bytes);
        const char *stringue = buffer;
        return stringue;
    }

#ifdef __cplusplus
}
#endif