package easportswrc

import (
	"encoding/binary"
	"fmt"
	"math"
)

const PacketEASportsWRCLength = 237

type PacketEASportsWRC struct {
	PacketUid                 uint64  // 0
	GameTotalTime             float32 // 1
	GameDeltaTime             float32 // 2
	GameFrameCount            uint64  // 3
	ShiftlightsFraction       float32 // 4
	ShiftlightsRpmStart       float32 // 5
	ShiftlightsRpmEnd         float32 // 6
	ShiftlightsRpmValid       bool    // 7
	VehicleGearIndex          uint8   // 8
	VehicleGearIndexNeutral   uint8   // 9
	VehicleGearIndexReverse   uint8   // 10
	VehicleGearMaximum        uint8   // 11
	VehicleSpeed              float32 // 12
	VehicleTransmissionSpeed  float32 // 13
	VehiclePositionX          float32 // 14
	VehiclePositionY          float32 // 15
	VehiclePositionZ          float32 // 16
	VehicleVelocityX          float32 // 17
	VehicleVelocityY          float32 // 18
	VehicleVelocityZ          float32 // 19
	VehicleAccelerationX      float32 // 20
	VehicleAccelerationY      float32 // 21
	VehicleAccelerationZ      float32 // 22
	VehicleLeftDirectionX     float32 // 23
	VehicleLeftDirectionY     float32 // 24
	VehicleLeftDirectionZ     float32 // 25
	VehicleForwardDirectionX  float32 // 26
	VehicleForwardDirectionY  float32 // 27
	VehicleForwardDirectionZ  float32 // 28
	VehicleUpDirectionX       float32 // 29
	VehicleUpDirectionY       float32 // 30
	VehicleUpDirectionZ       float32 // 31
	VehicleHubPositionBl      float32 // 32
	VehicleHubPositionBr      float32 // 33
	VehicleHubPositionFl      float32 // 34
	VehicleHubPositionFr      float32 // 35
	VehicleHubVelocityBl      float32 // 36
	VehicleHubVelocityBr      float32 // 37
	VehicleHubVelocityFl      float32 // 38
	VehicleHubVelocityFr      float32 // 39
	VehicleCpForwardSpeedBl   float32 // 40
	VehicleCpForwardSpeedBr   float32 // 41
	VehicleCpForwardSpeedFl   float32 // 42
	VehicleCpForwardSpeedFr   float32 // 43
	VehicleBrakeTemperatureBl float32 // 44
	VehicleBrakeTemperatureBr float32 // 45
	VehicleBrakeTemperatureFl float32 // 46
	VehicleBrakeTemperatureFr float32 // 47
	VehicleEngineRpmMax       float32 // 48
	VehicleEngineRpmIdle      float32 // 49
	VehicleEngineRpmCurrent   float32 // 50
	VehicleThrottle           float32 // 51
	VehicleBrake              float32 // 52
	VehicleClutch             float32 // 53
	VehicleSteering           float32 // 54
	VehicleHandbrake          float32 // 55
	StageCurrentTime          float32 // 56
	StageCurrentDistance      float64 // 57
	StageLength               float64 // 58
}

