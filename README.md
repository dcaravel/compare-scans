# Compare Scans

Ugly/messy utility written quickly to compare scans between two running StackRox instances - lots of room for improvement.

## Usage

### Saving scans from an StackRox environment

```sh
export ROX_ENDPOINT=...
export ROX_API_TOKEN=...

go run . dump

# Dump will be saved to a new data/<datetime> directory
```

<details>
<summary>Output</summary>

```sh
$ go run . dump
2024/12/17 16:48:57 Dumping files into: data/2024-12-17-164857
2024/12/17 16:48:57 Dumping: quay.io/dcaravel/sandbox/scanner:lin-fix-on-top-4.6.0-rc2
2024/12/17 16:48:57 Dumping: quay.io/dcaravel/temp:jdk-17.0.11.0.9-2.el8.x86_64
2024/12/17 16:48:57 Dumping: quay.io/dcaravel/temp:jdk-17.0.12.0.7-2.el8.x86_64
2024/12/17 16:48:57 Dumping: quay.io/openshift-release-dev/ocp-release@sha256:e8680baf0b44dc55accfe08c4ad298d508d5a19a371bc4747c2f6a92225aa38f
2024/12/17 16:48:57 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:010bfca35c0ee79ff1aff629a6351a057654c73cfbb236240b607597642ccca7
2024/12/17 16:48:57 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:0136aa5f8a73e9e0e9a9a2e18ad2378bb1f38e14f8f241b904adcf7669931389
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:09cd99d7439436633e16382deb8a4b164bae8986c63891f2ae5287656b5b4259
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:0b35c3489b4a47cda06c9f7ab17c70f1d3e4fd26b6b0542710a1d8c85d050632
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:137106528b4f7210e1eb60bb0d4f0b33a8614f67d29a66a3f2d336e2eb68b651
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:16a6db7159d6b4fdf0bf3509551d640ca1caf63930c542c9e29255806d222bd8
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:16ec84340f874f415a223333c8a7645f26b7a9749ab87335d59f2c12b1b25af9
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:1821408404d92846fe2a4a7ad852c9299041cfca907adffcef881ae891790947
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:1914a4674df5eb043e03476812e4942fd84538b5a352d7c7c0cb053369827ac0
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:1b3954ab7de376e369fbaa9773a19e828c678839166ba7ac9889c8498386eda1
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:1c4b13645d8edca0f1ed736822c09e6552c5e9354edf5e31026046bc848b34fe
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:228ec9aac0229f752c1245d3818a036fe6e83a316f7e4fa9c81d7f6c47ab6774
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:2569ad3a0fa872771e1629eb42007982b543fbcdbbdb9701d55bbfc450a778a0
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:25ed365f560ea1e856a0b6f9c13c995dc281aef7d28862e59aafb297cfc1face
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:28b7b528042590a8248df834e2c6b54b8bdfc6d1a373fe94fe17b06adf5956fa
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:2c36535004cd609a57268ceae24ef23ab9e0d292e95170326f5612e254dabd6e
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:2d3ac3bbe347de02564a49442f85c7cc692cdbd8082a17c3656cc5c31a2e2d5b
2024/12/17 16:48:58 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:2e7e52638641c80c353ffa69e39e7fb8b0e303608946fc9d9cd7f628311a9dd9
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:3185e547350fb8b52f0c695e826503de511bb9247b56711d7263b3758e8c9093
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:3d5c7b37400b26365519e6d6827192ec391d048a7ca9e42d1c1a88c782d31499
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:3dea0c7a80ff87e6f81c8224b103f43fafc7f64911578e4eb990d401736d4087
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:433a9ea9d40480f786c0dca505fdb97223c0386f8c692e261c4bc8c0cf38d62b
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:482129b30d7c45180d7ca8614a0471f5c8f4a065594e21070ba483e69240cf98
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:48550dd74fb0c774524d82f9d80f3bbd4243d25846297898fb2e091bd18d824a
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:49357c10877a1a733461d86bd5ca8bb522db735b7db8c958186ab245ce99bafb
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:50f230e8289d39895814b7d8c490f61338422b122235aab0dd70be86e5f5e767
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:526d11fb19f616850ff9b2b5b8a4780200c9d8d63cc43197d7c9ffcc30b3e004
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:549a6beaeac83094b210707f241ada1ea0cb13f55861b84f84778fd80d0c01d4
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:590295bd6c8855e1b52bf44e7d6808aa76517e2964606507fcd25a026c74fb15
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:5b7b1d5959c30918ef062e663b6bc61e3aec203155d54377c1fc6dd6defd7db6
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:609b2612ad66fcafbac82523679c0acc06462046ad33eb3906efceb89021c7a5
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:660ce72029aba1a934b6862eaca0f82b97dae3775393caa93586d68b99405740
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:6626f5567573f9a00620bc93443a4db0ea565cbb8b425de049f2d57619c8eed7
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:6726e36fd9346128c161f5481728e58dcc197d38177214242be7dc9c9b660b16
2024/12/17 16:48:59 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:6807b91a5e541682628bf47a9d343f6ba54a05c7ef045fcc9338b1e11b465b6e
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:69224ecce04c2f68b5834b934d5585cfbd45506bbf9add0fa575b7a00b1762a3
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:6c51cfc1eb477a99b8dcd6f032cafe9837e1ccc0eb2b45ce292e32068e2747f4
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:7143f23170f6650c54f4236c288d9f177a59dbd477f28e2e7682431102321b4d
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:7332280a5b7e91a38ba9b793ae4fa5ccdbac4b90839224d15114e87c55093551
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:7a6b1089536fe50ef73f21c7f90ae60c758e6e1fe1a83246ba5ea91cd79c1e3c
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:809e4b0acc8fa2184684db8baeb1cc67e9284d3da0c0768fde49d586f385266b
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:877742e075b364675dd821a77a0033750e854b923ce4651a68243a0f8190c288
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:8784049ad905eeeec8c39b11502edce90fa6cc5e43e3e00200c80cba59284afb
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:90b19de8a962e4b99cf336af1a51e6288ce493e35644f3fb8b9077b76e7ff98a
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:92916d9df4482358afdc582a79efe0b2a055584fa9da7ff96456f933cc38a3d6
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:962f024f81343fe18402a7f8f12d664f00f2388e95a1bd4b72c0e89244ce81fc
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:96678dfb7a22ce1282fe0aa85e7c72a51d13fed9ef9340ae0415483ce22ee8d0
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:9de5f4b385f2509859b6d5ceccac576d08b5058f992b55b7a6d52b30c43d2bf8
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a123333fefc13b204d2dfaff56e632d8ec01e7f59b8b069c698b14c427e33d57
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a1ee4fbd05d0876a828269e86803d4683dd844f35dabb363be45ab4bfe885e44
2024/12/17 16:49:00 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a8b45359686c1b84b9f0ced68cc5735655143190d8ca4abbfdae4cd75d92c84f
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a9de302775f058f771afa686f9268d6ae0ceca8c14e979dfb32ead7ffb5d528c
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:aa17a5e9afaec7c3f947361d5d74d89531b9babea54e460f8b3edd44dc6c5c67
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:abb7620fc503f2b7239e57292c6a65a6bddc38d3a0820b0efc7744cbafbbafb7
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:b0cb9964a798cd345070c09430796f5746981530ec605ae5ab95411747a67d55
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:b4d5874324a7b493e4815655e4c569a3137fbc3de47cabd03ec2abc2a9596949
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:b9a826c00b97bb89719b9df8b271c30c00af27671a59c43952ecda8eb9c9382c
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c1f8838079964de2ee4e9d7bc02ef71fa2b97dc706eb5d5f774e4c4e8dd2a4e3
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c463238757efc6cee597b0460ac7670b151c327e95a78638df4c3bba0c9ef0c8
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c4c5bbb0f88fccaa124666517ffa936ae3b951b5d9f9a2fba3bba653872a80c7
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c61578f0bae8f01aaaede511bc53fa7140c8b28ed68633b04b45496043f20cd6
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c678afbb9043dcaef2e5000a55ecaa1f06124d2215251113b6a9d31406328e90
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c737c185239ae3a47ecc568a0dec09642bf5841d70a5a5c31058cc0104b9d2f4
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c79df66f216072d54bc62c795bf7538c0b96aef87f53ef06402dc28f4e8a7e33
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ca33ab161180c31990d443d9fcb36c0b929e33a9129996a5d3945f3610fac660
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ce1d6f3172011651967257dc3da82dfbf9fc8609a7106887d0357b6d49148c34
2024/12/17 16:49:01 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:d7649f794d0cd572e18fa9b942a75dd799dad83d1500fd3a63c763e2bc1baf40
2024/12/17 16:49:02 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:df91f7f049fa8f808850dde2de749e159199954259c01967fa5b4e758cee26d5
2024/12/17 16:49:02 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e07518a890d9dddd554bf3f9faf331b9a67cb32b17dbcdd262a27250c3f5aae9
2024/12/17 16:49:02 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e204537c913ac9570043c8e7647949a0d66c44153f92df110813ff80ad611227
2024/12/17 16:49:02 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e22cfc52c9e6bd1d181c3edcb7d33c83da644b32cf13938a043d148ec39190c7
2024/12/17 16:49:02 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e458d234b9d7965d629707d4b136ceb423dc662da450390395cc24d66c3e192e
2024/12/17 16:49:02 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e75040362d9f243082dc9db8b2addf20db67d10399ed4548fcd81f20c2c4e4e0
2024/12/17 16:49:02 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e9484b715f54e372548cba5367004dc07d77743413bd76658f4e7e32c2c90f9e
2024/12/17 16:49:02 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:edfa7fb35f177ef83082d7f0d5a22214a58ab5e8e250b399ac09a6087297ee81
2024/12/17 16:49:02 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ef5cf312921e545c5f25166e50644596014d362d4e1586c659430be3ad24653f
2024/12/17 16:49:02 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ef70114a13362428e21b15bada2ddecf2e4dd9be28d79dd0cb360774ee42e724
2024/12/17 16:49:02 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:f03638bfcf12e418aac35b5b9193fae81ffd4ee574a8c68df4802fe57d9784c5
2024/12/17 16:49:02 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:f15a50f1780fbf3d1d2a013e315bc44b908483c4b6fb559d9596eb1b2343b259
2024/12/17 16:49:02 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:fc4bae5feec8cfea979bf45c7848a9b5127488624974461f50d6f8197faf0855
2024/12/17 16:49:02 Dumping: quay.io/rhacs-eng/central-db:4.6.0-rc.2
2024/12/17 16:49:02 Dumping: quay.io/rhacs-eng/collector:4.6.0-rc.2
2024/12/17 16:49:02 Dumping: quay.io/rhacs-eng/main:4.6.0-rc.2
2024/12/17 16:49:03 Dumping: quay.io/rhacs-eng/scanner-db:4.6.0-rc.2
2024/12/17 16:49:03 Dumping: quay.io/rhacs-eng/scanner-slim:4.6.0-rc.2
2024/12/17 16:49:03 Dumping: quay.io/rhacs-eng/scanner-v4-db:4.6.0-rc.2
2024/12/17 16:49:03 Dumping: quay.io/rhacs-eng/scanner-v4:4.6.0-rc.2
```

