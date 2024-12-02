#include "LevelSet.h"

LevelSet::LevelSet() {

}

void LevelSet::addLevel(int num) {
	set.push_back(num);
}

void LevelSet::show(std::ostream& stream) {	
	stream << (isSafeWithDampener() ? "Safe" : "Unsafe") << " set with: " << getCount() << " elements." << std::endl;
	stream << "{ ";

	for (int i = 0; i < getCount(); i++) {
		if (i == getCount() - 1) {
			stream << set[i] << " } ";
		}
		else {
			stream << set[i] << ", ";
		}
	}

	stream << std::endl;
}

bool LevelSet::isSafeWithDampener() {
	if (isSafe()) {
		return true;
	}


	for (int i = 0; i < set.size(); i++) {
		
		std::vector<int> tempSet = set;
		tempSet.erase(tempSet.begin() + i);

		if (isSafe(tempSet)) {
			return true;
		}
	}

	return false;
}

bool LevelSet::isSafe() const {
	return isSafe(set);
}

bool LevelSet::isSafe(std::vector<int> vectorToTest) const {
	bool patternHasBeenSet = false;
	bool patternIncreasing = false;
	int adjacentDifference = 0;
	bool safe = true;

	for (int i = 0; i < vectorToTest.size(); i++) {
		bool isLevelIncreasing = false;
		if (i == vectorToTest.size() - 1) {
			break;
		}

		// pos or neg?
		int difference = vectorToTest[i] - vectorToTest[i + 1];
		if (difference > 0) {
			isLevelIncreasing = true;
		}
		else {
			isLevelIncreasing = false;
		}

		// distance greater than three?
		if (abs(difference) < 1 || abs(difference) > 3) {
			safe = false;
			break;
		}

		if (patternIncreasing != isLevelIncreasing && patternHasBeenSet) {
			safe = false;
			break;
		}

		if (patternHasBeenSet == false) {
			patternIncreasing = isLevelIncreasing;
			patternHasBeenSet = true;
		}

	}

	return safe;
}
