package application

import "errors"

var (
	mes = map[string]string{
		"jan": "01",
		"fev": "02",
		"mar": "03",
		"abr": "04",
		"mai": "05",
		"jun": "06",
		"jul": "07",
		"ago": "08",
		"set": "09",
		"out": "10",
		"nov": "11",
		"dez": "12",
	}
	ErrLinhaJaExistente           = errors.New("linha ja existente")
	ErrMesAnoInvalido             = errors.New("mes/ano invalido")
	ErrVencimentoBasicoInvalido   = errors.New("vencimento basico invalido")
	ErrSomaInvalida               = errors.New("soma invalida")
	ErrPercentualInvalido         = errors.New("percentual invalido")
	ErrValorDevidoInvalido        = errors.New("valor devido invalido")
	ErrValorCorrigidoInvalido     = errors.New("valor corrigido invalido")
	ErrValorJurosMoraInvalido     = errors.New("valor juros mora invalido")
	ErrTotalDevidoInvalido        = errors.New("total devido invalido")
	ErrIndiceCorrecaoInvalido     = errors.New("indice correcao invalido")
	ErrJurosMoraInvalido          = errors.New("juros mora invalido")
	ErrProcessoNumeroInvalido     = errors.New("processo numero invalido")
	ErrProcessoPrincipalInvalido  = errors.New("processo principal invalido")
	ErrNumeroExecucaoInvalido     = errors.New("numero execucao invalido")
	ErrAjuizamentoInvalido        = errors.New("data de ajuizamento invalida")
	ErrCitacaoInvalido            = errors.New("data de citacao invalida")
	ErrDataCalculoInvalido        = errors.New("data de calculo invalida")
	ErrExequenteInvalido          = errors.New("exequente invalido")
	ErrCpfInvalido                = errors.New("cpf invalido")
	ErrIdUnicaInvalido            = errors.New("id unica invalido")
	ErrTabelaInvalida             = errors.New("tabela invalida")
	ErrTotalInvalido              = errors.New("total invalido")
	ErrTotalJurosInvalido         = errors.New("total juros invalido")
	ErrTotalPrincipalInvalido     = errors.New("total principal invalido")
	ErrTotalHonorarioInvalido     = errors.New("total honorario invalido")
	ErrTotalNaoBate               = errors.New("total nao bate")
	ErrTotalHonorarioNaoBate      = errors.New("total honorario nao bate")
	ErrLinhaInvalida              = errors.New("linha invalida")
	ErrLinhaNaoEhTotalGeral       = errors.New("linha nao eh total geral")
	ErrDesagio35Invalido          = errors.New("desagio 35 invalido")
	ErrTotalAposDesagio35Invalido = errors.New("total apos desagio 35 invalido")
	ErrHonorarioInvalido          = errors.New("honorario invalido")

	ErrProcessoNumeroNaoEncontrado      = errors.New("processo numero nao encontrado")
	ErrProcessoPrincipalNaoEncontrado   = errors.New("processo principal nao encontrado")
	ErrAjuizamentoNaoEncontrado         = errors.New("data de ajuizamento nao encontrada")
	ErrCitacaoNaoEncontrada             = errors.New("data de citacao nao encontrada")
	ErrCalculoNaoEncontrado             = errors.New("data de calculo nao encontrada")
	ErrExequenteNaoEncontrado           = errors.New("exequente nao encontrado")
	ErrDesagio35NaoEncontrado           = errors.New("desagio 35 nao encontrado")
	ErrTotalAposDesagio35NaoEncontrado  = errors.New("total apos desagio 35 nao encontrado")
	ErrValorCorrigidoTotalNaoEncontrado = errors.New("valor corrigido total nao encontrado")
	ErrTotalNaoEncontrado               = errors.New("total nao encontrado")
	ErrMesAnoNaoEncontrado              = errors.New("mes/ano nao encontrado")
	ErrMesAnoNaoExiste                  = errors.New("mes/ano nao existe")
	ErrVencimentoBasicoNaoEncontrado    = errors.New("vencimento basico nao encontrado")
	ErrSomaNaoEncontrada                = errors.New("soma nao encontrada")
	ErrValorDevidoNaoEncontrado         = errors.New("valor devido nao encontrado")
	ErrIndiceCorrecaoNaoEncontrado      = errors.New("indice correcao nao encontrado")
	ErrValorCorrigidoNaoEncontrado      = errors.New("valor corrigido nao encontrado")
	ErrJurosMoraNaoEncontrado           = errors.New("juros mora nao encontrado")
	ErrValorJurosMoraNaoEncontrado      = errors.New("valor juros mora nao encontrado")
	ErrTotalDevidoNaoEncontrado         = errors.New("total devido nao encontrado")
	ErrPercentualNaoEncontrado          = errors.New("percentual nao encontrado")

	ErrNumeroLocalDeExecucaoNaoEncontrado = errors.New("numero local de execucao nao encontrado")

	ErrArquivoNaoEncontrado    = errors.New("arquivo nao encontrado")
	ErrArquivoInvalido         = errors.New("arquivo invalido")
	ErrCalculoInconsistente    = errors.New("calculos inconsistentes")
	ErrNenhumArquivoEncontrado = errors.New("nenhum arquivo encontrado")
)
