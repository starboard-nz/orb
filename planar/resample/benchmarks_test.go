package resample_test

import (
	"encoding/json"
	"testing"

	"github.com/paulmach/orb/planar"
	"github.com/paulmach/orb/planar/resample"
)

func BenchmarkToMorePoints(b *testing.B) {
	ls := testLineString1()
	totalPoints := int(float64(len(ls)) * 1.616)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resample.Resample(ls, totalPoints)
	}
}

func BenchmarkLineStringResampleToLessPoints(b *testing.B) {
	ls := testLineString1()
	totalPoints := int(float64(len(ls)) / 1.616)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resample.Resample(ls, totalPoints)
	}
}

func testLineString1() planar.LineString {
	var data [][2]float64
	err := json.Unmarshal([]byte(`[[37.844221,-121.949564],[37.84419,-121.949493],[37.844166,-121.949419],[37.844141,-121.949346],[37.844118,-121.949277],[37.844091,-121.949209],[37.844061,-121.949145],[37.844032,-121.949083],[37.844004,-121.94902],[37.843976,-121.948959],[37.843946,-121.948897],[37.843917,-121.948836],[37.843883,-121.948775],[37.843849,-121.948722],[37.843817,-121.948664],[37.843784,-121.948611],[37.843748,-121.948559],[37.843714,-121.948506],[37.843678,-121.948451],[37.843638,-121.948396],[37.843603,-121.948342],[37.843567,-121.948284],[37.843535,-121.948223],[37.843504,-121.948161],[37.84347,-121.948102],[37.843441,-121.948042],[37.843419,-121.94798],[37.843396,-121.947924],[37.84336,-121.947881],[37.843313,-121.947863],[37.843263,-121.947867],[37.843224,-121.947895],[37.84319,-121.947936],[37.843159,-121.947979],[37.843126,-121.948024],[37.843097,-121.948072],[37.843067,-121.948122],[37.843038,-121.948173],[37.843006,-121.94822],[37.842968,-121.948258],[37.842925,-121.948297],[37.842875,-121.948324],[37.842819,-121.948334],[37.842758,-121.948328],[37.842698,-121.948298],[37.842636,-121.948259],[37.842566,-121.948214],[37.842499,-121.948172],[37.842439,-121.948133],[37.842391,-121.948082],[37.842356,-121.948021],[37.842341,-121.947954],[37.842336,-121.94788],[37.842345,-121.947797],[37.842362,-121.947718],[37.842383,-121.947647],[37.84241,-121.947576],[37.842442,-121.947505],[37.842474,-121.947438],[37.842509,-121.947372],[37.842545,-121.947308],[37.842582,-121.947243],[37.842621,-121.947178],[37.842665,-121.947118],[37.842712,-121.947062],[37.842764,-121.947012],[37.842815,-121.946964],[37.842861,-121.946919],[37.842905,-121.94687],[37.842949,-121.946818],[37.842993,-121.946766],[37.843034,-121.94671],[37.843072,-121.946652],[37.843111,-121.946592],[37.843147,-121.94653],[37.843179,-121.946467],[37.843209,-121.946406],[37.843241,-121.946353],[37.843277,-121.946304],[37.843313,-121.94625],[37.843351,-121.946198],[37.843385,-121.946146],[37.843422,-121.946089],[37.843451,-121.946032],[37.843459,-121.945973],[37.843457,-121.945911],[37.84344,-121.945848],[37.843391,-121.945815],[37.843326,-121.945783],[37.843271,-121.945754],[37.843224,-121.945719],[37.843181,-121.94567],[37.843139,-121.945622],[37.843103,-121.945561],[37.843068,-121.945506],[37.843035,-121.945455],[37.843005,-121.945407],[37.842972,-121.945359],[37.842938,-121.945315],[37.842904,-121.945271],[37.842871,-121.945224],[37.842837,-121.945179],[37.842806,-121.945135],[37.842771,-121.945089],[37.842733,-121.945039],[37.842694,-121.944988],[37.842659,-121.944931],[37.842628,-121.944876],[37.842595,-121.944826],[37.842562,-121.944777],[37.842529,-121.944722],[37.842496,-121.944663],[37.842468,-121.944604],[37.842437,-121.944536],[37.842403,-121.944469],[37.842366,-121.944404],[37.842334,-121.94434],[37.842306,-121.94428],[37.842278,-121.94422],[37.842249,-121.94416],[37.842215,-121.944101],[37.842177,-121.944038],[37.842141,-121.943981],[37.842107,-121.943927],[37.842074,-121.94387],[37.842045,-121.943812],[37.842024,-121.943751],[37.842025,-121.943675],[37.842038,-121.943592],[37.842061,-121.943517],[37.842092,-121.943447],[37.84213,-121.943386],[37.84217,-121.943321],[37.842216,-121.943252],[37.842263,-121.94318],[37.842315,-121.943109],[37.842371,-121.94304],[37.842432,-121.942972],[37.842494,-121.942902],[37.842557,-121.942826],[37.842616,-121.94275],[37.842672,-121.942677],[37.842729,-121.942599],[37.842779,-121.942517],[37.842824,-121.942434],[37.842872,-121.942357],[37.842924,-121.94228],[37.842975,-121.942208],[37.843024,-121.942139],[37.843074,-121.942077],[37.843127,-121.942022],[37.843185,-121.941961],[37.84324,-121.941892],[37.843292,-121.941813],[37.843342,-121.941735],[37.843388,-121.941659],[37.843436,-121.941592],[37.843474,-121.941527],[37.843513,-121.941465],[37.843543,-121.941406],[37.84353,-121.941343],[37.84349,-121.941308],[37.843423,-121.941303],[37.843351,-121.941309],[37.843279,-121.941308],[37.843208,-121.94129],[37.843146,-121.941248],[37.843094,-121.941188],[37.843038,-121.941128],[37.843002,-121.941044],[37.842978,-121.940955],[37.842967,-121.94086],[37.842956,-121.940766],[37.842946,-121.940665],[37.842966,-121.940573],[37.842996,-121.940491],[37.843046,-121.940423],[37.84311,-121.940377],[37.843167,-121.940321],[37.843215,-121.940252],[37.843249,-121.940177],[37.843249,-121.940102],[37.843222,-121.940057],[37.843174,-121.940053],[37.843112,-121.940079],[37.843056,-121.940104],[37.842996,-121.940131],[37.842936,-121.940169],[37.84288,-121.940202],[37.842825,-121.940229],[37.842766,-121.940264],[37.842713,-121.940304],[37.842654,-121.94035],[37.842593,-121.940392],[37.842538,-121.940437],[37.842487,-121.940482],[37.842441,-121.940525],[37.842386,-121.940561],[37.842334,-121.940589],[37.842273,-121.940609],[37.842212,-121.940613],[37.842153,-121.940607],[37.842093,-121.940593],[37.842031,-121.94057],[37.841976,-121.940535],[37.841921,-121.940499],[37.841869,-121.940465],[37.841824,-121.940427],[37.841782,-121.940387],[37.841742,-121.940346],[37.8417,-121.940305],[37.841653,-121.940263],[37.841605,-121.940216],[37.841559,-121.940167],[37.841515,-121.940103],[37.841481,-121.940032],[37.841444,-121.939965],[37.841416,-121.939891],[37.841393,-121.939818],[37.841378,-121.939736],[37.841366,-121.939653],[37.841355,-121.939571],[37.84134,-121.939489],[37.841324,-121.939412],[37.841305,-121.939338],[37.841289,-121.939269],[37.841273,-121.939203],[37.841249,-121.939144],[37.841223,-121.939087],[37.841168,-121.938969],[37.841158,-121.93894],[37.841145,-121.938902],[37.841123,-121.93883],[37.841102,-121.938752],[37.841086,-121.938669],[37.841077,-121.938581],[37.841081,-121.938488],[37.841089,-121.938399],[37.841099,-121.938317],[37.841109,-121.938237],[37.841121,-121.938157],[37.841144,-121.93808],[37.84117,-121.938007],[37.841199,-121.93794],[37.841232,-121.937876],[37.84127,-121.937816],[37.841308,-121.93776],[37.841352,-121.937705],[37.841398,-121.937652],[37.841442,-121.937602],[37.841489,-121.93755],[37.84154,-121.937501],[37.841595,-121.937456],[37.841654,-121.937414],[37.841712,-121.937375],[37.84177,-121.937339],[37.841829,-121.937303],[37.841885,-121.937268],[37.841942,-121.93724],[37.841999,-121.937206],[37.842059,-121.937177],[37.842119,-121.937158],[37.842178,-121.937142],[37.84224,-121.937128],[37.8423,-121.937117],[37.842361,-121.937107],[37.842421,-121.937096],[37.84248,-121.937077],[37.842539,-121.937057],[37.842594,-121.937034],[37.842644,-121.937009],[37.842697,-121.936986],[37.842751,-121.936965],[37.842803,-121.93695],[37.842861,-121.936942],[37.842919,-121.93694],[37.842976,-121.93694],[37.843019,-121.93694],[37.843065,-121.936943],[37.843117,-121.936944],[37.843173,-121.936939],[37.843231,-121.936938],[37.843287,-121.936941],[37.843347,-121.936949],[37.843407,-121.936957],[37.843465,-121.93697],[37.843518,-121.936996],[37.843573,-121.93702],[37.843629,-121.937045],[37.843682,-121.937069],[37.843736,-121.937084],[37.843795,-121.937084],[37.843848,-121.93708],[37.843905,-121.937071],[37.843961,-121.937061],[37.844014,-121.937053],[37.844067,-121.937043],[37.844117,-121.937033],[37.844167,-121.937023],[37.844217,-121.937011],[37.844264,-121.936996],[37.844308,-121.936972],[37.844341,-121.936941],[37.844361,-121.936901],[37.84437,-121.936856],[37.844368,-121.936806],[37.844355,-121.936759],[37.844332,-121.936713],[37.844297,-121.93666],[37.844258,-121.936604],[37.844219,-121.936554],[37.844173,-121.9365],[37.844132,-121.936446],[37.844087,-121.936394],[37.844039,-121.936344],[37.843994,-121.936295],[37.843956,-121.936242],[37.843909,-121.936196],[37.843865,-121.936147],[37.84382,-121.9361],[37.84377,-121.936059],[37.843728,-121.936014],[37.843696,-121.935956],[37.843669,-121.935899],[37.843645,-121.935842],[37.843635,-121.935781],[37.843626,-121.935719],[37.843622,-121.935648],[37.843619,-121.935573],[37.843625,-121.935493],[37.843637,-121.935407],[37.843641,-121.935325],[37.843645,-121.935243],[37.843645,-121.935163],[37.843641,-121.935081],[37.843638,-121.934999],[37.843634,-121.93492],[37.843634,-121.93484],[37.843632,-121.934764],[37.84363,-121.93469],[37.843629,-121.934615],[37.843626,-121.934538],[37.843625,-121.934459],[37.843623,-121.934379],[37.843624,-121.9343],[37.843629,-121.934221],[37.843644,-121.934145],[37.843666,-121.934078],[37.843699,-121.934022],[37.843739,-121.933972],[37.84378,-121.933929],[37.843821,-121.933892],[37.84386,-121.933853],[37.843895,-121.933811],[37.843928,-121.933768],[37.843955,-121.93373],[37.843981,-121.933689],[37.844008,-121.933649],[37.844031,-121.933603],[37.844049,-121.933551],[37.844057,-121.933498],[37.844053,-121.933441],[37.844036,-121.933387],[37.844008,-121.933345],[37.843977,-121.933305],[37.843937,-121.933263],[37.843895,-121.933217],[37.843852,-121.933178],[37.843806,-121.93314],[37.843761,-121.933106],[37.843711,-121.933081],[37.843667,-121.933065],[37.843622,-121.933058],[37.843572,-121.933059],[37.843519,-121.933064],[37.843463,-121.933066],[37.843403,-121.933062],[37.843343,-121.933043],[37.84328,-121.933017],[37.843218,-121.932992],[37.843163,-121.932967],[37.843111,-121.932942],[37.843063,-121.932915],[37.843014,-121.932893],[37.842964,-121.932865],[37.842911,-121.932843],[37.842858,-121.932826],[37.842803,-121.932811],[37.842751,-121.932795],[37.842701,-121.932776],[37.842654,-121.932751],[37.84261,-121.932721],[37.842567,-121.932686],[37.842527,-121.932648],[37.84249,-121.932605],[37.842452,-121.932559],[37.842416,-121.932507],[37.842386,-121.932452],[37.842359,-121.932396],[37.842326,-121.932341],[37.842285,-121.932294],[37.842243,-121.932252],[37.842199,-121.932218],[37.842152,-121.932197],[37.842099,-121.932182],[37.842048,-121.932166],[37.841997,-121.932161],[37.841945,-121.932156],[37.841896,-121.932149],[37.841847,-121.932133],[37.841799,-121.932113],[37.84175,-121.932096],[37.841703,-121.932081],[37.841657,-121.932061],[37.841606,-121.932042],[37.84156,-121.932024],[37.841518,-121.932001],[37.841485,-121.93197],[37.841464,-121.931921],[37.841456,-121.931857],[37.841472,-121.931784],[37.841489,-121.931721],[37.841511,-121.931658],[37.841548,-121.931604],[37.841597,-121.93156],[37.841651,-121.931523],[37.841708,-121.931489],[37.841767,-121.931466],[37.841825,-121.931444],[37.841888,-121.931416],[37.841949,-121.931376],[37.842007,-121.931327],[37.842055,-121.93126],[37.842095,-121.931183],[37.842126,-121.931101],[37.842147,-121.931012],[37.842161,-121.930926],[37.842173,-121.930843],[37.842182,-121.930762],[37.842191,-121.93068],[37.842198,-121.930595],[37.842205,-121.930512],[37.842214,-121.930432],[37.842217,-121.930354],[37.842215,-121.930274],[37.842213,-121.930196],[37.842208,-121.930123],[37.842204,-121.930052],[37.8422,-121.929982],[37.8422,-121.929912],[37.84221,-121.929845],[37.842222,-121.92978],[37.842232,-121.92971],[37.842238,-121.929638],[37.842237,-121.929567],[37.842227,-121.929499],[37.842204,-121.929431],[37.842176,-121.929366],[37.842142,-121.929302],[37.842103,-121.929241],[37.842061,-121.929181],[37.842021,-121.929123],[37.841977,-121.929068],[37.841934,-121.929013],[37.841896,-121.928955],[37.841866,-121.928888],[37.841842,-121.928813],[37.841829,-121.92874],[37.841841,-121.928657],[37.841862,-121.928576],[37.841885,-121.928497],[37.841905,-121.928419],[37.841927,-121.928337],[37.841944,-121.928253],[37.841963,-121.928172],[37.841982,-121.928091],[37.842003,-121.928012],[37.842023,-121.927938],[37.842051,-121.92787],[37.84209,-121.927811],[37.84214,-121.927764],[37.842193,-121.927722],[37.842246,-121.927681],[37.842289,-121.927633],[37.842333,-121.927576],[37.842373,-121.927516],[37.842406,-121.927455],[37.84243,-121.927392],[37.842442,-121.927324],[37.842448,-121.927256],[37.842454,-121.927183],[37.842457,-121.927107],[37.842462,-121.927032],[37.842465,-121.926956],[37.842471,-121.926879],[37.842483,-121.926806],[37.842508,-121.926744],[37.842549,-121.926707],[37.842592,-121.92668],[37.842639,-121.926656],[37.842677,-121.92662],[37.842715,-121.92658],[37.842745,-121.926528],[37.84277,-121.926469],[37.842779,-121.926403],[37.84278,-121.926341],[37.842774,-121.926279],[37.842757,-121.926215],[37.842736,-121.926155],[37.842711,-121.926097],[37.842684,-121.926041],[37.842653,-121.925986],[37.842614,-121.92594],[37.842569,-121.925902],[37.842518,-121.925872],[37.842469,-121.92585],[37.842417,-121.92583],[37.842361,-121.925811],[37.842305,-121.925793],[37.842248,-121.925767],[37.842187,-121.92574],[37.842131,-121.925706],[37.842077,-121.925667],[37.842027,-121.925619],[37.841979,-121.925567],[37.841929,-121.925514],[37.841878,-121.925467],[37.841827,-121.925423],[37.841776,-121.925384],[37.841729,-121.925341],[37.841681,-121.925296],[37.841635,-121.925251],[37.841588,-121.925206],[37.84154,-121.925161],[37.841492,-121.925114],[37.841444,-121.925072],[37.841396,-121.925028],[37.841345,-121.924988],[37.841294,-121.924946],[37.841242,-121.924904],[37.84119,-121.92486],[37.841131,-121.924817],[37.841077,-121.924763],[37.841031,-121.924702],[37.840996,-121.92463],[37.840977,-121.924551],[37.840977,-121.924458],[37.840975,-121.924367],[37.840975,-121.924275],[37.840969,-121.924186],[37.840958,-121.924098],[37.840948,-121.924013],[37.840934,-121.92393],[37.84092,-121.923847],[37.8409,-121.923765],[37.840874,-121.923686],[37.840848,-121.923609],[37.840824,-121.923529],[37.840801,-121.923444],[37.840779,-121.923358],[37.840757,-121.923267],[37.840736,-121.923173],[37.840715,-121.923076],[37.840698,-121.922975],[37.840692,-121.922872],[37.840696,-121.922769],[37.840702,-121.922667],[37.840714,-121.922563],[37.84072,-121.922466],[37.840731,-121.92237],[37.840747,-121.922274],[37.840765,-121.922179],[37.840786,-121.922084],[37.840818,-121.921987],[37.84086,-121.921902],[37.840913,-121.921823],[37.840975,-121.921753],[37.84104,-121.921682],[37.841096,-121.921602],[37.841152,-121.921527],[37.841211,-121.921457],[37.841271,-121.921398],[37.841339,-121.921358],[37.841415,-121.921334],[37.841487,-121.921317],[37.841565,-121.921329],[37.841642,-121.921361],[37.841715,-121.921401],[37.841788,-121.921441],[37.841863,-121.921485],[37.841936,-121.921523],[37.842006,-121.921561],[37.842079,-121.921599],[37.842154,-121.921637],[37.842232,-121.921673],[37.842309,-121.921705],[37.842382,-121.921731],[37.842451,-121.921743],[37.842508,-121.921715],[37.84257,-121.921671],[37.842612,-121.9216],[37.842623,-121.921503],[37.842617,-121.921386],[37.842613,-121.921269],[37.842611,-121.921158],[37.842616,-121.921048],[37.842623,-121.920932],[37.84265,-121.920817],[37.842673,-121.920713],[37.842708,-121.920613],[37.842762,-121.920528],[37.842825,-121.920462],[37.8429,-121.920414],[37.842986,-121.920399],[37.84307,-121.920412],[37.843143,-121.920456],[37.843201,-121.920529],[37.843255,-121.920608],[37.843301,-121.920691],[37.843348,-121.920777],[37.84339,-121.920868],[37.843429,-121.920962],[37.843465,-121.921058],[37.843495,-121.921156],[37.843519,-121.921245],[37.84355,-121.921331],[37.843572,-121.921417],[37.843598,-121.921499],[37.843628,-121.921583],[37.843657,-121.921673],[37.843685,-121.921759],[37.843714,-121.921839],[37.843744,-121.921917],[37.843782,-121.921999],[37.843824,-121.922078],[37.843869,-121.922153],[37.843923,-121.922225],[37.843985,-121.922302],[37.844048,-121.922386],[37.84411,-121.922471],[37.844163,-121.922556],[37.844217,-121.922651],[37.844266,-121.922753],[37.844303,-121.922854],[37.844331,-121.922957],[37.844349,-121.92307],[37.844371,-121.923179],[37.844384,-121.923283],[37.844404,-121.923385],[37.844425,-121.923491],[37.844447,-121.923592],[37.844466,-121.923691],[37.844478,-121.923788],[37.844499,-121.923886],[37.844527,-121.923985],[37.844546,-121.924081],[37.844575,-121.924178],[37.844614,-121.924271],[37.844652,-121.924356],[37.844686,-121.924437],[37.844718,-121.924519],[37.844752,-121.924598],[37.844785,-121.924679],[37.844827,-121.924758],[37.844864,-121.924846],[37.844897,-121.92494],[37.844922,-121.925045],[37.84494,-121.925153],[37.84494,-121.925262],[37.844933,-121.925371],[37.844931,-121.925483],[37.844937,-121.925596],[37.844956,-121.925709],[37.844975,-121.925819],[37.845008,-121.925935],[37.84503,-121.926044],[37.845053,-121.926149],[37.845079,-121.926245],[37.845107,-121.926336],[37.845135,-121.926428],[37.845153,-121.926513],[37.845184,-121.926599],[37.845221,-121.926686],[37.845257,-121.92677],[37.845301,-121.926856],[37.845343,-121.926932],[37.845389,-121.926992],[37.845446,-121.927039],[37.845524,-121.927061],[37.845624,-121.927052],[37.845681,-121.927007],[37.845734,-121.926945],[37.845793,-121.926888],[37.845855,-121.926839],[37.845931,-121.926814],[37.846014,-121.926816],[37.846093,-121.926827],[37.846169,-121.926835],[37.846242,-121.926848],[37.846316,-121.926854],[37.846385,-121.926865],[37.846453,-121.926886],[37.846522,-121.926915],[37.846587,-121.926955],[37.846646,-121.927017],[37.84669,-121.92711],[37.846724,-121.927214],[37.846763,-121.927323],[37.846812,-121.927436],[37.846881,-121.927558],[37.846956,-121.92768],[37.847025,-121.927809],[37.847089,-121.927944],[37.847142,-121.928091],[37.847173,-121.928238],[37.847193,-121.928388],[37.847206,-121.928544],[37.847225,-121.928698],[37.847248,-121.92884],[37.847294,-121.928968],[37.847351,-121.929078],[37.847433,-121.929181],[37.847516,-121.929275],[37.847596,-121.929383],[37.847666,-121.929502],[37.847722,-121.929628],[37.84776,-121.92977],[37.847799,-121.929907],[37.84785,-121.930041],[37.8479,-121.930166],[37.84795,-121.930294],[37.847996,-121.930424],[37.848041,-121.930555],[37.848089,-121.930689],[37.848136,-121.930821],[37.848178,-121.930954],[37.848225,-121.931092],[37.848276,-121.931225],[37.848334,-121.931357],[37.848392,-121.931491],[37.848451,-121.931626],[37.848518,-121.931758],[37.848589,-121.931882],[37.848668,-121.931994],[37.848747,-121.932099],[37.848827,-121.932205],[37.848903,-121.932312],[37.848983,-121.932414],[37.849063,-121.932505],[37.849153,-121.932588],[37.84925,-121.93266],[37.849347,-121.932734],[37.849449,-121.932801],[37.849563,-121.932858],[37.849666,-121.932897],[37.849763,-121.932909],[37.849857,-121.932904],[37.849952,-121.932895],[37.850059,-121.932878],[37.850179,-121.932854],[37.850301,-121.932846],[37.850422,-121.932844],[37.850534,-121.932859],[37.850647,-121.932862],[37.850757,-121.932852],[37.850862,-121.932802],[37.850955,-121.932705],[37.851037,-121.932573],[37.85109,-121.932454],[37.851122,-121.932328],[37.851143,-121.932206],[37.851149,-121.932087],[37.851146,-121.931985],[37.851134,-121.931908],[37.85109,-121.931864],[37.851014,-121.931838],[37.850897,-121.931801],[37.850783,-121.931721],[37.850682,-121.931629],[37.8506,-121.931516],[37.850558,-121.931414],[37.850554,-121.931299],[37.850568,-121.931191],[37.850611,-121.931104],[37.850687,-121.931032],[37.850758,-121.930961],[37.850824,-121.930894],[37.850893,-121.930835],[37.850965,-121.930786],[37.851041,-121.930739],[37.851106,-121.930679],[37.851153,-121.930597],[37.851187,-121.930515],[37.8512,-121.93043],[37.851203,-121.930341],[37.851187,-121.93026],[37.851147,-121.930206],[37.851072,-121.930169],[37.850992,-121.930148],[37.850908,-121.930152],[37.850834,-121.93016],[37.850757,-121.930179],[37.850675,-121.930208],[37.850586,-121.930241],[37.850499,-121.93025],[37.850425,-121.93023],[37.850366,-121.930188],[37.850325,-121.930122],[37.850306,-121.930042],[37.850313,-121.929961],[37.850346,-121.929896],[37.850383,-121.929832],[37.850447,-121.92978],[37.850509,-121.929737],[37.850575,-121.929707],[37.850645,-121.929688],[37.850716,-121.929667],[37.850791,-121.929652],[37.850867,-121.929636],[37.850948,-121.929614],[37.85102,-121.92958],[37.851075,-121.929531],[37.851139,-121.929481],[37.851196,-121.929429],[37.851251,-121.92938],[37.851306,-121.929321],[37.85136,-121.929261],[37.851411,-121.929204],[37.851459,-121.929138],[37.851501,-121.92907],[37.851539,-121.928999],[37.851572,-121.928928],[37.851604,-121.928857],[37.851636,-121.928786],[37.851667,-121.928714],[37.851703,-121.928639],[37.851734,-121.928565],[37.851765,-121.928495],[37.851791,-121.928427],[37.851812,-121.928356],[37.851822,-121.928291],[37.851813,-121.928231],[37.851792,-121.92818],[37.851754,-121.928136],[37.851712,-121.928097],[37.851667,-121.928053],[37.851624,-121.928016],[37.851581,-121.92797],[37.851541,-121.927915],[37.851505,-121.927854],[37.851485,-121.927773],[37.851476,-121.927688],[37.851486,-121.927601],[37.851513,-121.927517],[37.851545,-121.927439],[37.851579,-121.92736],[37.851611,-121.927279],[37.851636,-121.9272],[37.85166,-121.92712],[37.851672,-121.927041],[37.851669,-121.926963],[37.851662,-121.926893],[37.851646,-121.926831],[37.851624,-121.926774],[37.851589,-121.926722],[37.851551,-121.92667],[37.851505,-121.926617],[37.851459,-121.926562],[37.851417,-121.926505],[37.851395,-121.926431],[37.851396,-121.926353],[37.851427,-121.926284],[37.851472,-121.926224],[37.851524,-121.92617],[37.851576,-121.926116],[37.85163,-121.926062],[37.851685,-121.926007],[37.85174,-121.92595],[37.851796,-121.925889],[37.851847,-121.92583],[37.8519,-121.925775],[37.851952,-121.92573],[37.852008,-121.925704],[37.852074,-121.925696],[37.852143,-121.925718],[37.852206,-121.925769],[37.852263,-121.925844],[37.852303,-121.925937],[37.852333,-121.926032],[37.852362,-121.926136],[37.852392,-121.926245],[37.852422,-121.926353],[37.852453,-121.926466],[37.852485,-121.926579],[37.852513,-121.926696],[37.852542,-121.926812],[37.852571,-121.926927],[37.852607,-121.92704],[37.852647,-121.927148],[37.852682,-121.927259],[37.852716,-121.927368],[37.852742,-121.927478],[37.852772,-121.927587],[37.852801,-121.927697],[37.85283,-121.927808],[37.852855,-121.927912],[37.852875,-121.928013],[37.852886,-121.928111],[37.852892,-121.928209],[37.852893,-121.928309],[37.852895,-121.928408],[37.852904,-121.928504],[37.852907,-121.928599],[37.852906,-121.928689],[37.852903,-121.928777],[37.852902,-121.928867],[37.852899,-121.928954],[37.852901,-121.929043],[37.852909,-121.929135],[37.852922,-121.929227],[37.85295,-121.929318],[37.852986,-121.929397],[37.853034,-121.929466],[37.853089,-121.929536],[37.853137,-121.929614],[37.853186,-121.929693],[37.853232,-121.929784],[37.853267,-121.929877],[37.85329,-121.929978],[37.853312,-121.93008],[37.853334,-121.93018],[37.853362,-121.930278],[37.853398,-121.930371],[37.853435,-121.930461],[37.85348,-121.93054],[37.853535,-121.930604],[37.853591,-121.93066],[37.85365,-121.930718],[37.853724,-121.93077],[37.853795,-121.930815],[37.853862,-121.930861],[37.853926,-121.930911],[37.85399,-121.930959],[37.854051,-121.931011],[37.854112,-121.931061],[37.854173,-121.931104],[37.854236,-121.931148],[37.854295,-121.931187],[37.854351,-121.931223],[37.854409,-121.931269],[37.854471,-121.931322],[37.854528,-121.931374],[37.854581,-121.931434],[37.854632,-121.931497],[37.854683,-121.931562],[37.854733,-121.931625],[37.854779,-121.931685],[37.854826,-121.931741],[37.85488,-121.931804],[37.854935,-121.931884],[37.854993,-121.931964],[37.85505,-121.932039],[37.855105,-121.932113],[37.855164,-121.932189],[37.855214,-121.932263],[37.855268,-121.932329],[37.85531,-121.932389],[37.855353,-121.93247],[37.855399,-121.932568],[37.855435,-121.93265],[37.855463,-121.932729],[37.855489,-121.932809],[37.855505,-121.932884],[37.855514,-121.932969],[37.855521,-121.933055],[37.855529,-121.933145],[37.855534,-121.933234],[37.855545,-121.933328],[37.85556,-121.933423],[37.855576,-121.93351],[37.855593,-121.933583],[37.855613,-121.933652],[37.855643,-121.933715],[37.855672,-121.933776],[37.855705,-121.933832],[37.855739,-121.933877],[37.855776,-121.933927],[37.855812,-121.933975],[37.855847,-121.934028],[37.855878,-121.934082],[37.855911,-121.934138],[37.855945,-121.934201],[37.855982,-121.934266],[37.85602,-121.934332],[37.85606,-121.934392],[37.856104,-121.934439],[37.856155,-121.934473],[37.856214,-121.934502],[37.856271,-121.934527],[37.85633,-121.934545],[37.856393,-121.934557],[37.856455,-121.934567],[37.856519,-121.934581],[37.856586,-121.934594],[37.85665,-121.934615],[37.856716,-121.934652],[37.856787,-121.934699],[37.856849,-121.934764],[37.856907,-121.934833],[37.856961,-121.934908],[37.857009,-121.934994],[37.857055,-121.935084],[37.857104,-121.935177],[37.857148,-121.935265],[37.857186,-121.935355],[37.857232,-121.935449],[37.857278,-121.935536],[37.857332,-121.935606],[37.857402,-121.935663],[37.857488,-121.935694],[37.857579,-121.935703],[37.857673,-121.935692],[37.857769,-121.935675],[37.857867,-121.935657],[37.857962,-121.935643],[37.858064,-121.935628],[37.858168,-121.935619],[37.858277,-121.935599],[37.858388,-121.935578],[37.8585,-121.935559],[37.858611,-121.935533],[37.858718,-121.935502],[37.858819,-121.935472],[37.858911,-121.93544],[37.858991,-121.935399],[37.859055,-121.935352],[37.859111,-121.935302],[37.859164,-121.935248],[37.859216,-121.935196],[37.859267,-121.935148],[37.859318,-121.935097],[37.859365,-121.935048],[37.859414,-121.934998],[37.859461,-121.934942],[37.859506,-121.934888],[37.85955,-121.93484],[37.859593,-121.934794],[37.859637,-121.934752],[37.859685,-121.93471],[37.85973,-121.934669],[37.859775,-121.934627],[37.859821,-121.934587],[37.859868,-121.93455],[37.859915,-121.934511],[37.859962,-121.934476],[37.860012,-121.934441],[37.860062,-121.934405],[37.860112,-121.934379],[37.860166,-121.93435],[37.860226,-121.934322],[37.860284,-121.934297],[37.860338,-121.934272],[37.860394,-121.934251],[37.860448,-121.934226],[37.860502,-121.934204],[37.860558,-121.934184],[37.860618,-121.934159],[37.860676,-121.934131],[37.860731,-121.934098],[37.860788,-121.934064],[37.860843,-121.934029],[37.860899,-121.933996],[37.860957,-121.933962],[37.861014,-121.933925],[37.861071,-121.933884],[37.86113,-121.933847],[37.861194,-121.933808],[37.861253,-121.933774],[37.861308,-121.933739],[37.861367,-121.933701],[37.861421,-121.933663],[37.861473,-121.933622],[37.861524,-121.933583],[37.861575,-121.933543],[37.861629,-121.933502],[37.861681,-121.933459],[37.861735,-121.933416],[37.86179,-121.933371],[37.861847,-121.933326],[37.861905,-121.933284],[37.861961,-121.933242],[37.86202,-121.9332],[37.86208,-121.933161],[37.862137,-121.933123],[37.862197,-121.933087],[37.862257,-121.933048],[37.862313,-121.933007],[37.862369,-121.932969],[37.86243,-121.932932],[37.862485,-121.932893],[37.862537,-121.932849],[37.86259,-121.932807],[37.862641,-121.932767],[37.862695,-121.932731],[37.86275,-121.932705],[37.862807,-121.932684],[37.862866,-121.932668],[37.862921,-121.93267],[37.862978,-121.932685],[37.863035,-121.932703],[37.863097,-121.932727],[37.86316,-121.932754],[37.863223,-121.932783],[37.863288,-121.932806],[37.863352,-121.932832],[37.863414,-121.932864],[37.863478,-121.9329],[37.863538,-121.932939],[37.863595,-121.93298],[37.863654,-121.93302],[37.86371,-121.93306],[37.863769,-121.9331],[37.86383,-121.933131],[37.863883,-121.933149],[37.863934,-121.93315],[37.863982,-121.933137],[37.864029,-121.933107],[37.864074,-121.933061],[37.864106,-121.932999],[37.864122,-121.932928],[37.864146,-121.93286],[37.864171,-121.932789],[37.864195,-121.932728],[37.864225,-121.932672],[37.864263,-121.932628],[37.864303,-121.932593],[37.864343,-121.932575],[37.864384,-121.932577],[37.864414,-121.932608],[37.864435,-121.932653],[37.864448,-121.932703],[37.864463,-121.932761],[37.86448,-121.932829],[37.864506,-121.932894],[37.864538,-121.932951],[37.864586,-121.932995],[37.864637,-121.933028],[37.86469,-121.933054],[37.864742,-121.933077],[37.864789,-121.933103],[37.864836,-121.933132],[37.864878,-121.93316],[37.864913,-121.933191],[37.864951,-121.933221],[37.864991,-121.933242],[37.865029,-121.933242],[37.865065,-121.933219],[37.865103,-121.933184],[37.865137,-121.933138],[37.865168,-121.933083],[37.865204,-121.933036],[37.865238,-121.932983],[37.865274,-121.932915],[37.865303,-121.932852],[37.86533,-121.932788],[37.865358,-121.932727],[37.865385,-121.932665],[37.865408,-121.932597],[37.865439,-121.932528],[37.865471,-121.932463],[37.8655,-121.932405],[37.865536,-121.932339],[37.865571,-121.932274],[37.865601,-121.932213],[37.865634,-121.932159],[37.865659,-121.9321],[37.865682,-121.932041],[37.865706,-121.931984],[37.865731,-121.931927],[37.865756,-121.931875],[37.865789,-121.93183],[37.865838,-121.931801],[37.865883,-121.931782],[37.86593,-121.931782],[37.865974,-121.931797],[37.866024,-121.931824],[37.866075,-121.931852],[37.86613,-121.93188],[37.866181,-121.93191],[37.866232,-121.93193],[37.866285,-121.931954],[37.86634,-121.931974],[37.866391,-121.932006],[37.866437,-121.932038]]`), &data)
	if err != nil {
		panic(err)
	}

	ls := planar.NewLineStringFromYXData(data)
	lineString := planar.NewLineString()

	// this parameter can be used to thin to line to see if worst cases are impacting averages
	skip := 5
	for i := 0; i < len(ls); i += skip {
		lineString = append(lineString, ls[i])
	}

	return lineString
}