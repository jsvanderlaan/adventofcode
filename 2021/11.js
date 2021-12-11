const test = [
    '5483143223',
    '2745854711',
    '5264556173',
    '6141336146',
    '6357385478',
    '4167524645',
    '2176841721',
    '6882881134',
    '4846848554',
    '5283751526',
]

const input = [
    '4472562264',
    '8631517827',
    '7232144146',
    '2447163824',
    '1235272671',
    '5133527146',
    '6511372417',
    '3841841614',
    '8621368782',
    '3246336677',
]

const test2 = ['11111', '19991', '19191', '19991', '11111']

const getMap = (arr) => arr.map((row) => row.split('').map((cell) => parseInt(cell)))

const countFlashes = (map) => map.reduce((p, n) => p + n.reduce((prev, next) => prev + (next === 0 ? 1 : 0), 0), 0)

const step = (map) => {
    const sizeX = map[0].length
    const sizeY = map.length
    const increasedMap = map.map((row) => row.map((cell) => cell + 1))
    const flashed = increasedMap.map((row) => row.map((_) => false))
    const traverse = ([yT, xT]) => {
        if (!flashed[yT][xT] && ++increasedMap[yT][xT] > 9) {
            flashed[yT][xT] = true
            if (!(yT === 0)) traverse([yT - 1, xT])
            if (!(xT === 0)) traverse([yT, xT - 1])
            if (!(xT === sizeX - 1)) traverse([yT, xT + 1])
            if (!(yT === sizeY - 1)) traverse([yT + 1, xT])
            if (!(yT === 0) && !(xT === 0)) traverse([yT - 1, xT - 1])
            if (!(yT === 0) && !(xT === sizeX - 1)) traverse([yT - 1, xT + 1])
            if (!(yT === sizeY - 1) && !(xT === sizeX - 1)) traverse([yT + 1, xT + 1])
            if (!(yT === sizeY - 1) && !(xT === 0)) traverse([yT + 1, xT - 1])
        }
    }

    increasedMap.forEach((row, y) =>
        row.forEach((cell, x) => {
            if (!flashed[y][x] && cell > 9) {
                traverse([y, x])
            }
        })
    )

    return increasedMap.map((row) => row.map((cell) => (cell > 9 ? 0 : cell)))
}

const performSteps = (arr, count) =>
    [...Array(count)].reduce(
        (prev) => {
            console.log(prev.map.map((row) => row.join('')).join('\n'))
            console.log('\n')
            const newMap = step(prev.map)
            const flashes = countFlashes(newMap)
            return { map: newMap, flashes: prev.flashes + flashes }
        },
        { map: getMap(arr), flashes: 0 }
    )

console.log(performSteps(input, 100))

let currStep = 0
let currMap = getMap(input)
while (true) {
    currStep++
    console.log(currStep)
    currMap = step(currMap)
    if (countFlashes(currMap) === 100) {
        break
    }
}
