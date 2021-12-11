const testFish = [3, 4, 3, 1, 2]
const fish = [
    1, 1, 3, 5, 1, 1, 1, 4, 1, 5, 1, 1, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 2, 5, 1, 1, 1, 1, 1, 2, 1, 4, 1, 4, 1, 1, 1, 1, 1, 3, 1, 1,
    5, 1, 1, 1, 4, 1, 1, 1, 4, 1, 1, 3, 5, 1, 1, 1, 1, 4, 1, 5, 4, 1, 1, 2, 3, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 1,
    1, 1, 1, 1, 1, 1, 2, 2, 1, 1, 1, 1, 1, 5, 1, 1, 1, 3, 4, 1, 1, 1, 1, 3, 1, 1, 1, 1, 1, 4, 1, 1, 3, 1, 1, 3, 1, 1, 1, 1, 1, 3,
    1, 5, 2, 3, 1, 2, 3, 1, 1, 2, 1, 2, 4, 5, 1, 5, 1, 4, 1, 1, 1, 1, 2, 1, 5, 1, 1, 1, 1, 1, 5, 1, 1, 3, 1, 1, 1, 1, 1, 1, 4, 1,
    2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 2, 1, 1, 1, 1, 2, 2, 1, 2, 1, 1, 1, 5, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
    2, 1, 1, 4, 2, 1, 4, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1,
    3, 1, 1, 3, 3, 1, 1, 1, 3, 5, 1, 1, 4, 1, 1, 1, 1, 1, 4, 1, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1,
    1, 4, 1, 1, 1, 1,
]
const nextDay = (prevDay) => {
    const nextDayWithout = prevDay.map((fish) => (fish - 1 < 0 ? 6 : fish - 1))
    const numberOfNewFish = nextDayWithout.filter((fish, i) => fish === 6 && prevDay[i] === 0).length
    const nextDay = [...nextDayWithout, ...[...Array(numberOfNewFish)].map((_) => 8)]
    // console.log(nextDay)
    return nextDay
}
const recurseForN = (startingFish, n) =>
    [...Array(n)].reduce((prev, _, i) => {
        // console.log(i)
        return nextDay(prev)
    }, startingFish)

console.log(recurseForN(fish, 80).length)

/////////////

// fish tree
// get root fish start day
// end day -> for each cycle add fish with start day and see how many fish they add

const endDay = 256
const fishCache = {}
// fish: {startDay}

const getNumberOfChildren = (startDay) => {
    const cache = fishCache[startDay]
    if (cache !== undefined) {
        return cache
    }
    const daysToLive = endDay - startDay
    const numberOfDirectChildren = Math.floor(daysToLive / 7) + 1
    if (numberOfDirectChildren <= 0) {
        fishCache[startDay] = 0
        return 0
    }
    const directChildrenStartDays = [...Array(numberOfDirectChildren).keys()].map((i) => i * 7 + 9 + startDay)
    const numOfChilds = directChildrenStartDays.reduce((prev, childStartDay) => prev + getNumberOfChildren(childStartDay) + 1, 0)
    //console.log(fishCache)
    fishCache[startDay] = numOfChilds
    return numOfChilds
}

const getStartDays = (fishies) => fishies.map((fish) => fish + 1)
const getNumOfFish = (fish) =>
    getStartDays(fish)
        .map(getNumberOfChildren)
        .reduce((prev, next) => prev + next, 0) + fish.length
console.time('countFish')
console.log(getNumOfFish(fish))
console.timeEnd('countFish')
