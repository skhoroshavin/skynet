
export const latinLowerCase = 'abcdefghijklmnopqrstuvwxyz'
export const latinUpperCase = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
export const numeric = '0123456789'
export const alnum = latinLowerCase + latinUpperCase + numeric

export function upTo(n) {
    return Math.floor(Math.random() * n)
}

export function between(min, max) {
    return min + upTo(max + 1 - min)
}

export function oneOf(alphabet) {
    return alphabet[upTo(alphabet.length)]
}

export function stringOf(n, alphabet) {
    let s = ''
    for(let i=0; i<n; ++i) {
        s += oneOf(alphabet)
    }
    return s;
}

export function varStringOf(min, max, alphabet) {
    return stringOf(between(min, max), alphabet)
}
