const fs = require('fs');
const http = require('http');
const url = require('url');
const hostname = '127.0.0.1';
const port = 3000;
const wasm_exec = require("./wasm_exec.js")

const go = new Go();

async function fetchAndInstantiate() {
  const buf = fs.readFileSync('./smart_contract/main.wasm');
  const thing = await WebAssembly.instantiate(buf, go.importObject);
  go.run(thing.instance);
}
fetchAndInstantiate()

const server = http.createServer((req, res) => {
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
      const result = JSON.stringify(increaseCounter(data))

      res.setHeader('Content-Type', 'application/json;charset=utf-8');
      res.end(result)
      console.log(result)
    })
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});