</details>

### Comparing directories of saved scans

```sh
go run . compare data/<dir1> data/<dir2>

# Results will be printed to the screen in a non-standard format
```
<details>
<summary>Output</summary>

```sh
$ go run . compare data/left data/right
2024/12/17 16:47:21 Comparing data/left to data/right
2024/12/17 16:47:21 image "sha256:286bd104615e072a0c9b053b878978d8c92716af5367d989c5395d2fc1fc1593" (quay.io/dcaravel/sandbox/scanner:lin-fix-on-top-4.6.0-rc2) not in data/right
2024/12/17 16:47:21 image "sha256:35361e07c059cc00fb7f3fc274d9cb2f20e35cf6885cae518e61cdbb6addab37" (quay.io/rhacs-eng/scanner:4.6.0-rc.2) not in data/left
2024/12/17 16:47:21 Comparing image: sha256:efcb76c712b5b771848daec08e99303d660f5ef570b4ec7f4014aa736bd1da5e (quay.io/dcaravel/temp:jdk-17.0.12.0.7-2.el8.x86_64)
2024/12/17 16:47:21   Component counts differ (264 vs 189)
2024/12/17 16:47:21   CVE counts differ (167 vs 112)
2024/12/17 16:47:21   FixableCVE counts differ (9 vs 24)
2024/12/17 16:47:21   Deep Component counts differ (264 vs 189)
2024/12/17 16:47:21   Deep comp name ver "httpcomponents-core/4.4.13-8.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "curl/7.61.1-34.el8_10.2.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "pango/1.42.4-8.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "dejavu-sans-mono-fonts/2.35-7.el8.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libsoup/2.62.3-5.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "util-linux/2.32.1-46.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "device-mapper/8:1.02.181-14.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "glibc-common/2.28-251.el8_10.5.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libgusb/0.3.0-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libjpeg-turbo/1.5.3-12.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "ca-certificates/2024.2.69_v8.0.303-80.0.el8_10.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libXfixes/5.0.3-7.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "cairo-gobject/1.15.12-6.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "librepo/1.14.2-5.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "maven-shared-utils/3.3.4-6.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libfdisk/2.32.1-46.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "cdi-api/2.0.2-7.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "glibc-minimal-langpack/2.28-251.el8_10.5.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "dbus/1:1.12.8-26.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "findutils/1:4.6.0-23.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "maven-wagon/3.5.1-3.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "nss/3.101.0-7.el8_8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "bash/4.4.20-5.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "java-17-openjdk-headless/1:17.0.12.0.7-2.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "harfbuzz/1.7.5-4.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "redhat-release/8.10-0.3.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "cracklib-dicts/2.9.6-15.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "device-mapper-libs/8:1.02.181-14.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libxkbcommon/0.9.1-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "elfutils-debuginfod-client/0.190-2.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libdatrie/0.2.9-7.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libuuid/2.32.1-46.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "gpgme/1.13.1-12.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libssh/0.9.6-14.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "plexus-containers-component-annotations/2.1.1-3.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "dbus-tools/1:1.12.8-26.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "gdk-pixbuf2-modules/2.36.12-6.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "gtk3/3.22.30-12.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libnghttp2/1.33.0-6.el8_10.1.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libthai/0.1.27-2.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "nss-sysinit/3.101.0-7.el8_8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "nss-util/3.101.0-7.el8_8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "rest/0.8.1-2.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libmodman/2.0.1-17.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "nss-softokn/3.101.0-7.el8_8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "elfutils-default-yama-scope/0.190-2.el8.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "glib2/2.56.4-162.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "plexus-classworlds/2.6.0-13.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "abattis-cantarell-fonts/0.0.25-6.el8.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "plexus-sec-dispatcher/2.0-5.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "audit-libs/3.1.2-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "cairo/1.15.12-6.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "hicolor-icon-theme/0.17-2.el8.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "plexus-interpolation/1.26-13.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "gdk-pixbuf2/2.36.12-6.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libgcc/8.5.0-22.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libseccomp/2.5.2-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "dbus-daemon/1:1.12.8-26.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "openldap/2.4.46-19.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "httpcomponents-client/4.5.13-6.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "rpm-libs/4.14.3-31.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libxml2/2.9.7-18.el8_10.1.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "at-spi2-atk/2.26.2-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "dconf/0.28.0-4.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libpwquality/1.4.4-6.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "dbus-common/1:1.12.8-26.el8.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libwayland-cursor/1.21.0-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "colord-libs/1.4.2-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libsmartcols/2.32.1-46.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "p11-kit-trust/0.23.22-2.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libXdamage/1.1.4-14.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libtiff/4.0.9-32.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "file-libs/5.33-26.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "platform-python-setuptools/39.2.0-8.el8_10.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libepoxy/1.5.8-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "apache-commons-codec/1.15-8.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "jansi/2.4.0-7.module+el8.10.0+21301+657f54a3.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "jcl-over-slf4j/1.7.32-5.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "jsr-305/3.0.2-7.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libXft/2.3.3-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "avahi-libs/0.7-27.el8_10.1.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "guava/31.0.1-5.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libstdc++/8.5.0-22.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "slf4j/1.7.32-5.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "expat/2.2.5-15.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "glibc/2.28-251.el8_10.5.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libXrandr/1.5.2-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "fribidi/1.0.4-9.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "maven-openjdk17/1:3.8.5-6.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libX11/1.6.8-9.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "platform-python-pip/9.0.3-24.el8.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "python3-libs/3.6.8-67.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "platform-python/3.6.8-67.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "cryptsetup-libs/2.3.7-7.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libcurl/7.61.1-34.el8_10.2.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libutempter/1.1.6-14.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "rpm/4.14.3-31.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "gzip/1.9-13.el8_5.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libXcursor/1.1.15-3.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "glibc-gconv-extra/2.28-251.el8_10.5.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "gtk-update-icon-cache/3.22.30-12.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "sisu/1:0.3.5-3.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "plexus-utils/3.3.0-11.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "adwaita-icon-theme/3.28.0-3.el8.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "java-17-openjdk-devel/1:17.0.12.0.7-2.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "jbigkit-libs/2.1-14.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "atinject/1.0.5-5.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "shared-mime-info/1.9-4.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "at-spi2-core/2.28.0-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "krb5-libs/1.18.2-29.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "plexus-cipher/2.0-3.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libacl/2.2.53-3.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libX11-common/1.6.8-9.el8_10.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "pam/1.3.1-34.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "java-17-openjdk/1:17.0.12.0.7-2.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "kmod-libs/25-20.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libXinerama/1.1.4-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "glibc-langpack-en/2.28-251.el8_10.5.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "graphite2/1.3.10-10.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "jakarta-annotations/1.3.5-15.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "jasper-libs/2.0.14-5.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libwayland-client/1.21.0-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "maven-resolver/1:1.7.3-6.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "systemd-pam/239-82.el8_10.2.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "p11-kit/0.23.22-2.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "systemd-libs/239-82.el8_10.2.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "elfutils-libelf/0.190-2.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libwayland-egl/1.21.0-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "maven-lib/1:3.8.5-6.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "shadow-utils/2:4.6-22.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libdnf/0.63.0-20.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "gsettings-desktop-schemas/3.32.0-6.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "python3-setuptools-wheel/39.2.0-8.el8_10.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "lcms2/2.9-2.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "apache-commons-cli/1.5.0-5.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "atk/2.28.1-1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "acl/2.2.53-3.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "elfutils-libs/0.190-2.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libblkid/2.32.1-46.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libproxy/0.4.15-5.2.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "cracklib/2.9.6-15.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libtirpc/1.1.4-12.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "python3-pip-wheel/9.0.3-24.el8.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "adwaita-cursor-theme/3.28.0-3.el8.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "glib-networking/2.56.1-1.1.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libmount/2.32.1-46.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "pixman/0.38.4-4.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "apache-commons-io/1:2.11.0-3.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "diffutils/3.6-6.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "libssh-config/0.9.6-14.el8.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "nss-softokn-freebl/3.101.0-7.el8_8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "alsa-lib/1.2.10-2.el8.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "systemd/239-82.el8_10.2.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "apache-commons-lang3/3.12.0-8.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "xkeyboard-config/2.28-1.el8.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "cups-libs/1:2.2.6-60.el8_10.x86_64" not in data/right
2024/12/17 16:47:21   Deep comp name ver "google-guice/4.2.3-10.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "maven/1:3.8.5-6.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:47:21   Deep comp name ver "glibc-minimal-langpack/2.28-236.el8_9.13.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "librepo/1.14.2-4.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "systemd-libs/239-78.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "httpcomponents-core/4.4.13-7.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libX11/1.6.8-6.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libsmartcols/2.32.1-44.el8_9.1.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "apache-commons-codec/1.15-7.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "jcl-over-slf4j/1.7.32-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "platform-python-setuptools/39.2.0-7.el8.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "plexus-containers-component-annotations/2.1.1-2.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "java-17-openjdk-headless/1:17.0.11.0.9-2.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libacl/2.2.53-1.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libssh/0.9.6-13.el8_9.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "maven-lib/1:3.8.5-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "maven-shared-utils/3.3.4-5.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "python3-setuptools-wheel/39.2.0-7.el8.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "findutils/1:4.6.0-21.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libdnf/0.63.0-17.el8_9.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "shadow-utils/2:4.6-19.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libnghttp2/1.33.0-5.el8_9.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "glibc/2.28-236.el8_9.13.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libxml2/2.9.7-18.el8_9.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "maven-openjdk17/1:3.8.5-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "nss-util/3.90.0-6.el8_9.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libX11-common/1.6.8-6.el8.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "nss-sysinit/3.90.0-6.el8_9.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "plexus-cipher/2.0-2.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "expat/2.2.5-11.el8_9.1.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libssh-config/0.9.6-13.el8_9.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "nss/3.90.0-6.el8_9.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "bash/4.4.20-4.el8_6.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libstdc++/8.5.0-20.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "rpm-libs/4.14.3-28.el8_9.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "apache-commons-io/1:2.11.0-2.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "maven-resolver/1:1.7.3-5.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "openldap/2.4.46-18.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "google-guice/4.2.3-9.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "java-17-openjdk-devel/1:17.0.11.0.9-2.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "redhat-release/8.9-0.1.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "apache-commons-lang3/3.12.0-7.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "ca-certificates/2023.2.60_v7.0.306-80.0.el8_8.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "maven/1:3.8.5-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "cdi-api/2.0.2-6.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "p11-kit/0.23.22-1.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "plexus-sec-dispatcher/2.0-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "rpm/4.14.3-28.el8_9.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "avahi-libs/0.7-21.el8_9.1.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "gpgme/1.13.1-11.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "nss-softokn/3.90.0-6.el8_9.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "python3-libs/3.6.8-56.el8_9.3.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "alsa-lib/1.2.9-1.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "plexus-classworlds/2.6.0-12.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "glib2/2.56.4-161.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libgcc/8.5.0-20.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "sisu/1:0.3.5-2.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "audit-libs/3.0.7-5.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "guava/31.0.1-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "file-libs/5.33-25.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "java-17-openjdk/1:17.0.11.0.9-2.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libblkid/2.32.1-44.el8_9.1.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "maven-wagon/3.5.1-2.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "jsr-305/3.0.2-6.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libuuid/2.32.1-44.el8_9.1.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "nss-softokn-freebl/3.90.0-6.el8_9.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "apache-commons-cli/1.5.0-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "glibc-common/2.28-236.el8_9.13.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "httpcomponents-client/4.5.13-5.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libmount/2.32.1-44.el8_9.1.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libtirpc/1.1.4-8.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "python3-pip-wheel/9.0.3-23.el8_9.1.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "krb5-libs/1.18.2-26.el8_9.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "plexus-utils/3.3.0-10.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "slf4j/1.7.32-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "atinject/1.0.5-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "libcurl/7.61.1-33.el8_9.5.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "p11-kit-trust/0.23.22-1.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "cups-libs/1:2.2.6-54.el8_9.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "platform-python/3.6.8-56.el8_9.3.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "curl/7.61.1-33.el8_9.5.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "jansi/2.4.0-6.module+el8.8.0+18044+0a924b8f.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "plexus-interpolation/1.26-12.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21   Deep comp name ver "elfutils-libelf/0.189-3.el8.x86_64" not in data/left
2024/12/17 16:47:21   Deep comp name ver "jakarta-annotations/1.3.5-14.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:47:21 1 image only in  left dir "data/left"
2024/12/17 16:47:21 1 image only in right dir "data/right"
2024/12/17 16:47:21 89/90 images in common are identical
```
</details>


