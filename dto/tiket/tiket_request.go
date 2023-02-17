package tiketdto

type AddTiket struct {
	Jadwal         string `json:"jadwal" form:"jadwal"`
	TrainID        int    `json:"train" form:"train"`
	StasiunAsal    int    `json:"stasiunasal" form:"stasiunasal"`
	WaktuBerangkat string `json:"waktuberangkat" form:"waktuberangkat"`
	StasiunTujuan  int    `json:"stasiuntujuan" form:"stasiuntujuan"`
	WaktuTiba      string `json:"waktutiba" form:"waktutiba"`
	Harga          int    `json:"harga" form:"harga"`
	Stock          int    `json:"stock" form:"stock"`
}
type FilterTiket struct {
	Jadwal string `json:"jadwal" form:"jadwal"`
	Asal   string `json:"asal" form:"asal"`
	Tujuan string `json:"tujuan" form:"tujuan"`
}

type TransTiket struct {
	Qty int `json:"qty" form:"qty"`
}
