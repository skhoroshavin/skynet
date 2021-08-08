import auth from "../services/auth";
import useSWR from "swr"
import {useRouter} from "next/router";
import {useEffect} from "react";

export default function Home() {
    const userId = useSWR("me", auth.me).data
    const router = useRouter()

    useEffect(() => {
        const url = userId ? `/users/{userId}` : "/signin"
        router.push(url)
    }, [userId, router])

    return <div className="mx-auto text-xl text-primary-900">
        En route, calculating optimal destination...
    </div>
}