### Comparing results from two running StackRox instances

```sh

# Setup the 'left' StackRox instance
export ROX_ENDPOINT=...
export ROX_API_TOKEN=...

# Setup the 'right' StackRox instance (notice the 'R' at end of env var name)
export ROX_ENDPOINTR=...
export ROX_API_TOKENR=...

go run . livecompare
```

<details>
<summary>Output</summary>

```sh
$ go run . livecompare
2024/12/17 16:46:12 Dumping files into: data/left
2024/12/17 16:46:13 Dumping: quay.io/dcaravel/sandbox/scanner:lin-fix-on-top-4.6.0-rc2
2024/12/17 16:46:13 Dumping: quay.io/dcaravel/temp:jdk-17.0.11.0.9-2.el8.x86_64
2024/12/17 16:46:13 Dumping: quay.io/dcaravel/temp:jdk-17.0.12.0.7-2.el8.x86_64
2024/12/17 16:46:13 Dumping: quay.io/openshift-release-dev/ocp-release@sha256:e8680baf0b44dc55accfe08c4ad298d508d5a19a371bc4747c2f6a92225aa38f
2024/12/17 16:46:13 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:010bfca35c0ee79ff1aff629a6351a057654c73cfbb236240b607597642ccca7
2024/12/17 16:46:13 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:0136aa5f8a73e9e0e9a9a2e18ad2378bb1f38e14f8f241b904adcf7669931389
2024/12/17 16:46:13 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:09cd99d7439436633e16382deb8a4b164bae8986c63891f2ae5287656b5b4259
2024/12/17 16:46:13 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:0b35c3489b4a47cda06c9f7ab17c70f1d3e4fd26b6b0542710a1d8c85d050632
2024/12/17 16:46:13 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:137106528b4f7210e1eb60bb0d4f0b33a8614f67d29a66a3f2d336e2eb68b651
2024/12/17 16:46:13 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:16a6db7159d6b4fdf0bf3509551d640ca1caf63930c542c9e29255806d222bd8
2024/12/17 16:46:13 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:16ec84340f874f415a223333c8a7645f26b7a9749ab87335d59f2c12b1b25af9
2024/12/17 16:46:13 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:1821408404d92846fe2a4a7ad852c9299041cfca907adffcef881ae891790947
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:1914a4674df5eb043e03476812e4942fd84538b5a352d7c7c0cb053369827ac0
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:1b3954ab7de376e369fbaa9773a19e828c678839166ba7ac9889c8498386eda1
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:1c4b13645d8edca0f1ed736822c09e6552c5e9354edf5e31026046bc848b34fe
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:228ec9aac0229f752c1245d3818a036fe6e83a316f7e4fa9c81d7f6c47ab6774
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:2569ad3a0fa872771e1629eb42007982b543fbcdbbdb9701d55bbfc450a778a0
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:25ed365f560ea1e856a0b6f9c13c995dc281aef7d28862e59aafb297cfc1face
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:28b7b528042590a8248df834e2c6b54b8bdfc6d1a373fe94fe17b06adf5956fa
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:2c36535004cd609a57268ceae24ef23ab9e0d292e95170326f5612e254dabd6e
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:2d3ac3bbe347de02564a49442f85c7cc692cdbd8082a17c3656cc5c31a2e2d5b
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:2e7e52638641c80c353ffa69e39e7fb8b0e303608946fc9d9cd7f628311a9dd9
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:3185e547350fb8b52f0c695e826503de511bb9247b56711d7263b3758e8c9093
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:3d5c7b37400b26365519e6d6827192ec391d048a7ca9e42d1c1a88c782d31499
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:3dea0c7a80ff87e6f81c8224b103f43fafc7f64911578e4eb990d401736d4087
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:433a9ea9d40480f786c0dca505fdb97223c0386f8c692e261c4bc8c0cf38d62b
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:482129b30d7c45180d7ca8614a0471f5c8f4a065594e21070ba483e69240cf98
2024/12/17 16:46:14 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:48550dd74fb0c774524d82f9d80f3bbd4243d25846297898fb2e091bd18d824a
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:49357c10877a1a733461d86bd5ca8bb522db735b7db8c958186ab245ce99bafb
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:50f230e8289d39895814b7d8c490f61338422b122235aab0dd70be86e5f5e767
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:526d11fb19f616850ff9b2b5b8a4780200c9d8d63cc43197d7c9ffcc30b3e004
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:549a6beaeac83094b210707f241ada1ea0cb13f55861b84f84778fd80d0c01d4
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:590295bd6c8855e1b52bf44e7d6808aa76517e2964606507fcd25a026c74fb15
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:5b7b1d5959c30918ef062e663b6bc61e3aec203155d54377c1fc6dd6defd7db6
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:609b2612ad66fcafbac82523679c0acc06462046ad33eb3906efceb89021c7a5
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:660ce72029aba1a934b6862eaca0f82b97dae3775393caa93586d68b99405740
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:6626f5567573f9a00620bc93443a4db0ea565cbb8b425de049f2d57619c8eed7
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:6726e36fd9346128c161f5481728e58dcc197d38177214242be7dc9c9b660b16
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:6807b91a5e541682628bf47a9d343f6ba54a05c7ef045fcc9338b1e11b465b6e
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:69224ecce04c2f68b5834b934d5585cfbd45506bbf9add0fa575b7a00b1762a3
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:6c51cfc1eb477a99b8dcd6f032cafe9837e1ccc0eb2b45ce292e32068e2747f4
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:7143f23170f6650c54f4236c288d9f177a59dbd477f28e2e7682431102321b4d
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:7332280a5b7e91a38ba9b793ae4fa5ccdbac4b90839224d15114e87c55093551
2024/12/17 16:46:15 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:7a6b1089536fe50ef73f21c7f90ae60c758e6e1fe1a83246ba5ea91cd79c1e3c
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:809e4b0acc8fa2184684db8baeb1cc67e9284d3da0c0768fde49d586f385266b
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:877742e075b364675dd821a77a0033750e854b923ce4651a68243a0f8190c288
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:8784049ad905eeeec8c39b11502edce90fa6cc5e43e3e00200c80cba59284afb
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:90b19de8a962e4b99cf336af1a51e6288ce493e35644f3fb8b9077b76e7ff98a
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:92916d9df4482358afdc582a79efe0b2a055584fa9da7ff96456f933cc38a3d6
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:962f024f81343fe18402a7f8f12d664f00f2388e95a1bd4b72c0e89244ce81fc
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:96678dfb7a22ce1282fe0aa85e7c72a51d13fed9ef9340ae0415483ce22ee8d0
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:9de5f4b385f2509859b6d5ceccac576d08b5058f992b55b7a6d52b30c43d2bf8
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a123333fefc13b204d2dfaff56e632d8ec01e7f59b8b069c698b14c427e33d57
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a1ee4fbd05d0876a828269e86803d4683dd844f35dabb363be45ab4bfe885e44
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a8b45359686c1b84b9f0ced68cc5735655143190d8ca4abbfdae4cd75d92c84f
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a9de302775f058f771afa686f9268d6ae0ceca8c14e979dfb32ead7ffb5d528c
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:aa17a5e9afaec7c3f947361d5d74d89531b9babea54e460f8b3edd44dc6c5c67
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:abb7620fc503f2b7239e57292c6a65a6bddc38d3a0820b0efc7744cbafbbafb7
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:b0cb9964a798cd345070c09430796f5746981530ec605ae5ab95411747a67d55
2024/12/17 16:46:16 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:b4d5874324a7b493e4815655e4c569a3137fbc3de47cabd03ec2abc2a9596949
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:b9a826c00b97bb89719b9df8b271c30c00af27671a59c43952ecda8eb9c9382c
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c1f8838079964de2ee4e9d7bc02ef71fa2b97dc706eb5d5f774e4c4e8dd2a4e3
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c463238757efc6cee597b0460ac7670b151c327e95a78638df4c3bba0c9ef0c8
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c4c5bbb0f88fccaa124666517ffa936ae3b951b5d9f9a2fba3bba653872a80c7
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c61578f0bae8f01aaaede511bc53fa7140c8b28ed68633b04b45496043f20cd6
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c678afbb9043dcaef2e5000a55ecaa1f06124d2215251113b6a9d31406328e90
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c737c185239ae3a47ecc568a0dec09642bf5841d70a5a5c31058cc0104b9d2f4
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c79df66f216072d54bc62c795bf7538c0b96aef87f53ef06402dc28f4e8a7e33
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ca33ab161180c31990d443d9fcb36c0b929e33a9129996a5d3945f3610fac660
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ce1d6f3172011651967257dc3da82dfbf9fc8609a7106887d0357b6d49148c34
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:d7649f794d0cd572e18fa9b942a75dd799dad83d1500fd3a63c763e2bc1baf40
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:df91f7f049fa8f808850dde2de749e159199954259c01967fa5b4e758cee26d5
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e07518a890d9dddd554bf3f9faf331b9a67cb32b17dbcdd262a27250c3f5aae9
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e204537c913ac9570043c8e7647949a0d66c44153f92df110813ff80ad611227
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e22cfc52c9e6bd1d181c3edcb7d33c83da644b32cf13938a043d148ec39190c7
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e458d234b9d7965d629707d4b136ceb423dc662da450390395cc24d66c3e192e
2024/12/17 16:46:17 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e75040362d9f243082dc9db8b2addf20db67d10399ed4548fcd81f20c2c4e4e0
2024/12/17 16:46:18 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e9484b715f54e372548cba5367004dc07d77743413bd76658f4e7e32c2c90f9e
2024/12/17 16:46:18 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:edfa7fb35f177ef83082d7f0d5a22214a58ab5e8e250b399ac09a6087297ee81
2024/12/17 16:46:18 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ef5cf312921e545c5f25166e50644596014d362d4e1586c659430be3ad24653f
2024/12/17 16:46:18 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ef70114a13362428e21b15bada2ddecf2e4dd9be28d79dd0cb360774ee42e724
2024/12/17 16:46:18 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:f03638bfcf12e418aac35b5b9193fae81ffd4ee574a8c68df4802fe57d9784c5
2024/12/17 16:46:18 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:f15a50f1780fbf3d1d2a013e315bc44b908483c4b6fb559d9596eb1b2343b259
2024/12/17 16:46:18 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:fc4bae5feec8cfea979bf45c7848a9b5127488624974461f50d6f8197faf0855
2024/12/17 16:46:18 Dumping: quay.io/rhacs-eng/central-db:4.6.0-rc.2
2024/12/17 16:46:18 Dumping: quay.io/rhacs-eng/collector:4.6.0-rc.2
2024/12/17 16:46:18 Dumping: quay.io/rhacs-eng/main:4.6.0-rc.2
2024/12/17 16:46:18 Dumping: quay.io/rhacs-eng/scanner-db:4.6.0-rc.2
2024/12/17 16:46:18 Dumping: quay.io/rhacs-eng/scanner-slim:4.6.0-rc.2
2024/12/17 16:46:18 Dumping: quay.io/rhacs-eng/scanner-v4-db:4.6.0-rc.2
2024/12/17 16:46:18 Dumping: quay.io/rhacs-eng/scanner-v4:4.6.0-rc.2
2024/12/17 16:46:18 Dumping files into: data/right
2024/12/17 16:46:19 Dumping: quay.io/dcaravel/temp:jdk-17.0.11.0.9-2.el8.x86_64
2024/12/17 16:46:19 Dumping: quay.io/dcaravel/temp:jdk-17.0.12.0.7-2.el8.x86_64
2024/12/17 16:46:19 Dumping: quay.io/openshift-release-dev/ocp-release@sha256:e8680baf0b44dc55accfe08c4ad298d508d5a19a371bc4747c2f6a92225aa38f
2024/12/17 16:46:19 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:010bfca35c0ee79ff1aff629a6351a057654c73cfbb236240b607597642ccca7
2024/12/17 16:46:19 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:0136aa5f8a73e9e0e9a9a2e18ad2378bb1f38e14f8f241b904adcf7669931389
2024/12/17 16:46:19 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:09cd99d7439436633e16382deb8a4b164bae8986c63891f2ae5287656b5b4259
2024/12/17 16:46:19 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:0b35c3489b4a47cda06c9f7ab17c70f1d3e4fd26b6b0542710a1d8c85d050632
2024/12/17 16:46:19 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:137106528b4f7210e1eb60bb0d4f0b33a8614f67d29a66a3f2d336e2eb68b651
2024/12/17 16:46:19 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:16a6db7159d6b4fdf0bf3509551d640ca1caf63930c542c9e29255806d222bd8
2024/12/17 16:46:19 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:16ec84340f874f415a223333c8a7645f26b7a9749ab87335d59f2c12b1b25af9
2024/12/17 16:46:19 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:1821408404d92846fe2a4a7ad852c9299041cfca907adffcef881ae891790947
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:1914a4674df5eb043e03476812e4942fd84538b5a352d7c7c0cb053369827ac0
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:1b3954ab7de376e369fbaa9773a19e828c678839166ba7ac9889c8498386eda1
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:1c4b13645d8edca0f1ed736822c09e6552c5e9354edf5e31026046bc848b34fe
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:228ec9aac0229f752c1245d3818a036fe6e83a316f7e4fa9c81d7f6c47ab6774
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:2569ad3a0fa872771e1629eb42007982b543fbcdbbdb9701d55bbfc450a778a0
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:25ed365f560ea1e856a0b6f9c13c995dc281aef7d28862e59aafb297cfc1face
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:28b7b528042590a8248df834e2c6b54b8bdfc6d1a373fe94fe17b06adf5956fa
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:2c36535004cd609a57268ceae24ef23ab9e0d292e95170326f5612e254dabd6e
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:2d3ac3bbe347de02564a49442f85c7cc692cdbd8082a17c3656cc5c31a2e2d5b
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:2e7e52638641c80c353ffa69e39e7fb8b0e303608946fc9d9cd7f628311a9dd9
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:3185e547350fb8b52f0c695e826503de511bb9247b56711d7263b3758e8c9093
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:3d5c7b37400b26365519e6d6827192ec391d048a7ca9e42d1c1a88c782d31499
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:3dea0c7a80ff87e6f81c8224b103f43fafc7f64911578e4eb990d401736d4087
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:433a9ea9d40480f786c0dca505fdb97223c0386f8c692e261c4bc8c0cf38d62b
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:482129b30d7c45180d7ca8614a0471f5c8f4a065594e21070ba483e69240cf98
2024/12/17 16:46:20 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:48550dd74fb0c774524d82f9d80f3bbd4243d25846297898fb2e091bd18d824a
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:49357c10877a1a733461d86bd5ca8bb522db735b7db8c958186ab245ce99bafb
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:50f230e8289d39895814b7d8c490f61338422b122235aab0dd70be86e5f5e767
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:526d11fb19f616850ff9b2b5b8a4780200c9d8d63cc43197d7c9ffcc30b3e004
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:549a6beaeac83094b210707f241ada1ea0cb13f55861b84f84778fd80d0c01d4
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:590295bd6c8855e1b52bf44e7d6808aa76517e2964606507fcd25a026c74fb15
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:5b7b1d5959c30918ef062e663b6bc61e3aec203155d54377c1fc6dd6defd7db6
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:609b2612ad66fcafbac82523679c0acc06462046ad33eb3906efceb89021c7a5
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:660ce72029aba1a934b6862eaca0f82b97dae3775393caa93586d68b99405740
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:6626f5567573f9a00620bc93443a4db0ea565cbb8b425de049f2d57619c8eed7
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:6726e36fd9346128c161f5481728e58dcc197d38177214242be7dc9c9b660b16
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:6807b91a5e541682628bf47a9d343f6ba54a05c7ef045fcc9338b1e11b465b6e
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:69224ecce04c2f68b5834b934d5585cfbd45506bbf9add0fa575b7a00b1762a3
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:6c51cfc1eb477a99b8dcd6f032cafe9837e1ccc0eb2b45ce292e32068e2747f4
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:7143f23170f6650c54f4236c288d9f177a59dbd477f28e2e7682431102321b4d
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:7332280a5b7e91a38ba9b793ae4fa5ccdbac4b90839224d15114e87c55093551
2024/12/17 16:46:21 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:7a6b1089536fe50ef73f21c7f90ae60c758e6e1fe1a83246ba5ea91cd79c1e3c
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:809e4b0acc8fa2184684db8baeb1cc67e9284d3da0c0768fde49d586f385266b
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:877742e075b364675dd821a77a0033750e854b923ce4651a68243a0f8190c288
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:8784049ad905eeeec8c39b11502edce90fa6cc5e43e3e00200c80cba59284afb
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:90b19de8a962e4b99cf336af1a51e6288ce493e35644f3fb8b9077b76e7ff98a
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:92916d9df4482358afdc582a79efe0b2a055584fa9da7ff96456f933cc38a3d6
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:962f024f81343fe18402a7f8f12d664f00f2388e95a1bd4b72c0e89244ce81fc
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:96678dfb7a22ce1282fe0aa85e7c72a51d13fed9ef9340ae0415483ce22ee8d0
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:9de5f4b385f2509859b6d5ceccac576d08b5058f992b55b7a6d52b30c43d2bf8
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a123333fefc13b204d2dfaff56e632d8ec01e7f59b8b069c698b14c427e33d57
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a1ee4fbd05d0876a828269e86803d4683dd844f35dabb363be45ab4bfe885e44
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a8b45359686c1b84b9f0ced68cc5735655143190d8ca4abbfdae4cd75d92c84f
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a9de302775f058f771afa686f9268d6ae0ceca8c14e979dfb32ead7ffb5d528c
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:aa17a5e9afaec7c3f947361d5d74d89531b9babea54e460f8b3edd44dc6c5c67
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:abb7620fc503f2b7239e57292c6a65a6bddc38d3a0820b0efc7744cbafbbafb7
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:b0cb9964a798cd345070c09430796f5746981530ec605ae5ab95411747a67d55
2024/12/17 16:46:22 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:b4d5874324a7b493e4815655e4c569a3137fbc3de47cabd03ec2abc2a9596949
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:b9a826c00b97bb89719b9df8b271c30c00af27671a59c43952ecda8eb9c9382c
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c1f8838079964de2ee4e9d7bc02ef71fa2b97dc706eb5d5f774e4c4e8dd2a4e3
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c463238757efc6cee597b0460ac7670b151c327e95a78638df4c3bba0c9ef0c8
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c4c5bbb0f88fccaa124666517ffa936ae3b951b5d9f9a2fba3bba653872a80c7
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c61578f0bae8f01aaaede511bc53fa7140c8b28ed68633b04b45496043f20cd6
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c678afbb9043dcaef2e5000a55ecaa1f06124d2215251113b6a9d31406328e90
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c737c185239ae3a47ecc568a0dec09642bf5841d70a5a5c31058cc0104b9d2f4
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:c79df66f216072d54bc62c795bf7538c0b96aef87f53ef06402dc28f4e8a7e33
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ca33ab161180c31990d443d9fcb36c0b929e33a9129996a5d3945f3610fac660
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ce1d6f3172011651967257dc3da82dfbf9fc8609a7106887d0357b6d49148c34
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:d7649f794d0cd572e18fa9b942a75dd799dad83d1500fd3a63c763e2bc1baf40
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:df91f7f049fa8f808850dde2de749e159199954259c01967fa5b4e758cee26d5
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e07518a890d9dddd554bf3f9faf331b9a67cb32b17dbcdd262a27250c3f5aae9
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e204537c913ac9570043c8e7647949a0d66c44153f92df110813ff80ad611227
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e22cfc52c9e6bd1d181c3edcb7d33c83da644b32cf13938a043d148ec39190c7
2024/12/17 16:46:23 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e458d234b9d7965d629707d4b136ceb423dc662da450390395cc24d66c3e192e
2024/12/17 16:46:24 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e75040362d9f243082dc9db8b2addf20db67d10399ed4548fcd81f20c2c4e4e0
2024/12/17 16:46:24 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:e9484b715f54e372548cba5367004dc07d77743413bd76658f4e7e32c2c90f9e
2024/12/17 16:46:24 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:edfa7fb35f177ef83082d7f0d5a22214a58ab5e8e250b399ac09a6087297ee81
2024/12/17 16:46:24 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ef5cf312921e545c5f25166e50644596014d362d4e1586c659430be3ad24653f
2024/12/17 16:46:24 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:ef70114a13362428e21b15bada2ddecf2e4dd9be28d79dd0cb360774ee42e724
2024/12/17 16:46:24 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:f03638bfcf12e418aac35b5b9193fae81ffd4ee574a8c68df4802fe57d9784c5
2024/12/17 16:46:24 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:f15a50f1780fbf3d1d2a013e315bc44b908483c4b6fb559d9596eb1b2343b259
2024/12/17 16:46:24 Dumping: quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:fc4bae5feec8cfea979bf45c7848a9b5127488624974461f50d6f8197faf0855
2024/12/17 16:46:24 Dumping: quay.io/rhacs-eng/central-db:4.6.0-rc.2
2024/12/17 16:46:24 Dumping: quay.io/rhacs-eng/collector:4.6.0-rc.2
2024/12/17 16:46:24 Dumping: quay.io/rhacs-eng/main:4.6.0-rc.2
2024/12/17 16:46:24 Dumping: quay.io/rhacs-eng/scanner-db:4.6.0-rc.2
2024/12/17 16:46:24 Dumping: quay.io/rhacs-eng/scanner-slim:4.6.0-rc.2
2024/12/17 16:46:24 Dumping: quay.io/rhacs-eng/scanner-v4-db:4.6.0-rc.2
2024/12/17 16:46:24 Dumping: quay.io/rhacs-eng/scanner-v4:4.6.0-rc.2
2024/12/17 16:46:25 Dumping: quay.io/rhacs-eng/scanner:4.6.0-rc.2
2024/12/17 16:46:25 Comparing data/left to data/right
2024/12/17 16:46:25 image "sha256:286bd104615e072a0c9b053b878978d8c92716af5367d989c5395d2fc1fc1593" (quay.io/dcaravel/sandbox/scanner:lin-fix-on-top-4.6.0-rc2) not in data/right
2024/12/17 16:46:25 image "sha256:35361e07c059cc00fb7f3fc274d9cb2f20e35cf6885cae518e61cdbb6addab37" (quay.io/rhacs-eng/scanner:4.6.0-rc.2) not in data/left
2024/12/17 16:46:25 Comparing image: sha256:efcb76c712b5b771848daec08e99303d660f5ef570b4ec7f4014aa736bd1da5e (quay.io/dcaravel/temp:jdk-17.0.12.0.7-2.el8.x86_64)
2024/12/17 16:46:25   Component counts differ (264 vs 189)
2024/12/17 16:46:25   CVE counts differ (167 vs 112)
2024/12/17 16:46:25   FixableCVE counts differ (9 vs 24)
2024/12/17 16:46:25   Deep Component counts differ (264 vs 189)
2024/12/17 16:46:25   Deep comp name ver "gdk-pixbuf2-modules/2.36.12-6.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "gtk3/3.22.30-12.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libpwquality/1.4.4-6.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "librepo/1.14.2-5.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "elfutils-libelf/0.190-2.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "gpgme/1.13.1-12.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "lcms2/2.9-2.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "maven-openjdk17/1:3.8.5-6.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "elfutils-default-yama-scope/0.190-2.el8.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libdnf/0.63.0-20.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libxkbcommon/0.9.1-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "acl/2.2.53-3.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "apache-commons-codec/1.15-8.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "glibc-langpack-en/2.28-251.el8_10.5.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libuuid/2.32.1-46.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "platform-python-setuptools/39.2.0-8.el8_10.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "elfutils-libs/0.190-2.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libdatrie/0.2.9-7.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "nss-softokn-freebl/3.101.0-7.el8_8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "plexus-utils/3.3.0-11.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "expat/2.2.5-15.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "gtk-update-icon-cache/3.22.30-12.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libjpeg-turbo/1.5.3-12.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "rpm/4.14.3-31.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "glibc-minimal-langpack/2.28-251.el8_10.5.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "shared-mime-info/1.9-4.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "glibc-common/2.28-251.el8_10.5.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "apache-commons-io/1:2.11.0-3.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "plexus-containers-component-annotations/2.1.1-3.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "java-17-openjdk-headless/1:17.0.12.0.7-2.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libwayland-client/1.21.0-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "audit-libs/3.1.2-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libXrandr/1.5.2-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "openldap/2.4.46-19.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libwayland-egl/1.21.0-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "abattis-cantarell-fonts/0.0.25-6.el8.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "cracklib/2.9.6-15.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "kmod-libs/25-20.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "apache-commons-lang3/3.12.0-8.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libxml2/2.9.7-18.el8_10.1.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libstdc++/8.5.0-22.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "pixman/0.38.4-4.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libsmartcols/2.32.1-46.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "util-linux/2.32.1-46.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "at-spi2-atk/2.26.2-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "atk/2.28.1-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libXcursor/1.1.15-3.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libwayland-cursor/1.21.0-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "bash/4.4.20-5.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "glib2/2.56.4-162.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libtiff/4.0.9-32.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "slf4j/1.7.32-5.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libX11/1.6.8-9.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "cracklib-dicts/2.9.6-15.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "glib-networking/2.56.1-1.1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libepoxy/1.5.8-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libutempter/1.1.6-14.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "alsa-lib/1.2.10-2.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "httpcomponents-client/4.5.13-6.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "colord-libs/1.4.2-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libthai/0.1.27-2.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "systemd/239-82.el8_10.2.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "file-libs/5.33-26.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "jakarta-annotations/1.3.5-15.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "jansi/2.4.0-7.module+el8.10.0+21301+657f54a3.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libmodman/2.0.1-17.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "rest/0.8.1-2.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "gdk-pixbuf2/2.36.12-6.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libXinerama/1.1.4-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libacl/2.2.53-3.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libmount/2.32.1-46.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "maven-lib/1:3.8.5-6.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "systemd-libs/239-82.el8_10.2.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "dbus-common/1:1.12.8-26.el8.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libXdamage/1.1.4-14.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "pam/1.3.1-34.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "pango/1.42.4-8.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libX11-common/1.6.8-9.el8_10.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libfdisk/2.32.1-46.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "diffutils/3.6-6.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "cdi-api/2.0.2-7.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "p11-kit-trust/0.23.22-2.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "redhat-release/8.10-0.3.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "dbus/1:1.12.8-26.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libtirpc/1.1.4-12.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "p11-kit/0.23.22-2.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "atinject/1.0.5-5.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libnghttp2/1.33.0-6.el8_10.1.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "python3-setuptools-wheel/39.2.0-8.el8_10.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "rpm-libs/4.14.3-31.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libblkid/2.32.1-46.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libproxy/0.4.15-5.2.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libseccomp/2.5.2-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "maven-resolver/1:1.7.3-6.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "maven-shared-utils/3.3.4-6.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libcurl/7.61.1-34.el8_10.2.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "harfbuzz/1.7.5-4.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "httpcomponents-core/4.4.13-8.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "plexus-classworlds/2.6.0-13.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "adwaita-cursor-theme/3.28.0-3.el8.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "dbus-daemon/1:1.12.8-26.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "dbus-tools/1:1.12.8-26.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "curl/7.61.1-34.el8_10.2.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "findutils/1:4.6.0-23.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "elfutils-debuginfod-client/0.190-2.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libgusb/0.3.0-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "device-mapper-libs/8:1.02.181-14.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "platform-python-pip/9.0.3-24.el8.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "avahi-libs/0.7-27.el8_10.1.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "java-17-openjdk/1:17.0.12.0.7-2.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "plexus-sec-dispatcher/2.0-5.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "jcl-over-slf4j/1.7.32-5.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "jsr-305/3.0.2-7.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "python3-libs/3.6.8-67.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "graphite2/1.3.10-10.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "maven-wagon/3.5.1-3.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "jasper-libs/2.0.14-5.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libXfixes/5.0.3-7.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "nss-sysinit/3.101.0-7.el8_8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "shadow-utils/2:4.6-22.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "device-mapper/8:1.02.181-14.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "fribidi/1.0.4-9.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "google-guice/4.2.3-10.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "apache-commons-cli/1.5.0-5.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "glibc/2.28-251.el8_10.5.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "guava/31.0.1-5.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "cairo-gobject/1.15.12-6.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "dejavu-sans-mono-fonts/2.35-7.el8.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "jbigkit-libs/2.1-14.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libgcc/8.5.0-22.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "adwaita-icon-theme/3.28.0-3.el8.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "gzip/1.9-13.el8_5.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "krb5-libs/1.18.2-29.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libssh-config/0.9.6-14.el8.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "platform-python/3.6.8-67.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "cairo/1.15.12-6.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "glibc-gconv-extra/2.28-251.el8_10.5.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "ca-certificates/2024.2.69_v8.0.303-80.0.el8_10.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "python3-pip-wheel/9.0.3-24.el8.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "xkeyboard-config/2.28-1.el8.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "java-17-openjdk-devel/1:17.0.12.0.7-2.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "nss-softokn/3.101.0-7.el8_8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "systemd-pam/239-82.el8_10.2.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "cups-libs/1:2.2.6-60.el8_10.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "dconf/0.28.0-4.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "gsettings-desktop-schemas/3.32.0-6.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libXft/2.3.3-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libsoup/2.62.3-5.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "libssh/0.9.6-14.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "sisu/1:0.3.5-3.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "nss-util/3.101.0-7.el8_8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "maven/1:3.8.5-6.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "nss/3.101.0-7.el8_8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "hicolor-icon-theme/0.17-2.el8.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "plexus-cipher/2.0-3.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "at-spi2-core/2.28.0-1.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "cryptsetup-libs/2.3.7-7.el8.x86_64" not in data/right
2024/12/17 16:46:25   Deep comp name ver "plexus-interpolation/1.26-13.module+el8.10.0+21301+657f54a3.noarch" not in data/right
2024/12/17 16:46:25   Deep comp name ver "jsr-305/3.0.2-6.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "plexus-cipher/2.0-2.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "glibc-common/2.28-236.el8_9.13.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libtirpc/1.1.4-8.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "maven-resolver/1:1.7.3-5.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "expat/2.2.5-11.el8_9.1.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libssh-config/0.9.6-13.el8_9.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "maven-shared-utils/3.3.4-5.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "python3-pip-wheel/9.0.3-23.el8_9.1.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "rpm-libs/4.14.3-28.el8_9.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "elfutils-libelf/0.189-3.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "jcl-over-slf4j/1.7.32-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libxml2/2.9.7-18.el8_9.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "maven-wagon/3.5.1-2.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "plexus-classworlds/2.6.0-12.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "httpcomponents-client/4.5.13-5.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "java-17-openjdk-headless/1:17.0.11.0.9-2.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "redhat-release/8.9-0.1.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "file-libs/5.33-25.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "apache-commons-codec/1.15-7.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "java-17-openjdk-devel/1:17.0.11.0.9-2.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "apache-commons-lang3/3.12.0-7.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "alsa-lib/1.2.9-1.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "nss-sysinit/3.90.0-6.el8_9.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "cups-libs/1:2.2.6-54.el8_9.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "librepo/1.14.2-4.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "ca-certificates/2023.2.60_v7.0.306-80.0.el8_8.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libX11-common/1.6.8-6.el8.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "nss-softokn-freebl/3.90.0-6.el8_9.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "avahi-libs/0.7-21.el8_9.1.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "glibc/2.28-236.el8_9.13.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "openldap/2.4.46-18.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "audit-libs/3.0.7-5.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "glibc-minimal-langpack/2.28-236.el8_9.13.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "plexus-utils/3.3.0-10.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "curl/7.61.1-33.el8_9.5.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libgcc/8.5.0-20.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libnghttp2/1.33.0-5.el8_9.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "p11-kit-trust/0.23.22-1.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libX11/1.6.8-6.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "krb5-libs/1.18.2-26.el8_9.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "gpgme/1.13.1-11.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libsmartcols/2.32.1-44.el8_9.1.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "nss-util/3.90.0-6.el8_9.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libstdc++/8.5.0-20.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "platform-python-setuptools/39.2.0-7.el8.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "plexus-sec-dispatcher/2.0-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "jakarta-annotations/1.3.5-14.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "httpcomponents-core/4.4.13-7.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libdnf/0.63.0-17.el8_9.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "nss-softokn/3.90.0-6.el8_9.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "python3-libs/3.6.8-56.el8_9.3.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "slf4j/1.7.32-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "glib2/2.56.4-161.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "rpm/4.14.3-28.el8_9.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "apache-commons-io/1:2.11.0-2.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "guava/31.0.1-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "python3-setuptools-wheel/39.2.0-7.el8.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "systemd-libs/239-78.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "maven/1:3.8.5-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "atinject/1.0.5-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "plexus-containers-component-annotations/2.1.1-2.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "apache-commons-cli/1.5.0-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "google-guice/4.2.3-9.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libcurl/7.61.1-33.el8_9.5.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "p11-kit/0.23.22-1.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "platform-python/3.6.8-56.el8_9.3.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "cdi-api/2.0.2-6.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "sisu/1:0.3.5-2.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libblkid/2.32.1-44.el8_9.1.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libuuid/2.32.1-44.el8_9.1.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "plexus-interpolation/1.26-12.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libacl/2.2.53-1.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libmount/2.32.1-44.el8_9.1.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "jansi/2.4.0-6.module+el8.8.0+18044+0a924b8f.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "libssh/0.9.6-13.el8_9.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "java-17-openjdk/1:17.0.11.0.9-2.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "maven-lib/1:3.8.5-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "maven-openjdk17/1:3.8.5-4.module+el8.8.0+18044+0a924b8f.noarch" not in data/left
2024/12/17 16:46:25   Deep comp name ver "nss/3.90.0-6.el8_9.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "findutils/1:4.6.0-21.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "shadow-utils/2:4.6-19.el8.x86_64" not in data/left
2024/12/17 16:46:25   Deep comp name ver "bash/4.4.20-4.el8_6.x86_64" not in data/left
2024/12/17 16:46:25 1 image only in  left dir "data/left"
2024/12/17 16:46:25 1 image only in right dir "data/right"
2024/12/17 16:46:25 89/90 images in common are identical
```
</details>