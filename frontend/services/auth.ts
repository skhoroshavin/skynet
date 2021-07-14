
export default {
    signUp: async (id: string, password: string): Promise<string|null> => {
        const res = await fetch('/api/auth/signup', {
            method: "POST",
            body: JSON.stringify({id, password})
        })
        if (res.status == 200)
            return null
        const body = await res.json()
        return body.err
    },

    signIn: (id: string, password: string) => {
        return fetch('/api/auth/signup', {
            method: "POST",
            body: JSON.stringify({id, password})
        })
    },

    me: () => {
        return fetch('/api/auth/me')
    },
}