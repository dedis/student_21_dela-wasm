// emcc ed25519.c /Users/snufon/c/json-c/*.c /Users/snufon/libsodium/sodium/utils.c /Users/snufon/libsodium/randombytes/*.c /Users/snufon/libsodium/crypto_scalarmult/curve25519/*.c /Users/snufon/libsodium/crypto_scalarmult/curve25519/ref10/*.c /Users/snufon/libsodium/crypto_core/ed25519/ref10/*.c /Users/snufon/libsodium/crypto_core/ed25519/*.c  -o ed25519.js -I/Users/snufon/libsodium/include -I/Users/snufon/libsodium/include/sodium -I/Users/snufon/libsodium/include/sodium/private -I/Users/snufon/c/json-c -I/Users/snufon/libsodium/include/sodium/private -I/Users/snufon/c/json-c/json-c-build -I/Users/snufon/libsodium/crypto_core/ed25519/ref10/fe_25_5 -I/Users/snufon/libsodium/crypto_core/ed25519/ref10/fe_51 -I/Users/snufon/libsodium/crypto_scalarmult/curve25519/ref10 -s EXPORTED_FUNCTIONS='["_malloc", "_free"]' -s EXPORTED_RUNTIME_METHODS='["allocate", "UTF8ToString", "intArrayFromString", "ALLOC_NORMAL"]' -s MODULARIZE -s ALLOW_MEMORY_GROWTH=1

// /Users/snufon/deps/b64/*.c -I/Users/snufon/deps/b64
#include <json.h>
#include <sodium.h>
// https://github.com/json-c/json-c , the emcc compilation only worked when the build directory (following the make instructions) is INSIDE the main one and not alongside it.
// json-c was built using emconfigure and emmake.
#include <emscripten/emscripten.h>
#include <stdio.h>
#include <stdlib.h>
//#include "b64.h"
#ifdef __cplusplus
extern "C"
{
#endif

    // The various following functions are needed to avoid an inexplicable bug with Emscripten. Mostly redefining functions that used to exist elsewhere in libsodium.

    unsigned char *rand_bytes(size_t num_bytes)
    {
        unsigned char *stream = malloc(num_bytes);
        size_t i;

        for (i = 0; i < num_bytes; i++)
        {
            stream[i] = rand();
        }

        return stream;
    }

    int
    custom_is_zero(const unsigned char *n, const size_t nlen)
    {
        size_t i;
        volatile unsigned char d = 0U;

        for (i = 0U; i < nlen; i++)
        {
            d |= n[i];
        }
        return 1 & ((d - 1) >> 8);
    }

    int
    custom_is_canonical(const unsigned char s[32])
    {
        /* 2^252+27742317777372353535851937790883648493 */
        static const unsigned char L[32] = {
            0xed, 0xd3, 0xf5, 0x5c, 0x1a, 0x63, 0x12, 0x58, 0xd6, 0x9c, 0xf7,
            0xa2, 0xde, 0xf9, 0xde, 0x14, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10};
        unsigned char c = 0;
        unsigned char n = 1;
        unsigned int i = 32;

        do
        {
            i--;
            c |= ((s[i] - L[i]) >> 8) & n;
            n &= ((s[i] ^ L[i]) - 1) >> 8;
        } while (i != 0);

        return (c != 0);
    }

    unsigned char *
    rand_scalar(size_t num_bytes)
    {
        unsigned char *r = malloc(num_bytes);
        do
        {
            r = rand_bytes(crypto_core_ed25519_SCALARBYTES);
            r[crypto_core_ed25519_SCALARBYTES - 1] &= 0x1f;
        } while (custom_is_canonical(r) == 0 ||
                 custom_is_zero(r, crypto_core_ed25519_SCALARBYTES));
        return r;
    }


    EMSCRIPTEN_KEEPALIVE
    const char *cryptoOp(const char *str)
    {
        /* struct json_object *point1;
        struct json_object *point2;
        struct json_object *scalar;
        json_object_object_get_ex(jsonObj, "point1", &point1);
        json_object_object_get_ex(jsonObj, "point2", &point2);
        json_object_object_get_ex(jsonObj, "scalar", &scalar);
        const char *point1s = json_object_get_string(point1); */
        struct json_object *jsonObj = json_tokener_parse(str);
        json_object_object_add(jsonObj, "Accepted", json_object_new_string("true"));

        unsigned char point[crypto_box_PUBLICKEYBYTES];
        unsigned char pointf[crypto_box_PUBLICKEYBYTES];
        unsigned char *x = rand_bytes(32U); // {0x46, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30};
        unsigned char *y = rand_bytes(32U);
        //unsigned char *scalar = rand_scalar(32U);
        unsigned char px[crypto_core_ed25519_BYTES];
        crypto_core_ed25519_from_uniform(px, x);
        unsigned char py[crypto_core_ed25519_BYTES];
        crypto_core_ed25519_from_uniform(py, y);

        // mul 10k : 34
        // add 10k : 593

        int a;

        for (int i = 0; i < 1; ++i)
        {
            a = crypto_core_ed25519_add(point, px, py);
        }
        int length = snprintf(NULL, 0, "%d", a);
        char *value = malloc(length + 1);
        snprintf(value, length + 1, "%d", a);
        json_object_object_add(jsonObj, "result", json_object_new_string(value));

        /* unsigned char *client_secretkey = malloc(crypto_box_SECRETKEYBYTES);
        unsigned char bytes[] = {0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30, 0x43, 0x4d, 0x30, 0x30};

        a = crypto_scalarmult_curve25519(bytes, bytes, bytes);
        int b = crypto_core_ed25519_add(bytes, bytes, bytes);

        for (size_t i = 0; i < crypto_box_SECRETKEYBYTES; i++)
        {   
            client_secretkey[i] = 0x0D;
        }
        char buffer[50];
        //sprintf(buffer, "%s", point1s);
        const char *stringue = buffer;
        return stringue; */
        return json_object_to_json_string(jsonObj);
    }

#ifdef __cplusplus
}
#endif