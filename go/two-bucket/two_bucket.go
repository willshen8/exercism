package twobucket

import "errors"

type bucket struct {
	capacity     int
	currentLevel int
	startBucket  bool
}

// Solve resolves whether two buckets can reach a goal limit by pouring water in and out and the buckets
func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (goalBucket string, numSteps, otherBucketLevel int, e error) {
	if goalAmount <= 0 {
		return "", numSteps, 0, errors.New("invalid goal amount")
	}
	if sizeBucketOne <= 0 || sizeBucketTwo <= 0 {
		return "", numSteps, 0, errors.New("bucket size must be positive")
	}
	if sizeBucketTwo&sizeBucketOne == 0 {
		return "", numSteps, 0, errors.New("solution can't be found, two bucket sizes must be prime relative to each other")
	}

	bucketOne := &bucket{capacity: sizeBucketOne}
	bucketTwo := &bucket{capacity: sizeBucketTwo}
	switch startBucket {
	case "one":
		bucketOne.pourIn(bucketOne.capacity)
		bucketOne.startBucket = true
	case "two":
		bucketTwo.pourIn(bucketTwo.capacity)
		bucketTwo.startBucket = true
	default:
		return "", 0, 0, errors.New("invalid starting bucket")
	}
	numSteps++

	for (bucketOne.currentLevel != goalAmount) && (bucketTwo.currentLevel != goalAmount) {
		// if non starting bucket capacity equals to goal amount
		if bucketOne.startBucket && bucketTwo.capacity == goalAmount && bucketTwo.currentLevel == 0 {
			bucketTwo.fill()
			numSteps++
			break
		} else if bucketTwo.startBucket && bucketOne.capacity == goalAmount && bucketOne.currentLevel == 0 {
			bucketOne.fill()
			numSteps++
			break
		}

		// if none starting bucket is full and other is not empty nor full, then empty non-starting
		if !bucketOne.startBucket && bucketOne.currentLevel == bucketOne.capacity {
			bucketOne.empty()
			numSteps++
		} else if !bucketTwo.startBucket && bucketTwo.currentLevel == bucketTwo.capacity {
			bucketTwo.empty()
			numSteps++
		}

		// if starting bucket is full then empty into the other bucket
		if bucketOne.startBucket && bucketTwo.currentLevel == 0 || bucketOne.currentLevel == bucketOne.capacity {
			maxPourAmount := bucketTwo.capacity - bucketTwo.currentLevel
			if bucketOne.currentLevel < maxPourAmount {
				bucketTwo.pourIn(bucketOne.currentLevel)
				bucketOne.empty()
			} else if bucketOne.currentLevel >= maxPourAmount {
				bucketTwo.pourIn(maxPourAmount)
				bucketOne.pourOut(maxPourAmount)
			}
			numSteps++
		} else if bucketTwo.startBucket && bucketOne.currentLevel == 0 || bucketTwo.currentLevel == bucketTwo.capacity {
			maxPourAmount := bucketOne.capacity - bucketOne.currentLevel
			if bucketTwo.currentLevel < maxPourAmount {
				bucketOne.pourIn(bucketTwo.currentLevel)
				bucketTwo.empty()
			} else if bucketTwo.currentLevel >= maxPourAmount {
				bucketOne.pourIn(maxPourAmount)
				bucketTwo.pourOut(maxPourAmount)
			}
			numSteps++
		}

		// if starting bucket is empty then pour in full
		if bucketOne.startBucket && (bucketOne.currentLevel == 0) {
			bucketOne.fill()
			numSteps++
		} else if bucketTwo.startBucket && (bucketTwo.currentLevel == 0) {
			bucketTwo.fill()
			numSteps++
		}
	}

	if bucketOne.currentLevel == goalAmount {
		return "one", numSteps, bucketTwo.currentLevel, nil
	} else {
		return "two", numSteps, bucketOne.currentLevel, nil
	}

}

func (b *bucket) empty() {
	b.currentLevel = 0
}

func (b *bucket) fill() {
	b.currentLevel = b.capacity
}

func (b *bucket) pourIn(amount int) {
	b.currentLevel += amount
}

func (b *bucket) pourOut(amount int) {
	b.currentLevel -= amount
}
