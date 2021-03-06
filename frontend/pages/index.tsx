import {useRouter} from "next/router";
import {useEffect} from "react";
import {useCurrentUser} from "../services/auth";

export default function Home() {
    const userId = useCurrentUser()

    const router = useRouter()
    useEffect(() => {
        const url = userId ? `/users/${userId}` : "/signin"
        router.push(url)
    }, [userId, router])

    return null
}
