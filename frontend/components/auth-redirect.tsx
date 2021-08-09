import {PropsWithChildren, useEffect, useState} from "react";
import {useCurrentUser} from "../services/auth";
import {useRouter} from "next/router";

export const AuthRedirect = ({children}: PropsWithChildren<{}>) => {
    const [mounted, setMounted] = useState(false)
    const userId = useCurrentUser()

    const router = useRouter()
    useEffect(() => {
        if (userId) {
            router.push(`/users/${userId}`)
        } else {
            setMounted(true)
        }
    }, [userId, router])

    if (!mounted) {
        return null
    }

    return <>{children}</>
}
