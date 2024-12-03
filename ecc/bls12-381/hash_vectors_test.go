package bls12381

import (
	"testing"

	"github.com/Overclock-Validator/gnark-crypto/ecc/bls12-381/fp"
)

func TestG1IsogenyVectors(t *testing.T) {
	t.Parallel()

	// TODO @gbotrel fix me test vectors shouldn't set words directly
	p := G1Affine{
		fp.Element{
			3660217524291093078, 10096673235325531916, 228883846699980880, 13273309082988818590, 5645112663858216297, 1475745906155504807,
		},
		fp.Element{
			7179819451626801451, 8122998708501415251, 10493900036512999567, 8666325578439571587, 1547096619901497872, 644447436619416978,
		},
	}
	toMont(&p.X)
	toMont(&p.Y)

	ref := G1Affine{
		fp.Element{
			15068149172194637577, 9957346779704953421, 14194629579302688285, 14905041577284894537, 12723787027614029596, 1241178457703452833,
		},
		fp.Element{
			8713071345859776370, 18097455281831542002, 18193395493462724643, 6332597957331977118, 3845332352253397392, 1815350252291127063,
		},
	}

	toMont(&ref.X)
	toMont(&ref.Y)

	g1Isogeny(&p)

	if ref != p {
		t.Fail()
	}
}
func init() {
	encodeToG1Vector = encodeTestVector{
		dst: []byte("QUUX-V01-CS02-with-BLS12381G1_XMD:SHA-256_SSWU_NU_"),
		cases: []encodeTestCase{
			{
				msg: "", u: "0x156c8a6a2c184569d69a76be144b5cdc5141d2d2ca4fe341f011e25e3969c55ad9e9b9ce2eb833c81a908e5fa4ac5f03",
				P: point{"0x184bb665c37ff561a89ec2122dd343f20e0f4cbcaec84e3c3052ea81d1834e192c426074b02ed3dca4e7676ce4ce48ba", "0x04407b8d35af4dacc809927071fc0405218f1401a6d15af775810e4e460064bcc9468beeba82fdc751be70476c888bf3"},
				Q: point{"0x11398d3b324810a1b093f8e35aa8571cced95858207e7f49c4fd74656096d61d8a2f9a23cdb18a4dd11cd1d66f41f709", "0x19316b6fb2ba7717355d5d66a361899057e1e84a6823039efc7beccefe09d023fb2713b1c415fcf278eb0c39a89b4f72"},
			},
			{
				msg: "abc", u: "0x147e1ed29f06e4c5079b9d14fc89d2820d32419b990c1c7bb7dbea2a36a045124b31ffbde7c99329c05c559af1c6cc82",
				P: point{"0x009769f3ab59bfd551d53a5f846b9984c59b97d6842b20a2c565baa167945e3d026a3755b6345df8ec7e6acb6868ae6d", "0x1532c00cf61aa3d0ce3e5aa20c3b531a2abd2c770a790a2613818303c6b830ffc0ecf6c357af3317b9575c567f11cd2c"},
				Q: point{"0x1998321bc27ff6d71df3051b5aec12ff47363d81a5e9d2dff55f444f6ca7e7d6af45c56fd029c58237c266ef5cda5254", "0x034d274476c6307ae584f951c82e7ea85b84f72d28f4d6471732356121af8d62a49bc263e8eb913a6cf6f125995514ee"},
			},
			{
				msg: "abcdef0123456789", u: "0x04090815ad598a06897dd89bcda860f25837d54e897298ce31e6947378134d3761dc59a572154963e8c954919ecfa82d",
				P: point{"0x1974dbb8e6b5d20b84df7e625e2fbfecb2cdb5f77d5eae5fb2955e5ce7313cae8364bc2fff520a6c25619739c6bdcb6a", "0x15f9897e11c6441eaa676de141c8d83c37aab8667173cbe1dfd6de74d11861b961dccebcd9d289ac633455dfcc7013a3"},
				Q: point{"0x17d502fa43bd6a4cad2859049a0c3ecefd60240d129be65da271a4c03a9c38fa78163b9d2a919d2beb57df7d609b4919", "0x109019902ae93a8732abecf2ff7fecd2e4e305eb91f41c9c3267f16b6c19de138c7272947f25512745da6c466cdfd1ac"},
			},
			{
				msg: "q128_qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
				u:   "0x08dccd088ca55b8bfbc96fb50bb25c592faa867a8bb78d4e94a8cc2c92306190244532e91feba2b7fed977e3c3bb5a1f",
				P:   point{"0x0a7a047c4a8397b3446450642c2ac64d7239b61872c9ae7a59707a8f4f950f101e766afe58223b3bff3a19a7f754027c", "0x1383aebba1e4327ccff7cf9912bda0dbc77de048b71ef8c8a81111d71dc33c5e3aa6edee9cf6f5fe525d50cc50b77cc9"},
				Q:   point{"0x112eb92dd2b3aa9cd38b08de4bef603f2f9fb0ca226030626a9a2e47ad1e9847fe0a5ed13766c339e38f514bba143b21", "0x17542ce2f8d0a54f2c5ba8c4b14e10b22d5bcd7bae2af3c965c8c872b571058c720eac448276c99967ded2bf124490e1"},
			},
			{
				msg: "a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				u:   "0x0dd824886d2123a96447f6c56e3a3fa992fbfefdba17b6673f9f630ff19e4d326529db37e1c1be43f905bf9202e0278d",
				P:   point{"0x0e7a16a975904f131682edbb03d9560d3e48214c9986bd50417a77108d13dc957500edf96462a3d01e62dc6cd468ef11", "0x0ae89e677711d05c30a48d6d75e76ca9fb70fe06c6dd6ff988683d89ccde29ac7d46c53bb97a59b1901abf1db66052db"},
				Q:   point{"0x1775d400a1bacc1c39c355da7e96d2d1c97baa9430c4a3476881f8521c09a01f921f592607961efc99c4cd46bd78ca19", "0x1109b5d59f65964315de65a7a143e86eabc053104ed289cf480949317a5685fad7254ff8e7fe6d24d3104e5d55ad6370"},
			},
		},
	}

	hashToG1Vector = hashTestVector{
		dst: []byte("QUUX-V01-CS02-with-BLS12381G1_XMD:SHA-256_SSWU_RO_"),
		cases: []hashTestCase{
			{
				msg: "",
				u0:  "0x0ba14bd907ad64a016293ee7c2d276b8eae71f25a4b941eece7b0d89f17f75cb3ae5438a614fb61d6835ad59f29c564f",
				u1:  "0x019b9bd7979f12657976de2884c7cce192b82c177c80e0ec604436a7f538d231552f0d96d9f7babe5fa3b19b3ff25ac9",
				P:   point{"0x052926add2207b76ca4fa57a8734416c8dc95e24501772c814278700eed6d1e4e8cf62d9c09db0fac349612b759e79a1", "0x08ba738453bfed09cb546dbb0783dbb3a5f1f566ed67bb6be0e8c67e2e81a4cc68ee29813bb7994998f3eae0c9c6a265"},
				Q0:  point{"0x11a3cce7e1d90975990066b2f2643b9540fa40d6137780df4e753a8054d07580db3b7f1f03396333d4a359d1fe3766fe", "0x0eeaf6d794e479e270da10fdaf768db4c96b650a74518fc67b04b03927754bac66f3ac720404f339ecdcc028afa091b7"},
				Q1:  point{"0x160003aaf1632b13396dbad518effa00fff532f604de1a7fc2082ff4cb0afa2d63b2c32da1bef2bf6c5ca62dc6b72f9c", "0x0d8bb2d14e20cf9f6036152ed386d79189415b6d015a20133acb4e019139b94e9c146aaad5817f866c95d609a361735e"},
			},
			{
				msg: "abc",
				u0:  "0x0d921c33f2bad966478a03ca35d05719bdf92d347557ea166e5bba579eea9b83e9afa5c088573c2281410369fbd32951",
				u1:  "0x003574a00b109ada2f26a37a91f9d1e740dffd8d69ec0c35e1e9f4652c7dba61123e9dd2e76c655d956e2b3462611139",
				P:   point{"0x03567bc5ef9c690c2ab2ecdf6a96ef1c139cc0b2f284dca0a9a7943388a49a3aee664ba5379a7655d3c68900be2f6903", "0x0b9c15f3fe6e5cf4211f346271d7b01c8f3b28be689c8429c85b67af215533311f0b8dfaaa154fa6b88176c229f2885d"},
				Q0:  point{"0x125435adce8e1cbd1c803e7123f45392dc6e326d292499c2c45c5865985fd74fe8f042ecdeeec5ecac80680d04317d80", "0x0e8828948c989126595ee30e4f7c931cbd6f4570735624fd25aef2fa41d3f79cfb4b4ee7b7e55a8ce013af2a5ba20bf2"},
				Q1:  point{"0x11def93719829ecda3b46aa8c31fc3ac9c34b428982b898369608e4f042babee6c77ab9218aad5c87ba785481eff8ae4", "0x0007c9cef122ccf2efd233d6eb9bfc680aa276652b0661f4f820a653cec1db7ff69899f8e52b8e92b025a12c822a6ce6"},
			},
			{
				msg: "abcdef0123456789",
				u0:  "0x062d1865eb80ebfa73dcfc45db1ad4266b9f3a93219976a3790ab8d52d3e5f1e62f3b01795e36834b17b70e7b76246d4",
				u1:  "0x0cdc3e2f271f29c4ff75020857ce6c5d36008c9b48385ea2f2bf6f96f428a3deb798aa033cd482d1cdc8b30178b08e3a",
				P:   point{"0x11e0b079dea29a68f0383ee94fed1b940995272407e3bb916bbf268c263ddd57a6a27200a784cbc248e84f357ce82d98", "0x03a87ae2caf14e8ee52e51fa2ed8eefe80f02457004ba4d486d6aa1f517c0889501dc7413753f9599b099ebcbbd2d709"},
				Q0:  point{"0x08834484878c217682f6d09a4b51444802fdba3d7f2df9903a0ddadb92130ebbfa807fffa0eabf257d7b48272410afff", "0x0b318f7ecf77f45a0f038e62d7098221d2dbbca2a394164e2e3fe953dc714ac2cde412d8f2d7f0c03b259e6795a2508e"},
				Q1:  point{"0x158418ed6b27e2549f05531a8281b5822b31c3bf3144277fbb977f8d6e2694fedceb7011b3c2b192f23e2a44b2bd106e", "0x1879074f344471fac5f839e2b4920789643c075792bec5af4282c73f7941cda5aa77b00085eb10e206171b9787c4169f"},
			},
			{
				msg: "q128_qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
				u0:  "0x010476f6a060453c0b1ad0b628f3e57c23039ee16eea5e71bb87c3b5419b1255dc0e5883322e563b84a29543823c0e86",
				u1:  "0x0b1a912064fb0554b180e07af7e787f1f883a0470759c03c1b6509eb8ce980d1670305ae7b928226bb58fdc0a419f46e",
				P: point{"0x15f68eaa693b95ccb85215dc65fa81038d69629f70aeee0d0f677cf22285e7bf58d7cb86eefe8f2e9bc3f8cb84fac488",
					"0x1807a1d50c29f430b8cafc4f8638dfeeadf51211e1602a5f184443076715f91bb90a48ba1e370edce6ae1062f5e6dd38"},
				Q0: point{"0x0cbd7f84ad2c99643fea7a7ac8f52d63d66cefa06d9a56148e58b984b3dd25e1f41ff47154543343949c64f88d48a710", "0x052c00e4ed52d000d94881a5638ae9274d3efc8bc77bc0e5c650de04a000b2c334a9e80b85282a00f3148dfdface0865"},
				Q1: point{"0x06493fb68f0d513af08be0372f849436a787e7b701ae31cb964d968021d6ba6bd7d26a38aaa5a68e8c21a6b17dc8b579", "0x02e98f2ccf5802b05ffaac7c20018bc0c0b2fd580216c4aa2275d2909dc0c92d0d0bdc979226adeb57a29933536b6bb4"},
			},
			{
				msg: "a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				u0:  "0x0a8ffa7447f6be1c5a2ea4b959c9454b431e29ccc0802bc052413a9c5b4f9aac67a93431bd480d15be1e057c8a08e8c6",
				u1:  "0x05d487032f602c90fa7625dbafe0f4a49ef4a6b0b33d7bb349ff4cf5410d297fd6241876e3e77b651cfc8191e40a68b7",
				P: point{"0x082aabae8b7dedb0e78aeb619ad3bfd9277a2f77ba7fad20ef6aabdc6c31d19ba5a6d12283553294c1825c4b3ca2dcfe",
					"0x05b84ae5a942248eea39e1d91030458c40153f3b654ab7872d779ad1e942856a20c438e8d99bc8abfbf74729ce1f7ac8"},
				Q0: point{"0x0cf97e6dbd0947857f3e578231d07b309c622ade08f2c08b32ff372bd90db19467b2563cc997d4407968d4ac80e154f8", "0x127f0cddf2613058101a5701f4cb9d0861fd6c2a1b8e0afe194fccf586a3201a53874a2761a9ab6d7220c68661a35ab3"},
				Q1: point{"0x092f1acfa62b05f95884c6791fba989bbe58044ee6355d100973bf9553ade52b47929264e6ae770fb264582d8dce512a", "0x028e6d0169a72cfedb737be45db6c401d3adfb12c58c619c82b93a5dfcccef12290de530b0480575ddc8397cda0bbebf"},
			},
		},
	}

	encodeToG2Vector = encodeTestVector{
		dst: []byte("QUUX-V01-CS02-with-BLS12381G2_XMD:SHA-256_SSWU_NU_"),
		cases: []encodeTestCase{
			{
				P: point{"0x00e7f4568a82b4b7dc1f14c6aaa055edf51502319c723c4dc2688c7fe5944c213f510328082396515734b6612c4e7bb7,0x126b855e9e69b1f691f816e48ac6977664d24d99f8724868a184186469ddfd4617367e94527d4b74fc86413483afb35b",
					"0x0caead0fd7b6176c01436833c79d305c78be307da5f6af6c133c47311def6ff1e0babf57a0fb5539fce7ee12407b0a42,0x1498aadcf7ae2b345243e281ae076df6de84455d766ab6fcdaad71fab60abb2e8b980a440043cd305db09d283c895e3d"},
				msg: "",
				u:   "0x07355d25caf6e7f2f0cb2812ca0e513bd026ed09dda65b177500fa31714e09ea0ded3a078b526bed3307f804d4b93b04,0x02829ce3c021339ccb5caf3e187f6370e1e2a311dec9b75363117063ab2015603ff52c3d3b98f19c2f65575e99e8b78c",
				Q:   point{"0x18ed3794ad43c781816c523776188deafba67ab773189b8f18c49bc7aa841cd81525171f7a5203b2a340579192403bef,0x0727d90785d179e7b5732c8a34b660335fed03b913710b60903cf4954b651ed3466dc3728e21855ae822d4a0f1d06587", "0x00764a5cf6c5f61c52c838523460eb2168b5a5b43705e19cb612e006f29b717897facfd15dd1c8874c915f6d53d0342d,0x19290bb9797c12c1d275817aa2605ebe42275b66860f0e4d04487ebc2e47c50b36edd86c685a60c20a2bd584a82b011a"},
			},
			{
				P: point{"0x108ed59fd9fae381abfd1d6bce2fd2fa220990f0f837fa30e0f27914ed6e1454db0d1ee957b219f61da6ff8be0d6441f,0x0296238ea82c6d4adb3c838ee3cb2346049c90b96d602d7bb1b469b905c9228be25c627bffee872def773d5b2a2eb57d",
					"0x033f90f6057aadacae7963b0a0b379dd46750c1c94a6357c99b65f63b79e321ff50fe3053330911c56b6ceea08fee656,0x153606c417e59fb331b7ae6bce4fbf7c5190c33ce9402b5ebe2b70e44fca614f3f1382a3625ed5493843d0b0a652fc3f"},
				msg: "abc",
				u:   "0x138879a9559e24cecee8697b8b4ad32cced053138ab913b99872772dc753a2967ed50aabc907937aefb2439ba06cc50c,0x0a1ae7999ea9bab1dcc9ef8887a6cb6e8f1e22566015428d220b7eec90ffa70ad1f624018a9ad11e78d588bd3617f9f2",
				Q:   point{"0x0f40e1d5025ecef0d850aa0bb7bbeceab21a3d4e85e6bee857805b09693051f5b25428c6be343edba5f14317fcc30143,0x02e0d261f2b9fee88b82804ec83db330caa75fbb12719cfa71ccce1c532dc4e1e79b0a6a281ed8d3817524286c8bc04c", "0x0cf4a4adc5c66da0bca4caddc6a57ecd97c8252d7526a8ff478e0dfed816c4d321b5c3039c6683ae9b1e6a3a38c9c0ae,0x11cad1646bb3768c04be2ab2bbe1f80263b7ff6f8f9488f5bc3b6850e5a3e97e20acc583613c69cf3d2bfe8489744ebb"},
			},
			{
				P: point{"0x038af300ef34c7759a6caaa4e69363cafeed218a1f207e93b2c70d91a1263d375d6730bd6b6509dcac3ba5b567e85bf3,0x0da75be60fb6aa0e9e3143e40c42796edf15685cafe0279afd2a67c3dff1c82341f17effd402e4f1af240ea90f4b659b",
					"0x19b148cbdf163cf0894f29660d2e7bfb2b68e37d54cc83fd4e6e62c020eaa48709302ef8e746736c0e19342cc1ce3df4,0x0492f4fed741b073e5a82580f7c663f9b79e036b70ab3e51162359cec4e77c78086fe879b65ca7a47d34374c8315ac5e"},
				msg: "abcdef0123456789",
				u:   "0x18c16fe362b7dbdfa102e42bdfd3e2f4e6191d479437a59db4eb716986bf08ee1f42634db66bde97d6c16bbfd342b3b8,0x0e37812ce1b146d998d5f92bdd5ada2a31bfd63dfe18311aa91637b5f279dd045763166aa1615e46a50d8d8f475f184e",
				Q:   point{"0x13a9d4a738a85c9f917c7be36b240915434b58679980010499b9ae8d7a1bf7fbe617a15b3cd6060093f40d18e0f19456,0x16fa88754e7670366a859d6f6899ad765bf5a177abedb2740aacc9252c43f90cd0421373fbd5b2b76bb8f5c4886b5d37", "0x0a7fa7d82c46797039398253e8765a4194100b330dfed6d7fbb46d6fbf01e222088779ac336e3675c7a7a0ee05bbb6e3,0x0c6ee170ab766d11fa9457cef53253f2628010b2cffc102b3b28351eb9df6c281d3cfc78e9934769d661b72a5265338d"},
			},
			{
				P: point{"0x0c5ae723be00e6c3f0efe184fdc0702b64588fe77dda152ab13099a3bacd3876767fa7bbad6d6fd90b3642e902b208f9,0x12c8c05c1d5fc7bfa847f4d7d81e294e66b9a78bc9953990c358945e1f042eedafce608b67fdd3ab0cb2e6e263b9b1ad",
					"0x04e77ddb3ede41b5ec4396b7421dd916efc68a358a0d7425bddd253547f2fb4830522358491827265dfc5bcc1928a569,0x11c624c56dbe154d759d021eec60fab3d8b852395a89de497e48504366feedd4662d023af447d66926a28076813dd646"},
				msg: "q128_qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
				u:   "0x08d4a0997b9d52fecf99427abb721f0fa779479963315fe21c6445250de7183e3f63bfdf86570da8929489e421d4ee95,0x16cb4ccad91ec95aab070f22043916cd6a59c4ca94097f7f510043d48515526dc8eaaea27e586f09151ae613688d5a89",
				Q:   point{"0x0a08b2f639855dfdeaaed972702b109e2241a54de198b2b4cd12ad9f88fa419a6086a58d91fc805de812ea29bee427c2,0x04a7442e4cb8b42ef0f41dac9ee74e65ecad3ce0851f0746dc47568b0e7a8134121ed09ba054509232c49148aef62cda", "0x05d60b1f04212b2c87607458f71d770f43973511c260f0540eef3a565f42c7ce59aa1cea684bb2a7bcab84acd2f36c8c,0x1017aa5747ba15505ece266a86b0ca9c712f41a254b76ca04094ca442ce45ecd224bd5544cd16685d0d1b9d156dd0531"},
			},
			{
				P: point{"0x0ea4e7c33d43e17cc516a72f76437c4bf81d8f4eac69ac355d3bf9b71b8138d55dc10fd458be115afa798b55dac34be1,0x1565c2f625032d232f13121d3cfb476f45275c303a037faa255f9da62000c2c864ea881e2bcddd111edc4a3c0da3e88d",
					"0x043b6f5fe4e52c839148dc66f2b3751e69a0f6ebb3d056d6465d50d4108543ecd956e10fa1640dfd9bc0030cc2558d28,0x0f8991d2a1ad662e7b6f58ab787947f1fa607fce12dde171bc17903b012091b657e15333e11701edcf5b63ba2a561247"},
				msg: "a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				u:   "0x03f80ce4ff0ca2f576d797a3660e3f65b274285c054feccc3215c879e2c0589d376e83ede13f93c32f05da0f68fd6a10,0x006488a837c5413746d868d1efb7232724da10eca410b07d8b505b9363bdccf0a1fc0029bad07d65b15ccfe6dd25e20d",
				Q:   point{"0x19592c812d5a50c5601062faba14c7d670711745311c879de1235a0a11c75aab61327bf2d1725db07ec4d6996a682886,0x0eef4fa41ddc17ed47baf447a2c498548f3c72a02381313d13bef916e240b61ce125539090d62d9fbb14a900bf1b8e90", "0x1260d6e0987eae96af9ebe551e08de22b37791d53f4db9e0d59da736e66699735793e853e26362531fe4adf99c1883e3,0x0dbace5df0a4ac4ac2f45d8fdf8aee45484576fdd6efc4f98ab9b9f4112309e628255e183022d98ea5ed6e47ca00306c"},
			},
		},
	}

	hashToG2Vector = hashTestVector{
		dst: []byte("QUUX-V01-CS02-with-BLS12381G2_XMD:SHA-256_SSWU_RO_"),
		cases: []hashTestCase{
			{
				msg: "",
				u0:  "0x03dbc2cce174e91ba93cbb08f26b917f98194a2ea08d1cce75b2b9cc9f21689d80bd79b594a613d0a68eb807dfdc1cf8,0x05a2acec64114845711a54199ea339abd125ba38253b70a92c876df10598bd1986b739cad67961eb94f7076511b3b39a",
				u1:  "0x02f99798e8a5acdeed60d7e18e9120521ba1f47ec090984662846bc825de191b5b7641148c0dbc237726a334473eee94,0x145a81e418d4010cc027a68f14391b30074e89e60ee7a22f87217b2f6eb0c4b94c9115b436e6fa4607e95a98de30a435",
				P: point{
					"0x0141ebfbdca40eb85b87142e130ab689c673cf60f1a3e98d69335266f30d9b8d4ac44c1038e9dcdd5393faf5c41fb78a,0x05cb8437535e20ecffaef7752baddf98034139c38452458baeefab379ba13dff5bf5dd71b72418717047f5b0f37da03d",
					"0x0503921d7f6a12805e72940b963c0cf3471c7b2a524950ca195d11062ee75ec076daf2d4bc358c4b190c0c98064fdd92,0x12424ac32561493f3fe3c260708a12b7c620e7be00099a974e259ddc7d1f6395c3c811cdd19f1e8dbf3e9ecfdcbab8d6",
				},
				Q0: point{"0x019ad3fc9c72425a998d7ab1ea0e646a1f6093444fc6965f1cad5a3195a7b1e099c050d57f45e3fa191cc6d75ed7458c,0x171c88b0b0efb5eb2b88913a9e74fe111a4f68867b59db252ce5868af4d1254bfab77ebde5d61cd1a86fb2fe4a5a1c1d", "0x0ba10604e62bdd9eeeb4156652066167b72c8d743b050fb4c1016c31b505129374f76e03fa127d6a156213576910fef3,0x0eb22c7a543d3d376e9716a49b72e79a89c9bfe9feee8533ed931cbb5373dde1fbcd7411d8052e02693654f71e15410a"},
				Q1: point{"0x113d2b9cd4bd98aee53470b27abc658d91b47a78a51584f3d4b950677cfb8a3e99c24222c406128c91296ef6b45608be,0x13855912321c5cb793e9d1e88f6f8d342d49c0b0dbac613ee9e17e3c0b3c97dfbb5a49cc3fb45102fdbaf65e0efe2632", "0x0fd3def0b7574a1d801be44fde617162aa2e89da47f464317d9bb5abc3a7071763ce74180883ad7ad9a723a9afafcdca,0x056f617902b3c0d0f78a9a8cbda43a26b65f602f8786540b9469b060db7b38417915b413ca65f875c130bebfaa59790c"},
			},
			{
				msg: "abc",
				u0:  "0x15f7c0aa8f6b296ab5ff9c2c7581ade64f4ee6f1bf18f55179ff44a2cf355fa53dd2a2158c5ecb17d7c52f63e7195771,0x01c8067bf4c0ba709aa8b9abc3d1cef589a4758e09ef53732d670fd8739a7274e111ba2fcaa71b3d33df2a3a0c8529dd",
				u1:  "0x187111d5e088b6b9acfdfad078c4dacf72dcd17ca17c82be35e79f8c372a693f60a033b461d81b025864a0ad051a06e4,0x08b852331c96ed983e497ebc6dee9b75e373d923b729194af8e72a051ea586f3538a6ebb1e80881a082fa2b24df9f566",
				P: point{
					"0x02c2d18e033b960562aae3cab37a27ce00d80ccd5ba4b7fe0e7a210245129dbec7780ccc7954725f4168aff2787776e6,0x139cddbccdc5e91b9623efd38c49f81a6f83f175e80b06fc374de9eb4b41dfe4ca3a230ed250fbe3a2acf73a41177fd8",
					"0x1787327b68159716a37440985269cf584bcb1e621d3a7202be6ea05c4cfe244aeb197642555a0645fb87bf7466b2ba48,0x00aa65dae3c8d732d10ecd2c50f8a1baf3001578f71c694e03866e9f3d49ac1e1ce70dd94a733534f106d4cec0eddd16",
				},
				Q0: point{"0x12b2e525281b5f4d2276954e84ac4f42cf4e13b6ac4228624e17760faf94ce5706d53f0ca1952f1c5ef75239aeed55ad,0x05d8a724db78e570e34100c0bc4a5fa84ad5839359b40398151f37cff5a51de945c563463c9efbdda569850ee5a53e77", "0x02eacdc556d0bdb5d18d22f23dcb086dd106cad713777c7e6407943edbe0b3d1efe391eedf11e977fac55f9b94f2489c,0x04bbe48bfd5814648d0b9e30f0717b34015d45a861425fabc1ee06fdfce36384ae2c808185e693ae97dcde118f34de41"},
				Q1: point{"0x19f18cc5ec0c2f055e47c802acc3b0e40c337256a208001dde14b25afced146f37ea3d3ce16834c78175b3ed61f3c537,0x15b0dadc256a258b4c68ea43605dffa6d312eef215c19e6474b3e101d33b661dfee43b51abbf96fee68fc6043ac56a58", "0x05e47c1781286e61c7ade887512bd9c2cb9f640d3be9cf87ea0bad24bd0ebfe946497b48a581ab6c7d4ca74b5147287f,0x19f98db2f4a1fcdf56a9ced7b320ea9deecf57c8e59236b0dc21f6ee7229aa9705ce9ac7fe7a31c72edca0d92370c096"},
			},
			{
				msg: "abcdef0123456789",
				u0:  "0x0313d9325081b415bfd4e5364efaef392ecf69b087496973b229303e1816d2080971470f7da112c4eb43053130b785e1,0x062f84cb21ed89406890c051a0e8b9cf6c575cf6e8e18ecf63ba86826b0ae02548d83b483b79e48512b82a6c0686df8f",
				u1:  "0x1739123845406baa7be5c5dc74492051b6d42504de008c635f3535bb831d478a341420e67dcc7b46b2e8cba5379cca97,0x01897665d9cb5db16a27657760bbea7951f67ad68f8d55f7113f24ba6ddd82caef240a9bfa627972279974894701d975",
				P: point{
					"0x121982811d2491fde9ba7ed31ef9ca474f0e1501297f68c298e9f4c0028add35aea8bb83d53c08cfc007c1e005723cd0,0x190d119345b94fbd15497bcba94ecf7db2cbfd1e1fe7da034d26cbba169fb3968288b3fafb265f9ebd380512a71c3f2c",
					"0x05571a0f8d3c08d094576981f4a3b8eda0a8e771fcdcc8ecceaf1356a6acf17574518acb506e435b639353c2e14827c8,0x0bb5e7572275c567462d91807de765611490205a941a5a6af3b1691bfe596c31225d3aabdf15faff860cb4ef17c7c3be",
				},
				Q0: point{"0x0f48f1ea1318ddb713697708f7327781fb39718971d72a9245b9731faaca4dbaa7cca433d6c434a820c28b18e20ea208,0x06051467c8f85da5ba2540974758f7a1e0239a5981de441fdd87680a995649c211054869c50edbac1f3a86c561ba3162", "0x168b3d6df80069dbbedb714d41b32961ad064c227355e1ce5fac8e105de5e49d77f0c64867f3834848f152497eb76333,0x134e0e8331cee8cb12f9c2d0742714ed9eee78a84d634c9a95f6a7391b37125ed48bfc6e90bf3546e99930ff67cc97bc"},
				Q1: point{"0x004fd03968cd1c99a0dd84551f44c206c84dcbdb78076c5bfee24e89a92c8508b52b88b68a92258403cbe1ea2da3495f,0x1674338ea298281b636b2eb0fe593008d03171195fd6dcd4531e8a1ed1f02a72da238a17a635de307d7d24aa2d969a47", "0x0dc7fa13fff6b12558419e0a1e94bfc3cfaf67238009991c5f24ee94b632c3d09e27eca329989aee348a67b50d5e236c,0x169585e164c131103d85324f2d7747b23b91d66ae5d947c449c8194a347969fc6bbd967729768da485ba71868df8aed2"},
			},
			{
				msg: "q128_qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
				u0:  "0x025820cefc7d06fd38de7d8e370e0da8a52498be9b53cba9927b2ef5c6de1e12e12f188bbc7bc923864883c57e49e253,0x034147b77ce337a52e5948f66db0bab47a8d038e712123bb381899b6ab5ad20f02805601e6104c29df18c254b8618c7b",
				u1:  "0x0930315cae1f9a6017c3f0c8f2314baa130e1cf13f6532bff0a8a1790cd70af918088c3db94bda214e896e1543629795,0x10c4df2cacf67ea3cb3108b00d4cbd0b3968031ebc8eac4b1ebcefe84d6b715fde66bef0219951ece29d1facc8a520ef",
				P: point{
					"0x19a84dd7248a1066f737cc34502ee5555bd3c19f2ecdb3c7d9e24dc65d4e25e50d83f0f77105e955d78f4762d33c17da,0x0934aba516a52d8ae479939a91998299c76d39cc0c035cd18813bec433f587e2d7a4fef038260eef0cef4d02aae3eb91",
					"0x14f81cd421617428bc3b9fe25afbb751d934a00493524bc4e065635b0555084dd54679df1536101b2c979c0152d09192,0x09bcccfa036b4847c9950780733633f13619994394c23ff0b32fa6b795844f4a0673e20282d07bc69641cee04f5e5662",
				},
				Q0: point{"0x09eccbc53df677f0e5814e3f86e41e146422834854a224bf5a83a50e4cc0a77bfc56718e8166ad180f53526ea9194b57,0x0c3633943f91daee715277bd644fba585168a72f96ded64fc5a384cce4ec884a4c3c30f08e09cd2129335dc8f67840ec", "0x0eb6186a0457d5b12d132902d4468bfeb7315d83320b6c32f1c875f344efcba979952b4aa418589cb01af712f98cc555,0x119e3cf167e69eb16c1c7830e8df88856d48be12e3ff0a40791a5cd2f7221311d4bf13b1847f371f467357b3f3c0b4c7"},
				Q1: point{"0x0eb3aabc1ddfce17ff18455fcc7167d15ce6b60ddc9eb9b59f8d40ab49420d35558686293d046fc1e42f864b7f60e381,0x198bdfb19d7441ebcca61e8ff774b29d17da16547d2c10c273227a635cacea3f16826322ae85717630f0867539b5ed8b", "0x0aaf1dee3adf3ed4c80e481c09b57ea4c705e1b8d25b897f0ceeec3990748716575f92abff22a1c8f4582aff7b872d52,0x0d058d9061ed27d4259848a06c96c5ca68921a5d269b078650c882cb3c2bd424a8702b7a6ee4e0ead9982baf6843e924"},
			},
			{
				msg: "a512_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				u0:  "0x190b513da3e66fc9a3587b78c76d1d132b1152174d0b83e3c1114066392579a45824c5fa17649ab89299ddd4bda54935,0x12ab625b0fe0ebd1367fe9fac57bb1168891846039b4216b9d94007b674de2d79126870e88aeef54b2ec717a887dcf39",
				u1:  "0x0e6a42010cf435fb5bacc156a585e1ea3294cc81d0ceb81924d95040298380b164f702275892cedd81b62de3aba3f6b5,0x117d9a0defc57a33ed208428cb84e54c85a6840e7648480ae428838989d25d97a0af8e3255be62b25c2a85630d2dddd8",
				P: point{"0x01a6ba2f9a11fa5598b2d8ace0fbe0a0eacb65deceb476fbbcb64fd24557c2f4b18ecfc5663e54ae16a84f5ab7f62534,0x11fca2ff525572795a801eed17eb12785887c7b63fb77a42be46ce4a34131d71f7a73e95fee3f812aea3de78b4d01569",
					"0x0b6798718c8aed24bc19cb27f866f1c9effcdbf92397ad6448b5c9db90d2b9da6cbabf48adc1adf59a1a28344e79d57e,0x03a47f8e6d1763ba0cad63d6114c0accbef65707825a511b251a660a9b3994249ae4e63fac38b23da0c398689ee2ab52"},
				Q0: point{"0x17cadf8d04a1a170f8347d42856526a24cc466cb2ddfd506cff01191666b7f944e31244d662c904de5440516a2b09004,0x0d13ba91f2a8b0051cf3279ea0ee63a9f19bc9cb8bfcc7d78b3cbd8cc4fc43ba726774b28038213acf2b0095391c523e", "0x17ef19497d6d9246fa94d35575c0f8d06ee02f21a284dbeaa78768cb1e25abd564e3381de87bda26acd04f41181610c5,0x12c3c913ba4ed03c24f0721a81a6be7430f2971ffca8fd1729aafe496bb725807531b44b34b59b3ae5495e5a2dcbd5c8"},
				Q1: point{"0x16ec57b7fe04c71dfe34fb5ad84dbce5a2dbbd6ee085f1d8cd17f45e8868976fc3c51ad9eeda682c7869024d24579bfd,0x13103f7aace1ae1420d208a537f7d3a9679c287208026e4e3439ab8cd534c12856284d95e27f5e1f33eec2ce656533b0", "0x0958b2c4c2c10fcef5a6c59b9e92c4a67b0fae3e2e0f1b6b5edad9c940b8f3524ba9ebbc3f2ceb3cfe377655b3163bd7,0x0ccb594ed8bd14ca64ed9cb4e0aba221be540f25dd0d6ba15a4a4be5d67bcf35df7853b2d8dad3ba245f1ea3697f66aa"},
			},
		},
	}
}

var rSquare = fp.Element{
	17644856173732828998,
	754043588434789617,
	10224657059481499349,
	7488229067341005760,
	11130996698012816685,
	1267921511277847466,
}

// toMont converts z to Montgomery form
// sets and returns z = z * r²
func toMont(z *fp.Element) {
	z.Mul(z, &rSquare)
}
