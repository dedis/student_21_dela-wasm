const fs = require('fs');
const http = require('http');
const url = require('url');
const hostname = '127.0.0.1';
const port = 3000;
const wasm_exec = require("./wasm_exec.js")

// JavaScript glue code which runs the WASM binaries
const increaseCounterC = require('./c/increaseCounterC.js');
const ed25519C_add = require('./c/ed25519.js');
const ed25519C_mul = require('./c/ed25519_mul.js');


//// GO INSTANTIATIONS (& OPTIONAL WARMUPS)

// Increase Counter
const go = new Go();
async function fetchAndInstantiate() {
  var buf = fs.readFileSync('./go/increaseCounter/main.wasm');
  var thing = await WebAssembly.instantiate(buf, go.importObject);
  go.run(thing.instance);
  //increaseCounter('{"counter":50}')
}
fetchAndInstantiate();

// Addition(s) of 2 points on Ed25519
const go2 = new Go();
async function fetchAndInstantiate2() {
  var buf = fs.readFileSync('./go/ed25519/main.wasm');
  var thing = await WebAssembly.instantiate(buf, go2.importObject);
  go2.run(thing.instance);
  //cryptoOp('{"point1":"Q6Fi2A7Ot69+ApLGfdjWyStCM2sHg5NnCzCuRmzm3ic=","point2":"JkKvN3MQYcmQxFGwOtpsD5zSHS5qFYEtM949b+Z3XMc=","scalar":"/koEUcby5r3S3U1t+1IBCyY9USOSKP2SfHEOoc3C/Q4="}')
}
fetchAndInstantiate2();


/* const go3 = new Go();
async function fetchAndInstantiate3() {
  var buf = fs.readFileSync('./go/randFloatMul/randFloatMul.wasm');
  var thing = await WebAssembly.instantiate(buf, go3.importObject);
  go3.run(thing.instance);
  console.log(JSON.stringify(randFloatMul('{"rand1":"3.14159265359","rand2":"3.14159265359"}')))
}
fetchAndInstantiate3(); */

// Scalar multiplication(s) of the generator
const go4 = new Go();
async function fetchAndInstantiate4() {
  var buf = fs.readFileSync('./go/simpleEC/simpleEC.wasm');
  var thing = await WebAssembly.instantiate(buf, go4.importObject);
  go4.run(thing.instance);
  //simpleEC('{"scalar":"/koEUcby5r3S3U1t+1IBCyY9USOSKP2SfHEOoc3C/Q4="}');
}
fetchAndInstantiate4();

// Scalar multiplication(s) of a point on the ed25519 curve
const go5 = new Go();
async function fetchAndInstantiate5() {
  var buf = fs.readFileSync('./go/ed25519_mul/ed25519_mul.wasm');
  var thing = await WebAssembly.instantiate(buf, go5.importObject);
  go5.run(thing.instance);
  //ed25519_mul('{"point1":"Q6Fi2A7Ot69+ApLGfdjWyStCM2sHg5NnCzCuRmzm3ic=", "scalar":"/koEUcby5r3S3U1t+1IBCyY9USOSKP2SfHEOoc3C/Q4="}')
}
fetchAndInstantiate5();

//// C INSTANTIATIONS
increaseCounterC().then((instanceCounter) => {
  ed25519C_add().then((instance_ed25519_add) => {
    ed25519C_mul().then((instance_ed25519_mul) => {

      // Simple Rest API, with help from https://medium.com/bb-tutorials-and-thoughts/how-to-write-simple-nodejs-rest-api-with-core-http-module-dcedd2c1256

      const server = http.createServer((req, res) => {
        res.setHeader('Content-Type', 'application/json;charset=utf-8');
        const size = parseInt(req.headers['content-length'], 10)
        const buffer = Buffer.allocUnsafe(size)
        var pos = 0
        req
          .on('data', (chunk) => {
            const offset = pos + chunk.length
            if (offset > size) {
              reject(413, 'Too Large', res)
              return
            }
            chunk.copy(buffer, pos)
            pos = offset
          })
          .on('end', () => {
            if (pos !== size) {
              reject(400, 'Bad Request', res)
              return
            }
            const data = buffer.toString()
            const jsonObj = JSON.parse(data)

            switch (jsonObj.contractLanguage) {
              
              case "go":
                switch (jsonObj.contractName) {
                  case "increaseCounter":
                    var result = JSON.stringify(increaseCounter(data))
                    res.end(result)
                    break;
                  case "ed25519":
                    var result = JSON.stringify(cryptoOp(data))
                    res.end(result)
                    break;
                  case "ed25519_mul":
                    var result = JSON.stringify(ed25519_mul(data))
                    res.end(result)
                    break;
                  case "simpleEC":
                    var result = JSON.stringify(simpleEC(data))
                    res.end(result)
                    break;
                }
                break;

              case "c":
                switch (jsonObj.contractName) {
                  case "increaseCounter":
                    var ptr = instanceCounter.allocate(instanceCounter.intArrayFromString(data), instanceCounter.ALLOC_NORMAL)
                    var result = instanceCounter.UTF8ToString(instanceCounter._increaseCounter(ptr));
                    instanceCounter._free(ptr);
                    res.end(result)
                    //console.log(result)
                    break;
                  case "ed25519":
                    var ptr = instance_ed25519_add.allocate(instance_ed25519_add.intArrayFromString(data), instance_ed25519_add.ALLOC_NORMAL)
                    var result = instance_ed25519_add.UTF8ToString(instance_ed25519_add._cryptoOp(ptr));
                    instance_ed25519_add._free(ptr);
                    res.end(result)
                    //console.log(result)
                    break;
                  case "ed25519_mul":
                    var ptr = instance_ed25519_mul.allocate(instance_ed25519_mul.intArrayFromString(data), instance_ed25519_mul.ALLOC_NORMAL)
                    var result = instance_ed25519_mul.UTF8ToString(instance_ed25519_mul._cryptoOp(ptr));
                    instance_ed25519_mul._free(ptr);
                    res.end(result)
                    //console.log(result)
                    break;
                }
                break;
            }
          })
      });

      server.listen(port, hostname, () => {
        console.log(`Server running at http://${hostname}:${port}/`);
      });
    });
  });
});
