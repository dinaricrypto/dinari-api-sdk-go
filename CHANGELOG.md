# Changelog

## 0.1.0-alpha.2 (2025-06-03)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/dinaricrypto/dinari-api-sdk-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### ⚠ BREAKING CHANGES

* **client:** rename resp package
* **client:** improve core function names

### Features

* **api:** api update ([341e362](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/341e3628a4d4072e13a1dc1070611ceba2405f46))
* **api:** api update ([43d6d32](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/43d6d326dfccf8e835f59811924b582438e5df4f))
* **api:** api update ([fd976b6](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/fd976b6fc06ab7d87e51b5bdb48a817a20ba238b))
* **api:** generated missing apis ([24b7578](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/24b7578490ae47877fa432ff407951e074ca3645))
* **api:** manual updates ([fb134a6](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/fb134a6bb923006edd5f169ea79a0ff81728d4ca))
* **api:** manual updates ([d24b262](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/d24b262d2e863093b3fc7d04aa0f5f70ef4a2508))
* **api:** manual updates ([5a12d3a](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/5a12d3a6b5440ca6b55b8f904edfea1ed04ce84d))
* **api:** Update package names ([ead2fb3](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/ead2fb35663d079af30b5ba8cba40af323ffeb52))
* **client:** add support for endpoint-specific base URLs in python ([ea86e45](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/ea86e45a81e13950f679fa089010d21e6d294d80))
* **client:** experimental support for unmarshalling into param structs ([bae87c0](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/bae87c0a9e818d91c75a6db7058cc50f34ea299d))
* **client:** rename resp package ([1f6157e](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/1f6157e68c56ceb5ef1f60d4367e99fd42ac4082))


### Bug Fixes

* **client:** clean up reader resources ([53b2197](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/53b2197e72dc7ad7b28726e7a5e95a9d744ca76e))
* **client:** correctly set stream key for multipart ([24ea18b](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/24ea18b83cfb8862c102069033028eedc0f5b91e))
* **client:** correctly update body in WithJSONSet ([8a6c113](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/8a6c11330293c57c7ae7c8f667218cd933e2aa8c))
* **client:** don't panic on marshal with extra null field ([e42f237](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/e42f237e0a9acf28f3bbe34460b0c76041897d0c))
* **client:** improve core function names ([e1bba23](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/e1bba236633d9f9c04a661a68bb15806085cb935))
* **client:** improve docs ([98c0e99](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/98c0e996b1a856bf7e092a5674eb943ce99e7419))
* **client:** resolve issue with optional multipart files ([d80fcdf](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/d80fcdfb1c45424be7a70e3ecb662560d1f690d2))
* **client:** unmarshal responses properly ([2280fdc](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/2280fdc080af0c6949ed552d7f9d822fcdfeb1cb))
* correct unmarshalling of root body params ([9f65f86](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/9f65f8688dce2c068360ce5c3186512ac64e7106))
* fix error ([f55c8e4](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/f55c8e40c8560911383beaf9d3dc9bdcda1d9f3f))
* handle empty bodies in WithJSONSet ([943c242](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/943c242118101852c9f843c2f38520d5f075f52b))


### Chores

* **docs:** grammar improvements ([e5c38d4](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/e5c38d427928d706a5a2590cd462e4f1a2165f06))
* **docs:** update respjson package name ([4e9f47c](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/4e9f47c2c6e35299a861d59a80e6d6e41c5d6bab))
* improve devcontainer setup ([6e3483a](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/6e3483a9c75465d713edceb1354e5679eb324399))
* make go mod tidy continue on error ([4e83bb9](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/4e83bb972210c801c2863e0f44e5f4ca71bf32aa))
* update SDK settings ([876c733](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/876c733bc4110fb1da961a93bba03482eadc9b92))

## 0.1.0-alpha.1 (2025-04-24)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/dinaricrypto/dinari-api-sdk-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** manual updates ([0d7ecb9](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/0d7ecb9c12e6c76cc7e526eaba516529b1a66ef4))


### Chores

* configure new SDK language ([cf99344](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/cf99344dd4870e78048632ce270b8a65af3bcb78))
* go live ([8c8d015](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/8c8d0156ff5ced63ec205c150e4ef6d214c79357))
