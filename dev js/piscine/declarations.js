const escapeStr = '`, \\, /, " and \'.';

const arr = Object.freeze([4, '2']);

const obj = Object.freeze({
  str: 'string value',
  num: 42,
  bool: true,
  undef: undefined
});
 
const nested = Object.freeze({
  arr: Object.freeze([4, undefined, '2']),
  obj: Object.freeze({
    str: 'nested string value',
    num: 3.14,
    bool: false
  })
});