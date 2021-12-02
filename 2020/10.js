const data = [
  84,
  60,
  10,
  23,
  126,
  2,
  128,
  63,
  59,
  69,
  127,
  73,
  140,
  55,
  154,
  133,
  36,
  139,
  4,
  70,
  110,
  97,
  153,
  105,
  41,
  106,
  79,
  145,
  35,
  134,
  146,
  148,
  13,
  77,
  49,
  107,
  46,
  138,
  88,
  152,
  83,
  120,
  52,
  114,
  159,
  158,
  53,
  76,
  16,
  28,
  89,
  25,
  42,
  66,
  119,
  3,
  17,
  67,
  94,
  99,
  7,
  56,
  85,
  122,
  18,
  20,
  43,
  160,
  54,
  113,
  29,
  130,
  19,
  135,
  30,
  80,
  116,
  91,
  161,
  115,
  141,
  102,
  37,
  157,
  129,
  34,
  147,
  142,
  151,
  68,
  78,
  24,
  90,
  121,
  123,
  33,
  98,
  1,
  40,
  0,
];

const example = [0, 28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3];

const [, one, three] = data
  .sort((a, b) => a - b)
  .reduce(
    ([prev, one, three], next) => {
      if (next - prev === 1) one++;
      if (next - prev === 3) three++;
      return [next, one, three];
    },
    [0, 0, 1]
  );

console.log(one * three);

const getEdges = (list) =>
  list
    .sort((a, b) => b - a)
    .map((x, i) => ({
      num: x,
      parents: list.slice(i + 1, i + 4).filter((a) => x - a <= 3),
      childs: list.slice(Math.max(i - 3, 0), Math.max(i, 0)).filter((a) => x - a <= 3),
    }));

const edges = getEdges(data);
const max = Math.max(...edges.map((edge) => edge.num));
const map = new Map();
edges.sort((a, b) => a.num - b.num).forEach((edge) => map.set(edge.num, edge.parents.reduce((acc, next) => acc + map.get(next), 0) || 1));

console.log(map.get(max));
