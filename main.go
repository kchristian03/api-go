package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type mahasiswa struct {
	Nim   string `json:"nim"`
	Name  string `json:"name"`
	Prodi string `json:"prodi"`
}

var initData = []mahasiswa{
	{Nim: "0706012110011", Name: "Agus", Prodi: "IMT"},
	{Nim: "0706012110022", Name: "Cahyo", Prodi: "IMT"},
	{Nim: "0706012110033", Name: "Nugraha", Prodi: "IMT"},
}

func getMahasiswa(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, initData)
}

func getMahasiswa2(context *gin.Context) {
	nim := context.Param("nim")
	mahasiswa, err := getMahasiswaByNim(nim)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound,
			gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, mahasiswa)
}

func getMahasiswaByNim(nim string) (*mahasiswa, error) {
	for i, t := range initData {
		if t.Nim == nim {
			return &initData[i], nil
		}
	}
	return nil, errors.New("data not found")
}

func addMahasiswa(context *gin.Context) {
	var newData mahasiswa
	if err := context.BindJSON(&newData); err != nil {
		return
	}
	initData = append(initData, newData)
	context.IndentedJSON(http.StatusCreated, newData)
}

func editMahasiswa_Prodi(context *gin.Context) {
	nim := context.Param("nim")
	prodi := context.Param("prodi")

	mahasiswa, err := getMahasiswaByNim(nim)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound,
			gin.H{"message": err.Error()})
		return
	}
	mahasiswa.Prodi = prodi
	context.IndentedJSON(http.StatusOK, mahasiswa)
}

func main() {
	router := gin.Default()
	router.GET("/mahasiswa", getMahasiswa)
	router.POST("/mahasiswa", addMahasiswa)
	router.GET("/mahasiswa/:nim", getMahasiswa2)
	router.PATCH("/mahasiswa/:nim/:prodi", editMahasiswa_Prodi)
	router.Run("localhost:9090")
}
