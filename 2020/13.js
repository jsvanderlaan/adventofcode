const dataT = 1008169;
const data =
  "29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,37,x,x,x,x,x,653,x,x,x,x,x,x,x,x,x,x,x,x,13,x,x,x,17,x,x,x,x,x,23,x,x,x,x,x,x,x,823,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19";

const exampleT = 939;
const example = "7,13,x,x,59,x,31,19";

const timeTileDeparture = data
  .split(",")
  .filter((b) => b !== "x")
  .map((b) => +b)
  .map((b) => ({ id: b, time: b - (dataT % b) }));
console.log(timeTileDeparture);

const ans = timeTileDeparture.reduce((min, next) => (next.time < min.time ? next : min), { id: 0, time: Number.MAX_VALUE });
console.log(ans.time * ans.id);

const timesAndOffset = example
  .split(",")
  .map((t, i) => ({ id: t, offset: i }))
  .filter((x) => x.id !== "x")
  .map((x) => ({ id: +x.id, offset: x.offset }));

console.log(timesAndOffset);

timesAndOffset.reduce(
  ({ min, last }, next) => {
    let minT = min;
    while (minT % next.id !== next.offset) {
      minT += min * last;
    }
    console.log(minT);
    return { min: Math.max(minT, next.id), last: next.id };
  },
  { min: 0, last: 0 }
);
//console.log(minTime);