func (p *PacketEASportsWRC) UnmarshalBinary(b []byte) error {
	if len(b) < PacketEASportsWRCLength {
		return fmt.Errorf("invalid packet size: %d", len(b))
	}
	p.PacketUid = binary.LittleEndian.Uint64(b[0:8])
	p.GameTotalTime = math.Float32frombits(binary.LittleEndian.Uint32(b[8:12]))
	p.GameDeltaTime = math.Float32frombits(binary.LittleEndian.Uint32(b[12:16]))
	p.GameFrameCount = binary.LittleEndian.Uint64(b[16:24])
	p.ShiftlightsFraction = math.Float32frombits(binary.LittleEndian.Uint32(b[24:28]))
	p.ShiftlightsRpmStart = math.Float32frombits(binary.LittleEndian.Uint32(b[28:32]))
	p.ShiftlightsRpmEnd = math.Float32frombits(binary.LittleEndian.Uint32(b[32:36]))
	p.ShiftlightsRpmValid = b[36] != 0
	p.VehicleGearIndex = b[37]
	p.VehicleGearIndexNeutral = b[38]
	p.VehicleGearIndexReverse = b[39]
	p.VehicleGearMaximum = b[40]
	p.VehicleSpeed = math.Float32frombits(binary.LittleEndian.Uint32(b[41:45]))
	p.VehicleTransmissionSpeed = math.Float32frombits(binary.LittleEndian.Uint32(b[45:49]))
	p.VehiclePositionX = math.Float32frombits(binary.LittleEndian.Uint32(b[49:53]))
	p.VehiclePositionY = math.Float32frombits(binary.LittleEndian.Uint32(b[53:57]))
	p.VehiclePositionZ = math.Float32frombits(binary.LittleEndian.Uint32(b[57:61]))
	p.VehicleVelocityX = math.Float32frombits(binary.LittleEndian.Uint32(b[61:65]))
	p.VehicleVelocityY = math.Float32frombits(binary.LittleEndian.Uint32(b[65:69]))
	p.VehicleVelocityZ = math.Float32frombits(binary.LittleEndian.Uint32(b[69:73]))
	p.VehicleAccelerationX = math.Float32frombits(binary.LittleEndian.Uint32(b[73:77]))
	p.VehicleAccelerationY = math.Float32frombits(binary.LittleEndian.Uint32(b[77:81]))
	p.VehicleAccelerationZ = math.Float32frombits(binary.LittleEndian.Uint32(b[81:85]))
	p.VehicleLeftDirectionX = math.Float32frombits(binary.LittleEndian.Uint32(b[85:89]))
	p.VehicleLeftDirectionY = math.Float32frombits(binary.LittleEndian.Uint32(b[89:93]))
	p.VehicleLeftDirectionZ = math.Float32frombits(binary.LittleEndian.Uint32(b[93:97]))
	p.VehicleForwardDirectionX = math.Float32frombits(binary.LittleEndian.Uint32(b[97:101]))
	p.VehicleForwardDirectionY = math.Float32frombits(binary.LittleEndian.Uint32(b[101:105]))
	p.VehicleForwardDirectionZ = math.Float32frombits(binary.LittleEndian.Uint32(b[105:109]))
	p.VehicleUpDirectionX = math.Float32frombits(binary.LittleEndian.Uint32(b[109:113]))
	p.VehicleUpDirectionY = math.Float32frombits(binary.LittleEndian.Uint32(b[113:117]))
	p.VehicleUpDirectionZ = math.Float32frombits(binary.LittleEndian.Uint32(b[117:121]))
	p.VehicleHubPositionBl = math.Float32frombits(binary.LittleEndian.Uint32(b[121:125]))
	p.VehicleHubPositionBr = math.Float32frombits(binary.LittleEndian.Uint32(b[125:129]))
	p.VehicleHubPositionFl = math.Float32frombits(binary.LittleEndian.Uint32(b[129:133]))
	p.VehicleHubPositionFr = math.Float32frombits(binary.LittleEndian.Uint32(b[133:137]))
	p.VehicleHubVelocityBl = math.Float32frombits(binary.LittleEndian.Uint32(b[137:141]))
	p.VehicleHubVelocityBr = math.Float32frombits(binary.LittleEndian.Uint32(b[141:145]))
	p.VehicleHubVelocityFl = math.Float32frombits(binary.LittleEndian.Uint32(b[145:149]))
	p.VehicleHubVelocityFr = math.Float32frombits(binary.LittleEndian.Uint32(b[149:153]))
	p.VehicleCpForwardSpeedBl = math.Float32frombits(binary.LittleEndian.Uint32(b[153:157]))
	p.VehicleCpForwardSpeedBr = math.Float32frombits(binary.LittleEndian.Uint32(b[157:161]))
	p.VehicleCpForwardSpeedFl = math.Float32frombits(binary.LittleEndian.Uint32(b[161:165]))
	p.VehicleCpForwardSpeedFr = math.Float32frombits(binary.LittleEndian.Uint32(b[165:169]))
	p.VehicleBrakeTemperatureBl = math.Float32frombits(binary.LittleEndian.Uint32(b[169:173]))
	p.VehicleBrakeTemperatureBr = math.Float32frombits(binary.LittleEndian.Uint32(b[173:177]))
	p.VehicleBrakeTemperatureFl = math.Float32frombits(binary.LittleEndian.Uint32(b[177:181]))
	p.VehicleBrakeTemperatureFr = math.Float32frombits(binary.LittleEndian.Uint32(b[181:185]))
	p.VehicleEngineRpmMax = math.Float32frombits(binary.LittleEndian.Uint32(b[185:189]))
	p.VehicleEngineRpmIdle = math.Float32frombits(binary.LittleEndian.Uint32(b[189:193]))
	p.VehicleEngineRpmCurrent = math.Float32frombits(binary.LittleEndian.Uint32(b[193:197]))
	p.VehicleThrottle = math.Float32frombits(binary.LittleEndian.Uint32(b[197:201]))
	p.VehicleBrake = math.Float32frombits(binary.LittleEndian.Uint32(b[201:205]))
	p.VehicleClutch = math.Float32frombits(binary.LittleEndian.Uint32(b[205:209]))
	p.VehicleSteering = math.Float32frombits(binary.LittleEndian.Uint32(b[209:213]))
	p.VehicleHandbrake = math.Float32frombits(binary.LittleEndian.Uint32(b[213:217]))
	p.StageCurrentTime = math.Float32frombits(binary.LittleEndian.Uint32(b[217:221]))
	p.StageCurrentDistance = math.Float64frombits(binary.LittleEndian.Uint64(b[221:229]))
	p.StageLength = math.Float64frombits(binary.LittleEndian.Uint64(b[229:237]))
	return nil
}

