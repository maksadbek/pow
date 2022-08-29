# Proof-of-work using Hashcash algorithm

## How to build and run

Clone the source code
```
git clone https://github.com/maksadbek/pow.git && cd pow/pow
```

Build

```
make
```

Run server
```
export ADDR=:1313

./build/server
```

Run client
```
export ADDR=:1313
export ID=hashcash@gmail.com

./build/client
```