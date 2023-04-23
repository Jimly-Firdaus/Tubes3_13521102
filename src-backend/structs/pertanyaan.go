package structs

type Pertanyaan struct {
  IdPertanyaan int64   `json:"id_pertanyaan"`
  Pertanyaan   string  `json:"pertanyaan"`
  Jawaban      string  `json:"jawaban"`
}
