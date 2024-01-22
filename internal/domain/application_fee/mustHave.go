package application_fee

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"

	"github.com/pericles-luz/go-base/pkg/utils"
)

const (
	MustHaveNotFound    = "não encontrado"
	MustHaveFound       = "encontrado"
	MustHaveAlreadyPaid = "já pago"
)

type MustHave struct {
	ProcessNumber  string `json:"processNumber"`
	Execution      uint16 `json:"execution"`
	Sequence       uint16 `json:"sequence"`
	Name           string `json:"name"`
	CPF            string `json:"cpf"`
	classification string
}

func NewMustHave() *MustHave {
	return &MustHave{}
}

func (m *MustHave) SetClassification(classification string) {
	m.classification = classification
}

func (m *MustHave) Classification() string {
	return m.classification
}

type MustHaveList struct {
	mustHaveList []*MustHave
}

func NewMustHaveList() *MustHaveList {
	return &MustHaveList{}
}

func (m *MustHaveList) FromJSONFile(path string) error {
	if !utils.FileExists(path) {
		return errors.New("arquivo não encontrado")
	}
	source, err := os.Open(path)
	if err != nil {
		return err
	}
	defer source.Close()
	sourceRaw := json.NewDecoder(bufio.NewReader(source))
	if err := sourceRaw.Decode(&m.mustHaveList); err != nil {
		return err
	}
	for _, mustHave := range m.mustHaveList {
		mustHave.CPF = utils.GetOnlyNumbers(mustHave.CPF)
	}
	return nil
}

func (m *MustHaveList) List() []*MustHave {
	return m.mustHaveList
}

func (m *MustHaveList) Has(cpf string) bool {
	cpf = utils.GetOnlyNumbers(cpf)
	for _, mustHave := range m.mustHaveList {
		if mustHave.CPF == cpf {
			return true
		}
	}
	return false
}

func (m *MustHaveList) Get(cpf string) *MustHave {
	cpf = utils.GetOnlyNumbers(cpf)
	for _, mustHave := range m.mustHaveList {
		if mustHave.CPF == cpf {
			return mustHave
		}
	}
	return nil
}

func (m *MustHaveList) SetClassification(cpf, classification string) {
	if mustHave := m.Get(cpf); mustHave != nil {
		mustHave.SetClassification(classification)
	}
}
