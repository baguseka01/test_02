package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type TransaksiStruct struct {
	ID                 int       `json:"id"`
	Tanggal_Transaksi  time.Time `json:"tgl_transaksi"`
	No_Transaksi       string    `json:"no_transaksi"`
	Nama               string    `json:"nama"`
	Jenis_Kelamin      string    `json:"jenis_kelamin"`
	Usia               string    `json:"usia"`
	Email              string    `json:"email"`
	Perokok            string    `json:"perokok"`
	Nominal            float64   `json:"nominal"`
	Lama_Investasi     string    `json:"lama_investasi"`
	Periode_Pembayaran string    `json:"periode_pembayaran"`
	Metode_Bayar       string    `json:"metode_bayar"`
	Total_Bayar        float64   `json:"total_bayar"`
}

var (
	Database = map[int]*TransaksiStruct{}
	seq      = 1
)


// --------------------------------------------------------SOAL 2----------------------------------------------------------------

func GetController(e echo.Context) error {
	return e.JSON(SuccessResponseWithData(Database))
}

func InputController(e echo.Context) error {
	transaksi := &TransaksiStruct{
		ID: seq,
	}

	if err := e.Bind(transaksi); err != nil {
		return err
	}

	for _, v := range Database {
		if v.Email == transaksi.Email {
			return e.JSON(http.StatusBadRequest, "Email ditolak!")
		}
	}

	transaksi.Usia = transaksi.Usia + " " + "tahun"
	transaksi.Lama_Investasi = transaksi.Lama_Investasi + " " + "tahun"
	transaksi.Tanggal_Transaksi = time.Now()
	transaksi.No_Transaksi = "TRX000001"

	if transaksi.Periode_Pembayaran == "tahunan" {

		transaksi.Total_Bayar = transaksi.Nominal - (transaksi.Nominal / 12)
	}

	Database[transaksi.ID] = transaksi
	seq++

	return e.JSON(SuccessResponseWithData(transaksi))

}

// --------------------------------------------------SOAL NO 3 DAN 4----------------------------------------------------------------

func UpdateController(e echo.Context) error {
	u := new(TransaksiStruct)

	if err := e.Bind(u); err != nil {
		return err
	}

	id, _ := strconv.Atoi(e.Param("id"))

	Database[id].No_Transaksi = u.No_Transaksi
	Database[id].Nama = u.Nama
	Database[id].Jenis_Kelamin = u.Jenis_Kelamin
	Database[id].Usia = u.Usia
	Database[id].Email = u.Email
	Database[id].Perokok = u.Perokok
	Database[id].Nominal = u.Nominal
	Database[id].Lama_Investasi = u.Lama_Investasi
	Database[id].Periode_Pembayaran = u.Periode_Pembayaran
	Database[id].Metode_Bayar = u.Metode_Bayar

	return e.JSON(SuccessResponseWithData(Database[id]))
}

func DeleteController(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	delete(Database, id)

	return e.JSON(SuccessResponseWithoutData())

}

func main() {
	e := echo.New()

	e.POST("", InputController)
	e.GET("", GetController)
	e.PUT("/:id", UpdateController)
	e.DELETE("/:id", DeleteController)

	e.Start(":7000")
}

var Success string = "200"

type SuccessResponseSpec struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponseWithData(data interface{}) (int, *SuccessResponseSpec) {
	return http.StatusOK, &SuccessResponseSpec{
		Code:    Success,
		Message: "success",
		Data:    data,
	}
}

func SuccessResponseWithoutData() (int, *SuccessResponseSpec) {
	return http.StatusOK, &SuccessResponseSpec{
		Code:    Success,
		Message: "success",
		Data:    map[string]interface{}{},
	}
}
