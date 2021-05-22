# student_21_dela-wasm

# Webassembly execution environment for Dela #

Generating WASM binaries can be hard and depend on several factors but is not necessary as the binaries are included along the original high-level contracts in C/C++ and Go.

In the case where you want to compile to WASM yourself, you should open the relevant folder (inside the "wasm_env" folder) containing the smart contract as a seperate workspace. The commands are included as comments at the top of each original contract.

The environment needs to be manually launched by launching the "node app.js" command from the "wasm_env" folder before it can handle POST requests from the Dela framework, which should be opened as a separate workspace.
