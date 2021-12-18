export const sum = (arr: number[]): number => arr.reduce((prev, nex) => prev + nex, 0);
export const forEachNeighbour = <T>(arr: T[][], [y, x]: number[], callback: (arr: T[][], pos: number[]) => void) => {
    const sizeY = arr.length;
    const sizeX = arr[0].length;
    if (!(y === 0)) callback(arr, [y - 1, x]);
    if (!(x === 0)) callback(arr, [y, x - 1]);
    if (!(x === sizeX - 1)) callback(arr, [y, x + 1]);
    if (!(y === sizeY - 1)) callback(arr, [y + 1, x]);
    if (!(y === 0) && !(x === 0)) callback(arr, [y - 1, x - 1]);
    if (!(y === 0) && !(x === sizeX - 1)) callback(arr, [y - 1, x + 1]);
    if (!(y === sizeY - 1) && !(x === sizeX - 1)) callback(arr, [y + 1, x + 1]);
    if (!(y === sizeY - 1) && !(x === 0)) callback(arr, [y + 1, x - 1]);
};

export const distinct = <T>(value: T, index: number, self: T[]) => self.indexOf(value) === index;
export const customDistinct =
    <T>(equals: (a: T, b: T) => boolean) =>
    (value: T, index: number, self: T[]) =>
        self.findIndex((x) => equals(value, x)) === index;
export const copies = <T>(value: T, index: number, self: T[]) => self.indexOf(value) !== index;
export const getGrid = <T>(sizeX: number, sizeY: number, cellF: (x: number, y: number) => T) =>
    [...Array(sizeY)].map((_, y) => [...Array(sizeX)].map((_, x) => cellF(x, y)));
