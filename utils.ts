import { Cell, Grid } from './types';

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
export const getOrthogonalNeighbours = <T>([x, y]: Cell, arr: T[][]): Cell[] => {
    const sizeY = arr.length;
    const sizeX = arr[0].length;
    const neighbours = [] as Cell[];
    if (!(y === 0)) neighbours.push([x, y - 1]);
    if (!(x === 0)) neighbours.push([x - 1, y]);
    if (!(x === sizeX - 1)) neighbours.push([x + 1, y]);
    if (!(y === sizeY - 1)) neighbours.push([x, y + 1]);
    return neighbours;
};

export const distinct = <T>(value: T, index: number, self: T[]) => self.indexOf(value) === index;
export const customDistinct =
    <T>(equals: (a: T, b: T) => boolean) =>
    (value: T, index: number, self: T[]) =>
        self.findIndex(x => equals(value, x)) === index;
export const copies = <T>(value: T, index: number, self: T[]) => self.indexOf(value) !== index;

export const getArray = <T>(size: number, itemF: (i: number) => T) => [...Array(size)].map((_, i) => itemF(i));
export const getEmptyArray = (size: number) => [...Array(size)].map(_ => null);
export const getGrid = <T>(sizeX: number, sizeY: number, cellF: (x: number, y: number) => T) =>
    getArray(sizeY, y => getArray(sizeX, x => cellF(x, y)));
export const getEmptyGrid = (sizeX: number, sizeY: number) => getArray(sizeY, _ => getEmptyArray(sizeX));

export const mapGrid = <T, S>(grid: Grid<T>, cellF: (value: T, cell: Cell) => S) =>
    grid.map((row, y) => row.map((cell, x) => cellF(cell, [x, y])));

export const convertArrayToObject = <K, S>(array: K[], keyF: (k: K) => string, valueF: (i: K) => S) => {
    const initialValue = {} as { [x: string]: S };
    return array.reduce((obj, item) => {
        return {
            ...obj,
            [keyF(item)]: valueF(item),
        };
    }, initialValue);
};
