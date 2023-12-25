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
			"Stora",
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
			"Arauco",
			"Coronel",
			"Lota",
			"Santa Juana",
			"Los Angeles",
			"San Rosendo",
			"Laja",
			"Yumbel",
			"Florida",
			"Hualqui",
			"Reputo",
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
		32702.908203125:   {0, 0},
		8306.2373046875:   {0, 9},
		21768.318359375:   {1, 0},
		3630.523193359375: {1, 8},
		21780.54296875:    {1, 1},
		11371.871094:      {1, 2},
		27065.39453125:    {2, 0},
		6154.95751953125:  {2, 10},
		7242.689453125:    {2, 11},
		25884.58203125:    {3, 0},
		7798.4951171875:   {3, 10},
		8101.09228515625:  {3, 8},
		9099.501953125:    {3, 10},
		30647.3671875:     {4, 0},
		31854.994140625:   {5, 0},
		7790.3369140625:   {5, 10},
		7818.212890625:    {5, 11},
		17430.73828125:    {7, 0},
		11414.5859375:     {8, 0},
		24990.927734375:   {9, 0},
		5884.07763671875:  {9, 9},
		7982.541015625:    {13, 5},
		7023.32177734375:  {15, 2},
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
