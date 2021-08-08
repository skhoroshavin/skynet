import {SKYNET_API} from "./_constants";

export const auth = {
    signUp: async (id: string, password: string): Promise<string|null> => {
        const res = await fetch(`${SKYNET_API}/auth/signup`, {
            method: "POST",
            body: JSON.stringify({id, password})
        })
        if (res.status == 200)
            return null
        const body = await res.json()
        return body.err
    },

    signIn: async (id: string, password: string): Promise<string|null> => {
        const res = await fetch(`${SKYNET_API}/auth/signin`, {
            method: "POST",
            body: JSON.stringify({id, password})
        })
        if (res.status == 200)
            return null
        const body = await res.json()
        return body.err
    },

    signOut: async (): Promise<void> => {
        await fetch(`${SKYNET_API}/auth/signout`, {
            method: "POST"
        })
        return
    },

    me: async (): Promise<string|null> => {
        const res = await fetch(`${SKYNET_API}/auth/me`)
        if (res.status != 200)
            return null
        const body = await res.json()
        return body.id
    },
}
