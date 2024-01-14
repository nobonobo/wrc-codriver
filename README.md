# wrc-logger

「EA Sports WRC」向けのペースノートロガー兼ペースノート読み上げツール

## WRCテレメトリについて

このツールはWRCが出力できるテレメトリUDPパケットを受けることで動作します。あらかじめ、UDPパケット出力を有効にしておいてください。

[公式テレメトリ出力ガイド](https://answers.ea.com/t5/Guides-Documentation/EA-SPORTS-WRC-How-to-use-User-Datagram-Protocol-UDP-on-PC/m-p/13178407/thread-id/1?attachment-id=757119)

このツールはデフォルトのリッスンポートは「127.0.0.1:20777」です。コマンドラインオプションで変更できます。

「wrc-logger.exe -listen=127.0.0.1:20778」

また、同じパケット内容を転送する機能もあります。既存で同じパケットを利用している場合は以下のように転送オプションを使ってください。

「wrc-logger.exe -forward="192.168.1.1:20777"」

## ロギングの仕掛けについて

「-logging」オプションをつけるとOSDのペースノートを読み取ってログを記録する機能が有効になります。

- OBSでゲーム画面をフルHD（1920x1080）で取り込み、「仮想カメラ」にて出力
- このツールのロギングを有効にしたら「仮想カメラ」出力のコドライバーOSD表示のマークを画像解析してログを記録します
- ツールのlogフォルダ配下には「ステージ長.log」というファイルが記録されます
- もし既存のファイルがあれば「ステージ長.log.1」「ステージ長.log.2」と番号が付与されて残ります

「wrc-logger.exe -logging -camera 3」というオプションで仮想カメラの選択を変更できます。

## コドライバー読み上げについて

- VoiceVoxというランタイムを利用して音声合成を行います
- 音声生成の基本辞書が「log/base.json」にあります

### 生成のオプション

起動オプションにて以下のパラメータを変更できます。

- 「-actor アクターID」: VoiceVoxの話者変更
- 「-speed 0.5～2.0」: しゃべる速度
- 「-pitch 0.5～2.0」: 声の高さを変更
- 「-volume 0.5～2.0」: 音量変更
- 「-offset -50..50」: 読み上げのタイミング変更（メートル単位）

## ログのフォーマット
ファイル名：「log/ステージ長.log」

```
0,straight,unknown,50
47.01781152725222,3-right,unknown,140
199.68337484359745,slight-left,twisty,60
...
32032.167246975354,finish,unknown,0
```

「距離,単語１,単語２,単語３」の羅列です。
距離順にソートされている必要があります。

- 単語１には基本辞書のmarksキーワードかコンマを含まない日本語が直接書けます
- 単語２には基本辞書のiconsキーワードかコンマを含まない日本語が直接書けます、「unknown」は発音無し
- 単語３には基本辞書のdists距離数かコンマを含まない日本語が直接書けます「0」は発音無し

## アクターIDについて

|キャラクター名|スタイル|ID|
|-------|----|-----|
|四国めたん|ノーマル|2|
||あまあま|0|
||ツンツン|6|
||セクシー|4|
||ささやき|36|
||ヒソヒソ|37|
|ずんだもん|ノーマル|3|
||あまあま|1|
||ツンツン|7|
||セクシー|5|
||ささやき|22|
||ヒソヒソ|38|
|春日部つむぎ|ノーマル|8|
|雨晴はう|ノーマル|10|
|波音リツ|ノーマル|9|
|玄野武宏|ノーマル|11|
||喜び|39|
||ツンギレ|40|
||悲しみ|41|
|白上虎太郎|ふつう|12|
||わーい|32|
||びくびく|33|
||おこ|34|
||びえーん|35|
|青山龍星|ノーマル|13|
|冥鳴ひまり|ノーマル|14|
|九州そら|ノーマル|16|
||あまあま|15|
||ツンツン|18|
||セクシー|17|
||ささやき|19|
|もち子さん|ノーマル|20|
|剣崎雌雄|ノーマル|21|
|WhiteCUL|ノーマル|23|
||たのしい|24|
||かなしい|25|
||びえーん|26|
|後鬼|人間ver.|27|
||ぬいぐるみver.|28|
|No.7|ノーマル|29|
||アナウンス|30|
||読み聞かせ|31|
|ちび式じい|ノーマル|42|
|櫻歌ミコ|ノーマル|43|
||第二形態|44|
||ロリ|45|
|小夜/SAYO|ノーマル|46|
|ナースロボ＿タイプＴ|ノーマル|47|
||楽々|48|
||恐怖|49|
||内緒話|50|

# 既知の問題点

- ロギングモードが環境依存が大きいかも（自分の環境がワイドモニタからフルHD切り出しをしている）
- 画像からペースノートを作ったので実際の読み上げよりも情報が少ない（ショートやロングなどはアイコンには無い）
- 画像処理も確度が９０％ちょいしかなく、結構読み取り間違いや重複が入っちゃってる（特に空が見えにくいステージに弱い）
- コールが忙しい区間で早めに読み始めるという処理が無いので、読み遅れが発生したりする
- ただし以上の問題はログの手修正で直せるので地道に頑張っていこうと思います

# ログの収集状況

- newは最新の分類処理（正確さ95%程度）によるログ。
- oldは精度の悪い分類処理(正確さ70～85%)によるログ。
- fixedは手修正したログ
- noneやコメントの無いものはまだログが無い

- [x] Rallye Monte-Carlo
	- [x] La Bollène-Vésubie - Peïra Cava
	- [x] Peïra Cava - La Bollène-Vésubie
	- [x] La Bollène-Vésubie - Col de Turini
	- [x] Pra d'Alart
	- [x] La Maïris
	- [x] Baisse de Patronel
	- [x] Saint-Léger-les-Mélèzes - La Bâtie-Neuve
	- [x] La Bâtie-Neuve - Saint-Léger-les-Mélèzes
	- [x] Moissière
	- [x] Ancelle
	- [x] Ravin de Coste Belle
	- [x] Les Borels
- [x] Rally Sweden
	- [x] Hof-Finnskog
	- [x] Åsnes
	- [x] Spikbrenna
	- [x] Lauksjøen
	- [x] Åslia
	- [x] Knapptjernet
	- [x] Vargasen
	- [x] Lövstaholm
	- [x] Älgsjön
	- [x] Ekshärad
	- [x] Stora Jangen
	- [x] Sunne
- [x] Guanajuato Rally México
	- [x] El Chocolate
	- [x] Otates
	- [x] Ortega
	- [x] Las Minas
	- [x] Ibarrilla
	- [x] Derramadero
	- [x] El Brinco
	- [x] Guanajuatito
	- [x] Alfaro
	- [x] Mesa Cuata
	- [x] San Diego
	- [x] El Mosquito
- [x] Croatia Rally
	- [x] Bliznec
	- [x] Trakošćan
	- [x] Vrbno
	- [x] Zagorska Sela
	- [x] Kumrovec
	- [x] Grdanjci
	- [x] Stojdraga
	- [x] Mali Lipovec
	- [x] Hartje
	- [x] Kostanjevac
	- [x] Krašić
	- [x] Petruš Vrh
- [x] Vodafone Rally de Portugal
	- [x] Baião
	- [x] Caminha
	- [x] Fridão
	- [x] Marão
	- [x] Ponte de Lima
	- [x] Viana do Castelo
	- [x] Ervideiro
	- [x] Celeiro
	- [x] Touca
	- [x] Vila Boa
	- [x] Carrazedo
	- [x] Anjos
- [x] Rally Italia Sardegna
	- [x] Rena Majore
	- [x] Monte Olia
	- [x] Littichedda
	- [x] Ala del Sardi
	- [x] Mamone
	- [x] Li Pinnenti
	- [x] Malti
	- [x] Bassacutena
	- [x] Bortigiadas
	- [x] Sa Mela
	- [x] Monte Muvri
	- [x] Monte Acuto
- [x] Safari Rally Kenya
  - [x] Malewa
  - [x] Tarambete
  - [x] Moi North
  - [x] Marula
  - [x] Wileli
  - [x] Kingono
  - [x] Soysambu
  - [x] Mbaruk
  - [x] Sugunoi
  - [x] Nakuru
  - [x] Kanyawa
  - [x] Kanyawa - Nakura
- [x] Rally Estonia
  - [x] Otepää
  - [x] Rebaste
  - [x] Nüpli
  - [x] Truuta
  - [x] Koigu
  - [x] Kooraste
  - [x] Elva
  - [x] Metsalaane
  - [x] Vahessaare
  - [x] Külaaseme
  - [x] Vissi
  - [x] Vellavere
- [x] SECTO Rally Finland
  - [x] Leustu
  - [x] Lahdenkyla
  - [x] Saakoski
  - [x] Maahi
  - [x] Painna
  - [x] Peltola
  - [x] Paijala
  - [x] Ruokolahti
  - [x] Honkanen
  - [x] Venkajarvi
  - [x] Vehmas
  - [x] Hatanpaa
- [ ] EKO ACROPOLIS Rally Greece
  - [ ] Gravia
  - [x] Prosilio
  - [ ] Mariolata
  - [ ] Karoutes
  - [ ] Viniani
  - [ ] Delphi
  - [ ] Eptalofos
  - [x] Lilea
  - [x] Parnassós
  - [x] Bauxites
  - [x] Drosochori
  - [x] Amfissa
- [ ] BIO BÍO Rally Chile
  - [ ] Bio Bío
  - [ ] Pulpería
  - [x] Río Lía
  - [ ] María Las Cruces
  - [ ] Las Paraguas
  - [ ] Rere
  - [ ] El Poñen
  - [ ] Laja
  - [x] Yumbel
  - [x] Río Claro
  - [x] Hualqui
  - [x] Chivilingo
- [ ] Central Europe Rally
  - [ ] Rouské
  - [x] Lukoveček
  - [ ] Raztoka
  - [x] Žabárna
  - [x] Provodovice
  - [ ] Chvalčov
  - [x] Vítová
  - [x] Brusné
  - [x] Libosváry
  - [x] Rusava
  - [x] Osíčko
  - [x] Příkazy
- [ ] Forum8 Rally Japan
  - [x] Lake Mikawa
  - [x] Kudarisawa
  - [ ] Oninotaira
  - [ ] Okuwacho
  - [ ] Habu Dam
  - [x] Habucho
  - [x] Nenoue Plateau
  - [ ] Tegano
  - [ ] Higashino
  - [ ] Hokono Lake
  - [ ] Nenoue Highlands
  - [x] Nakatsugawa
- [ ] Rally Mediterraneo
  - [x] Asco
  - [ ] Ponte
  - [ ] Monte Cinto
  - [ ] Albarello
  - [x] Capannace
  - [x] Serra Di Cuzzioli
  - [ ] Maririe
  - [ ] Poggiola
  - [ ] Monte Alloradu
  - [x] Ravin de Finelio
  - [ ] Cabanella
  - [ ] Moltifao
- [ ] Agon By AOC Rally Pacifico
  - [ ] Talao
  - [ ] Talanghilirair
  - [ ] SungaiKunit
  - [ ] Sangir Balai Janggo
  - [ ] South Solok
  - [ ] Kebun Raya Solok
  - [ ] Batukangkung
  - [x] Abai
  - [x] Moearaikoer
  - [x] Bidaralam
  - [ ] Loeboekmalaka
  - [x] Gunung Tujuh
- [ ] Fanatec Rally Oceania
  - [x] Oakleigh
  - [ ] Doctors Hill
  - [ ] Mangapai
  - [x] Brynderwyn
  - [ ] Taipuha
  - [ ] Mareretu
  - [ ] Waiwera
  - [ ] Tahekeroa
  - [ ] Noakes Hill
  - [x] Orewa
  - [x] Tahekeroa - Orewa
  - [ ] Makarau
- [ ] Rally Scandia
  - [x] Holtjønn
  - [x] Hengeltjønn
  - [x] Fyresvatn
  - [ ] Russvatn
  - [ ] Tovsli
  - [ ] Kottjønn
  - [ ] Fordol
  - [x] Fyresdal
  - [ ] Ljosdalstjønn
  - [ ] Dagtrolltjønn
  - [ ] Tovslioytjorn
  - [ ] Bergsøytjønn
- [ ] Rally Iberia
  - [x] Santes Creus
  - [ ] Valldossera
  - [x] Campdasens
  - [x] Pontils
  - [ ] Montagut
  - [ ] Aiguamúrcia
  - [ ] Alforja
  - [ ] Les Irles
  - [ ] L'Argentera
  - [ ] Les Voltes
  - [ ] Montclar
  - [x] Botareli
