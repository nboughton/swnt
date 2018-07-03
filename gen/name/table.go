package name

import (
	"strings"

	"github.com/nboughton/rollt"
	"github.com/nboughton/swnt/gen/culture"
)

// Names represents rollable tables for Name generation
var Names = Tables{
	Table{
		Culture: culture.Arabic,
		Male:    rollt.List{Items: []string{"Aamir", "Ayub", "Binyamin", "Efraim", "Ibrahim", "Ilyas", "Ismail", "Jibril", "Jumanah", "Kazi", "Lut", "Matta", "Mohammed", "Mubarak", "Mustafa", "Nazir", "Rahim", "Reza", "Sharif", "Taimur", "Usman", "Yakub", "Yusuf", "Zakariya", "Zubair"}},
		Female:  rollt.List{Items: []string{"Aisha", "Alimah", "Badia", "Bisharah", "Chanda", "Daliya", "Fatimah", "Ghania", "Halah", "Kaylah", "Khayrah", "Layla", "Mina", "Munisa", "Mysha", "Naimah", "Nissa", "Nura", "Parveen", "Rana", "Shalha", "Suhira", "Tahirah", "Yasmin", "Zulehka"}},
		Surname: rollt.List{Items: []string{"Abdel", "Awad", "Dahhak", "Essa", "Hanna", "Harbi", "Hassan", "Isa", "Kasim", "Katib", "Khalil", "Malik", "Mansoor", "Mazin", "Musa", "Najeeb", "Namari", "Naser", "Rahman", "Rasheed", "Saleh", "Salim", "Shadi", "Sulaiman", "Tabari"}},
		Place:   rollt.List{Items: []string{"Adan", "Magrit", "Ahsa", "Masqat", "Andalus", "Misr", "Asmara", "Muruni", "Asqlan", "Qabis", "Baqubah", "Qina", "Basit", "Rabat", "Baysan", "Ramlah", "Baytlahm", "Riyadh", "Bursaid", "Sabtah", "Dahilah", "Salalah", "Darasalam", "Sana", "Dawhah", "Sinqit", "Ganin", "Suqutrah", "Gebal", "Sur", "Gibuti", "Tabuk", "Giddah", "Tangah", "Harmah", "Tarifah", "Hartum", "Tarrakunah", "Hibah", "Tisit", "Hims", "Uman", "Hubar", "Urdunn", "Karbala", "Wasqah", "Kut", "Yaburah", "Lacant", "Yaman"}},
	},
	Table{
		Culture: culture.Chinese,
		Male:    rollt.List{Items: []string{"Aiguo", "Bohai", "Chao", "Dai", "Dawei", "Duyi", "Fa", "Fu", "Gui", "Hong", "Jianyu", "Kang", "Li", "Niu", "Peng", "Quan", "Ru", "Shen", "Shi", "Song", "Tao", "Xue", "Yi", "Yuan", "Zian"}},
		Female:  rollt.List{Items: []string{"Biyu", "Changying", "Daiyu", "Huidai", "Huiliang", "Jia", "Jingfei", "Lan", "Liling", "Liu", "Meili", "Niu", "Peizhi", "Qiao", "Qing", "Ruolan", "Shu", "Suyin", "Ting", "Xia", "Xiaowen", "Xiulan", "Ya", "Ying", "Zhilan"}},
		Surname: rollt.List{Items: []string{"Bai", "Cao", "Chen", "Cui", "Ding", "Du", "Fang", "Fu", "Guo", "Han", "Hao", "Huang", "Lei", "Li", "Liang", "Liu", "Long", "Song", "Tan", "Tang", "Wang", "Wu", "Xing", "Yang", "Zhang"}},
		Place:   rollt.List{Items: []string{"Andong", "Luzhou", "Anqing", "Ningxia", "Anshan", "Pingxiang", "Chaoyang", "Pizhou", "Chaozhou", "Qidong", "Chifeng", "Qingdao", "Dalian", "Qinghai", "Dunhuang", "Rehe", "Fengjia", "Shanxi", "Fengtian", "Taiyuan", "Fuliang", "Tengzhou", "Fushun", "Urumqi", "Gansu", "Weifang", "Ganzhou", "Wugang", "Guizhou", "Wuxi", "Hotan", "Xiamen", "Hunan", "Xian", "Jinan", "Xikang", "Jingdezhen", "Xining", "Jinxi", "Xinjiang", "Jinzhou", "Yidu", "Kunming", "Yingkou", "Liaoning", "Yuxi", "Linyi", "Zigong", "Lushun", "Zoige"}},
	},
	Table{
		Culture: culture.English,
		Male:    rollt.List{Items: []string{"Adam", "Albert", "Alfred", "Allan", "Archibald", "Arthur", "Basil", "Charles", "Colin", "Donald", "Douglas", "Edgar", "Edmund", "Edward", "George", "Harold", "Henry", "Ian", "James", "John", "Lewis", "Oliver", "Philip", "Richard", "William"}},
		Female:  rollt.List{Items: []string{"Abigail", "Anne", "Beatrice", "Blanche", "Catherine", "Charlotte", "Claire", "Eleanor", "Elizabeth", "Emily", "Emma", "Georgia", "Harriet", "Joan", "Judy", "Julia", "Lucy", "Lydia", "Margaret", "Mary", "Molly", "Nora", "Rosie", "Sarah", "Victoria"}},
		Surname: rollt.List{Items: []string{"Barker", "Brown", "Butler", "Carter", "Chapman", "Collins", "Cook", "Davies", "Gray", "Green", "Harris", "Jackson", "Jones", "Lloyd", "Miller", "Roberts", "Smith", "Taylor", "Thomas", "Turner", "Watson", "White", "Williams", "Wood", "Young"}},
		Place:   rollt.List{Items: []string{"Aldington", "Kedington", "Appleton", "Latchford", "Ashdon", "Leigh", "Berwick", "Leighton", "Bramford", "Maresfield", "Brimstage", "Markshall", "Carden", "Netherpool", "Churchill", "Newton", "Clifton", "Oxton", "Colby", "Preston", "Copford", "Ridley", "Cromer", "Rochford", "Davenham", "Seaford", "Dersingham", "Selsey", "Doverdale", "Stanton", "Elsted", "Stockham", "Ferring", "Stoke", "Gissing", "Sutton", "Heydon", "Thakeham", "Holt", "Thetford", "Hunston", "Thorndon", "Hutton", "Ulting", "Inkberrow", "Upton", "Inworth", "Westhorpe", "Isfield", "Worcester"}},
	},
	Table{
		Culture: culture.Greek,
		Male:    rollt.List{Items: []string{"Alexander", "Alexius", "Anastasius", "Christodoulos", "Christos", "Damian", "Dimitris", "Dysmas", "Elias", "Giorgos", "Ioannis", "Konstantinos", "Lambros", "Leonidas", "Marcos", "Miltiades", "Nestor", "Nikos", "Orestes", "Petros", "Simon", "Stavros", "Theodore", "Vassilios", "Yannis"}},
		Female:  rollt.List{Items: []string{"Alexandra", "Amalia", "Callisto", "Charis", "Chloe", "Dorothea", "Elena", "Eudoxia", "Giada", "Helena", "Ioanna", "Lydia", "Melania", "Melissa", "Nika", "Nikolina", "Olympias", "Philippa", "Phoebe", "Sophia", "Theodora", "Valentina", "Valeria", "Yianna", "Zoe"}},
		Surname: rollt.List{Items: []string{"Andreas", "Argyros", "Dimitriou", "Floros", "Gavras", "Ioannidis", "Katsaros", "Kyrkos", "Leventis", "Makris", "Metaxas", "Nikolaidis", "Pallis", "Pappas", "Petrou", "Raptis", "Simonides", "Spiros", "Stavros", "Stephanidis", "Stratigos", "Terzis", "Theodorou", "Vasiliadis", "Yannakakis"}},
		Place:   rollt.List{Items: []string{"Adramyttion", "Kallisto", "Ainos", "Katerini", "Alikarnassos", "Kithairon", "Avydos", "Kydonia", "Dakia", "Lakonia", "Dardanos", "Leros", "Dekapoli", "Lesvos", "Dodoni", "Limnos", "Efesos", "Lykia", "Efstratios", "Megara", "Elefsina", "Messene", "Ellada", "Milos", "Epidavros", "Nikaia", "Erymanthos", "Orontis", "Evripos", "Parnasos", "Gavdos", "Petro", "Gytheio", "Samos", "Ikaria", "Syros", "Ilios", "Thapsos", "Illyria", "Thessalia", "Iraia", "Thira", "Irakleio", "Thiva", "Isminos", "Varvara", "Ithaki", "Voiotia", "Kadmeia", "Vyvlos"}},
	},
	Table{
		Culture: culture.Indian,
		Male:    rollt.List{Items: []string{"Amrit", "Ashok", "Chand", "Dinesh", "Gobind", "Harinder", "Jagdish", "Johar", "Kurien", "Lakshman", "Madhav", "Mahinder", "Mohal", "Narinder", "Nikhil", "Omrao", "Prasad", "Pratap", "Ranjit", "Sanjay", "Shankar", "Thakur", "Vijay", "Vipul", "Yash"}},
		Female:  rollt.List{Items: []string{"Amala", "Asha", "Chandra", "Devika", "Esha", "Gita", "Indira", "Indrani", "Jaya", "Jayanti", "Kiri", "Lalita", "Malati", "Mira", "Mohana", "Neela", "Nita", "Rajani", "Sarala", "Sarika", "Sheela", "Sunita", "Trishna", "Usha", "Vasanta"}},
		Surname: rollt.List{Items: []string{"Achari", "Banerjee", "Bhatnagar", "Bose", "Chauhan", "Chopra", "Das", "Dutta", "Gupta", "Johar", "Kapoor", "Mahajan", "Malhotra", "Mehra", "Nehru", "Patil", "Rao", "Saxena", "Shah", "Sharma", "Singh", "Trivedi", "Venkatesan", "Verma", "Yadav"}},
		Place:   rollt.List{Items: []string{"Ahmedabad", "Jaisalmer", "Alipurduar", "Jharonda", "Alubari", "Kadambur", "Anjanadri", "Kalasipalyam", "Ankleshwar", "Karnataka", "Balarika", "Kutchuhery", "Bhanuja", "Lalgola", "Bhilwada", "Mainaguri", "Brahmaghosa", "Nainital", "Bulandshahar", "Nandidurg", "Candrama", "Narayanadri", "Chalisgaon", "Panipat", "Chandragiri", "Panjagutta", "Charbagh", "Pathankot", "Chayanka", "Pathardih", "Chittorgarh", "Porbandar", "Dayabasti", "Rajasthan", "Dikpala", "Renigunta", "Ekanga", "Sewagram", "Gandhidham", "Shakurbasti", "Gollaprolu", "Siliguri", "Grahisa", "Sonepat", "Guwahati", "Teliwara", "Haridasva", "Tinpahar", "Indraprastha", "Villivakkam"}},
	},
	Table{
		Culture: culture.Japanese,
		Male:    rollt.List{Items: []string{"Akira", "Daisuke", "Fukashi", "Goro", "Hiro", "Hiroya", "Hotaka", "Katsu", "Katsuto", "Keishuu", "Kyuuto", "Mikiya", "Mitsunobu", "Mitsuru", "Naruhiko", "Nobu", "Shigeo", "Shigeto", "Shou", "Shuji", "Takaharu", "Teruaki", "Tetsushi", "Tsukasa", "Yasuharu"}},
		Female:  rollt.List{Items: []string{"Aemi", "Airi", "Ako", "Ayu", "Chikaze", "Eriko", "Hina", "Kaori", "Keiko", "Kyouka", "Mayumi", "Miho", "Namiko", "Natsu", "Nobuko", "Rei", "Ririsa", "Sakimi", "Shihoko", "Shika", "Tsukiko", "Tsuzune", "Yoriko", "Yorimi", "Yoshiko"}},
		Surname: rollt.List{Items: []string{"Abe", "Arakaki", "Endo", "Fujiwara", "Goto", "Ito", "Kikuchi", "Kinjo", "Kobayashi", "Koga", "Komatsu", "Maeda", "Nakamura", "Narita", "Ochi", "Oshiro", "Saito", "Sakamoto", "Sato", "Suzuki", "Takahashi", "Tanaka", "Watanabe", "Yamamoto", "Yamasaki"}},
		Place:   rollt.List{Items: []string{"Bando", "Mitsukaido", "Chikuma", "Moriya", "Chikusei", "Nagano", "Chino", "Naka", "Hitachi", "Nakano", "Hitachinaka", "Ogi", "Hitachiomiya", "Okaya", "Hitachiota", "Omachi", "Iida", "Ryugasaki", "Iiyama", "Saku", "Ina", "Settsu", "Inashiki", "Shimotsuma", "Ishioka", "Shiojiri", "Itako", "Suwa", "Kamisu", "Suzaka", "Kasama", "Takahagi", "Kashima", "Takeo", "Kasumigaura", "Tomi", "Kitaibaraki", "Toride", "Kiyose", "Tsuchiura", "Koga", "Tsukuba", "Komagane", "Ueda", "Komoro", "Ushiku", "Matsumoto", "Yoshikawa", "Mito", "Yuki"}},
	},
	Table{
		Culture: culture.Latin,
		Male:    rollt.List{Items: []string{"Agrippa", "Appius", "Aulus", "Caeso", "Decimus", "Faustus", "Gaius", "Gnaeus", "Hostus", "Lucius", "Mamercus", "Manius", "Marcus", "Mettius", "Nonus", "Numerius", "Opiter", "Paulus", "Proculus", "Publius", "Quintus", "Servius", "Tiberius", "Titus", "Volescus"}},
		Female:  rollt.List{Items: []string{"Appia", "Aula", "Caesula", "Decima", "Fausta", "Gaia", "Gnaea", "Hosta", "Lucia", "Maio", "Marcia", "Maxima", "Mettia", "Nona", "Numeria", "Octavia", "Postuma", "Prima", "Procula", "Septima", "Servia", "Tertia", "Tiberia", "Titia", "Vibia"}},
		Surname: rollt.List{Items: []string{"Antius", "Aurius", "Barbatius", "Calidius", "Cornelius", "Decius", "Fabius", "Flavius", "Galerius", "Horatius", "Julius", "Juventius", "Licinius", "Marius", "Minicius", "Nerius", "Octavius", "Pompeius", "Quinctius", "Rutilius", "Sextius", "Titius", "Ulpius", "Valerius", "Vitellius"}},
		Place:   rollt.List{Items: []string{"Abilia", "Lucus", "Alsium", "Lugdunum", "Aquileia", "Mediolanum", "Argentoratum", "Novaesium", "Ascrivium", "Patavium", "Asculum", "Pistoria", "Attalia", "Pompeii", "Barium", "Raurica", "Batavorum", "Rigomagus", "Belum", "Roma", "Bobbium", "Salernum", "Brigantium", "Salona", "Burgodunum", "Segovia", "Camulodunum", "Sirmium", "Clausentum", "Spalatum", "Corduba", "Tarraco", "Coriovallum", "Treverorum", "Durucobrivis", "Verulamium", "Eboracum", "Vesontio", "Emona", "Vetera", "Florentia", "Vindelicorum", "Lactodurum", "Vindobona", "Lentia", "Vinovia", "Lindum", "Viroconium", "Londinium", "Volubilis"}},
	},
	Table{
		Culture: culture.Nigerian,
		Male:    rollt.List{Items: []string{"Adesegun", "Akintola", "Amabere", "Arikawe", "Asagwara", "Chidubem", "Chinedu", "Chiwetei", "Damilola", "Esangbedo", "Ezenwoye", "Folarin", "Genechi", "Idowu", "Kelechi", "Ketanndu", "Melubari", "Nkanta", "Obafemi", "Olatunde", "Olumide", "Tombari", "Udofia", "Uyoata", "Uzochi"}},
		Female:  rollt.List{Items: []string{"Abike", "Adesuwa", "Adunola", "Anguli", "Arewa", "Asari", "Bisola", "Chioma", "Eduwa", "Emilohi", "Fehintola", "Folasade", "Mahparah", "Minika", "Nkolika", "Nkoyo", "Nuanae", "Obioma", "Olafemi", "Shanumi", "Sominabo", "Suliat", "Tariere", "Temedire", "Yemisi"}},
		Surname: rollt.List{Items: []string{"Adegboye", "Adeniyi", "Adeyeku", "Adunola", "Agbaje", "Akpan", "Akpehi", "Aliki", "Asuni", "Babangida", "Ekim", "Ezeiruaku", "Fabiola", "Fasola", "Nwokolo", "Nzeocha", "Ojo", "Okonkwo", "Okoye", "Olaniyan", "Olawale", "Olumese", "Onajobi", "Soyinka", "Yamusa"}},
		Place:   rollt.List{Items: []string{"Abadan", "Jere", "Ador", "Kalabalge", "Agatu", "Katsina", "Akamkpa", "Knoduga", "Akpabuyo", "Konshishatse", "Ala", "Kukawa", "Askira", "Kwande", "Bakassi", "Kwayakusar", "Bama", "Logo", "Bayo", "Mafa", "Bekwara", "Makurdi", "Biase", "Nganzai", "Boki", "Obanliku", "Buruku", "Obi", "Calabar", "Obubra", "Chibok", "Obudu", "Damboa", "Odukpani", "Dikwa", "Ogbadibo", "Etung", "Ohimini", "Gboko", "Okpokwu", "Gubio", "Otukpo", "Guzamala", "Shani", "Gwoza", "Ugep", "Hawul", "Vandeikya", "Ikom", "Yala"}},
	},
	Table{
		Culture: culture.Russian,
		Male:    rollt.List{Items: []string{"Aleksandr", "Andrei", "Arkady", "Boris", "Dmitri", "Dominik", "Grigory", "Igor", "Ilya", "Ivan", "Kiril", "Konstantin", "Leonid", "Nikolai", "Oleg", "Pavel", "Petr", "Sergei", "Stepan", "Valentin", "Vasily", "Viktor", "Yakov", "Yegor", "Yuri"}},
		Female:  rollt.List{Items: []string{"Aleksandra", "Anastasia", "Anja", "Catarina", "Devora", "Dima", "Ekaterina", "Eva", "Irina", "Karolina", "Katlina", "Kira", "Ludmilla", "Mara", "Nadezdha", "Nastassia", "Natalya", "Oksana", "Olena", "Olga", "Sofia", "Svetlana", "Tatyana", "Vilma", "Yelena"}},
		Surname: rollt.List{Items: []string{"Abelev", "Bobrikov", "Chemerkin", "Gogunov", "Gurov", "Iltchenko", "Kavelin", "Komarov", "Korovin", "Kurnikov", "Lebedev", "Litvak", "Mekhdiev", "Muraviov", "Nikitin", "Ortov", "Peshkov", "Romasko", "Shvedov", "Sikorski", "Stolypin", "Turov", "Volokh", "Zaitsev", "Zhukov"}},
		Place:   rollt.List{Items: []string{"Amur", "Omsk", "Arkhangelsk", "Orenburg", "Astrakhan", "Oryol", "Belgorod", "Penza", "Bryansk", "Perm", "Chelyabinsk", "Pskov", "Chita", "Rostov", "Gorki", "Ryazan", "Irkutsk", "Sakhalin", "Ivanovo", "Samara", "Kaliningrad", "Saratov", "Kaluga", "Smolensk", "Kamchatka", "Sverdlovsk", "Kemerovo", "Tambov", "Kirov", "Tomsk", "Kostroma", "Tula", "Kurgan", "Tver", "Kursk", "Tyumen", "Leningrad", "Ulyanovsk", "Lipetsk", "Vladimir", "Magadan", "Volgograd", "Moscow", "Vologda", "Murmansk", "Voronezh", "Novgorod", "Vyborg", "Novosibirsk", "Yaroslavl"}},
	},
	Table{
		Culture: culture.Spanish,
		Male:    rollt.List{Items: []string{"Alejandro", "Alonso", "Amelio", "Armando", "Bernardo", "Carlos", "Cesar", "Diego", "Emilio", "Estevan", "Felipe", "Francisco", "Guillermo", "Javier", "Jose", "Juan", "Julio", "Luis", "Pedro", "Raul", "Ricardo", "Salvador", "Santiago", "Valeriano", "Vicente"}},
		Female:  rollt.List{Items: []string{"Adalina", "Aleta", "Ana", "Ascencion", "Beatriz", "Carmela", "Celia", "Dolores", "Elena", "Emelina", "Felipa", "Inez", "Isabel", "Jacinta", "Lucia", "Lupe", "Maria", "Marta", "Nina", "Paloma", "Rafaela", "Soledad", "Teresa", "Valencia", "Zenaida"}},
		Surname: rollt.List{Items: []string{"Arellano", "Arispana", "Borrego", "Carderas", "Carranzo", "Cordova", "Enciso", "Espejo", "Gavilan", "Guerra", "Guillen", "Huertas", "Illan", "Jurado", "Moretta", "Motolinia", "Pancorbo", "Paredes", "Quesada", "Roma", "Rubiera", "Santoro", "Torrillas", "Vera", "Vivero"}},
		Place:   rollt.List{Items: []string{"Aguascebas", "Loreto", "Alcazar", "Lujar", "Barranquete", "Marbela", "Bravatas", "Matagorda", "Cabezudos", "Nacimiento", "Calderon", "Niguelas", "Cantera", "Ogijares", "Castillo", "Ortegicar", "Delgadas", "Pampanico", "Donablanca", "Pelado", "Encinetas", "Quesada", "Estrella", "Quintera", "Faustino", "Riguelo", "Fuentebravia", "Ruescas", "Gafarillos", "Salteras", "Gironda", "Santopitar", "Higueros", "Taberno", "Huelago", "Torres", "Humilladero", "Umbrete", "Illora", "Valdecazorla", "Isabela", "Velez", "Izbor", "Vistahermosa", "Jandilla", "Yeguas", "Jinetes", "Zahora", "Limones", "Zumeta"}},
	},
}

