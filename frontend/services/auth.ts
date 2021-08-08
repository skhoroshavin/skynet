import useSWR, {mutate} from "swr";
import {auth} from "../api/auth";

const _CURRENT_USER_KEY = "/auth/me"

export const useCurrentUser = () => {
    const { data } = useSWR(_CURRENT_USER_KEY, auth.me)

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
}
