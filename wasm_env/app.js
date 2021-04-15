const fs = require('fs');
const http = require('http');
const url = require('url');
const hostname = '127.0.0.1';
const port = 3000;
const wasm_exec = require("./wasm_exec.js")
const factory = require('./c/increaseCounterC.js');

const go = new Go();

async function fetchAndInstantiate() {
  var buf = fs.readFileSync('./go/increaseCounter/main.wasm');
  var thing = await WebAssembly.instantiate(buf, go.importObject);
  go.run(thing.instance);
  increaseCounter('{"contractLanguage":"go","contractName":"increaseCounter","counter":50}')
}
fetchAndInstantiate();

const go2 = new Go();

async function fetchAndInstantiate2() {
  var buf = fs.readFileSync('./go/ed25519/main.wasm');
  var thing = await WebAssembly.instantiate(buf, go2.importObject);
  go2.run(thing.instance);
  cryptoOp('{"contractLanguage":"go","contractName":"ed25519","point1":"Q6Fi2A7Ot69+ApLGfdjWyStCM2sHg5NnCzCuRmzm3ic=","point2":"JkKvN3MQYcmQxFGwOtpsD5zSHS5qFYEtM949b+Z3XMc=","scalar":"/koEUcby5r3S3U1t+1IBCyY9USOSKP2SfHEOoc3C/Q4="}')
}
fetchAndInstantiate2();

factory().then((instance) => {
  var ptr = instance.allocate(instance.intArrayFromString("{ \"counter\" : 0}"), instance.ALLOC_NORMAL)
  result = instance.UTF8ToString(instance._increaseCounter(ptr));
  instance._free(ptr);
});

const server = http.createServer((req, res) => {
  res.setHeader('Content-Type', 'application/json;charset=utf-8');
  // partly from https://medium.com/bb-tutorials-and-thoughts/how-to-write-simple-nodejs-rest-api-with-core-http-module-dcedd2c1256

  const size = parseInt(req.headers['content-length'], 10)
  const buffer = Buffer.allocUnsafe(size)
  var pos = 0

  console.log(req.headers)

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
      console.log(data)
      const jsonObj = JSON.parse(data)

      switch (jsonObj.contractLanguage) {
        case "go":
          switch (jsonObj.contractName) {
            case "increaseCounter":
              var result = JSON.stringify(increaseCounter(data))
              res.end(result)
              console.log(result)
              break;
            case "ed25519":
              console.log(cryptoOp(data))
              var result = JSON.stringify(cryptoOp(data))
              res.end(result)
              break;

          }
          break;
        case "c":
          switch (jsonObj.contractName) {
            case "increaseCounter":
              factory().then((instance) => {
                var ptr = instance.allocate(instance.intArrayFromString(data), instance.ALLOC_NORMAL)
                var result = instance.UTF8ToString(instance._increaseCounter(ptr));
                instance._free(ptr);
                res.end(result)
                console.log(result)
              });
              break;
          }
          break;
      }
    })
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});















