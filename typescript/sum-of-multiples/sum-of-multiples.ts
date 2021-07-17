const SumOfMultiples = (divisors: number[]): { to: (to: number) => number } => {
    return {
        to: (to: number) => {
            const factors: Set<number> = new Set<number>()
            divisors.forEach(divisor => {
                for (let i = 2; i < to; i++) {
                    if (i % divisor == 0) {
                        factors.add(i)
                    }
                }
            })
            const sumReducer = (sum: number, currVal: number): number => sum + currVal
            return [...factors].reduce(sumReducer, 0)
        }
    }

}

export default SumOfMultiples