// System is a list of possible star system names taken from a wikipedia page on fictional planet names.
var System = rollt.List{Items: []string{"Abyormen", "Acheron", "Aegus", "Aether", "Ahnooie-4", "Aiur", "Aka", "Alcarinque", "Aldeian", "Alderaan", "Alkarinque", "Alpha", "Altair IV", "Alwas", "Amanga", "Amazo", "Ambar", "Anarres", "Anubis", "Aquarius", "Aquas", "Arcadia", "Arda", "Arisia", "Ark", "Arlia", "Armaghast", "Arrakis", "Astra", "Athos", "Athshe", "Atlantis", "Aurelia", "Auron", "Axturias", "Azeroth", "Baab", "Bajor", "Balaho", "Baloris", "Baltan", "Barathrum", "Barrayar", "Barsoom", "Bas-Lag", "Bazoik", "Beezee", "Belzagor", "Beowulf", "Berlin", "Bespin", "Beulah", "Bismoll", "Black Star", "Blue Sands", "Bog", "Bop", "Boskone", "Braal", "Brontitall", "Bryyo", "Cadwal", "Caladan", "Calafia", "Camazotz", "Caprica", "Carnil", "Carnivalia", "Centauri", "Cetaganda", "Chel", "Chthon", "Churchill", "Claire", "Clin", "Corneria", "Coruscant", "Covenant", "Crematoria", "Crete", "Cybertron", "Cygnus Alpha", "Cyteen", "Dada", "Dagobah", "Darkover", "Dar Sai", "Daxam", "Dyan", "Deemi", "Demeter", "Denzi", "Dezoris", "Dhrawn", "Diso", "Doisac", "Dorsai", "Dosadi", "Dragon's Egg", "Dragon", "Dres", "Druidia", "Dryad", "Dump", "Duna", "Durdane", "E", "Ea", "Earendil", "Eayn", "Echronedal", "Eeloo", "Elemmire", "Ellicoore 2", "Elysia", "Emerald", "Empyrrean", "Endor", "Epsilon 3", "Erna", "Eve", "Expel", "Exxilon", "Eylor", "Famille", "Fanbelt", "Far Away", "Fargett", "Fhloston", "Fichina", "Fiorina", "Flash", "Fortuna", "Freeza 79", "Fribbulus Xax", "Friedland", "Frystaat", "Galaxian 3", "Gallifrey", "Gamilon", "Gamilus", "Garissa", "Garrota", "Garth", "Gauda Prime", "Gaul", "Gelidus", "Genesis", "Gethen", "Giedi Prime", "Giganda", "Girath", "Gloob", "Gnarlach", "Gnosticus IV", "Gobotron", "God's Grove", "Golgota", "Gor", "Gorgona", "Gork", "Grayson", "Groth", "Gurun", "Gyodai", "Hades", "Hain", "Halvmork", "Harvest", "Hazard", "Heath", "Hebron", "Hegira", "Hekla", "Helicon", "Helion Prime", "Helliconia", "Hiigara", "Hikari", "Home", "Hope", "Horizon", "Hoth", "Houston", "Htrae", "Hummard", "Hyaita 4", "Hydros", "Hydross", "Hyperion", "Iga", "Ilu", "Imbar", "Incandescent", "Interchange", "Ireta", "Irk", "Ivy", "Ix", "Ixchel", "Janjur Qom", "Jijo", "Jinx", "Jobis", "Jool", "Jophekka", "Jurai", "Kaitain", "Kalimdor", "Kamino", "Kanassa", "Karn", "Kashyyyk", "Katina", "Kelewan", "Kerbin", "Kharak", "King", "Kithrup", "Klaus", "Klendathu", "Kobaia", "Kobol", "Komarr", "Koozebane", "Korath III", "Kosmos", "Krelar", "Krikkit", "Krishna", "Kronos", "Krypton", "Kukulkan", "Lagash", "La Maetelle", "Lamarckia", "Lave", "Laythe", "Leeds", "Leesti", "Legis XV", "Leonida", "Leslie", "Londinium", "London", "Luinil", "Lumbar", "Lumen", "Lumiere", "Lusch", "Lusitania", "LV-426", "MacBeth", "Maetel", "Magma", "Magrathea", "Majipoor", "Manhattan", "Mare Infinitus", "Marshmalia", "Marune", "Maske", "Maui-Covenant", "Medea", "Meiji", "Mejerr", "Melmac", "Mer", "Meridian", "Merle", "Mesklin", "Metaluna", "Midkemia", "Milokeenia", "Minbar", "Miranda", "Mogo", "Moho", "Mok", "Mondas", "Monea", "Mongo", "Mons", "Mor-Tax", "Morthrai", "Motavia", "Mote Prime", "Mount", "Mustafar", "Naboo", "Nackle", "Nacre", "Namek", "Narn", "Navi", "Nebula 71", "Nebula L-77", "Nebula Z95", "Nede", "Nemesis", "Nenar", "New Amazonia", "New Chicago", "New Earth", "New Namek", "New Terra", "New Vegeta", "Nihil", "Nirn", "Nopalgarth", "Norfolk", "Norion", "Nova Kong", "Nuliajuk", "Nyvan", "Oa", "Oban", "Omicron", "Omnivarium", "Onyx", "Optera", "Ork", "Ormazd", "Orthe", "Osiris", "Pacem", "Palain IX", "Palamok", "Palma", "Palshife", "Pandarve", "Pandora", "Pant", "Parvati", "Peladon", "Pern", "Petaybee", "Phaaze", "Pittsburgh", "Arb", "Plateau", "Plootarg", "Pluto", "Pol", "Pyrrus", "Q-13", "Qar'To", "Qom-Riyadh", "Qo'noS", "Quintessa", "Rain", "Rainbow", "Raxicor", "Reach", "Red Star", "Regis III", "Remulak", "Remus", "Rentus", "Requiem", "Resurgam", "Reverie", "Riedquat", "Rigel", "Rigel 7", "Rime", "River", "Roak", "Roche", "Romulus", "Rougpelt", "Rubanis", "Ruzhena", "Rykros", "Rylos", "Salusa Secundus", "Sanctuary", "Sanghelios", "Sangre", "Saraksh", "Sarris", "Saula", "Sauria", "Sauron", "Second Miltia", "Sedon", "Sedra", "Sergyar", "Serpo", "Sesharrim", "Shadow", "Shaggai", "Shikasta", "Shora", "Sigma Octanus IV", "Signo", "Sihnon", "Skaro", "Sky's Edge", "Solaria", "Solaris", "Solbrecht", "Sol Draconi", "Soror", "Sparta", "Spherus Magna", "Spider", "Spira", "SR-388", "Star One", "St. Ekaterina", "Stroggos", "Styx", "Synnax", "Tagora", "Talark", "Tallon IV", "Tamaran", "Tanis", "Tanith", "Tatooine", "Taurus", "Te", "Tek", "Tellus Secundus", "Tellus Tertius", "Temblor", "Tenebra", "Tergiverse IV", "Terminal", "Terminus", "Thalassa", "Thalassean", "Thallon", "Thel", "Thermia", "The Smoke Ring", "Thra", "Tiamat", "T'ien Shan", "Timbl", "Tirol", "Tissa", "Titan", "Titania", "Tleilax", "Tokyo", "Torto", "Traal", "Tralfamadore", "Trantor", "Trenco", "Tribute", "Trullion", "Tschai", "T'vao", "Twinsun", "Tylo", "U40", "Ummo", "Undomiel", "Unicron", "Uriel", "Urth", "Urtraghus", "Vall", "Vanguard 3", "Vega", "Vegandon", "Vegeta", "Velantia", "Velux", "Venom", "Vindine", "Vladislava", "Vorticon VI", "Vortis", "Vulcan", "Vusstra", "Wallach IX", "Waterloo", "Wegthor", "Wormwood", "Wyst", "X", "X-13", "Xenex", "Xenon", "Xindus", "Yaila", "Yavin 4", "Yellowstone", "Yugopotamia", "Zahir", "Zark", "Zarkon", "Zartron-9", "Zebes", "Zedelbrock 473", "Zedelbrock 475", "Zeelich", "Z'ha'dum", "Zog", "Zok", "Zokk", "Zoness", "Zorg", "Zutinma", "Zyrgon"}}

// Name generator. This is a primitive first effort that will be refined over time
var vl = rollt.Table{
	Name: "vowels",
	Dice: "1d4,1d3",
	Items: []rollt.Item{
		{Match: []int{2}, Text: "y"},
		{Match: []int{3}, Text: "a"},
		{Match: []int{4}, Text: "e"},
		{Match: []int{5}, Text: "i"},
		{Match: []int{6}, Text: "o"},
		{Match: []int{7}, Text: "u"},
	},
}

var con = rollt.List{
	Name: "consonants",
	Items: []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "z",
		"ch", "gh", "kh", "ph", "rh", "sh", "th", "qu", "ck"},
}

// Generate creates a random name by combining alternating vowels and consonants
func Generate(ln int) string {
	name := ""
	for i := 0; i <= ln; i++ {
		if i%2 != 0 {
			name += con.Roll()
		} else {
			name += vl.Roll()
		}
	}

	return strings.ToUpper(string(name[0])) + string(name[1:])
}
