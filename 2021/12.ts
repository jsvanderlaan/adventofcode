import { copies, distinct } from '../utils';

type Caves = { [key: string]: string[] };
const test: string[] = ['start-A', 'start-b', 'A-c', 'A-b', 'b-d', 'A-end', 'b-end'];
const test2 = ['dc-end', 'HN-start', 'start-kj', 'dc-start', 'dc-HN', 'LN-dc', 'HN-end', 'kj-sa', 'kj-HN', 'kj-dc'];
const test3 = [
    'fs-end',
    'he-DX',
    'fs-he',
    'start-DX',
    'pj-DX',
    'end-zg',
    'zg-sl',
    'zg-pj',
    'pj-he',
    'RW-he',
    'fs-DX',
    'pj-RW',
    'zg-RW',
    'start-pj',
    'he-WI',
    'zg-he',
    'pj-fs',
    'start-RW',
];

const input = [
    'pn-TY',
    'rp-ka',
    'az-aw',
    'al-IV',
    'pn-co',
    'end-rp',
    'aw-TY',
    'rp-pn',
    'al-rp',
    'end-al',
    'IV-co',
    'end-TM',
    'co-TY',
    'TY-ka',
    'aw-pn',
    'aw-IV',
    'pn-IV',
    'IV-ka',
    'TM-rp',
    'aw-PD',
    'start-IV',
    'start-co',
    'start-pn',
];

const getPaths = (strings: string[]) => strings.map((str) => str.split('-'));

const getCaves = (paths: string[][]) =>
    paths.reduce((prev: Caves, next: string[]) => {
        const setCave = (cave1: string, cave2: string) => (prev[cave1] = prev[cave1] ? [...prev[cave1], cave2] : [cave2]);
        setCave(next[0], next[1]);
        setCave(next[1], next[0]);
        return prev;
    }, {});

const traverse = (caves: Caves, current: string, end: string, visited: string[]): string[][] => {
    if (current === end) return [[current]];
    if (current === current.toLocaleLowerCase() && visited.includes(current)) return [];
    return caves[current].map((path) => traverse(caves, path, end, [...visited, current]).map((p) => [current, ...p])).flat(1);
};

console.time('traverse');
console.log(traverse(getCaves(getPaths(input)), 'start', 'end', []).length);
console.timeEnd('traverse');

const traverse2 = (caves: Caves, current: string, start: string, end: string, visited: string[]): string[][] => {
    if (current === end) return [[current]];
    if (
        current === current.toLocaleLowerCase() &&
        visited.includes(current) &&
        (current === start || visited.filter((cave) => cave === cave.toLocaleLowerCase()).filter(copies).length > 0)
    )
        return [];
    return caves[current]
        .map((path) => traverse2(caves, path, start, end, [...visited, current]).map((p) => [current, ...p]))
        .flat(1);
};

console.time('traverse2');
console.log(traverse2(getCaves(getPaths(input)), 'start', 'start', 'end', []).length);
console.timeEnd('traverse2');