type Location struct {
	Name   string
	Stages []string
}

type StageID struct {
	Location int
	Stage    int
}

type Stage struct {
	Location string
	Stage    string
}

var (
	LocationKeys = []string{
		"monte-carlo",
		"sweden",
		"mexico",
		"croatia",
		"portugal",
		"sardegna",
		"kenya",
		"estonia",
		"finland",
		"greece",
		"chile",
		"europe",
		"japan",
		"mediterraneo",
		"pacifico",
		"oceania",
		"scandia",
		"iberia",
	}
	Locations = []Location{
		// 0
		{"Rallye Monte-Carlo", []string{
			"La Bollène-Vésubie - Peïra Cava",
			"Peïra Cava - La Bollène-Vésubie",
			"La Bollène-Vésubie - Col de Turini",
			"Pra d'Alart",
			"La Maïris",
			"Baisse de Patronel",
			"Saint-Léger-les-Mélèzes - La Bâtie-Neuve",
			"La Bâtie-Neuve - Saint-Léger-les-Mélèzes",
			"Moissière",
			"Ancelle",
			"Ravin de Coste Belle",
			"Les Borels",
		}},
		// 1
		{"Rally Sweden", []string{
			"Hof-Finnskog",
			"Åsnes",
			"Spikbrenna",
			"Lauksjøen",
			"Åslia",
			"Knapptjernet",
			"Vargasen",
			"Lövstaholm",
			"Älgsjön",
			"Ekshärad",
			"Stora Jangen",
			"Sunne",
		}},
		// 2
		{"Guanajuato Rally México", []string{
			"El Chocolate",
			"Otates",
			"Ortega",
			"Las Minas",
			"Ibarrilla",
			"Derramadero",
			"El Brinco",
			"Guanajuatito",
			"Alfaro",
			"Mesa Cuata",
			"San Diego",
			"El Mosquito",
		}},
		// 3
		{"Croatia Rally", []string{
			"Bliznec",
			"Trakošćan",
			"Vrbno",
			"Zagorska Sela",
			"Kumrovec",
			"Grdanjci",
			"Stojdraga",
			"Mali Lipovec",
			"Hartje",
			"Kostanjevac",
			"Krašić",
			"Petruš Vrh",
		}},
		// 4
		{"Vodafone Rally de Portugal", []string{
			"Baião",
			"Caminha",
			"Fridão",
			"Marão",
			"Ponte de Lima",
			"Viana do Castelo",
			"Ervideiro",
			"Celeiro",
			"Touca",
			"Vila Boa",
			"Carrazedo",
			"Anjos",
		}},
		// 5
		{"Rally Italia Sardegna", []string{
			"Rena Majore",
			"Monte Olia",
			"Littichedda",
			"Ala del Sardi",
			"Mamone",
			"Li Pinnenti",
			"Malti",
			"Bassacutena",
			"Bortigiadas",
			"Sa Mela",
			"Monte Muvri",
			"Monte Acuto",
		}},
		// 6
		{"Safari Rally Kenya", []string{
			"Malewa",
			"Tarambete",
			"Moi North",
			"Marula",
			"Wileli",
			"Kingono",
			"Soysambu",
			"Mbaruk",
			"Sugunoi",
			"Nakuru",
			"Kanyawa",
			"Kanyawa - Nakura",
		}},
		// 7
		{"Rally Estonia", []string{
			"Otepää",
			"Rebaste",
			"Nüpli",
			"Truuta",
			"Koigu",
			"Kooraste",
			"Elva",
			"Metsalaane",
			"Vahessaare",
			"Külaaseme",
			"Vissi",
			"Vellavere",
		}},
		// 8
		{"SECTO Rally Finland", []string{
			"Leustu",
			"Lahdenkyla",
			"Saakoski",
			"Maahi",
			"Painna",
			"Peltola",
			"Paijala",
			"Ruokolahti",
			"Honkanen",
			"Venkajarvi",
			"Vehmas",
			"Hatanpaa",
		}},
		// 9
		{"EKO ACROPOLIS Rally Greece", []string{
			"Gravia",
			"Prosilio",
			"Mariolata",
			"Karoutes",
			"Viniani",
			"Delphi",
			"Eptalofos",
			"Lilea",
			"Parnassós",
			"Bauxites",
			"Drosochori",
			"Amfissa",
		}},
		// 10
		{"BIO BIO Rally Chile", []string{
			"Bio Bío",
			"Pulpería",
			"Río Lía",
			"María Las Cruces",
			"Las Paraguas",
			"Rere",
			"El Poñen",
			"Laja",
			"Yumbel",
			"Río Claro",
			"Hualqui",
			"Chivilingo",
		}},
		// 11
		{"Central Europe Rally", []string{
			"Rouské",
			"Lukoveček",
			"Raztoka",
			"Žabárna",
			"Provodovice",
			"Chvalčov",
			"Vítová",
			"Brusné",
			"Libosváry",
			"Rusava",
			"Osíčko",
			"Příkazy",
		}},
		// 12
		{"Forum8 Rally Japan", []string{
			"Lake Mikawa",
			"Kudarisawa",
			"Oninotaira",
			"Okuwacho",
			"Habu Dam",
			"Habucho",
			"Nenoue Plateau",
			"Tegano",
			"Higashino",
			"Hokono Lake",
			"Nenoue Highlands",
			"Nakatsugawa",
		}},
		// 13
		{"Rally Mediterraneo", []string{
			"Asco",
			"Ponte",
			"Monte Cinto",
			"Albarello",
			"Capannace",
			"Serra Di Cuzzioli",
			"Maririe",
			"Poggiola",
			"Monte Alloradu",
			"Ravin de Finelio",
			"Cabanella",
			"Moltifao",
		}},
		// 14
		{"Agon By AOC Rally Pacifico", []string{
			"Talao",
			"Talanghilirair",
			"SungaiKunit",
			"Sangir Balai Janggo",
			"South Solok",
			"Kebun Raya Solok",
			"Batukangkung",
			"Abai",
			"Moearaikoer",
			"Bidaralam",
			"Loeboekmalaka",
			"Gunung Tujuh",
		}},
		// 15
		{"Fanatec Rally Oceania", []string{
			"Oakleigh",
			"Doctors Hill",
			"Mangapai",
			"Brynderwyn",
			"Taipuha",
			"Mareretu",
			"Waiwera",
			"Tahekeroa",
			"Noakes Hill",
			"Orewa",
			"Tahekeroa - Orewa",
			"Makarau",
		}},
		// 16
		{"Rally Scandia", []string{
			"Holtjønn",
			"Hengeltjønn",
			"Fyresvatn",
			"Russvatn",
			"Tovsli",
			"Kottjønn",
			"Fordol",
			"Fyresdal",
			"Ljosdalstjønn",
			"Dagtrolltjønn",
			"Tovslioytjorn",
			"Bergsøytjønn",
		}},
		// 17
		{"Rally Iberia", []string{
			"Santes Creus",
			"Valldossera",
			"Campdasens",
			"Pontils",
			"Montagut",
			"Aiguamúrcia",
			"Alforja",
			"Les Irles",
			"L'Argentera",
			"Les Voltes",
			"Montclar",
			"Botareli",
		}},
	}
	Stages = map[float64]StageID{
		// 0:Monte-Carlo ----------------------------------------------------------
		18799.8984375:    {0, 0},  // fixed: La Bollène-Vésubie - Peïra Cava
		18606.03125:      {0, 1},  // new: Peïra Cava - La Bollène-Vésubie
		12349.2734375:    {0, 2},  // fixed: La Bollène-Vésubie - Col de Turini
		12167.2060546875: {0, 3},  // new: Pra d'Alart
		6745.568359375:   {0, 4},  // new: La Maïris
		6680.1611328125:  {0, 5},  // new: Baisse de Patronel
		17064.154296875:  {0, 6},  // new: Saint-Léger-les-Mélèzes - La Bâtie-Neuve
		16908.458984375:  {0, 7},  // new: La Bâtie-Neuve - Saint-Léger-les-Mélèzes
		8478.833984375:   {0, 8},  // new: Moissière
		8306.2373046875:  {0, 9},  // new: Ancelle
		8924.6201171875:  {0, 10}, // new: Ravin de Coste Belle
		8922.3984375:     {0, 11}, // new: Les Borels
		// 1:Sweden ---------------------------------------------------------------
		21768.318359375:   {1, 0},  // new: Hof-Finnskog
		21780.54296875:    {1, 1},  // new: Åsnes
		11371.87109375:    {1, 2},  // new: Spikbrenna
		11270.384765625:   {1, 3},  // new: Lauksjøen
		10706.1689453125:  {1, 4},  // new: Åslia
		10775.3662109375:  {1, 5},  // new: Knapptjernet
		8551.2998046875:   {1, 6},  // new: Vargasen
		8549.8896484375:   {1, 7},  // new: Lövstaholm
		3630.523193359375: {1, 8},  // new: Älgsjön
		3678.771240234375: {1, 9},  // new: Ekshärad
		5182.29833984375:  {1, 10}, // new: Stora Jangen
		5088.5087890625:   {1, 11}, // new: Sunne
		// 2:México ---------------------------------------------------------------
		27065.39453125:   {2, 0},  // new: El Chocolate
		25112.0078125:    {2, 1},  // new: Otates
		13419.46875:      {2, 2},  // new: Ortega
		11845.1259765625: {2, 3},  // new: Las Minas
		13308.2275390625: {2, 4},  // new: Ibarrilla
		7556.85693359375: {2, 5},  // new: Derramadero
		10915.162109375:  {2, 6},  // new: El Brinco
		10996.3623046875: {2, 7},  // new: Guanajuatito
		8367.2353515625:  {2, 8},  // new: Alfaro
		9197.359375:      {2, 9},  // new: Mesa Cuata
		6154.95751953125: {2, 10}, // new: San Diego
		7242.689453125:   {2, 11}, // new: El Mosquito
		// 3:Croatia --------------------------------------------------------------
		25884.58203125:   {3, 0},  // new: Bliznec
		25880.095703125:  {3, 1},  // new: Trakošćan
		13017.4873046875: {3, 2},  // new: Vrbno
		13012.927734375:  {3, 3},  // new: Zagorska Sela
		13264.982421875:  {3, 4},  // new: Kumrovec
		13185.1201171875: {3, 5},  // new: Grdanjci
		10568.0625:       {3, 6},  // new: Stojdraga
		10559.8603515625: {3, 7},  // new: Mali Lipovec
		8101.09228515625: {3, 8},  // new: Hartje
		9022.259765625:   {3, 9},  // new: Kostanjevac
		9099.501953125:   {3, 10}, // new: Krašić
		9101.0771484375:  {3, 11}, // new: Petruš Vrh
		// 4:Portugal -------------------------------------------------------------
		30647.3671875:   {4, 0},  // new: Baião
		31512.115234375: {4, 1},  // new: Caminha
		17035.876953125: {4, 2},  // new: Fridão
		15447.84765625:  {4, 3},  // new: Marão
		15045.11328125:  {4, 4},  // new: Ponte de Lima
		8186.74609375:   {4, 5},  // new: Viana do Castelo
		7591.076171875:  {4, 6},  // new: Ervideiro
		8477.583984375:  {4, 7},  // new: Celeiro
		7806.734375:     {4, 8},  // new: Touca
		7703.224609375:  {4, 9},  // new: Vila Boa
		7798.4951171875: {4, 10}, // new: Carrazedo
		7733.7841796875: {4, 11}, // new: Anjos
		// 5:Sardegna -------------------------------------------------------------
		31854.994140625:  {5, 0},  // new: Rena Majore
		31971.994140625:  {5, 1},  // new: Monte Olia
		13663.78515625:   {5, 2},  // new: Littichedda
		18540.404296875:  {5, 3},  // new: Ala del Sardi
		16802.18359375:   {5, 4},  // new: Mamone
		7913.38134765625: {5, 5},  // new:  Li Pinnenti
		8093.1669921875:  {5, 6},  // new: Malti
		7856.53857421875: {5, 7},  // new: Bassacutena
		9376.2978515625:  {5, 8},  // new: Bortigiadas
		9421.0478515625:  {5, 9},  // new: Sa Mela
		7818.212890625:   {5, 10}, // new: Monte Muvri
		7790.3369140625:  {5, 11}, // new: Monte Acuto
		// 6:Kenya ----------------------------------------------------------------
		10021.7666015625: {6, 0},  // new: Malewa
		9891.7412109375:  {6, 1},  // new: Tarambete
		5753.6005859375:  {6, 2},  // new: Moi North
		5739.994140625:   {6, 3},  // new: Marula
		4848.55517578125: {6, 4},  // fixed: Wileli
		4649.8076171875:  {6, 5},  // new: Kingono
		20541.1796875:    {6, 6},  // new: Soysambu
		20521.3984375:    {6, 7},  // new: Mbaruk
		10031.7802734375: {6, 8},  // new: Sugunoi
		9844.90234375:    {6, 9},  // new: Nakuru
		11013.4697265625: {6, 10}, // new: Kanyawa
		11013.076171875:  {6, 11}, // new: Kanyawa - Nakura
		// 7:Estonia --------------------------------------------------------------
		17430.73828125:   {7, 0},  // new: Otepää
		17420.412109375:  {7, 1},  // new: Rebaste
		8934.5380859375:  {7, 2},  // new: Nüpli
		8952.447265625:   {7, 3},  // new: Truuta
		8832.642578125:   {7, 4},  // new: Koigu
		9093.1376953125:  {7, 5},  // new: Kooraste
		12149.255859375:  {7, 6},  // new: Elva
		11939.3037109375: {7, 7},  // new: Metsalaane
		6549.94677734375: {7, 8},  // new: Vahessaare
		6237.77734375:    {7, 9},  // new: Külaaseme
		5973.14990234375: {7, 10}, // new: Vissi
		6022.7451171875:  {7, 11}, // new: Vellavere
		// 8:Finland --------------------------------------------------------------
		11414.5859375:    {8, 0},  // new: Leustu
		11329.416015625:  {8, 1},  // new: Lahdenkylä
		5151.962890625:   {8, 2},  // new: Saakoski
		5057.02197265625: {8, 3},  // new: Maahi
		6737.29248046875: {8, 4},  // new: Painna
		6467.689453125:   {8, 5},  // new: Peltola
		23354.720703125:  {8, 6},  // new: Päijälä
		23216.017578125:  {8, 7},  // new: Ruokolahti
		10862.580078125:  {8, 8},  // new: Honkanen
		10670.9384765625: {8, 9},  // new: Venkajärvi
		12889.9365234375: {8, 10}, // new: Vehmas
		12827.0439453125: {8, 11}, // new: Hatanpää
		// 9:Greece ---------------------------------------------------------------
		24990.927734375:  {9, 0},  // new: Gravia
		24989.751953125:  {9, 1},  // new: Prosilio
		13848.80859375:   {9, 2},  // new: Mariolata
		13832.6533203125: {9, 3},  // new: Karoutes
		11475.8349609375: {9, 4},  // new: Viniani
		11468.4091796875: {9, 5},  // new: Delphi
		10721.888671875:  {9, 6},  // new: Eptalofos
		10703.537109375:  {9, 7},  // new: Lilea
		5906.15625:       {9, 8},  // new: Parnassós
		5884.07763671875: {9, 9},  // new: Bauxites
		9025.0712890625:  {9, 10}, // new: Drosochori
		9025.2080078125:  {9, 11}, // new: Amfissa
		// 10:Chile ---------------------------------------------------------------
		35043.18359375:   {10, 0},  // new: Bio Bío
		35115.52734375:   {10, 1},  // new: Pulpería
		18300.140625:     {10, 2},  // new: Río Lía
		17057.689453125:  {10, 3},  // new: María Las Cruces
		17205.98828125:   {10, 4},  // new: Las Paraguas
		11114.083984375:  {10, 5},  // new: Rere
		10402.248046875:  {10, 6},  // new: El Poñen
		8197.9619140625:  {10, 7},  // new: Laja
		8075.86572265625: {10, 8},  // new: Yumbel
		8551.7421875:     {10, 9},  // new: Río Claro
		8425.1728515625:  {10, 10}, // new: Hualqui
		8840.3115234375:  {10, 11}, // new: Chivilingo
		// 11:Central Europe ------------------------------------------------------
		32702.908203125:  {11, 0},  // new: Rouské
		32679.244140625:  {11, 1},  // new: Lukoveček
		15779.5947265625: {11, 2},  // new: Raztoka
		15770.38671875:   {11, 3},  // new: Žabárna
		17328.599609375:  {11, 4},  // new: Provodovice
		17310.33203125:   {11, 5},  // new: Chvalčov
		9173.345703125:   {11, 6},  // new: Vítová
		9098.77734375:    {11, 7},  // new: Brusné
		15078.583984375:  {11, 8},  // new: Libosváry
		14987.3271484375: {11, 9},  // new: Rusava
		9267.7421875:     {11, 10}, // new: Osíčko
		8979.5126953125:  {11, 11}, // new: Příkazy
		// 12:Japan ---------------------------------------------------------------
		20209.443359375:  {12, 0},  // new: Lake Mikawa
		20237.0234375:    {12, 1},  // new: Kudarisawa
		11782.9990234375: {12, 2},  // new: Oninotaira
		11723.8271484375: {12, 3},  // new: Okuwacho
		10608.0771484375: {12, 4},  // new: Habu Dam
		10629.9638671875: {12, 5},  // new：Habucho
		13664.8837890625: {12, 6},  // new: Nenoue Plateau
		14124.6884765625: {12, 7},  // new: Tegano
		7321.4169921875:  {12, 8},  // new: Higashino
		7312.6826171875:  {12, 9},  // new: Hokono Lake
		6734.7861328125:  {12, 10}, // new: Nenoue Highlands
		7184.89013671875: {12, 11}, // new: Nakatsugawa
		// 13:Mediterraneo --------------------------------------------------------
		29517.841796875: {13, 0}, // new: Asco
		//1 Ponte
		15444.12109375:   {13, 2},  // new: Monte Cinto
		16482.353515625:  {13, 3},  // new: Albarello
		20774.0390625:    {13, 4},  // new: Canpannace
		7982.541015625:   {13, 5},  // new: Serra Di Cuzzioli
		8828.4140625:     {13, 6},  // new: Maririe
		8782.9814453125:  {13, 7},  // new: Poggiola
		11075.619140625:  {13, 8},  // new: Monte Alloradu
		9752.8134765625:  {13, 9},  // old: Ravin de Finelio
		10414.5029296875: {13, 10}, // new: Cabanella
		11520.50390625:   {13, 11}, // new: Moltifao
		// 14:Pacifico ------------------------------------------------------------
		//0 Talao
		//1 Talanghilirair
		//2 SungaiKunit
		//3 Sangir Balai Janggo
		//4 South Solok
		//5 Kebun Raya Solok
		//6 Batukangkung
		5712.67041015625: {14, 7}, // old: Abai
		6709.298828125:   {14, 8}, // fixed: Moearaikoer
		8058.00634765625: {14, 9}, // new: Bidaralam
		//10 Loeboekmalaka
		9444.4287109375: {14, 11}, // new: Gunung Tujuh
		// 15:Oceania -------------------------------------------------------------
		11336.53125: {15, 0}, // new: Oakleigh
		//1 Doctors Hill
		7023.32177734375: {15, 2}, // none: Mangapai
		6983.908203125:   {15, 3}, // new: Brynderwyn
		//4 Taipuha
		//5 Mareretu
		//6 Waiwera
		//7 Tahekeroa
		//8 Noakes Hill
		9625.2822265625: {15, 9},  // old: Orewa
		8901.7470703125: {15, 10}, // new: Tahekeroa - Orewa
		//11 Makarau
		// 16:Scandia -------------------------------------------------------------
		31230.755859375: {16, 0}, // old: Holtjønn
		32164.1796875:   {16, 1}, // new: Hengeltjønn
		17404.24609375:  {16, 2}, // new: Fyresvatn
		//3 Russvatn
		//4 Tovsli
		//5 Kottjønn
		//6 Fordol
		5756.9423828125: {16, 7}, // new: Fyresdal
		//8 Ljosdalstjønn
		//9 Dagtrolltjønn
		//10 Tovslioytjorn
		//11 Bergsøytjønn
		// 17:Iberia --------------------------------------------------------------
		19315.458984375: {17, 0}, // new: Santes Creus
		//1 Valldossera
		10071.61328125:  {17, 2}, // new: Campdasens
		10075.623046875: {17, 3}, // old: Pontils
		//4 Montagut
		//5 Aiguamúrcia
		//6 Alforja
		//7 Les Irles
		//8 L'Argentera
		//9 Les Voltes
		//10 Montclar
		7663.49072265625: {17, 11}, // new: Botareli
	}
)

func GetStage(sd float64) Stage {
	s, ok := Stages[sd]
	if ok {
		loc := Location{Name: "unknown", Stages: nil}
		if s.Location >= 0 && s.Location < len(Locations) {
			loc = Locations[s.Location]
		}
		stage := "unknown"
		if s.Stage >= 0 && s.Stage < len(loc.Stages) {
			stage = loc.Stages[s.Stage]
		}
		return Stage{
			Location: loc.Name,
			Stage:    stage,
		}
	}
	return Stage{
		Location: "unknown",
		Stage:    "unknown",
	}
}
