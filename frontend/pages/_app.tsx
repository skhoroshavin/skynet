import '../styles/globals.css'
import type { AppProps } from 'next/app'
import Link from "next/link";
import useSWR, {mutate} from "swr";
import auth from "../services/auth";
import React from "react";

const UserSignedOut = () => {
  return <>
    <Link href="/signin" passHref>
      Sign in
    </Link>
    <Link href="/signup" passHref>
      Sign up
    </Link>
  </>
}

const UserSignedIn = (props: {userId: string|null|undefined}) => {
  const signOut = async (e: React.MouseEvent<HTMLElement>) => {
    e.preventDefault()
    await auth.signOut()
    await mutate("me")
  }

  return <div id="user-id">
    <a href="#/signout" onClick={signOut}>
      Sign out
    </a>
  </div>
}

const UserInfo = (props: {userId: string|null|undefined}) => {
  if (!props.userId) {
    return <UserSignedOut/>
  }
  return <UserSignedIn {...props}/>
}

function SkyNet({ Component, pageProps }: AppProps) {
  const userId = useSWR("me", auth.me).data

  return <>
    <div>
      <div>SkyNET</div>
      <form>
        <input type="search"/>
        <button>Search</button>
      </form>
      <div>
        <UserInfo userId={userId}/>
      </div>
    </div>
    <div style={{paddingTop: "60px"}}>
      <Component {...pageProps}/>
    </div>
  </>
}
export default SkyNet
