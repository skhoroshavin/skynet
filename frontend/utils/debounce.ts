
export function debounce<F extends (...args: any[]) => void>(func: F, wait?: number) {
    wait = wait || 500
    let timeoutID: any = null
    return function(this: any, ...args: Parameters<F>) {
        clearTimeout(timeoutID)
        timeoutID = setTimeout(() => func.apply(this, args), wait)
    } as F;
}

export function asyncDebounce<F extends (...args: any[]) => Promise<any>>(func: F, wait?: number) {
    const debounced = debounce((resolve: any, reject: any, args: Parameters<F>) => {
        func(...args).then(resolve).catch(reject);
    }, wait);

    return (...args: Parameters<F>): ReturnType<F> =>
        new Promise((resolve, reject) => {
            debounced(resolve, reject, args);
        }) as ReturnType<F>;
}
