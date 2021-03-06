# iota-eu-cma

Demonstration that IOTA fails the EU-CMA-GAME (because @Come-From-Beyond didn't
believe that the MIT DCI private keys were generated at random).

## Usage

```
dep ensure # Install Curl version of giota, instead of Kerl version
CGO_ENABLED=0 go run main.go # Build without cgo because the PoW LDFLAGS breaks the build
```

## Example output

Seed and address will change each time

```
Generated seed: ATHRNEGCPHLJEAIFLBF9HKBOKA9IHGKPCILS9TBZTQCRO9RANDRHWABCLCZA9PVLWOXJWDLCZECLDMFZY
Generated address: FOZBLDJDYPZJXIDAREPMW9LFZLJXPDFQBXESWAHBLXSSF9XOYFGZSMYWZPPMOFPYRQFYUDTYARYWCHDNV
Signing bundle
Fragment 0: IDWZIZIFTBXKSAQERKDGBAVYOTONHUHFZLXZFSNJVGRWORZYIQSDFZRBH9CZYEATFNOFFEZLYSANOIYPFAXAZOKYSTJSQAUVBPRUYVIIDRJBTTZTIMMXUVLU9GBNCPFXXMHTEFHTJ9UYLLCWPTLIFZCMLPKWHGMUEYMZSCAMQFIP9NNBMTZWVKWCYVGIUIKOBOJOQSSDEGUGHLSKXIHJE9AVLXDHCVRLXZJVDZBDJBLUZKLWIITAPBCDOECKBXAXPTIXTVTWCVFI9QFGTOCNFVIURXQMEGIFGKMU9TCHCPCWJDYQEYNBJHQAXWUMJYWN9JGDKQIZTRALCLWBCC9KXRUZEFFPKOYPVTK9KDEPQAWVJGSAQKPHVCEPYEIGQP9PGKBSHNJ9BUAWVM9CQNDXSYWJXIHOTGUPQX9BUMGURUIPEFUKAFCBFVHHBMFBQKORWOJRHKGFHIODP99TQPDLVIAIBE9OIIUQFDNLGKZUKDYOPSXLQSMTXFQJPTGPPTFUGDL9TVGLUJPEQQOJRHZUBYIMKPOMUSZEMWZYGRRDFLVQZVRPUHYABMWLZOXLNILVQHRLWWHDWIKRYABSNIVJRUFQPJHJUWMDWXDGQXEC9NCJURVSCUDGQR9PWAEXDHCXSEZFVU9FUTQJBUPABOJZNCBQFGFFNWOZYYUOK9POQXUIGXBSVXHWQHRJWYOW9UAFPBN9CWUCYFOICKZUTARLNIPOSZO9YKMGJLDSVDGZBXZATCHAAABIMOTEGZDBGQTPWX9TRYZOIPICHPJPGWMHRUENLYDKQKQHYZNGLNJEAEZLZWIX9DEFCZIDFOCMGIJIVT9IGUYCT9ZREETDLNBUEPFAWWISINOPZVYYSGDBZDZEGRCFHXEZAQDOHGUTZBYXDKXYBYROYTEVARKDRQJDRPDAJXR9XGESDLV9GKBTSYAYBKXJGYPLJYKYJAKHMDKHJRLZHAXRGOEAWZUZ9CXJFBF9MLKXTDRTPTSVYFLGLZWPEKXLJMFXVFQODUOVBDQIPSFOTYSMZFVDCQPBBVORFVKDYHOMILSDCFZPYZGPBUAMJFTEHWITKOONZWNWCLUEBMOQODMWYIDIXTVUCADEQEMXMGJPULRZVDGNWWXRPEYCALMLOFOQNFXQXOPK9CZTUXGKXHGZUCQQYHZRKCVGTIUOSLQT9XOUTHEHLQ9I9JIZLX9EEOCZCAYINFQOXJSELJBUQMUTSZAQWDOFRYOXEDYBBEE9XIWXQ9YQFEAAUZW9WXIBMMBEGYRPIQCBVDZTZOAPU9RQ9XOMEWNAOB9VVKZGCQWB9RAFQUGCBWQS9INRSZDQWXJOPPPXC9RRQE9GDWEBJNXNRZ9COSDRDZKAEWIULDYRIZK9OLIDHKFSCHBOFPGJYKUNGSZZMNCDL9PEHWYG9UNVWVYEVFSGJEZNFHYUIPYSAG9BZQEVTEKME9LXLZDNRPTSN9IOHAJTLTLGYSLLSSZRAHMSBYGWBMUXUSQMOFUKPWODVCSLFAIDWNIRLSLJLMRLMLDSDD9RWZAXKCBFXSYAMBKYKCTTXPWSROHYQYHXOFA9PLR9RMZFMOWRFOWGHCVCTCLQSCLONWBPCVDPSZFJDAWHVJDYUHBTTNZDBJPPDBPOFDUEYDSPNEQU9LSD9VIE9GIF9IFLYV9LDAEDJNXOLYHGAVEYEXDYFZV9KUXTJAPMEHVOTOFUGYJQDSGFXNBPWFDOGTHONHBWRMQYSZBTOZGJYP9KXOLNNPMRXMPK9FMMYBUSVRDCZKOVHMRX9BVQCKOGOOINUIWFCVFUKKJIMOGCO9ULHOQ9BNAV9PDQWOMQVVYURZ9RNWUQONSPOH9MSRWPQIODQSPIWMSBNOOISTQTPQNJLQHPIVRVOMQPUSTCNLDHEMETTMWLDGXKTMZBIRHULAMODREANQZIZJFXCYKVJFWZOELHCOZQBLVH9LINGTYCUPZRDBJIOYDFGIGYZIDPFJTAAYAKTAHUVWACUOWZTDVYNJYPIRTFVBROFQI9SVKDWTPTHXTHE9ENGUKPSK9LRTJOX9QAYYSUBSHAEWRUETOGETLJAMH9HLSGEDFVOICFWCROFHLCLXHAYSBYZDUXBAOUFJYJKTYCTHKOM99C9NVNNJBPDSDPWUHINPIBQGXDZLTHVMBHFDSHBLOBYADPDWTONKLJVCAUQWCPAUT
Fragment 1: II9WTVPAPPJDPHXBCOVHIMFYX9EKDUVKSSRJUOIHRDHUSEHBGLWYFQKFBAWHVDRYVY99JXGPYBIOSPZLPUHBNMEISNPCOKAPHGWCISALLDOOHV9COLOPKVXGIKXHJSOIFJIRHOPKA9SZBEJJHBUMXBNYTJKEDNRMEINHUUOPIASBAOCHSLHIPETJQIDZJNNVMBKHXTSLFOKKZYJQOMJJAEGPYZUWKSPLACOYOYIQSKXM9FZ9K9BRWMKPVTBIHKYYZXXLC99FUAULHVIJQ9RZJJVZ9HWZEOLYLDGEFQRBDQCVACVGJGLWIKAMKAUBMAMWJDRGRRPKSQB9KG9BKGYR9QIGFIJVLAAXYIMUZDDCEDXZFODTCTOAHCSPR9IEKPVTGLJLYHUIDWLLZBRNENJGNJLNIAVAGXWJ9GMKHWOBUXGJWDQOJRRO9XNTNOAONQBHHNVQLRXVNELWMOW9GUAN9WMOQKATISYVSZCOAMJBHBGAWAFWHJPPWDJJKHETGMQFXXQQGPYGAEBMBXSCBATRJBQCDIFPQQSKQTXRKMVRRHODYUESXPMSUO9SYOBFZGNXP9GNMPVVBHHL9AKAUGEMRKKTTUYUTUUPSUBCZXJKP9INGLCPVTQ9AGPRNTTISFWNWXPUJYVCVUZHXXJFUSULEGCOUGWZLAVNYLBORWHXRTCOTPLETSKWPTETHDUEKGTMONLCUQBUUAIGRTAGDHCKYCVKQYZBOYRBEAIVLUWOZBSE9SUAOLOYPMIVXQLZDDNM9QABFSRGROVTQ9WTSDMVWQJWJUWXRUCSM9XSQWLLLEJPY9ZMNFB9TN9BPMPGCX9CBQIBAGK9KBSSTHLQWGCQZE9BZ9QFRVDTLMEESZV9U9QIZRZORAGMWUAM9RDH9NRNIWTJWIETX9FQAVASEP9IEPLXZZBMOMNKIFTBTTMQTUZ9ZXCLVNDUETJSWKIOGQMDEMNFETZLPOUSKCDRXQGJALCXOCMYIWMBC9PWRUBCGHMWMWHFAUVSE9BA9JHVACEQDIXCLEQZTTUETECYHVZZ9WDKJBVNALMJJZM9ECXZUNUR9VLH9QOXUPLSEF9CQITNPATRGPBXFPMGAIZIPGCKSKVBPXPQVOWABFESEXBUMKYTIAQ9NIXDZBTMZCWF9GSTPXAFMMPINZIAHPTLE9ILWHQNRXOSTKFJFLQGYXJQQNCGTCPANBXNZOFDSBWARWGWGDL9PGMJCUFYAWFMHWTFTHDWMRBHTJOXP9ASNVI9FMFCGOMLLLDDSGSPWLXCTZ9JEBJIEZJXTVKFG9XEGBDP99EUCHVQCSEEETNJHKIJDGFVPNPOSKOHXYA9WISDCPUT9DLUSWLOIWBLFJXAOTGYNQGEURZVTBJPSCPRHWBFCMMKVNUTDOAVXZYSBFKEAZGLYCDRENRXHRULFBOZCIQTRFFBKS9QPYCWMBODRVIWXAVMZYVNOQWV9RZBBBHUQKXARRCLSGUQQMRDDJOVV9HWSQNMLYHPBQYNDCJYAYPLRBUUUPUJHVYBQYQRDKQNJRJHISGHAYUITKKJRZMFI9ECSNAEMQE9UKYSSJDKMECI9SUKHLFFYL99JPTHVDM9DYSCSUSLFQKKYKFDMELVWEADOYUHXF9EPHSYSWOSEFQIVHNSTNYKUMFXNFRLOHFWU9SUNN9WWNTYDQWIBLQCILFWTPUPGLCVBDACWYVSLO9BVSRYGCMUFBJKTZDXQTRLEGTHGZVA9ABE9ET9ACRQILGKDKBLWJEVJJIESHKPOODJDCEMCENFUJOAP9OOJJAYGCUSYFHAWPRJFZXBEHRLRCSL9AWOSTUZKAHTEDBDIGADNYXLJDNRQA9YKLPAWFVICGPTDBDSSQJFJCM9EEWVMESPKBQJQCVCYNQLZXYH9QZCLXTIQWHL9PKKMTQHIIX9ZZBLUCBYAER9JSVOWSWD9OXJCIRWFCVKYSIRIGMZEUUUGLIBSXKEZP9IPKVUXBBNNMKXUKFWYEPPCMBTXFTFYFFYJZXJHTVJJIYEFPRALTNRXGE9ULHDUQRAMQWDTZYRKNECYDOYAPOUQMJRQAKOUZWKLPTIQ9CCWRHUIOXTKRWEWQHJPCOGHMWJX9OPPHHRDHTECLMCPZUDJVTIYAKQTYXAZFWVLTQRJXYMFJCN9VGTUFCUPRILFDGCJPQ9CZITXFYAVWSBKPDHURG

Verifying signed bundle:
SHA-256 hash is 3baef7059404e926f917bbf0dd2b0562dbea0d6a471c6815f3205385c11216a7
Curl-P hash is QLQGXIBGDZMQP9OYZELAPXFHECTSGTDLDAABBYCUWMZHSDTQOOVXZUPVPLOSPHEBIQ9VVHCN9GBYKXZVN
Signature is valid!

Verifying forged bundle:
SHA-256 hash is 1e43a3bfff1c1d6a676d337a12db2fadfc5c4d6ee437fb99b04042e31a634207
Curl-P hash is QLQGXIBGDZMQP9OYZELAPXFHECTSGTDLDAABBYCUWMZHSDTQOOVXZUPVPLOSPHEBIQ9VVHCN9GBYKXZVN
Signature is valid!
```
