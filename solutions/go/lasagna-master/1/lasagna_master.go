package lasagnamaster

func PreparationTime(layers []string, time int) int {
    if time == 0 {
        time = 2
    }
    return len(layers) * time
}

func Quantities(layers []string) (noodles int, sauce float64) {
    for _, layer := range layers {
        if layer == "noodles" {
            noodles += 50
        } else if layer == "sauce" {
            sauce += 0.2
        }
    }
    return
}

func AddSecretIngredient(list1 []string, list2 []string) {
    list2[len(list2)-1] = list1[len(list1)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
    result := make([]float64, len(quantities))
    factor := float64(portions) / 2.0
    for i, val := range quantities {
        result[i] = val * factor
    }
    return result
}

