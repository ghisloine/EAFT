package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"log"

	"github.com/MaxHalford/eaopt"
	"github.com/manifoldco/promptui"
)
var Pc = ManuelConfiguration()
func selectMainAlgorithm(obj *GeneticObject) {

	prompt := promptui.Select{
		Label: "Select Main Algorithm",
		Items: MainAlgorithms,
	}

	_, MainAlgorithm, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// ? Mode selection based on algorithm
	obj.ObjectType = MainAlgorithm
	selectModelSettings(MainAlgorithm, obj)

}

func selectModelSettings(MainAlgorithm string, obj *GeneticObject) {
	if MainAlgorithm == "Genetic Algorithm" {
		prompt := promptui.Select{
			Label: "Select Model",
			Items: Models,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
		}

		if result == "Generational" {
			obj.ObjectStruct.Model = eaopt.ModGenerational{
				Selector:  selectSelectionMethod(obj),
				MutRate:   selectIntResults("Please Enter Mutation Rate ?", true),
				CrossRate: selectIntResults("Please Enter Crossover Rate ?", true),
			}
		} else if result == "Steady State" {
			obj.ObjectStruct.Model = eaopt.ModSteadyState{
				Selector:  selectSelectionMethod(obj),
				MutRate:   selectIntResults("Please Enter Mutation Rate ?", true),
				CrossRate: selectIntResults("Please Enter Crossover Rate ?", true),
				KeepBest:  true,
			}
		} else if result == "Down to Size" {
			obj.ObjectStruct.Model = eaopt.ModDownToSize{
				SelectorA:   selectSelectionMethod(obj),
				SelectorB:   selectSelectionMethod(obj),
				MutRate:     selectIntResults("Please Enter Mutation Rate ?", true),
				CrossRate:   selectIntResults("Please Enter Crossover Rate ?", true),
				NOffsprings: uint(selectIntResults("Please Enter Number of Offsprings : ", false)),
			}
		} else if result == "Ring" {
			obj.ObjectStruct.Model = eaopt.ModRing{
				Selector: selectSelectionMethod(obj),
				MutRate:  selectIntResults("Please Enter Mutation Rate ?", true),
			}
		} else if result == "Mutation Only" {
			obj.ObjectStruct.Model = eaopt.ModMutationOnly{
				Strict: true,
			}
		}

	} else if MainAlgorithm == "Particle Swarm Optimization" {
		fmt.Println()
	}
}

func selectSelectionMethod(obj *GeneticObject) eaopt.Selector {
	var SelectionMethod eaopt.Selector
	prompt := promptui.Select{
		Label: "Select Selection Methods",
		Items: Selections,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Println(err)
	}

	if result == "Tournament" {
		SelectionMethod = eaopt.SelTournament{
			NContestants: uint(selectIntResults("Please Enter Number of Contestants : ", false)),
		}
	} else if result == "Roulette" {
		SelectionMethod = eaopt.SelRoulette{}
	} else if result == "Elitism" {
		SelectionMethod = eaopt.SelElitism{}
	}

	return SelectionMethod
}

func selectStringResults(question string) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New("string is required")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     question,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func selectIntResults(question string, restriction bool) float64 {
	validate := func(input string) error {
		result, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("can't convert string to uint")
		}

		if restriction {
			if (result >= 0) && (result <= 1) {
				return nil
			} else {
				return errors.New("value must be between [0, 1]")
			}
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     question,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	int_result, _ := strconv.ParseFloat(result, 64)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return int_result
}

// TODO Tum bu konfigurasyonlar icin test yaz.
// TODO Bu fonksiyonun ciktisi her yerden erisilebilir olmali.

func SelectConfigurations() GeneticObject {
	
	
	GenObj := GeneticObject{}
	selectMainAlgorithm(&GenObj)
	GenObj.ResultFolderName = selectStringResults("What will be the result folder name ?")
	GenObj.GccShortcut = selectStringResults("How do you call GCC in your CLI ?")
	GenObj.ObjectStruct.NGenerations = uint(selectIntResults("Please Enter Number of Generation ?", false))
	GenObj.ObjectStruct.PopSize = uint(selectIntResults("Please Enter Population Size?", false))
	GenObj.ObjectStruct.NPops = 1
	GenObj.ObjectStruct.HofSize = 1
	GenObj.ObjectStruct.ParallelEval = true

	return GenObj
}

func ManuelConfiguration() GeneticObject {
	GenObj := GeneticObject{}

	GenObj.ObjectType = "Genetic Algorithm"
	GenObj.ResultFolderName = "2mm"
	GenObj.GccShortcut = "gcc-12"

	GenObj.ObjectStruct = eaopt.GAConfig{
		NPops:        1,
		PopSize:      30,
		HofSize:      1,
		NGenerations: 50,
		Model: eaopt.ModGenerational{
			Selector: eaopt.SelTournament{
				NContestants: 3,
			},
			MutRate:   0.5,
			CrossRate: 0.7,
		},
		ParallelEval: true,
	}

	log.Println(GenObj)
	return GenObj
}
