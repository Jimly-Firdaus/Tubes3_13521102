package structs

type PertanyaanTopik struct {
  IdPertanyaan int64      `json:"id_pertanyaan"`
  IdTopik      int64      `json:"id_topik"`
  Pertanyaan   Pertanyaan `json:"pertanyaan"`
}
