export const sum = (arr) => arr.reduce((prev, nex) => prev + nex, 0)
export const forEachNeighbour = (arr, [y, x], callback) => {
    const sizeY = arr.length
    const sizeX = arr[0].length
    if (!(y === 0)) callback(arr, [y - 1, x])
    if (!(x === 0)) callback(arr, [y, x - 1])
    if (!(x === sizeX - 1)) callback(arr, [y, x + 1])
    if (!(y === sizeY - 1)) callback(arr, [y + 1, x])
    if (!(y === 0) && !(x === 0)) callback(arr, [y - 1, x - 1])
    if (!(y === 0) && !(x === sizeX - 1)) callback(arr, [y - 1, x + 1])
    if (!(y === sizeY - 1) && !(x === sizeX - 1)) callback(arr, [y + 1, x + 1])
    if (!(y === sizeY - 1) && !(x === 0)) callback(arr, [y + 1, x - 1])
}
