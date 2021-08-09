import useSWR, {mutate} from "swr";
import {auth} from "../api/auth";
import {IS_SERVER} from "../utils/constants";

const _CURRENT_USER_KEY = "/auth/me"

if (!IS_SERVER) {
    const userId = localStorage.getItem(_CURRENT_USER_KEY)
    if (userId) {
        mutate(_CURRENT_USER_KEY, userId)
    }
}

export const useCurrentUser = () => {
    const { data } = useSWR(_CURRENT_USER_KEY, auth.me, {
        onSuccess(userId) {
            if (userId)
                localStorage.setItem(_CURRENT_USER_KEY, userId)
            else
                localStorage.removeItem(_CURRENT_USER_KEY)
        },
        onError() {
            localStorage.removeItem(_CURRENT_USER_KEY)
        }
    })

    return data;
}

export const signIn = async (userId: string, password: string) => {
    const err = await auth.signIn(userId, password)
    if (!err) {
        await mutate(_CURRENT_USER_KEY, userId, false)
    }
    return err
}

export const signUp = async (userId: string, password: string) => {
    const err = await auth.signUp(userId, password)
    if (!err) {
        await mutate(_CURRENT_USER_KEY, userId, false)
    }
    return err
}

export const signOut = async () => {
    await auth.signOut()
    await mutate(_CURRENT_USER_KEY, undefined)
    localStorage.removeItem(_CURRENT_USER_KEY)
}
