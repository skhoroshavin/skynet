import {SKYNET_API} from "./_constants";

export type UserData = {
    firstName: string,
    lastName: string,
    birthday: string,
    city: string,
    interests: string
}

export const user = {
    exists: async (id: string): Promise<boolean> => {
        // TODO: Use head request
        const res = await fetch(`${SKYNET_API}/users/${id}`)
        return res.status == 200;
    },

    data: async (id: string): Promise<UserData | Error> => {
        const res = await fetch(`${SKYNET_API}/users/${id}`)
        const body = await res.json()

        if (res.status != 200) {
            return new Error(body.message)
        }

        return {
            firstName: body.first_name || null,
            lastName: body.last_name || null,
            birthday: body.birthday || null,
            city: body.city || null,
            interests: body.interests || null
        }
    }
}
