package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	//panic("RemainingOvenTime not implemented")
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	//panic("PreparationTime not implemented")
	return numberOfLayers * 2 // 2 minutes per layer
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	//panic("ElapsedTime not implemented")
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
