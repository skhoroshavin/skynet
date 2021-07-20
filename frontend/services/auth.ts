
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

    signIn: async (id: string, password: string): Promise<string|null> => {
        const res = await fetch('/api/auth/signin', {
            method: "POST",
            body: JSON.stringify({id, password})
        })
        if (res.status == 200)
            return null
        const body = await res.json()
        return body.err
    },

    me: async (): Promise<string|null> => {
        const res = await fetch('/api/auth/me')
        if (res.status != 200)
            return null
        const body = await res.json()
        return body.id
    },
}
