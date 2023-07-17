package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Moreira-Henrique-Pedro/sorteador-fantasy/utils"
)

type Time struct {
	Nome string
	GM   string
}

func main() {
	listaTimes := []Time{
		{Nome: utils.TIME_GRAJA_CITY_ROCKETS, GM: utils.GM_FERNANDO_RODRIGUES},
		{Nome: utils.TIME_FANTASY_LEGEND, GM: utils.GM_RODRIGO_MOREIRA},
		{Nome: utils.TIME_POSEIDON_TRIDENTS, GM: utils.GM_LUCCA_JANNUZZI},
		{Nome: utils.TIME_THE_CLUCHT, GM: utils.GM_GABRIEL_MEDEIROS},
		{Nome: utils.TIME_TEAM_PICCAGLI, GM: utils.GM_ANDRE_PICCAGLI},
		{Nome: utils.TIME_MINHA_MAE_DO_CEU, GM: utils.GM_ESTANISLAU_CORREA},
		{Nome: utils.TIME_FAUSTINHO_TEAMS_MVP, GM: utils.GM_GUILHERME_ALI_BASSANI},
		{Nome: utils.TIME_MEAN_MACHINE, GM: utils.GM_DENIS_CORREA},
		{Nome: utils.TIME_ITAPARK_WAR_MACHINES, GM: utils.GM_CARLOS_EDUARDO_DOS_SANTOS},
		{Nome: utils.TIME_RITMO_RAGATANGA_INSANO, GM: utils.GM_DANIEL_SOUZA},
		{Nome: utils.TIME_CAMPINAS_WARRIORS, GM: utils.GM_ARTUR_LIMA},
		{Nome: utils.TIME_SUPERBOWL_DO_VIKINGS, GM: utils.GM_DOUGLAS_MARINHO},
		{Nome: utils.TIME_GREEN_BAY_PORCOS, GM: utils.GM_GIL_NASCIMENTO},
		{Nome: utils.TIME_NICIOLI, GM: utils.GM_GUSTAVO_NICIOLI},
		{Nome: utils.TIME_Z_OH_SHIT_A_MVP, GM: utils.GM_GUILHERME_ZANINETTI},
		{Nome: utils.TIME_MENSAGEIRO_DO_CAOS, GM: utils.GM_LUIZ_CARLOS_LIMA_ALVES},
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(listaTimes), func(i, j int) { listaTimes[i], listaTimes[j] = listaTimes[j], listaTimes[i] })

	numGrupos := 4
	timesPorGrupo := len(listaTimes) / numGrupos
	grupos := make([][]Time, numGrupos)
	for i := 0; i < numGrupos; i++ {
		grupos[i] = make([]Time, timesPorGrupo)
		copy(grupos[i], listaTimes[i*timesPorGrupo:(i+1)*timesPorGrupo])
	}

	for i, grupo := range grupos {
		fmt.Printf("Grupo %d:\n", i+1)
		for _, time := range grupo {
			fmt.Println(time.Nome, time.GM)
		}
		fmt.Println()
	}
}
