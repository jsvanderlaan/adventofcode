import { convertArrayToObject, distinct, getEmptyArray } from '../utils';

const tpl = 'PHVCVBFHCVPFKBNHKNBO';
const map = {
    HK: 'F',
    VN: 'S',
    NB: 'F',
    HF: 'B',
    CK: 'N',
    VP: 'B',
    HO: 'P',
    NH: 'N',
    CC: 'N',
    FC: 'P',
    OK: 'S',
    OO: 'P',
    ON: 'C',
    VF: 'B',
    NN: 'O',
    KS: 'P',
    FK: 'K',
    HB: 'V',
    SH: 'O',
    OB: 'K',
    PB: 'V',
    BO: 'O',
    NV: 'K',
    CV: 'H',
    PH: 'H',
    KO: 'B',
    BC: 'B',
    KC: 'B',
    SO: 'P',
    CF: 'V',
    VS: 'F',
    OV: 'N',
    NS: 'K',
    KV: 'O',
    OP: 'O',
    HH: 'C',
    FB: 'S',
    CO: 'K',
    SB: 'K',
    SN: 'V',
    OF: 'F',
    BN: 'F',
    CP: 'C',
    NC: 'H',
    VH: 'S',
    HV: 'V',
    NF: 'B',
    SS: 'K',
    FO: 'F',
    VO: 'H',
    KK: 'C',
    PF: 'V',
    OS: 'F',
    OC: 'H',
    SK: 'V',
    FF: 'H',
    PK: 'N',
    PC: 'O',
    SP: 'B',
    CB: 'B',
    CH: 'H',
    FN: 'V',
    SV: 'O',
    SC: 'P',
    NP: 'B',
    BB: 'S',
    PV: 'S',
    VB: 'P',
    SF: 'H',
    VC: 'O',
    HN: 'V',
    BF: 'O',
    NO: 'O',
    HP: 'N',
    VV: 'K',
    HS: 'P',
    FH: 'N',
    KB: 'F',
    KF: 'B',
    PN: 'K',
    KH: 'K',
    CN: 'S',
    PP: 'O',
    BP: 'O',
    OH: 'B',
    FS: 'O',
    BK: 'B',
    PO: 'V',
    CS: 'C',
    BV: 'N',
    KP: 'O',
    KN: 'B',
    VK: 'F',
    HC: 'O',
    BH: 'B',
    FP: 'H',
    NK: 'V',
    BS: 'C',
    FV: 'F',
    PS: 'P',
};

const testTpl = 'NNCB';
const testMap = {
    CH: 'B',
    HH: 'N',
    CB: 'H',
    NH: 'C',
    HB: 'C',
    HC: 'B',
    HN: 'C',
    NN: 'C',
    BH: 'H',
    NC: 'B',
    NB: 'B',
    BN: 'B',
    BB: 'N',
    BC: 'B',
    CC: 'N',
    CN: 'C',
};

const getInserter = (map: { [key: string]: string }) => (str: string) =>
    str.split('').reduce((prev, _, i) => {
        if (i === str.length - 1) return prev;
        const first = str[i];
        const second = str[i + 1];
        const pair = `${first}${second}`;
        return `${prev}${i === 0 ? first : ''}${map[pair]}${second}`;
    }, '');

const insertMany = (map: { [key: string]: string }, times: number, start: string) => {
    const inserter = getInserter(map);
    return getEmptyArray(times).reduce((prev, _) => inserter(prev), start);
};
console.time('insert');
const afterInserts = insertMany(map, 10, tpl);
console.timeEnd('insert');

const countChars = (str: string) =>
    str.split('').reduce((prev, next) => {
        if (prev[next] === undefined) {
            prev[next] = 1;
        } else {
            prev[next]++;
        }
        return prev;
    }, {} as { [key: string]: number });

console.time('count');
const count = countChars(afterInserts);
console.log(count);
const values = Object.values(count);
const max = Math.max(...values);
const min = Math.min(...values);
console.log(max - min);
console.timeEnd('count');

const getStartingCache = (map: { [key: string]: string }) =>
    convertArrayToObject(
        Object.keys(map),
        x => x,
        x => [
            convertArrayToObject(
                x.split(''),
                x => x,
                _ => (x.split('').filter(distinct).length === 2 ? 1 : 2) as number
            ),
        ]
    );

const combineValues = (left: { [x: string]: number }, right: { [x: string]: number }, insert: string) =>
    Object.entries(left).reduce((prev, next) => ({ ...prev, [next[0]]: next[1] + (prev[next[0]] ?? 0) }), {
        ...right,
        [insert]: Math.max(right[insert] - 1, 0),
    });

const getCounts = (map: { [key: string]: string }, str: string, times: number) => {
    const cache = getEmptyArray(times).reduce((curr, _, i) => {
        Object.entries(curr).forEach(([key, value]) => {
            const insert = map[key];
            const left = `${key[0]}${insert}`;
            const right = `${insert}${key[1]}`;

            const valuesLeft = curr[left][i];
            const valuesRight = curr[right][i];
            const values = combineValues(valuesLeft, valuesRight, insert);

            curr[key] = [...value, values];
        });

        return curr;
    }, getStartingCache(map));

    return str.split('').reduce((prev, curr, i) => {
        return i === str.length - 1
            ? { ...prev, [curr]: prev[curr] + 1 }
            : combineValues(prev, cache[`${curr}${str[i + 1]}`][times], str[i + 1]);
    }, {} as { [key: string]: number });
};
console.time('count');
const counts = getCounts(map, tpl, 40);
const countValues = Object.values(counts);
console.log(Math.max(...countValues) - Math.min(...countValues));
console.timeEnd('count');
