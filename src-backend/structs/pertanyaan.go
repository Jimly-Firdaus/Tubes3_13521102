package structs

type Pertanyaan struct {
  idPertanyaan int64   `json:"id_pertanyaan"`
  pertanyaan   string  `json:"pertanyaan"`
  jawaban      string  `json:"jawaban"`
}
