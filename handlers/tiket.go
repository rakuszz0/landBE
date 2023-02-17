package handlers

import (
	"encoding/json"
	"fmt"
	dto "landtick/dto/result"
	tiketdto "landtick/dto/tiket"
	"landtick/models"
	"landtick/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerTiket struct {
	TiketRepository repositories.TiketRepository
}

func HandlerTiket(tiketRepository repositories.TiketRepository) *handlerTiket {
	return &handlerTiket{tiketRepository}
}
func (h *handlerTiket) CreateTiket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	train, _ := strconv.Atoi(r.FormValue("train"))
	stasiunasal, _ := strconv.Atoi(r.FormValue("stasiunasal"))
	stasiuntujuan, _ := strconv.Atoi(r.FormValue("stasiuntujuan"))
	hargaTiket, _ := strconv.Atoi(r.FormValue("harga"))
	stokTiket, _ := strconv.Atoi(r.FormValue("stock"))

	request := tiketdto.AddTiket{
		Jadwal:         r.FormValue("jadwal"),
		TrainID:        train,
		StasiunAsal:    stasiunasal,
		WaktuBerangkat: r.FormValue("waktuberangkat"),
		StasiunTujuan:  stasiuntujuan,
		WaktuTiba:      r.FormValue("waktutiba"),
		Harga:          hargaTiket,
		Stock:          stokTiket,
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	jadwal, _ := time.Parse("2006-01-02", request.Jadwal)
	TimeStart, _ := time.Parse("2006-01-02T15:04", request.WaktuBerangkat)
	TimeEnd, _ := time.Parse("2006-01-02T15:04", request.WaktuTiba)

	fmt.Println(jadwal)
	fmt.Println(TimeStart)
	fmt.Println(TimeEnd)
	fmt.Println(request.WaktuTiba)

	dataTiket := models.Tiket{
		Jadwal:          jadwal.Local(),
		TrainID:         request.TrainID,
		StasiunAsalID:   request.StasiunAsal,
		WaktuBerangkat:  TimeStart,
		StasiunTujuanID: request.StasiunTujuan,
		WaktuTiba:       TimeEnd,
		Harga:           request.Harga,
		Stock:           request.Stock,
		Kode:            "KAI987",
	}

	data, err := h.TiketRepository.CreateTiket(dataTiket)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Cek Data Tiket"}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerTiket) FilterTiket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	request := new(tiketdto.FilterTiket)
	fmt.Println(request)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Cek Dto Request"}
		json.NewEncoder(w).Encode(response)
		return
	}
	fmt.Println(request.Jadwal)

	jadwal, _ := time.Parse("2006-01-02", request.Jadwal)
	fmt.Println(jadwal.UTC())

	KotaAsal, err := h.TiketRepository.FilterKotaStasiun(request.Asal)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Cek Kota Asal"}
		json.NewEncoder(w).Encode(response)
		return
	}

	KotaTujuan, err := h.TiketRepository.FilterKotaStasiun(request.Tujuan)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Cek Kota Tujuan"}
		json.NewEncoder(w).Encode(response)
		return
	}

	Tiket, err := h.TiketRepository.FilterTiket(KotaAsal.ID, KotaTujuan.ID, jadwal)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Cek Data Tiket"}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: Tiket}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerTiket) GetAllTiket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tiket, err := h.TiketRepository.FindTiket()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Cek Tiket All"}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: tiket}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerTiket) CreateTransTiket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	request := new(tiketdto.TransTiket)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Cek Dto Request"}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	idTiket, _ := strconv.Atoi(mux.Vars(r)["id"])

	tiket, err := h.TiketRepository.GetTiket(idTiket)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Not Found Tiket"}
		json.NewEncoder(w).Encode(response)
		return
	}

	TotalHarga := tiket.Harga * request.Qty

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	IdUser := int(userInfo["id"].(float64))

	transID := int(time.Now().Unix())

	requestTrans := models.Transaction{
		ID:        transID,
		Status:    "pending",
		TiketID:   idTiket,
		UserID:    IdUser,
		Qty:       request.Qty,
		Total:     TotalHarga,
		DateTrans: time.Now(),
	}
	fmt.Println(requestTrans.UserID)

	trans, err := h.TiketRepository.CreateTrans(requestTrans)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Cek Request Transaksi"}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trans}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerTiket) DeleteTiket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	tiket, err := h.TiketRepository.GetTiket(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TiketRepository.Deletetiket(tiket)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
