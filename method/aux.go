package method

// To validate shoe size by comparing against rage per sex
func InRange(shoe Shoes) bool{
    switch {
    case shoe.Size < Limits[shoe.Sex].Min:
        return false
    case shoe.Size > Limits[shoe.Sex].Max:
        return false
    default:
        return true
    }
}

// Initialize cubby with random numbers
func CubbyStart(max int){
    men := Limits[M].Max
    women := Limits[W].Max
    rand.Seed(time.Now().UnixNano())

    for size := men; size >= Limits[M].Min; size--{
        ShoeReturn[Shoes{size, M}] = rand.Intn(max)
    }
    for size := women; size >= Limits[W].Min; size--{
        ShoeReturn[Shoes{size, W}] = rand.Intn(max)
    }
}

