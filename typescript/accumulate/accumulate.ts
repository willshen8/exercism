type GenericCallBack<T> = (input: T) => T

const accumulate = <T>(input:T[], accumulator:GenericCallBack<any>): T[] => {
    return input.map(item => accumulator(item))
}

export default accumulate