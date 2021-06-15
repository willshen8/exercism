import { List } from "./sublist"

// tests written myself using TDD approach

describe('Sublist', () => {
    test('Two empty list are equal', () => {
        const listA = new List();
        const listB =  new List();
        const result = listA.compare(listB);
        expect(result).toEqual('EQUAL')
    })

    test('Two list are equal', () => {
        const listA = new List([1,2,3,4,5]);
        const listB =  new List([1,2,3,4,5]);
        const result = listA.compare(listB);
        expect(result).toEqual('EQUAL')
    })

    test('Two list are not equal', () => {
        const listA = new List([1]);
        const listB =  new List([2]);
        const result = listA.compare(listB);
        expect(result).toEqual('UNEQUAL')
    })

    test('First list is a sub list of second list', () => {
        const listA = new List([1,2]);
        const listB =  new List([1,2,3]);
        const result = listA.compare(listB);
        expect(result).toEqual('SUBLIST')
    })

    test('First list is a super list of second list', () => {
        const listA = new List([1,2]);
        const listB =  new List([1]);
        const result = listA.compare(listB);
        expect(result).toEqual('SUPERLIST')
    })

    test('First list is a sub list of second list', () => {
        const listA = new List([1, 2, 3]);
        const listB =  new List([1, 2, 3, 4, 5]);
        const result = listA.compare(listB);
        expect(result).toEqual('SUBLIST')
    })


    test('A is a sublist of B', () => {
        const listA = new List([3, 4]);
        const listB =  new List([1, 2, 3, 4, 5]);
        const result = listA.compare(listB);
        expect(result).toEqual('SUBLIST')
    })

    test('First list is a suerlist list of second list, not in order', () => {
        const listA = new List([1, 2, 3, 4, 5]);
        const listB =  new List([3, 4, 5]);
        const result = listA.compare(listB);
        expect(result).toEqual('SUPERLIST')
    })

    test('False start', () => {
        const listA = new List([1, 2, 3]);
        const listB =  new List([1, 1, 2, 3]);
        const result = listA.compare(listB);
        expect(result).toEqual('SUBLIST')
    })

    test('Not Equal', () => {
        const listA = new List([1, 2, 4]);
        const listB =  new List([1, 2, 3, 4, 5]);
        const result = listA.compare(listB);
        expect(result).toEqual('UNEQUAL')
    })

    test('Not Equal', () => {
        const listA = new List([1, 2, 3]);
        const listB =  new List([1, 2, 3]);
        const result = listA.compare(listB);
        expect(result).toEqual('EQUAL')
    })
})