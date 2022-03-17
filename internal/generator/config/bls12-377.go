package config

var BLS12_377 = Curve{
	Name:         "bls12-377",
	CurvePackage: "bls12377",
	EnumID:       "BLS12_377",
	FrModulus:    "8444461749428370424248824938781546531375899335154063827935233455917409239041",
	FpModulus:    "258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458177",
	G1: Point{
		CoordType:        "fp.Element",
		CoordExtDegree:   1,
		PointName:        "g1",
		GLV:              true,
		CofactorCleaning: true,
		CRange:           defaultCRange(),
	},
	G2: Point{
		CoordType:        "fptower.E2",
		CoordExtDegree:   2,
		CoordExtRoot:     -5,
		PointName:        "g2",
		GLV:              true,
		CofactorCleaning: true,
		CRange:           defaultCRange(),
		Projective:       true,
	},
	// 2-isogeny
	HashE1: &HashSuite{
		A: []string{"0x1ae3a4617c510ea34b3c4687866d1616212919cefb9b37e860f40fde03873fc0a0bf847bffffff8b9857ffffffffff2"},
		B: []string{"0x16"},
		Z: []int{5},
		Isogeny: &Isogeny{
			XMap: RationalPolynomial{
				Num: [][]string{
					{"0x142abb491d3ccb00d65810beba93dbb0a661fd85974d6aa82c4bb2e1a3c84ffdd6ef419b80000000000000000000000"},
					{"0x4d9d782ee8a7b7630cd57be9a2ca555e2f689a3cb86f60022910be6480000004284600000000001"},
					{"0x142abb491d3ccb014ac44505178f6ec539a237640b7ceab573689a3cb86f600114885f32400000063c6900000000001"},
				},
				Den: [][]string{
					{"0x13675e0bba29edd8c3355efa68b295578bda268f2e1bd8008a442f99200000010a11800000000004"},
				},
			},
			YMap: RationalPolynomial{
				Num: [][]string{
					{"0x142abb491d3ccb014ac44505178f6ec539a237640b7ceab573689a3cb86f600114885f32400000063c68fffffffffff"},
					{"0x35c748c2f8a21d6af848e30c1b78229a46644922460e73f6faf06c327b438084815848140000010a11800000000002"},
					{"0xd71d230be288756a6446249c205dced645709767bd81c863eb7f8d8e4f15003f5f407b84000000a64af00000000002"},
					{"0x17872fd54cc6ecd6d73a5085f0d2013b6de7eb4a0d6711d3b14f5e9c2c81f001429f19baa0000007467a80000000001"},
				},
				Den: [][]string{
					{"0x1ae3a4617c510eac63b05c06ca1493b1a22d9f300f5138f1ef3622fba094800170b5d44300000008508bffffffffff9"},
					{"0x746c34465cfb9314934039de742f800d471ce75b14a710033d991d96c00000063c6900000000000c"},
					{"0x3a361a232e7dc98a49a01cef3a17c006a38e73ad8a5388019ecc8ecb600000031e3480000000000c"},
				},
			},
		},
	},

	// 23-isogeny
	HashE2: &HashSuite{
		A: []string{"0x152964189f4c623685ae0423eb10294ce6458c064f093208504005b37d04d5d336dc9d66a97093f84d62e778f8c82be", "0x735c455387ab435839e5a5dbc1a30510070300f4becac797642fe56985064e95f7d6521a1a6e71004047f835c1f957"},
		B: []string{"0x19e38372e0d4bf401d2fa5f2261e1e3fc95d51a3857fc23b1385d51ea9c973a89c22148a93dff96447700bf1c3aebac", "0x1579ddb5c1c595b7c08c3a3cef5626143c25757c6b67d0a2677b22fc0c890d8b2b1a17895d047a98c49047069f725"},
		Z: []int{12, 1},
		Isogeny: &Isogeny{
			XMap: RationalPolynomial{
				Num: [][]string{
					{"0x113b0abb7ba48832ffb7aaaa7ce085078312d4bf0bf8882e8f4a0a6e24d91b535b6c81277ad9369cacc733de5cf86d9",
						"0x11e62d211119d8108bbf13b3b4b79a1aabbe2d6004858109139667b300e5097370015d42782ca1bf7652dba1d4cac3c"},
					{"0x51a19546efd3c582ec16ee049bafb8ccbdd87d98f753e654e0b93ae62a627a6f610b30f76016baa8d18dc4677dc401",
						"0xddc3aac50aa5af86f8c1d31db787f4b6ac0262a8ea40c07d61389c2a70b06c51fb59520394e05c55c72e363b9607a1"},
					{"0x93198927f09c55a596756ab909a06ab3c0127be6538fb275f2b7b5bfaa1ab85b12a45ef345d628d5c6e69effd7e76a",
						"0x2c02fb38b1440c7b2d4fae061763cd734c40b1eddf1ca9552554d2859906f8d6290038a485a655f8178f99b1fe43e1"},
					{"0x19091208a40d61bd9a55359da28b532617da69a58addac35288930ccd5083166d791cc7dd3c6533e01f786000c532f7",
						"0xcf913d369130273a81f982f469bc8fd98a8983af247d28845d76729468a777836e1d3cdf1e1adea6b08566db5b514f"},
					{"0x181a36b975099a83b4773c6e62d93365d5fd3787c2b62d0a315245f6226ab13cd462f731b73a436a598b34dc0ec58d0",
						"0xdeebd58f0234fa272439c69d9f0ed559b955d416256645c2ab1857d988b49109ff460206e2b978d877fc9274fe4840"},
					{"0x10cd4f85b581fa9bbd614f001a2e64839a4635373187b57373f17f77e11ea30ce505e6ebab860ee2aa46c9e6b327add",
						"0xc1f11111a68f3346dfeb807985c61491498efd3439c5c7a0b77e9d819a48be950e43c9f30cd21e5e439815a2ab6a9a"},
					{"0x1813ccf2b39896b88561d57592c5e7d7594146f9eb341852cdbc192e88792058059bd79b248c5b8c415d6149d51130a",
						"0xeb958c5ed40e2b75a170536e42368207df82fcc26518f6f83f40cea65efdd81dcc248290c660107e0a2d650a38d526"},
					{"0xa143ff2d94ba2d2322d63cb3769874b17c00fe76fff4b53fb5fa4744da3114899adaca240bde7d88357aa80b144bb1",
						"0xa25054cd91b2b06cdad439cdf86265929d35738f09e759a56164861f45586602c630ef182e9262f0413374b154dd20"},
					{"0x1f0a897c94b01589c9eb5cdd999c580c284f39c6fa6f2ad0c569c3b219cc245484ed7310b858f980b2917508d94f59",
						"0x5c56eb37beca858c993f39fbf09b73efbd98934b5b1d8de0c21fc75d20075e6d7720601832565b1823339bffc17a6d"},
					{"0x13a21ac110c90240e49de23070e86bcb64b8c500d8f9d792095a0a7f9133f4d8d94c518f9eb1359c8c82926e41d004c",
						"0x18471a982b5d8d7756ff4971fbba3945be37cea5cda74e17674b76cbf527f8ae51a27236a95d815a4f36a2461df5b67"},
					{"0x380e386dce2ad72755e1d21461a3d33ce275b036c2a11a755df7e948161f3318531311b8dfaad26a2b6af699e42f07",
						"0xc86f279fd500d0594160f87a462ff89f4b0db7b7fe6bc809871f11183f13235d343e0a8bae13c9f5791a2821ce8bf"},
					{"0xd73c34c8f904e327cd1fc22e8a90e04bed13e46eb98e3d47d18cde142b9e0f3a03eaac69214acac03042aa9b9b6495",
						"0x76ea3cf6b7185d7d6e3e73ac03a21b645e5aaee7a99dccecc485620fdf64e7486c7b1e4f8b6e069240b7a1e76b0609"},
					{"0xca983f9f6ea728c5b4bb9a6dc72d24ecab7833e1a8242bf1a84140deb996176a8f7c7d6fe34fdb4660afd6dfa5fd74",
						"0x7e4badbdba86487f1b87daccf8d14f62bcb01bca058d8e318ad40ae8343a47d40f6b18af1e0bf9860eeb32db4f02aa"},
					{"0x1552beaa2371b711181b528b4ea3e801e3a71c2bc0bde70a2babef6d5702ff12af4f5f2823599093d5485b94c873d36",
						"0x11783b91dc5ce3d03283278116bb901408dc61ef8bad40991144de2f3509787d83b566820fd315b558ccd76c9debd5d"},
					{"0x175374ccff26ce81965dce12856efd5cb587251e70a588e2e4633110d65bf17f2991df664e727ebc65b42d74a938e4b",
						"0xe75b2e26120b6428e51285251602a89ca9f865d444853ac3527b940d7b218df89806e725ec52e75b92dfcf63616483"},
					{"0xecfa3729a48226a477e72d4067aa372a29fa8943324495f2f39f754bd1d114fc2d9c62453fde06e3272bf09faea06f",
						"0xe95c673c49bda6dacee18f07327cee3902582326fbe658d3e8cb15b40327b8b7c1b04e0160071f1e8c617361754705"},
					{"0x155dd556a15b55b766cc98f3455788f3ca450ac5c613b2e1464b53f0e5412a626e423cf9dd5127edaa6deeaa5638f93",
						"0x15ac434170f220d28273b44c48eca01f2cfaae20e2342429ca6f3668e0825821a3d7f1227db23b1a7a36e7f8abf89e"},
					{"0xe785dde291472723e95d737c7bc0bb72a8125cfbd91b122beedd3f45dedefc1480e9350bbfb6a5ca2fb9ec905573fa",
						"0xe01e2c8ab4ddf439c8d4028ccc68b13de8099e3aa7d5120867366605682a4db4ec308cf07e043cf7256e35c23ecbc7"},
					{"0x4e2c05f403fa31d03e91053fb52a3fc12ae58977575f53cd8a52b53d1067cc61de65daada7c7c1bf5bdb07f3719b14",
						"0x18de34a92f81eba3a57d55a9d4bd5cf7da7b1794a0f9ece37abd537ddcdb29edbb01aa88baef416408d34a227874c33"},
					{"0xe4a28f5d9cdcb181099868140cf3064cd573f892aceb79e3df63546686df19a3164efedc851333bbe232be84e3e627",
						"0xd31f8027ef5f6b916364d998dbc4640196191cf0ca44e2c2c4825bf4babd706351d21f1f6177ee4c3f5d16928d059b"},
					{"0x602985c2bf7472638e59af721dc3e635d08ede15ad23084d54441a493d515ab9d727832a06110b8a5e4cdf978980d1",
						"0x159cb19fac8df89c704efce7423d8965e1477606471d21f77988cd744d0ea0938a85eb4a21d41a34c97d64a409f7314"},
					{"0x12f6020b2f6dc691f96e6a7c5e42a221a861c1b27223e30adbb75022ff42f318e8d96e0a964dfcb9dd0d20fd3a7c247",
						"0x1766ffa6c449787005d03e90251c271190a3c8e6db8383ce2925b91c87ae29c989d503705ed7566fe5c7bfefb6902c6"},
					{"0x471dda72677bf4df7027fff15bb88083362b56b41dbbeb3e3f9afed6c7fd9b8e03a7b66f7dde784cdb94ef12714ed1",
						"0xb95429ba2726547f58f812d8516a9c7545fb43d2080d1b154e4e8aaa689a378141fd28ad946cf12fef100e180c9762"},
					{"0x1ac99df3bf98db5235d2892b0f00a86db204cc8f367535d37ff878fd945bcd6f4991d08e3d7598788e2f860c950d110",
						"0x0"},
				},
				Den: [][]string{
					{"0x146e50faa64a7651c3b0b61836501a66805927dcb051b32662d0d261c3b5458776918917726e7ab20050fb3f1e9d32c",
						"0xb1ea777f7008d8bcafa4799bac2e33475d6b287b6c831af5c06bb778507412d8aa8d347a14ea074b7983518d668616"},
					{"0x79f0dba8c34b89c8d1fd2363421e1f97f7f12579ab42f520b86c9e7b44e822fadf62f6e98bfebed3065ef7e6d873e2",
						"0x40dd9ff7e21f5f6802ef1e47c240880f84fbe0232f659145f22f9c842736890572e9a0018ff427005b45140ad2597d"},
					{"0xde82c1e45e55039d5d8ae4684e278f230e94285b4bc97eae01f84acf258e9be8888a96dd0fdccde825922ec5517ec5",
						"0x101b161db02dadd38c1b219d253f01728d2b3ba65e3dc05ebf102ff9fab2f2ef1cc95047570148228b4915820b0301d"},
					{"0x19c6b7b7a90b244a9cc8e9cf77c1fc40ffeae17efe57efaad45c6b85105effe955b8aec25cd5a33b5fe1b403099cc87",
						"0x1658235a28d3dcd62a783eb6e135a65f19ee1ded96bf717f26fb9e81763b6d1d0d32ef686998776504a5c6426ba2ddc"},
					{"0x1374e714d24649d3de48ea7e809ed20ba8de79a7a95b59e26babe22856e57c5f9e6667da8107e27b1215afe2846471",
						"0x9fa434f0a714ff9b2bb5ea51856efce45659377666fe4377774d04ca63a1ae957c5cc11376defad113c94ae6ce4652"},
					{"0x11796b770f7504fb0a25336ae91aca083bb600d5523a27885e44f3bedff57966411346d44cce4bbaf50770a210445b7",
						"0xa5dc732bb3f95325df484711467e89dc69a7986d2439a692fb9667c22512d87f968d452c1a7a4af94629d6fbfab126"},
					{"0x36fc8fd543e00229e3e56ca9655db01cb161ad493c35d5225acea1924f2673890b0981feb9170a2b73e1ed5852276a",
						"0x11e47853a548367d9473b4251c2641a1a05ded1721fb74dfd4455bb41894cd35ffa77ac1a1dda4171e8e3c598cc0cf"},
					{"0xec3aecb4916b2621d73c22adf37fd2bd067cfcd09f38cfa7060c1dc08573f6e6da62ae459d00c719b153b23eb30380",
						"0x33b7e5c5174d637291e17246bd01864eaf2b12d27ab74b84e029ed82c9cf3551faa44b79bdbb079a9f57f598e8c35"},
					{"0x14d8b5c3fe739ff9bca51db2433814bd39a0ba8f082801125801599b81682fe51896aeed82936cfc3d7850d2f9622a3",
						"0x7adf53096d8b637caa65a6703d654842a9b7d4d82a1ebd6ae4eb4cf7df07912638d359dbfe6f8443fcae492cb150c5"},
					{"0x16e8eda7dfd1927fe1dc22dacaa78237dfa0714315f535e804649326acfe5c67a86dcbee4d78a5d662ff122b8580245",
						"0x40de374242af4c470b0952cc13dae780d77d43729e9598ba3173ac586624ebe305c8c9e6e5e7245f98a970af1ee60d"},
					{"0x16c94723b3eacaa13fef87169f4eca0959eecf2184d6a7468da9d8f32acb6a36c5254bf31d733f5b1c053dd88e54282",
						"0x19c508401f6e0cdcf9e2b694c12550eba4247c85c52b8aba5991f9aeecbefee61355271663b842b0b23c0de66bcb8f1"},
					{"0xf3ffd87f27eee739719fa368e3d9bebdb8f06a192b5ae980b629ad78421034b2f919d658fbcd1490fc7692e1e916ce",
						"0x1468f37ed6477c421170d67235b0370aa3388db7e4c0cb3dbcf1fef55b2c7cd6989451199b411d7aaf8a5b24436f6e3"},
					{"0x1035e4ec93dc6bc358f961e7dc4c82087605cbe2bfdbe713fa5e7b21d75c1ad45088bdbc6daaa2e12d1f0bcfb4b84f7",
						"0x11e842bd544d0bbee76808f0a1ba264be31aedda57ceb64330b587d5a89d03598cdb8d2b528fe56a02dbfd48e8aaa83"},
					{"0xe57bf7b799080766c07fa8b997f65da74043a005d785b60215c0875d8d80a2a55de8c81dc081ee75cafdcb6d9f5bab",
						"0x4f2c024aca0ee3f8be88ed1450ac2448d3bd8b540d7cb9b5d93d68eedc20119213d260c02c7dbaa180aa008407b0dc"},
					{"0xe0b435cad88cfd17508b5b3da274b76521f715b9129c3141f70a3c7dd22e14c3fe58a5db7c168f47c673d91a05c02f",
						"0x16ba8b632d08a5abe396f990cfc790f6d29725220c89e7dacb3be59699757c6245bc0793bfeb5b2f917508cdd949a79"},
					{"0x17c5e37674a25c3385908a02be0a23c1623eb710deb01632301277cb968ae2bccd3e5b38e144e424a9da459c4b8600a",
						"0x503556d7e47b59d5888f787cb7cec46d917d02f8118e01a63ffc3b01398260d7c722dc26c7cd03d7c05e1a996e5471"},
					{"0x890e4e58ded34996cf252dc184d7aec1c1866ad2da0b50ffc208a1ef31dd9e277b3f5e2fd1f3cf9422b81c526bbd0d",
						"0x108938787e3377815ba98065f7c57b955fb7bc2247a21fd516fcad09652e92819c4bd34cf9be483e261639e540ce577"},
					{"0x665b1acc1e271d15bf9416f014ae1ddc76838a60230be0cd17cc7505d6b5f7cb01037545fd2c3b9997b0745ad8ad9c",
						"0x1ae0fc9e185ac85323228f01780e30ba86767c59cc90d2d05ddbba8d6cf35bbd737b7fcd09b3850af513ebb6fc3ced9"},
					{"0x16db2dd779c2cc1699abf727df65339443007994b70c43a8d9767e265dfe45809093cc631ba5b79f1280236ac93fddf",
						"0x43671e62ee377b90bbd80c91a3dc1eb53ed37ea8ae6ccc0154e2f1cf498198acc92b2ce3c07e292b2b5f038c6c04c9"},
					{"0x93b2b87c160c52f6b056e4c085930add9aeb706709ee013eb54ca60bc820b0da42fbca8e9c9a579585bdb8af77acbe",
						"0xcb5381cf50932ff1dd7fed460732ae9311b24ab543d98304b80963cf706598ef147c2b047d6c2e49357f5704e568ba"},
					{"0x731b19b58014b17a0be1915dfee088ceb8ca7964deed74d985adfbe0e9aa99cac1453ae4210abf949afc9c7425f2b6",
						"0x3f906b21602c1254dc293fbddc3974f24bef13d0291760083c7a125c8ad5c49ed9ba755743e9344387d1e712d7ede2"},
					{"0xbee29453de653858b3b37cd3b85e7941177c94de27ab781cb8ea619a1b9668a0b0ddd01318b6699bc582051c23dd8a",
						"0x1793e129728346d60828550ad395de628b4f4ab2952bd0e95003c382bebd49daa486d981fb51dee163ded1fb204d09f"},
				},
			},
			YMap: RationalPolynomial{
				Num: [][]string{
					{"0x19474819a519191689e0816c7f84b2a58a6a5aaf5adfd7fb0204d8623f67c19e9a64d522ce4c963e32de8edfd367284",
						"0x100e238102de877617418a44564d9379372b6f69c721b205c60af2e298bfa546a7c7a39883e433f947601b24266a306"},
					{"0xb58b352e8f9bc6162e40b4ab27f066cb71ed844e9b5e376cdef8e8be14e38d2f13392f19da19ca1faa758bbb083dc6",
						"0x3293cacae9e6251abad8ac4d8263754f1301062d54ff721d183d201f9133db07341cd48c721d8b9ec40e41d4bd0db0"},
					{"0xdfec4d8ecb1cb0adbe13f68ccb2e070bcbf5ee3ee8328cb6613a63a543c252ef4ecb525fbf65ca12f33b425da04d5",
						"0x177fa15abb4e47ced85ae0765a6aef971131743de887073556fe92f142b83a8355edc8544530ed7296a3e21c4f9828a"},
					{"0x17bbcf8bf3802b3f868ec94dcbb32a8b0cfd5f62daa2079cf8370b487cf440c664863e710ffeafc40cf374a58263139",
						"0x16f6e49903922180e44d4bc45aeb65f093fe4fad10cf29583c3664cfbd76b9ca5b9a7ef2337a62cb2e45623cee7ced3"},
					{"0xc1fa03051d6150ed02ce658113c5b9de9232db80782740741c13e95afb61ddc78d32e0af5fa67e2df4203ea7c7f720",
						"0x1a20dfca20c821fa263e1b04009e6e0f5a373e665afe7f603f0bb9703ce6864cc9ce09aa1e2d694a16082973591eed7"},
					{"0xed3c8528803a2e8ea1d51fbc40888c93aacdbabaf485d5a8a037472849450ed254f2bf0876516c81b466d03adc0b37",
						"0xe0b3c359ae53e4901f777b107add779bd384d8da8a46c02f044cbeb099743b6888cc191d97bd035de9136783ac7866"},
					{"0x108e5399fb0e462641e346430020b47260782a557c8476b06aff3fa95b6251bfbd40dec8c213b3355c50a24b96d95d5",
						"0x1aa7ad61ba7c96aaf8b72db59b48309bd72b50ec5781988ba9a55a43b4e4a595e36aceec3a51a1804a788848e9f96cb"},
					{"0xb16e02b02354b93cbf74b3551a94aabd0acfea5d6f20c56eb2f46433003d2dc4225912600eda79dc7e3adf0702e763",
						"0xec85210b0e638ec662fc76335f470cbe21a165a66aad036058932f6d52e77f1b6e3269cce0432e1404b0810c053c2c"},
					{"0xc0b22a4dda3741b4fcd5b6c2abb9d2e8481ebd442e00d81865d70c8717d4c6c78a41e6f85c97b10fd010d8a966cbaa",
						"0x1396ca7d90b82b12adcfcc5d81a850203097d56d9ef35381dc666cbaf594b4e033e6a525460e4ff6ba10bbf1afd94ce"},
					{"0x14c78f774a2896fa64a052004df6fb736bf806e662e9d7097f589b0df03fee5aee9609407005296ea99b879fa0ab8ac",
						"0xcd6467d031974a892ff5ecff9f2f8d46a53e674500111351806a15432a6cf8744725b9c101161d59c792807b253a3e"},
					{"0x1526589e372650ab85ccf35382db0e7cf5d02f0de09e848caa73396338a01b9f6064695efe06eabef61098ea1c0c638",
						"0x68acae9cd1355b5c9781656519838b804fbeeb960df4bff77dd2477ef674d7240676b62cb03a51326f8f47f14377e"},
					{"0x167a5b9940feddfc5ea3bae9c03a52e10fdb3b7055e547d27fa804baab5c74796d2f1cf56675c664229ecd622f32ed3",
						"0x44f0cdf3744beae8718ae72373f6a01838b210bf23bbc13d3522d5595f45980057324c394fdd5f4a588716b2c0cf2e"},
					{"0x1511bd492587501195e9a7df7360c8e8fe72d53b57cec62eed296b202b156724087121691d7e7ac6d80656b40604602",
						"0x51219c71bb96909a005e363c5ec60f36c77e407fcc05e801c3962f83566d18ae9ff4a658d0dd08e1d90de6c758bbe3"},
					{"0x48208860fc71ea0cb4a266e00c3ba02f8c272902831fc7d0a26d4734eb1fde8d496dec19ef8b076fde5c77b534f4ff",
						"0x6707d246dda7a0a59b2b054307653dca9144172aa6281893bc9f681368f71325baca1eba0229dc19f9a7853088ca8c"},
					{"0x32a1895b110a91f8fb9e1bd01a312d28281131018a0ad9470c3c3ea1952f63ce90112b8d5ea7bc933d4f1cb426abe",
						"0x6fced059c7cb1819994a13bcb3d805cb1170eb16e35561c90c1454c4af849c5eff11b66554ee507fec5c0ebc4eebf0"},
					{"0x182a668115cbca00c180b580798d261e58b4d2c140ccb68342ee5a859b51af3076e871060c103c2a0b4ca8a80a43a08",
						"0x458656746821dc18aacaf7533ef33077889590c4eb36a545e120d40c812dad8bf3c8b3425a427a0f58e49460a4f7f0"},
					{"0x1371823e71e71d670cd8509978788f70e4aa0ad9e88dd3983c8b52a688c11865b8d1a183a82c866cebeaa40f4f87295",
						"0x3fafdb6deb055c9eaa8a59bb2fafee94632d4c538c8837cca7743fec86a8f9ce24bdbf3e34595d1bb9e43571cf2370"},
					{"0xda0900c121f7a9a289faa9bd55ceaffe9d93a38a2b0890875289abe4a4491757eb96eebf7c2babb6169994d7eb414",
						"0xfaa7103de621a04fad6afeb3e1334e1fe69a19029bf77e32833684a5e6c9f1a9fd6cbd6efdb6496b95fa4b20521f9"},
					{"0xdf001ab5b668caa34f373dea333476adfb1243e35453af80a6e66dabd825256f185763cb22112686723a50c47acaba",
						"0xffdc67df2a8983f611da939f6b81abc4db37ecbbdb4409abf9e1eb6f763aaa51fce86f4ea92e4500079313a249d7a3"},
					{"0x2967da20443829494af8becbaac23650aab235202a7a20e64e52e0bfa17d045a066167931e3dbd41d55eb956d2ed10",
						"0xc7f02fdee5f78a0a03e3abe964cb50d03313df2c8fe42b99d16f3af8a2a0bcc1a88449924ebe0acd21bd3ce63128fe"},
					{"0x13cac271335756e85815b2d299b18f25f34c16d33b8ff67e923b74f5c468196cc31ab24a3a7ee56ab9e4e10676fc0eb",
						"0x18a1b7092e497c2b1ce9e73f447618dcb50b43e20f6115931833d0d1e79da6f3d728e4a1b651c0ff536cf2549919282"},
					{"0xd4ec1a90744458ece499140238ae77c1de76f242a8a05f1482eb221783eead306e1ee6c15da37f4363df82b6cbfd56",
						"0x7eee2877ada79fb65b13e1af4d49a98ac24dbe7965a7dd9d4150274a31032ceb4a3a0aff0369ce7b15235d91cbbc25"},
					{"0x2b277ffeee377436109a9f871027362a333ae2d49c65c7452be93c50a2220673b21006879b1fa9b10c5c5f0d933f38",
						"0x9b2ffda579da5430137b22b26fdaa13f52ff11c13f4f88fac0bfa0c30b5c1bb190d8728a60153ec7986dcd7dc5640b"},
					{"0x949aedd81f20923b0b0e5f5cbfc2c0cf7ba155ed488e756404f6adf9ad0c85e8bbe102a83693e805aff42b7f21617d",
						"0x176910539df9c32d55b92b8f843c668db490d8905e108f18d1b136d45989a2d2ef4cfae817975c2729cc92032d905eb"},
					{"0xdd974b42d37d78da2e8761e99596fd113add085fc115b9ca57a94c218b278addb80c8430bab38b8d53738fbb33b75c",
						"0x163fb82f31df72b309c3422ad404ce953f5fb5292fe68cb61b73100bc262a987e46299e183fc582cca8bc848dd7c621"},
					{"0x15d0fae7eabe20ca577fc60964b17e021e699c1f8ad28ecae4f3766dfcc924a5a89dfc48fa0945e9e650c45980affc3",
						"0x16d1223b7d1b6311f2462306650b490d11b3772de34b66fd06810fd3dc56e90d15382767477c4952d25d0da2656d6e1"},
					{"0xb34a5e0ece8efa676afb297e8a81b8c5af019a878bee197bc4ecfabfbdaa46da1ed1a51cf3b0f1ed90691dcfd2d594",
						"0x164c4047b6cc6e62f7c7a00c215157206c7b50adfc52435469ca2846701fd8857ca14c5ce17fde980afdf2b8ea58d2d"},
					{"0x167fdf7ff361330674041cc850f227daa736a1b450928a621273de03cb104998d76a529c84a9cd93a91ccde3cb4623d",
						"0x18a6f2eb4840007bb2157cd6809ac05f522ae2e4047f77f4de06732445c7f03290de120acdbfe3d4590edf09ac3d0c"},
					{"0xf8649870cd6866641205263c10df90d9ba3c8b56d8ceb178c4897ead46bc3be88ef7e0e31952dd1396ac6f7905b8c7",
						"0x9a05a5b1cba4d191bc5098efa547c070297e57748416d8386d72875a5cc5127d3a9109b88bb74bd7552b4eb3106175"},
					{"0x1378329c9fadb2e63f34a6f55ab5335ff1b5d913c2d7edbbba48295caa230f00150d65d25dd44880cb562ca912495c9",
						"0xe5f0b2d55f8713e807c8fb1f96f13e0b5caa3de8607f04bd93a7260e3d8810a48860a9626c7b4a424bdaad21741308"},
					{"0x11386a65a011fb28d612fbc7d0f76b2dec1330512b504f541d09517a0cc00f0f49b6fc1fa52a8725965385dbae1f4b8",
						"0x170632bca9a590287a0ffac955344cc8f15e59118f03adbf86e07bb82f5bb76cfb7fbf1e42a1eb384ff9c11f21b910c"},
					{"0xded86746533417ce10e14535b4811d2246eae38429763e4427ffa9dd12120df30c8c36497d089a06066626724c6cac",
						"0x40ad9ca1fa683544a978c718fae74751b37a5dac2b57dcac813479f9ebb387e9ac1eafb90b12483b5282391738f286"},
					{"0xdfda607fa565d6becc5842deda97367bbf14db84a0cae4783dfb89d2ea2ccedefca2ec5883577294bba68a29815d6",
						"0x137601d38808323444ab0018632b9e567be332256b624b8ce929756c42eab72fca278590b52071b38b711be22f5b0d5"},
					{"0x13dec721e4606cc9691488c3d003366a1462ba6a713277de0503f1131d1d61fab5c79cba5bb7332c6abc80435464be0",
						"0x0"},
				},
				Den: [][]string{
					{"0x126e7ae82b984bc44d0e37b08e50643a99f06396f1ae9e20117c4bd86b9b1939fd92c3d95ebd307fce9080a95674e4f",
						"0x18561a650deca0ad7df837e9ffaa27f0b8e85ff347362f1259d3b1e529e892327aa7c6ccf4a3ad6e365226b8d875fe1"},
					{"0x1a9963059b91cc5aab175a23fb536129d851da9e2f8c5bb7165e05937bdca1db1508e98843f1f62c047ce3d6ec8fc60",
						"0x1fb73b0a8e121906f51d20eec7cb8491cb61bbb3b3ecaded3df7283d9128fc3635307f3c3ca133af720d9775729e91"},
					{"0x8b9ff20c9eba7b3dc948e79ce661704cf236ce6bd8361d146491031d3fc9e58b338a67ddc8548f4a36460fb5dceb40",
						"0x3c27d81172a310833e5ab3bf81225b01c4068b13026ca1deeebad80a558665100bad779466ea0bc4ccf4c4ba3d9c7d"},
					{"0x1642ab99157e12e44ad3e54b1fa7a4ec2d80b545653fa336b5aa5deb0452f3d47824cdc51de8d83775d658426e24e3e",
						"0x7d42c28d3d9fae9c3286c2dcb8826e91d7b36fc45468e8f64b8a029c91863cebd4add2412ee5f3bd090bcccda03ab6"},
					{"0xb233cdfc79172767df7df3731177932ce335fcff5e0b1579c01cf6d34d3747bebe1bfe3b111ef9176fbbba581e459e",
						"0xbe45fabc97a7c23a215db4b2aae647fa5749ad332d9a66eccdcd7195a412def4d28ecb256f45d5879cb6b4e6e1d396"},
					{"0x4684f66af9512c51edcd89002de77e666b71ff7151c341319a44095d8113addf41f718b54842a7c29432ea9e2b8b88",
						"0x88b603e95e15fff2d8bd1f77ceaf19ee06b7a9301faf29d983a8962bf7969a6b3ddbdfea1ee16f30cfb81cf7e21715"},
					{"0x354c69013ae211b621a899ca7a65f4bcbea4d39ea633002758cb777ba71f570995d52adceb2fc931d92a12398eb12c",
						"0x175ca20c6e2e1b0b8305dc90cced8b1e4c3f4493d341e46e39b6e33741b604f1ae9f8b2be59f4a9a52e948c802a108c"},
					{"0x58a2ae5eeb2f349aa139ef3ddffd1df269c2cf99e23d6f64ed7ee9bd2cbfde0027449827b1e2b116da86ec051e0f07",
						"0x47776d799a1f49c7bfa7feebaf8012282bb0dbc4b1ce41448474e41e53f64a098c71c79ecf21bece5867750a030982"},
					{"0x1a0a1098c733a0ba9c9ed018ecb794dfb3b84b048ad03850b46e6cfd1e1018fd25ecd9924e4cc4f61e490d92fe1ed89",
						"0xdb1cc6b3b535c13494ba5c4ce330f401298118fb2c774b057daf91e1919b84ea0982828f8a07fa58cb3ea8f0201cfa"},
					{"0x1a19656718399a35e838587e3ca3aad4337c715ec6baf72f3c7e3fca08fce3651140c73a45031117e9e0ab0f7496b3c",
						"0x11d225e689b1c5533a4da5615ccfd27032d2a33de261bc8c83e1d938637efe7b0612df05f42d2d858376b499c652d75"},
					{"0x1a97aec63b6f352e2a2494186b69f0276247a857d82f7879e7d6c01e4a38f6e9d97e347e87ce507a5b94f2b3773eac6",
						"0x8256cc2663dc6938ae7605931d1a8f441ef85d2cef47445416cb593d2860d2eefcfd9a3c35d76d25d17109a764ebce"},
					{"0x9deeddd201ebe394bc84236d884b755378a169e31bb81e389ed375576782cb5fd30a4cbbcecabdacb808eab7c6d4fe",
						"0x174a2760a9f8253d0539dec3d626f269841b51a989ce5205fcc09366eebf6535bab520fc56d623cd04fcde78d685a0b"},
					{"0x1e3c14cdeaa4ed3e266a36616cd141dcb32072e63e4b9c1b19eefc2d6b490397977352e865b5016f148f74d1234df6",
						"0x5d1a7ce417ff6c955a2d06c4c840b541f386fe4a4370228afae2be534e2a5417fade1ead66031d64458270725699d6"},
					{"0xb44bcb83f10d65248dcf486aa2bdfbedc610b3f42aaea5a4581cacaad83d83a5bc9af704bbfd377d916268ed5be0ca",
						"0x154a6f2aebf4c09b1e69a0a4d435958eab0b85a6f405246a1c6af4a4432e13ba0ae3e631b4cbb9958962a2bc05dadce"},
					{"0x43ec129619c24c903d1e8ced9e08af5e09deaaad2f85f5ff1a65a070055bda662d6b0ec13b2a12dce07dca1ca0d34e",
						"0xc696b6c3b4b65ba99f95b0419f35d9f5a13e81f259173b4babc0bf290a0e2d9979511b5a6e1122aa7a572eb04c937b"},
					{"0x5f8d9df1cd3ea471f926003dae6e4cbe4eac8676d97693f76b082468922496ab47cf0a6fdbab2717328446ddecdc17",
						"0x588223c67657eb350a8aba1f5ac04ccee4c02ab4a5c7d038c8a3eb8c8ebc0233950f6c7f12cdbb6fa6a423e97cd661"},
					{"0xe23936ef578d86d6e629e2146599905c9a3cfc4d158bb35f4513ade1031830de7ca355bfb3ec74c487f7e91dbaeb22",
						"0x1168faccfe32c916ceece34bd697bfa7aad4048bf1ef00409f0c7ad20564f46eac4e7710083011e85e8543b1b8b3c3e"},
					{"0xce947ebe5063aafb2a3e9796bf3677237ecd980026889351d1120fa73b24a1a04e971deab8ad4c7f189e9559b0d1b1",
						"0x156b554c56db14b68a4ed5367a442fac6d95d2714fc59e9eaa235dba01a804277302625ce66f99e4eead7b7c7137a90"},
					{"0xeb0854aaaa2035ec173a3419093ef828c51df43b492dfad54fa62fcdcf40621c21d040f678a3eab5831e12e4a0d02f",
						"0x380b52ae6bdef0a4cfeab3cbef0efd4017286319f2e1dd8022fd896cb0cb8be986c4abb1ec2bd6aa90c45b1a61b6aa"},
					{"0x19cee627a52ee5d83bfba584b81e8da3e9fee36a4785cc5e277a0635ad178c03863d708258214e80851cec3bab077b0",
						"0x19a62cec87d6fa62a59d35af20772f84611a11d62e573e008cf6cc4e6e473a93a89b0b2b4579827b835109ffae07696"},
					{"0x193224b2c9149c504131017f21f51adb859b41abcf51ad61eca54273f7962d47bbe4c33a96e0938b31d4d475e20c472",
						"0x17c9844cbfaebc9866e792216a887a4efe3c85a55891e1f8e0f0c5bcd1276e607087e9a67b9ca05b30577a70e0b9cd3"},
					{"0x1a381bff3a99144474ed68d432a55cc7a87a77ec8ef9ec7dc8207bebe95d55bda2ec40d34992be5e68691717ad25e87",
						"0x1998a69751dae99d28fafd96b81cc341482051957408078e8ea74617fba8242075a8f8b079f1e18a274a9defcf12cf9"},
					{"0x1436817e660a5107b3d259cb3f2af5d373b951ad6a54ad6831ff9d7053cd857ae5abcdf0e44bb1a45dc4dc43be82a24",
						"0x7d2fe1f55c31673b19f9fc000ffa74819d7001a3f1bfcd87bad06d6642b6ebd2d69836d88e9e43fa6af8225a380ed6"},
					{"0x321b8a5f187405be1d90830e71a287eadc2722bc45bb9f5e84a921331835127361f4544d86b4ace8ceed8886063f02",
						"0x58cd6cd7c128e53ec5f4b55e8cb0356e36b4f4673890bd57600018e658100e9ebe25311b6c0a9b21b95a6414a88ad1"},
					{"0xa88afc381c5826291bd27ee2a6d88bbf10068154b09f1df0b3cbe9478e2f2f835133d5486b0797d9ade242e0d442c7",
						"0xf76768c2ed3f4e85913f870de4ef972299114025ab5bd97a70dff551d3fdb6e780876e8d5979631e170590f104b3c6"},
					{"0x6bca067bf3c3fda8fe8060531d6f5550fef83c2769c7686682abf9585ca3fc04d64a75061c9965a8782864126167f9",
						"0xa0e802a11d90c37cd105a8c597cb844efb81706dedf9a6a19254dc7db8b0bec5127ab383370ddf3ea8426be5d2d9d"},
					{"0x46d4e78de408979012bcc8d12b78d6d85d8beca1b7700c410fa6a3b78de2b181a3b2d464d3d1745f5a4f722bcf667c",
						"0xdcb51648143fc79a24093a797cda524c95512a51d521ec38321e624b1cbcdb9dd278c71eb0fa345eb67639e9563dd8"},
					{"0x1620bcb2404787ee39cae59ba2c85a2062aed25694677b198fed619d6558b859da3aaad5d38e773ca645cd710379803",
						"0x10620b57ed7f896194775ed566c2e01276e4f7bcdf233a23ad1d3e9b77c3fa0cdf84c564c1021f5e584458ab26bd8b8"},
					{"0x6dabe74a1d15321e50aabfd662991b784c84a9b72e3340eee8efe520e68624376814038f2d1d683df0cc4f81c37a3",
						"0x130be04e20fed0aabac1be1a13614e262414e7492e8701816e89b9982a435e2cf80b11474358c9bf758393f8fe05022"},
					{"0x83889427d535bd6411a28c2a0ac1d40396ee0e686f0052f0bf59abb62bd3004f1ff3a76bf93b70926638b5c05f0a3c",
						"0x131b916656d906aeda904a98d7664ed693a9a793a051a824c4ed8c815835eed2bcf1d686414c859af676f850e3dfd71"},
					{"0x160c9ea53e9e85745773c4faccea8a4bc12d558b01068c2f3d88ca4d43b6f5e338a0ab7230fb044757ec2666f769dbb",
						"0xb40e44b0581a642829ebdc74a3579c83018737cb432db5f3d692411082f2139753f6122e53566737d4ef53ac06ad1"},
					{"0x186657fc77b2f47a24d4ce9ff9978c8b856ebb1d23e75468ac95d73168cdd6f7ae9ed767ad52bc9a3ea835a37c34aa3",
						"0xa216aa737f6a614825bc46fbbcd767235927731a2c14c71ac6e2a38da8479785028a30559d92f544c2cf5d24b542fd"},
					{"0x11e53de7dcd97d4850d8d3b3d948db5e1a33adf4d3b81342b155f926729619cf1094cb81ca5119e69a84307aa35cc4f",
						"0x15ebff8d6d9c62eada64518cd85683baffe02073d8191ce5006a93c64dd1aec73e6f5c2178face4ded883af8b0738ee"},
				},
			},
		},
	},
}

var tBLS12_77 = TwistedEdwardsCurve{
	Name:     BLS12_377.Name,
	Package:  "twistededwards",
	EnumID:   BLS12_377.EnumID,
	A:        "-1",
	D:        "3021",
	Cofactor: "4",
	Order:    "2111115437357092606062206234695386632838870926408408195193685246394721360383",
	BaseX:    "717051916204163000937139483451426116831771857428389560441264442629694842243",
	BaseY:    "882565546457454111605105352482086902132191855952243170543452705048019814192",
}

func init() {
	addCurve(&BLS12_377)
	addTwistedEdwardCurve(&tBLS12_77)
}
