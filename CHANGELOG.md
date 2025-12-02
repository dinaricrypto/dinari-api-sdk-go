# Changelog

## 0.7.1 (2025-12-02)

Full Changelog: [v0.7.0...v0.7.1](https://github.com/dinaricrypto/dinari-api-sdk-go/compare/v0.7.0...v0.7.1)

### Bug Fixes

* **client:** use repeat for array parameters ([fac9e0d](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/fac9e0d7c69a0064736e9c5599c316b01f005d9b))

## 0.7.0 (2025-11-25)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/dinaricrypto/dinari-api-sdk-go/compare/v0.6.0...v0.7.0)

### Features

* **api:** api update ([d72558f](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/d72558f8f16064aeabbab646c482214a117ebf5a))


### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([d6e6a19](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/d6e6a19630ff71569351d540776e77665a7eb873))

## 0.6.0 (2025-11-14)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/dinaricrypto/dinari-api-sdk-go/compare/v0.5.0...v0.6.0)

### Features

* **api:** api update ([032ae56](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/032ae5646d2a985e25f7daa1f09503d6be6d8c41))
* **api:** api update ([75351d9](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/75351d991c5ff13791b770c117e792faef282701))
* **api:** api update ([8604f1f](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/8604f1fd9ac5a47984ab231732f578b556318e3a))
* **api:** api update ([7663bfb](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/7663bfb35d599fc60595c595cab436ca550faad9))
* **api:** api update ([bb40438](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/bb40438184fa98d9fee5e76bd26a7daf221303af))
* **api:** api update ([426a14a](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/426a14a07620b525b237e94ae1266d8e2c23a352))
* **api:** api update ([e16bb08](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/e16bb08e01c9e151c133ebed4622f87c0835f411))
* **api:** api update ([1903ae4](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/1903ae45a6e321bd2a171db0945e1a1ab28ea602))
* **api:** api update ([6f1790b](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/6f1790be9d1da7622a9a5bf17abd5501d48ee6f9))
* **api:** permit, permit transaction, batch cancel ([ef1bb76](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/ef1bb7677441c38db16c66ec6c8eca6885ad06d0))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([a8a1b18](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/a8a1b1803d07dcbdc63269e0976fc59f20ae7474))
* close body before retrying ([0ac41e7](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/0ac41e7602d68fd52250ac21c37505d56afcddd0))
* **internal:** unmarshal correctly when there are multiple discriminators ([5f601e1](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/5f601e1b73e4994187bfb64006c17b7520f4090f))
* remove null from release please manifest ([b25d240](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/b25d240b5e8c83e95a0f80c12717d3a12b9784e4))
* use release please annotations on more places ([466b9cf](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/466b9cf8d1ccc914db0f9440c90836c83c844530))
* use slices.Concat instead of sometimes modifying r.Options ([db0be55](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/db0be5566d7a54b36a78488848359d3b1a7db9d7))


### Chores

* bump gjson version ([ab3b8e8](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/ab3b8e85e17070179335dec6e81cc90932691706))
* bump minimum go version to 1.22 ([50dbf8d](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/50dbf8d647c2d4adea685802cc51a7fad516072a))
* do not install brew dependencies in ./scripts/bootstrap by default ([adec94e](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/adec94ed65d6316057a351032c6d38c9c0af3a11))
* **internal:** grammar fix (it's -&gt; its) ([aec89eb](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/aec89eb2b79df985ba6366a59748d3892427b3b7))
* update more docs for 1.22 ([5b17404](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/5b1740437764de9b743331693b048f4fe89ba5a9))

## 0.5.0 (2025-08-20)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/dinaricrypto/dinari-api-sdk-go/compare/v0.4.0...v0.5.0)

### Features

* **api:** api update ([bbe0d2e](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/bbe0d2e9d8483c8c766efd5a3d16cb42342a6ef4))
* **api:** api update ([2e3b1bd](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/2e3b1bdd2d64f9973f9dda818eeaa7a32c8c6dc7))
* **api:** manual updates ([cbd9a4f](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/cbd9a4f4943c221d25586cb2502cf3ce8e1b362f))
* **client:** support optional json html escaping ([ee0a82a](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/ee0a82a6948d8a5eab8af8b89e166463a6c4d6c2))


### Bug Fixes

* **client:** process custom base url ahead of time ([29e51da](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/29e51dafd42e05b782e33e11c312356c1eae2bca))


### Chores

* **internal:** update comment in script ([f9e5c6c](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/f9e5c6cd83d91896478b79f46c025af4b56becef))
* **internal:** update test skipping reason ([2e22de3](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/2e22de3268f2c5b4b3aeac614266beacd516c150))
* update @stainless-api/prism-cli to v5.15.0 ([bff825e](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/bff825efaaa9bc4c863240bc3c432b3ce80bcdae))

## 0.4.0 (2025-07-15)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/dinaricrypto/dinari-api-sdk-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** api update ([9069221](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/906922130b6bbaba652ac94f8d6ba1b36bf0381b))
* **api:** api update ([7e330a2](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/7e330a21da5d216d6bfc7726da450447a07fe036))
* **api:** api update ([42f4e74](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/42f4e74f5820dc9690fadc791265af897e6a6aca))
* **api:** manual updates ([e1365e8](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/e1365e8ca31cc7316989634199039f1c121a9570))


### Chores

* **ci:** only run for pushes and fork pull requests ([8a7cb33](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/8a7cb33eb06e709001d278fa7e543c0c765b7cc2))
* **internal:** fix lint script for tests ([7ea8765](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/7ea8765d85c0772fdb3beb68112edb91b3757ef7))
* lint tests ([933d1e2](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/933d1e23811f7560a19501631ec51056b673a69c))
* lint tests in subpackages ([db4c246](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/db4c246f937cf66128f32775d4baa4b57c95e560))

## 0.3.0 (2025-06-27)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/dinaricrypto/dinari-api-sdk-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** api update ([2225bce](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/2225bced5626e4cd834fc1245c806eacdc37ef9d))
* **api:** api update ([71bae00](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/71bae000c777fb878dd2bd5d3f5bfcc82e70dd0b))
* **api:** api update ([fe64a2a](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/fe64a2a588dbdaaa54176d1a2a3c4830ed75bc05))
* **api:** api update ([f63812f](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/f63812f7e3ea7f54e472f20a1b9fa937ed30a964))
* **api:** remove quote api ([93b68cd](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/93b68cd54227ac39c438608cecbfaf15c7951442))


### Bug Fixes

* don't try to deserialize as json when ResponseBodyInto is []byte ([1211cf4](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/1211cf4b054df77f03e35b505298d6efafae82d2))

## 0.2.0 (2025-06-24)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/dinaricrypto/dinari-api-sdk-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** 20250624 new updates ([cc4f77b](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/cc4f77b740f9225ed4b1646e10754f2774eaa1ae))
* **api:** Add link to docs ([25b3c06](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/25b3c06476e1a8e65825864dcacf9ecef11f0f10))
* **api:** api update ([9d3a919](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/9d3a919a4cb2c9774f6409f631b50ad34b50660a))
* **api:** api update ([341c8d8](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/341c8d8e20e8d4193f115d388085d9e20f9ae6f5))
* **api:** api update ([23b16f4](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/23b16f4a5f6e52ca82f59149ae2f06ae2cbcce07))
* **api:** api update ([f8c80c8](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/f8c80c875fcb0d89d33af49ceb6d695592696278))
* **client:** add debug log helper ([08d1699](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/08d16996f15377e88a1b68ff07f6c90073fa49eb))
* **client:** add escape hatch for null slice & maps ([a50563e](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/a50563eaac3261c9cbd794c407a89fcd170917b1))


### Chores

* **ci:** enable for pull requests ([4c1197b](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/4c1197b9ca41481fdda0b4d957a072a5ea68506a))
* fix documentation of null map ([d561882](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/d5618820d5b954f2e2b61cee2205eb567b144665))

## 0.1.0 (2025-06-10)

Full Changelog: [v0.1.0-alpha.2...v0.1.0](https://github.com/dinaricrypto/dinari-api-sdk-go/compare/v0.1.0-alpha.2...v0.1.0)

### Features

* **api:** api update ([b074b4e](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/b074b4e602837e286c9bc4b56800b854d11d6c60))
* **api:** update routes ([3d26304](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/3d263042d1c26ee39a1c0f3657896c8e1c6dbe83))
* **client:** allow overriding unions ([f22a9d1](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/f22a9d1103c3d2ed1ba741a3559f4639d4d3a286))


### Bug Fixes

* **client:** cast to raw message when converting to params ([b1f19f3](https://github.com/dinaricrypto/dinari-api-sdk-go/commit/b1f19f3de2f05dc28f406e3f7e2fa6ab2369acce))

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
